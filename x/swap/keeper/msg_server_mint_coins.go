package keeper

import (
	"context"
	"coswap/x/swap/types"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k msgServer) MintCoins(goCtx context.Context, msg *types.MsgMintCoins) (*types.MsgMintCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return nil, err
	}

	isAdmin := k.IsAdmin(ctx, admin)
	if !isAdmin {
		return nil, errors.New(msg.Admin + " is not admin")
	}

	err = k.Keeper.bankKeeper.MintCoins(ctx, banktypes.ModuleName, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Admin),
		),
	)

	return &types.MsgMintCoinsResponse{}, nil
}
