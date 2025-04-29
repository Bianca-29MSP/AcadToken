package types

import (
	"testing"

	"github.com/Bianca-29MSP/AcademicToken/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterInstitution_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterInstitution
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRegisterInstitution{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRegisterInstitution{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
