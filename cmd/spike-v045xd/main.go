package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"
	cmdcfg "spike-v045x/cmd/config"
	"spike-v045x/cmd/spike-v045xd/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"spike-v045x/app"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
}
