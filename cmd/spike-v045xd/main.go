package main

import (
	cmdcfg "github.com/Fueav/spike-v045x/cmd/config"
	"github.com/Fueav/spike-v045x/cmd/spike-v045xd/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"

	"github.com/Fueav/spike-v045x/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
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
