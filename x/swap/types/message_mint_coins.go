package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintCoins = "mint_coins"

var _ sdk.Msg = &MsgMintCoins{}

func NewMsgMintCoins(admin string, amount sdk.Coins) *MsgMintCoins {
	return &MsgMintCoins{
		Admin:  admin,
		Amount: amount,
	}
}

func (msg *MsgMintCoins) Route() string {
	return RouterKey
}

func (msg *MsgMintCoins) Type() string {
	return TypeMsgMintCoins
}

func (msg *MsgMintCoins) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgMintCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}
