package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "ravi-ecommerce-test/testutil/keeper"
	"ravi-ecommerce-test/x/raviecommercetest/keeper"
	"ravi-ecommerce-test/x/raviecommercetest/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RaviecommercetestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
