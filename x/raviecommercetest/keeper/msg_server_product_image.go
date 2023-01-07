package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"ravi-ecommerce-test/x/raviecommercetest/types"
)

func (k msgServer) ProductImage(goCtx context.Context, msg *types.MsgProductImage) (*types.MsgProductImageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var post = types.Post{
		Creator: msg.Creator,
		ProductImage: msg.ProductImage,
		ProductPrice: msg.ProductPrice,
	}

	id := k.AppendPost(ctx, post)

	return &types.MsgProductImageResponse{Id: id}, nil
}
