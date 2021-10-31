package main

import (
	"os"

	"github.com/sachinsmc/risk-prediction-module/cmd/risk-prediction-moduled/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sachinsmc/risk-prediction-module/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
