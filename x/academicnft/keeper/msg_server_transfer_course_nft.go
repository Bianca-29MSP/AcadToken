package keeper

import (
	"context"
	
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferCourseNft(goCtx context.Context, msg *types.MsgTransferCourseNft) (*types.MsgTransferCourseNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Verificar se o NFT existe
	nft, found := k.GetCourseNft(ctx, msg.NftId)
	if !found {
		return nil, types.ErrNFTNotFound
	}
	
	// Verificar se o remetente é o proprietário
	if nft.Owner != msg.Creator {
		return nil, types.ErrNotOwner
	}
	
	// Atualizar o proprietário
	nft.Owner = msg.Recipient
	
	// Salvar o NFT atualizado
	k.SetCourseNft(ctx, nft)
	
	// Emitir evento de transferência
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferCourseNFT,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.NftId),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
	)
	
	return &types.MsgTransferCourseNftResponse{}, nil
}