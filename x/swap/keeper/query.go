package keeper

import (
	"coswap/x/swap/types"
)

var _ types.QueryServer = Keeper{}
