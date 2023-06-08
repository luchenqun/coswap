package keeper

import (
	"context"

	"coswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DistributeCoins(goCtx context.Context, msg *types.MsgDistributeCoins) (*types.MsgDistributeCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDistributeCoinsResponse{}, nil
}
