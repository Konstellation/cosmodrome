package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/iavl/common"

	cfg "github.com/tendermint/tendermint/config"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	clientFlags "github.com/cosmos/cosmos-sdk/client/flags"
	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	common2 "github.com/konstellation/cosmodrome/common"
	"github.com/konstellation/cosmodrome/types"
	"github.com/konstellation/kn-sdk/common/utils"
	"github.com/konstellation/kn-sdk/crypto/keybase"
	kntypes "github.com/konstellation/kn-sdk/types"
)

const nodeDirPerm = 0755

var (
	flagOutputDir      = "output-dir"
	flagNodeDaemonHome = "node-daemon-home"
	flagNodeCliHome    = "node-cli-home"
	flagKeyStorageFile = "key-storage"
	flagNetConfigFile  = "net-config"

	outDir             = ""
	gentxsDir          = ""
	configDir          = ""
	chainID            = ""
	nodeDaemonHomeName = ""
	nodeCliHomeName    = ""
)

func GenNetCmd(
	ctx *server.Context,
	cdc *codec.JSONMarshaler,
	mbm module.BasicManager,
	gus kntypes.GenesisUpdaters,
	genBalIterator banktypes.GenesisBalancesIterator,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate-network",
		Aliases: []string{"gen-net", "gn"},
		Short:   "Initialize files for a Konstellation network",
		Long: `Command will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	cosmodrome gn --chain-id darchub -n ./config/localnet.json  -o ./localnet
	cosmodrome generate-network --chain-id darchub --net-config ./config/testnet.json  --output-dir ./testnet
	`,
		RunE: func(_ *cobra.Command, _ []string) error {
			return generateNetwork(ctx, cdc, mbm, gus, genBalIterator)
		},
	}

	cmd.Flags().StringP(flagOutputDir, "o", "./net",
		"Directory to store initialization data for the network",
	)
	cmd.Flags().StringP(flagNodeDaemonHome, "d", ".konstellation",
		"Home directory of the node's daemon configuration",
	)
	cmd.Flags().StringP(flagNodeCliHome, "c", ".konstellationcli",
		"Home directory of the node's cli configuration",
	)
	cmd.Flags().StringP(flagNetConfigFile, "n", "./config/net.json",
		"Net configuration file",
	)
	cmd.Flags().StringP(flagKeyStorageFile, "k", "./config/keys.json",
		"Keys file",
	)
	cmd.Flags().String(
		server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", kntypes.StakeDenom),
		"Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01apple,0.001darc)",
	)

	return cmd
}

func parseKeyStorage(keyStorageFile string) (*types.KeyStorage, error) {
	var keyStorage types.KeyStorage
	if err := utils.ReadJson(keyStorageFile, &keyStorage); err != nil {
		return nil, err
	}

	return &keyStorage, nil
}

func saveKey(cliDir string, key *types.Key) error {
	_, secret, err := keybase.SaveCoinKey(cliDir, key.Name, key.Password, key.Mnemonic, true)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return err
	}

	info := map[string]string{"secret": secret}
	cliPrint, err := json.Marshal(info)
	if err != nil {
		return err
	}

	if err := utils.WriteFile(fmt.Sprintf("%v.json", "key_seed"), cliDir, cliPrint); err != nil {
		return err
	}

	return nil
}

func parseNetConfig(netConfigFile string) (*types.NetConfig, error) {
	var netConfig types.NetConfig

	if err := utils.ReadJson(netConfigFile, &netConfig); err != nil {
		return nil, err
	}

	return &netConfig, nil
}

func genAccounts(genaccs []*types.GenAccount) ([]*authtypes.GenesisAccount, error) {
	accs := make([]*authtypes.GenesisAccount, 0)
	for _, ga := range genaccs {
		addr, err := sdk.AccAddressFromBech32(ga.Address)
		if err != nil {
			return nil, err
		}

		genacc := authtypes.GenesisAccount{
			Address: addr,
			Coins: sdk.NewCoins(
				sdk.NewCoin(kntypes.StakeDenom, sdk.NewInt(ga.CoinGenesis)),
			),
		}

		accs = append(accs, genacc)
	}

	return accs, nil
}

func clientConfig(config *cfg.Config, configFile *srvconfig.Config, validators []*types.Validator) (err error) {
	var addressesIPs []string

	for _, validator := range validators {
		addressesIPs = append(addressesIPs, validator.Memo)
	}

	sort.Strings(addressesIPs)
	config.SetRoot(configDir)
	config.Moniker = ""
	config.RPC.CORSAllowedOrigins = []string{"*"}
	config.P2P.PersistentPeers = strings.Join(addressesIPs, ",")

	if err := os.MkdirAll(filepath.Join(configDir, "config"), nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return err
	}

	configFilePath := filepath.Join(configDir, "config/app.toml")
	srvconfig.WriteConfigFile(configFilePath, configFile)
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	return nil
}

