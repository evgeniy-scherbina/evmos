package keeper

import (
	"evmos/x/evm/types"
)

var _ types.QueryServer = Keeper{}
