package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDistributeCoins = "distribute_coins"

var _ sdk.Msg = &MsgDistributeCoins{}

func NewMsgDistributeCoins(admin string, toAddress string, amount sdk.Coins) *MsgDistributeCoins {
	return &MsgDistributeCoins{
		Admin:     admin,
		ToAddress: toAddress,
		Amount:    amount,
	}
}

func (msg *MsgDistributeCoins) Route() string {
	return RouterKey
}

func (msg *MsgDistributeCoins) Type() string {
	return TypeMsgDistributeCoins
}

func (msg *MsgDistributeCoins) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgDistributeCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDistributeCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}