func configValidator(config *cfg.Config, configFile *srvconfig.Config,
	valInfo *types.ValidatorInfo,
	key *types.Key,
	genAccount *authtypes.GenesisAccount,
) (*types.Validator, error) {
	nodeDir := filepath.Join(outDir, valInfo.Name, nodeDaemonHomeName)
	cliDir := filepath.Join(outDir, valInfo.Name, nodeCliHomeName)
	if err := os.MkdirAll(cliDir, nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	nodeConfig := types.NodeConfig{
		DirName:   valInfo.Name,
		DaemonDir: nodeDir,
		CliDir:    cliDir,
	}

	configFilePath := filepath.Join(nodeDir, "config/app.toml")
	srvconfig.WriteConfigFile(configFilePath, configFile)

	config.SetRoot(nodeDir)
	config.Moniker = valInfo.Description.Moniker
	if valInfo.Config != nil {
		if err := mapstructure.Decode(valInfo.Config, config); err != nil {
			return nil, err
		}
	}

	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	memo := fmt.Sprintf("%s@%s:26656", nodeID, valInfo.IP)

	if err := saveKey(nodeConfig.CliDir, key); err != nil {
		return nil, err
	}

	return &types.Validator{
		Index:       valInfo.Index,
		Moniker:     valInfo.Description.Moniker,
		NodeConfig:  nodeConfig,
		GenFile:     config.GenesisFile(),
		Memo:        memo,
		ID:          nodeID,
		ChainID:     chainID,
		Cors:        valInfo.Cors,
		ValPubKey:   valPubKey,
		IP:          valInfo.IP,
		Key:         valInfo.Key,
		Description: valInfo.Description,
		GenAccount:  genAccount,
	}, nil
}

func configValidators(config *cfg.Config, configFile *srvconfig.Config,
	keyStorage *types.KeyStorage,
	genAccounts []*authtypes.GenesisAccount,
	netConfig *types.NetConfig) (validators []*types.Validator, err error) {
	if netConfig.GlobalConfig != nil {
		if err := mapstructure.Decode(netConfig.GlobalConfig, config); err != nil {
			return nil, err
		}
	}

	for _, valInfo := range netConfig.Validators {
		key, err := keyStorage.GetKey(valInfo.Key.Address)
		if err != nil {
			return nil, err
		}

		addr, err := sdk.AccAddressFromBech32(valInfo.Key.Address)
		if err != nil {
			return nil, err
		}

		var genAccount *authtypes.GenesisAccount
		for _, gacc := range genAccounts {
			if gacc.Address.Equals(addr) {
				genAccount = gacc
			}
		}

		node, err := configValidator(config, configFile, valInfo, key, genAccount)
		if err != nil {
			return nil, err
		}

		validators = append(validators, node)
	}

	return
}

func initGenFiles(cdc *codec.JSONMarshaler, mbm module.BasicManager, gus kntypes.GenesisUpdaters, validators []*types.Validator, accs []*authtypes.GenesisAccount, config *cfg.Config) error {
	appGenState := mbm.DefaultGenesis(*cdc)

	appGenState[authtypes.ModuleName] = (*cdc).MustMarshalJSON(accs)

	// Update default genesis
	for _, gu := range gus {
		gu.UpdateGenesis(*cdc, appGenState)
	}

	// todo - nil??
	if err := mbm.ValidateGenesis(*cdc, nil, appGenState); err != nil {
		return fmt.Errorf("error validating genesis: %s", err.Error())
	}

	appState := (*cdc).MustMarshalJSON(appGenState)

	genDoc := &tmtypes.GenesisDoc{}
	genDoc.ChainID = chainID
	genDoc.Validators = nil
	genDoc.AppState = appState

	// generate empty genesis files for each validator and save
	for _, validator := range validators {
		if err := genutil.ExportGenesisFile(genDoc, validator.GenFile); err != nil {
			return err
		}

		toPrint := common2.NewPrintInfo(validator.Moniker, chainID, validator.ID, "", appState)
		if err := common2.DisplayInfo(cdc, toPrint); err != nil {
			return err
		}
	}

	return nil
}

func genTxs(
	cdc *codec.JSONMarshaler,
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	validators []*types.Validator,
	keyStorage *types.KeyStorage,
) error {
	for _, validator := range validators {
		genDoc, err := tmtypes.GenesisDocFromFile(validator.GenFile)
		if err != nil {
			return err
		}

		var genesisState map[string]json.RawMessage
		if err = (*cdc).UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
			return err
		}

		// todo - nil??
		if err = mbm.ValidateGenesis(*cdc, nil, genesisState); err != nil {
			return err
		}

		kb, err := clkeys.NewLegacyKeyBaseFromDir(validator.NodeConfig.CliDir)
		if err != nil {
			return err
		}

		valKey, err := keyStorage.GetKey(validator.Key.Address)
		if err != nil {
			return err
		}

		key, err := kb.Get(valKey.Name)
		if err != nil {
			return err
		}

		c := sdk.NewCoin(kntypes.StakeDenom, sdk.NewInt(validator.Key.CoinDelegate))
		coins := sdk.NewCoins(c)
		if err := genutil.ValidateAccountInGenesis(genesisState, genBalIterator, key.GetAddress(), coins, *cdc); err != nil {
			return err
		}

		msg, err := stakingtypes.NewMsgCreateValidator(
			sdk.ValAddress(*validator.GenAccount.Address),
			validator.ValPubKey,
			c,
			stakingtypes.NewDescription(
				validator.Description.Moniker,
				validator.Description.Identity,
				validator.Description.Website,
				validator.Description.SecurityContact,
				validator.Description.Details,
			),
			stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			sdk.OneInt(),
		)
		if err != nil {
			return err
		}

		tx := auth.NewStdTx([]sdk.Msg{msg}, auth.StdFee{}, []auth.StdSignature{}, validator.Memo)
		txBldr := auth.NewTxBuilderFromCLI().WithChainID(chainID).WithMemo(validator.Memo).WithKeybase(kb)
		signedTx, err := txBldr.SignStdTx(valKey.Name, valKey.Password, tx, false)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		txBytes, err := (*cdc).MarshalJSON(signedTx)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		if err := utils.WriteFile(fmt.Sprintf("%v.json", validator.Moniker), gentxsDir, txBytes); err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}
	}

	return nil
}

