package keeper_test

import (
	"context"
	"testing"

	keepertest "coswap/testutil/keeper"
	"coswap/x/swap/keeper"
	"coswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SwapKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
