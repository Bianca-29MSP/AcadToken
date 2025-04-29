package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterInstitution{}

func NewMsgRegisterInstitution(creator string, address string, name string) *MsgRegisterInstitution {
	return &MsgRegisterInstitution{
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgRegisterInstitution) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
