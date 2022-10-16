package keeper

import (
	"exchange/x/exchange/types"
)

var _ types.QueryServer = Keeper{}
