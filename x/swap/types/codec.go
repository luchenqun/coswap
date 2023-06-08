package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintCoins{}, "swap/MintCoins", nil)
	cdc.RegisterConcrete(&MsgDistributeCoins{}, "swap/DistributeCoins", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "swap/AddLiquidity", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintCoins{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDistributeCoins{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLiquidity{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
