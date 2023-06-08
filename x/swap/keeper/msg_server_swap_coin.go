package keeper

import (
	"context"

	"coswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SwapCoin(goCtx context.Context, msg *types.MsgSwapCoin) (*types.MsgSwapCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapCoinResponse{}, nil
}
