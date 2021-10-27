package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/celestiaorg/demochain/testutil/keeper"
	"github.com/celestiaorg/demochain/x/demochain/keeper"
	"github.com/celestiaorg/demochain/x/demochain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DemochainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
