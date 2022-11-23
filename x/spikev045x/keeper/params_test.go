package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "spike-v045x/testutil/keeper"
	"spike-v045x/x/spikev045x/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Spikev045xKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
