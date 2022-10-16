package keeper

import (
	"exchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
