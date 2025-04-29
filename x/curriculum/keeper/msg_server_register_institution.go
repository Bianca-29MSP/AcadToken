package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterInstitution(goCtx context.Context, msg *types.MsgRegisterInstitution) (*types.MsgRegisterInstitutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterInstitutionResponse{}, nil
}
