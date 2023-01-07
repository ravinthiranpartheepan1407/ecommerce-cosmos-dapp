package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProductImage = "product_image"

var _ sdk.Msg = &MsgProductImage{}

func NewMsgProductImage(creator string, productPrice string) *MsgProductImage {
	return &MsgProductImage{
		Creator:      creator,
		ProductPrice: productPrice,
	}
}

func (msg *MsgProductImage) Route() string {
	return RouterKey
}

func (msg *MsgProductImage) Type() string {
	return TypeMsgProductImage
}

func (msg *MsgProductImage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProductImage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProductImage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
