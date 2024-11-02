package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "evmos/testutil/keeper"
	"evmos/x/evm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EvmKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
