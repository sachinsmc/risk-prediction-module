package keeper

import (
	"github.com/sachinsmc/risk-prediction-module/x/riskpredictionmodule/types"
)

var _ types.QueryServer = Keeper{}
