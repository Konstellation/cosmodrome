package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/konstellation/cosmodrome/types"
	"github.com/konstellation/kn-sdk/crypto/keybase"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	common2 "github.com/konstellation/cosmodrome/common"
	"github.com/konstellation/kn-sdk/common/utils"
	kntypes "github.com/konstellation/kn-sdk/types"
)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCliHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"
	flagKeyStorageFile    = "key-storage"
	flagNetConfigFile     = "net-config"

	outDir             = ""
	gentxsDir          = ""
	configDir          = ""
	chainID            = ""
	nodeDaemonHomeName = ""
	nodeCliHomeName    = ""
)

const nodeDirPerm = 0755

// get cmd to initialize all files for tendermint localnet and application
func LocalnetCmd(
	ctx *server.Context,
	cdc *codec.Codec,
	mbm module.BasicManager,
	gus kntypes.GenesisUpdaters,
	_ genutilcli.StakingMsgBuildingHelpers,
	genAccIterator genutiltypes.GenesisAccountsIterator,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "localnet",
		Short: "Initialize files for a Konstellation localnet",
		Long: `localnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	cosmodrome localnet --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(_ *cobra.Command, _ []string) error {
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

			chainID = viper.GetString(client.FlagChainID)
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

			validators, err := configValidators(config, configFile, netConfig.Validators, keyStorage, accs)
			if err != nil {
				return err
			}

			if err := initGenFiles(cdc, mbm, gus, validators, accs, config); err != nil {
				return err
			}

			if err := genTxs(cdc, mbm, genAccIterator, validators, keyStorage); err != nil {
				return err
			}

			if err := clientConfig(config, configFile, validators); err != nil {
				return err
			}

			if err := collectGenFiles(cdc, config, genaccounts.AppModuleBasic{}, validators); err != nil {
				return err
			}

			fmt.Printf("Successfully initialized %d node directories\n", len(validators))
			return nil
		},
	}

	cmd.Flags().StringP(flagOutputDir, "o", "./localnet",
		"Directory to store initialization data for the localnet",
	)
	cmd.Flags().String(flagNodeDirPrefix, "node",
		"Prefix the directory name for each node with (node results in node0, node1, ...)",
	)
	cmd.Flags().String(flagNodeDaemonHome, "konstellation",
		"Home directory of the node's daemon configuration",
	)
	cmd.Flags().String(flagNodeCliHome, "konstellationcli",
		"Home directory of the node's cli configuration",
	)
	cmd.Flags().String(flagNetConfigFile, "./config/localnet.json",
		"Net configuration file",
	)
	cmd.Flags().String(flagKeyStorageFile, "./config/keys.json",
		"Keys file",
	)
	cmd.Flags().String(flagStartingIPAddress, "testnode",
		"Starting IP address (testnode results in persistent peers list ID0@testnode-0:26656, ID1@testnode-1:26656, ...)")

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

func genAccounts(genaccs []*types.GenAccount) ([]*genaccounts.GenesisAccount, error) {
	accs := make([]*genaccounts.GenesisAccount, 0)
	for _, ga := range genaccs {
		addr, err := sdk.AccAddressFromBech32(ga.Address)
		if err != nil {
			return nil, err
		}

		genacc := &genaccounts.GenesisAccount{
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

func configValidator(config *cfg.Config, configFile *srvconfig.Config, info *types.ValidatorInfo, key *types.Key, genAccount *genaccounts.GenesisAccount) (validator *types.Validator, err error) {
	nodeDir := filepath.Join(outDir, info.Name, nodeDaemonHomeName)
	cliDir := filepath.Join(outDir, info.Name, nodeCliHomeName)
	nodeConfig := types.NodeConfig{
		DirName:   info.Name,
		DaemonDir: nodeDir,
		CliDir:    cliDir,
	}

	config.SetRoot(nodeDir)
	config.Moniker = info.Name
	if info.Cors != "" {
		config.RPC.CORSAllowedOrigins = strings.Split(info.Cors, ",")
	} else {
		config.RPC.CORSAllowedOrigins = []string{}
	}

	if err := os.MkdirAll(cliDir, nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	configFilePath := filepath.Join(nodeDir, "config/app.toml")
	srvconfig.WriteConfigFile(configFilePath, configFile)

	if err := saveKey(nodeConfig.CliDir, key); err != nil {
		return nil, err
	}

	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	memo := fmt.Sprintf("%s@%s:26656", nodeID, info.IP)

	return &types.Validator{
		Index:       info.Index,
		Moniker:     info.Description.Moniker,
		NodeConfig:  nodeConfig,
		GenFile:     config.GenesisFile(),
		Memo:        memo,
		ID:          nodeID,
		ChainID:     chainID,
		Cors:        info.Cors,
		ValPubKey:   valPubKey,
		IP:          info.IP,
		Key:         info.Key,
		Description: info.Description,
		GenAccount:  genAccount,
	}, nil
}

func configValidators(config *cfg.Config, configFile *srvconfig.Config,
	validatorInfos []*types.ValidatorInfo,
	keyStorage *types.KeyStorage,
	genAccounts []*genaccounts.GenesisAccount) (validators []*types.Validator, err error) {
	for _, valInfo := range validatorInfos {
		key, err := keyStorage.GetKey(valInfo.Key.Address)
		if err != nil {
			return nil, err
		}

		addr, err := sdk.AccAddressFromBech32(valInfo.Key.Address)
		if err != nil {
			return nil, err
		}

		var genAccount *genaccounts.GenesisAccount
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

func initGenFiles(cdc *codec.Codec, mbm module.BasicManager, gus kntypes.GenesisUpdaters, validators []*types.Validator, accs []*genaccounts.GenesisAccount, config *cfg.Config) error {
	appGenState := mbm.DefaultGenesis()

	appGenState[genaccounts.ModuleName] = cdc.MustMarshalJSON(accs)

	// Update default genesis
	for _, gu := range gus {
		gu.UpdateGenesis(cdc, appGenState)
	}

	if err := mbm.ValidateGenesis(appGenState); err != nil {
		return fmt.Errorf("error validating genesis: %s", err.Error())
	}

	appState := cdc.MustMarshalJSON(appGenState)

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
	cdc *codec.Codec,
	mbm module.BasicManager,
	genAccIterator genutiltypes.GenesisAccountsIterator,
	validators []*types.Validator,
	keyStorage *types.KeyStorage,
) error {
	for _, validator := range validators {
		genDoc, err := tmtypes.GenesisDocFromFile(validator.GenFile)
		if err != nil {
			return err
		}

		var genesisState map[string]json.RawMessage
		if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
			return err
		}
		if err = mbm.ValidateGenesis(genesisState); err != nil {
			return err
		}

		kb, err := client.NewKeyBaseFromDir(validator.NodeConfig.CliDir)
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
		if err := genutil.ValidateAccountInGenesis(genesisState, genAccIterator, key.GetAddress(), coins, cdc); err != nil {
			return err
		}

		msg := staking.NewMsgCreateValidator(
			sdk.ValAddress(validator.GenAccount.Address),
			validator.ValPubKey,
			c,
			staking.NewDescription(
				validator.Description.Moniker,
				validator.Description.Identity,
				validator.Description.Website,
				validator.Description.Details,
			),
			staking.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			sdk.OneInt(),
		)

		tx := auth.NewStdTx([]sdk.Msg{msg}, auth.StdFee{}, []auth.StdSignature{}, validator.Memo)
		txBldr := auth.NewTxBuilderFromCLI().WithChainID(chainID).WithMemo(validator.Memo).WithKeybase(kb)
		signedTx, err := txBldr.SignStdTx(valKey.Name, valKey.Password, tx, false)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		txBytes, err := cdc.MarshalJSON(signedTx)
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
	cdc *codec.Codec,
	config *cfg.Config,
	genAccIterator genutiltypes.GenesisAccountsIterator,
	validators []*types.Validator,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for _, validator := range validators {
		config.SetRoot(validator.NodeConfig.DaemonDir)
		config.Moniker = validator.Moniker
		initCfg := genutil.NewInitConfig(chainID, gentxsDir, validator.Moniker, validator.ID, validator.ValPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(validator.GenFile)
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc, genAccIterator)
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
