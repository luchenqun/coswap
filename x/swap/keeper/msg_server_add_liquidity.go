package keeper

import (
	"context"
	"coswap/x/swap/types"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if msg.TokenA.Amount.IsZero() || msg.TokenB.Amount.IsZero() {
		return nil, errors.New("token must great than 0")
	}

	err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.Coins{msg.TokenA, msg.TokenB}.Sort())
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	)

	return &types.MsgAddLiquidityResponse{}, nil
}
