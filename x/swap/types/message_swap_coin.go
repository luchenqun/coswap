package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwapCoin = "swap_coin"

var _ sdk.Msg = &MsgSwapCoin{}

func NewMsgSwapCoin(creator string, amountIn sdk.Coin, amountOutMin sdk.Coin) *MsgSwapCoin {
	return &MsgSwapCoin{
		Creator:      creator,
		AmountIn:     amountIn,
		AmountOutMin: amountOutMin,
	}
}

func (msg *MsgSwapCoin) Route() string {
	return RouterKey
}

func (msg *MsgSwapCoin) Type() string {
	return TypeMsgSwapCoin
}

func (msg *MsgSwapCoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwapCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwapCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
