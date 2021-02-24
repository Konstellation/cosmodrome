package main

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/konstellation/cosmodrome/cmd"
	"github.com/konstellation/kn-sdk/types"
	"github.com/konstellation/konstellation/app"
)

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	types.RegisterNativeCoinUnits()
	types.RegisterBech32Prefix()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "cosmodrome",
		Short:             "cosmodrome",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		cmd.GenNetCmd(ctx, cdc, app.ModuleBasics, app.GenesisUpdaters, genaccounts.AppModuleBasic{}),
	)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, app.EnvPrefixNode, app.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
