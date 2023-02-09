package iotrustchain_test

import (
	"testing"

	keepertest "IOTrustChain/testutil/keeper"
	"IOTrustChain/testutil/nullify"
	"IOTrustChain/x/iotrustchain"
	"IOTrustChain/x/iotrustchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IotrustchainKeeper(t)
	iotrustchain.InitGenesis(ctx, *k, genesisState)
	got := iotrustchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