func collectGenFiles(
	cdc *codec.JSONMarshaler,
	config *cfg.Config,
	genBalIterator banktypes.GenesisBalancesIterator,
	validators []*types.Validator,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for _, validator := range validators {
		config.SetRoot(validator.NodeConfig.DaemonDir)
		config.Moniker = validator.Moniker
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, validator.Moniker, validator.ValPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(validator.GenFile)
		if err != nil {
			return err
		}

		// TODO - nil?
		nodeAppState, err := genutil.GenAppStateFromConfig(*cdc, nil, config, initCfg, *genDoc, genBalIterator)
		if err != nil {
			return err
		}

		// set the canonical application state (they should not differ)
		if appState == nil {
			appState = nodeAppState
		}

		// overwrite each validator's genesis file to have a canonical genesis time
		genFile := config.GenesisFile()
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	// genesis client
	config.SetRoot(configDir)
	return genutil.ExportGenesisFileWithTime(config.GenesisFile(), chainID, nil, appState, genTime)
}

func generateNetwork(ctx *server.Context,
	cdc *codec.JSONMarshaler,
	mbm module.BasicManager,
	gus kntypes.GenesisUpdaters,
	genBalIterator banktypes.GenesisBalancesIterator,
) error {
	config := ctx.Config
	configFile := srvconfig.DefaultConfig()
	configFile.MinGasPrices = viper.GetString(server.FlagMinGasPrices)
	netConfigFile := viper.GetString(flagNetConfigFile)
	keyStorageFile := viper.GetString(flagKeyStorageFile)

	nodeDaemonHomeName = viper.GetString(flagNodeDaemonHome)
	nodeCliHomeName = viper.GetString(flagNodeCliHome)

	outDir = viper.GetString(flagOutputDir)
	gentxsDir = filepath.Join(outDir, "gentxs")
	configDir = filepath.Join(outDir)

	chainID = viper.GetString(clientFlags.FlagChainID)
	if chainID == "" {
		chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
	}

	keyStorage, err := parseKeyStorage(keyStorageFile)
	if err != nil {
		return err
	}

	netConfig, err := parseNetConfig(netConfigFile)
	if err != nil {
		return err
	}

	accs, err := genAccounts(netConfig.GenAccounts)
	if err != nil {
		return err
	}

	validators, err := configValidators(config, configFile, keyStorage, accs, netConfig)
	if err != nil {
		return err
	}

	if err := initGenFiles(cdc, mbm, gus, validators, accs, config); err != nil {
		return err
	}

	if err := genTxs(cdc, mbm, genBalIterator, validators, keyStorage); err != nil {
		return err
	}

	if err := clientConfig(config, configFile, validators); err != nil {
		return err
	}

	if err := collectGenFiles(cdc, config, genBalIterator, validators); err != nil {
		return err
	}

	fmt.Printf("Successfully initialized %d node directories\n", len(validators))
	return nil
}
