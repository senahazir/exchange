package exchange_test

import (
	"testing"

	keepertest "exchange/testutil/keeper"
	"exchange/testutil/nullify"
	"exchange/x/exchange"
	"exchange/x/exchange/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PoolList: []types.Pool{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PoolCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExchangeKeeper(t)
	exchange.InitGenesis(ctx, *k, genesisState)
	got := exchange.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	require.Equal(t, genesisState.PoolCount, got.PoolCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
