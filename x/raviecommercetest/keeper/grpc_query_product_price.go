package keeper

import (
	"context"

    "ravi-ecommerce-test/x/raviecommercetest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProductPrice(goCtx context.Context,  req *types.QueryProductPriceRequest) (*types.QueryProductPriceResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Process the query
    _ = ctx

	return &types.QueryProductPriceResponse{}, nil
}
