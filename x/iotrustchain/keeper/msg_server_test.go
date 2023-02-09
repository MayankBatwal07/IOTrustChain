package keeper_test

import (
	"context"
	"testing"

	keepertest "IOTrustChain/testutil/keeper"
	"IOTrustChain/x/iotrustchain/keeper"
	"IOTrustChain/x/iotrustchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IotrustchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
