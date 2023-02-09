package keeper_test

import (
	"testing"

	testkeeper "IOTrustChain/testutil/keeper"
	"IOTrustChain/x/iotrustchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IotrustchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
