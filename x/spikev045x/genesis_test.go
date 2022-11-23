package spikev045x_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "spike-v045x/testutil/keeper"
	"spike-v045x/testutil/nullify"
	"spike-v045x/x/spikev045x"
	"spike-v045x/x/spikev045x/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Spikev045xKeeper(t)
	spikev045x.InitGenesis(ctx, *k, genesisState)
	got := spikev045x.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
