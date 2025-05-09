package types

import (
	"testing"

	"github.com/Bianca-29MSP/AcademicToken/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgApproveEquivalence_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgApproveEquivalence
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveEquivalence{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgApproveEquivalence{
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
