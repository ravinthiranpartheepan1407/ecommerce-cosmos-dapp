package raviecommercetest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ravi-ecommerce-test/testutil/keeper"
	"ravi-ecommerce-test/testutil/nullify"
	"ravi-ecommerce-test/x/raviecommercetest"
	"ravi-ecommerce-test/x/raviecommercetest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RaviecommercetestKeeper(t)
	raviecommercetest.InitGenesis(ctx, *k, genesisState)
	got := raviecommercetest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
