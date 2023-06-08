package keeper_test

import (
	"testing"

	testkeeper "coswap/testutil/keeper"
	"coswap/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Admin, k.Admin(ctx))
	require.EqualValues(t, params.CommissionRate, k.CommissionRate(ctx))
}
