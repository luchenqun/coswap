package keeper

import (
	"context"
	"coswap/x/swap/types"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SwapCoin(goCtx context.Context, msg *types.MsgSwapCoin) (*types.MsgSwapCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if !(msg.AmountIn.Denom == "eth" && msg.AmountOutMin.Denom == "btc" || msg.AmountIn.Denom == "btc" && msg.AmountOutMin.Denom == "eth") {
		return nil, errors.New("Only supports the exchange of ETH and BTC coin pairs")
	}

	var amountOut sdk.Coin

	if msg.AmountIn.Denom == "eth" {
		if msg.AmountIn.Amount.GTE(sdk.NewInt(10)) {
			amountOut = sdk.NewCoin("btc", msg.AmountIn.Amount.QuoRaw(10))
		} else {
			return nil, errors.New("amount in eth amount must >= 10")
		}
	}

	if msg.AmountIn.Denom == "btc" {
		if msg.AmountIn.Amount.GT(sdk.NewInt(0)) {
			amountOut = sdk.NewCoin("eth", msg.AmountIn.Amount.MulRaw(10))
		} else {
			return nil, errors.New("amount in btc amount must > 0")
		}
	}

	if amountOut.IsGTE(msg.AmountOutMin) {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.Coins{amountOut})
		if err != nil {
			return nil, err
		}

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.Coins{msg.AmountIn})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("amount out less than amount out min")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute("amount-out", amountOut.String()),
		),
	)

	return &types.MsgSwapCoinResponse{}, nil
}
