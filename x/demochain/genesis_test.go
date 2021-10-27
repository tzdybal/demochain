package demochain_test

import (
	"testing"

	keepertest "github.com/celestiaorg/demochain/testutil/keeper"
	"github.com/celestiaorg/demochain/x/demochain"
	"github.com/celestiaorg/demochain/x/demochain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DemochainKeeper(t)
	demochain.InitGenesis(ctx, *k, genesisState)
	got := demochain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
