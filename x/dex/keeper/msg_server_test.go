package keeper_test

import (
	"context"
	"testing"

	keepertest "exchange/testutil/keeper"
	"exchange/x/dex/keeper"
	"exchange/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
