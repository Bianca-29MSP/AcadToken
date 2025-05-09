package keeper

import (
    "context"
    "time"
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintCourseNft(goCtx context.Context, msg *types.MsgMintCourseNft) (*types.MsgMintCourseNftResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Criar um ID único para o NFT
    nftId := msg.CourseId + "_" + msg.Institution
    
    // Verificar se o NFT já existe
    _, found := k.GetCourseNft(ctx, nftId)
    if found {
        return nil, types.ErrNFTAlreadyExists
    }
    
    // Criar o NFT
    var courseNft = types.CourseNft{
        NftId:                    nftId,
        Creator:                  msg.Creator,
        Owner:                    msg.Creator, // Inicialmente o criador é o proprietário
        CourseId:                 msg.CourseId,
        Institution:              msg.Institution,
        Title:                    msg.Title,
        Code:                     msg.Code,
        WorkloadHours:            msg.WorkloadHours,
        Credits:                  msg.Credits,
        Description:              msg.Description,
        Objectives:               msg.Objectives,
        TopicUnits:               msg.TopicUnits,
        Methodologies:            msg.Methodologies,
        EvaluationMethods:        msg.EvaluationMethods,
        BibliographyBasic:        msg.BibliographyBasic,
        BibliographyComplementary: msg.BibliographyComplementary,
        Keywords:                 msg.Keywords,
        ContentHash:              msg.ContentHash,
        CreatedAt:                int32(time.Now().Unix()), // Convertendo int64 para int32
        ApprovedEquivalences:     []string{},
    }
    
    // Armazenar o NFT
    k.SetCourseNft(ctx, courseNft)
    
    // Adicionar ao índice de NFTs por proprietário
    k.SetNFTByOwnerIndex(ctx, courseNft)
    
    // Emitir evento de NFT criado
    ctx.EventManager().EmitEvent(
        sdk.NewEvent(
            types.EventTypeMintCourseNFT,
            sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
            sdk.NewAttribute(types.AttributeKeyNFTID, nftId),
            sdk.NewAttribute(types.AttributeKeyCourseID, msg.CourseId),
            sdk.NewAttribute(types.AttributeKeyInstitution, msg.Institution),
            sdk.NewAttribute(types.AttributeKeyContentHash, msg.ContentHash),
        ),
    )
    
    // Como a resposta MsgMintCourseNftResponse não tem campos, retornamos um objeto vazio
    return &types.MsgMintCourseNftResponse{}, nil
}