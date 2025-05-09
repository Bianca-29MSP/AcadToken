package keeper

import (
	"context"
	
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveEquivalence(goCtx context.Context, msg *types.MsgApproveEquivalence) (*types.MsgApproveEquivalenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Verificar se os NFTs existem
	nft1, found1 := k.GetCourseNft(ctx, msg.NftId1)
	if !found1 {
		return nil, types.ErrNFTNotFound
	}
	
	nft2, found2 := k.GetCourseNft(ctx, msg.NftId2)
	if !found2 {
		return nil, types.ErrNFTNotFound
	}
	
	// Adicionar à lista de equivalências de cada NFT
	
	// Verificar se a equivalência já existe no NFT1
	equivalenceExists1 := false
	for _, eq := range nft1.ApprovedEquivalences {
		if eq == msg.NftId2 {
			equivalenceExists1 = true
			break
		}
	}
	
	// Verificar se a equivalência já existe no NFT2
	equivalenceExists2 := false
	for _, eq := range nft2.ApprovedEquivalences {
		if eq == msg.NftId1 {
			equivalenceExists2 = true
			break
		}
	}
	
	// Adicionar equivalências se não existirem
	if !equivalenceExists1 {
		nft1.ApprovedEquivalences = append(nft1.ApprovedEquivalences, msg.NftId2)
		k.SetCourseNft(ctx, nft1)
	}
	
	if !equivalenceExists2 {
		nft2.ApprovedEquivalences = append(nft2.ApprovedEquivalences, msg.NftId1)
		k.SetCourseNft(ctx, nft2)
	}
	
	// Emitir evento de equivalência aprovada
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeApproveEquivalence,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyNFTID1, msg.NftId1),
			sdk.NewAttribute(types.AttributeKeyNFTID2, msg.NftId2),
			sdk.NewAttribute(types.AttributeKeyScore, msg.Score),
			sdk.NewAttribute(types.AttributeKeyJustification, msg.Justification),
		),
	)
	
	return &types.MsgApproveEquivalenceResponse{}, nil
}