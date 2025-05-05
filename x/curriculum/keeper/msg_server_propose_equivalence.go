package keeper

import (
	"context"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProposeEquivalence(goCtx context.Context, msg *types.MsgProposeEquivalence) (*types.MsgProposeEquivalenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProposeEquivalenceResponse{}, nil
}
