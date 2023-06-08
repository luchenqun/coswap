package keeper

import (
	"coswap/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Admin(ctx),
		k.CommissionRate(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Admin returns the Admin param
func (k Keeper) Admin(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyAdmin, &res)
	return
}

// CommissionRate returns the CommissionRate param
func (k Keeper) CommissionRate(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyCommissionRate, &res)
	return
}
