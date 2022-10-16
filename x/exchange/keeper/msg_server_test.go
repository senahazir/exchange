package keeper_test

import (
	"context"
	"testing"

	keepertest "exchange/testutil/keeper"
	"exchange/x/exchange/keeper"
	"exchange/x/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ExchangeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
