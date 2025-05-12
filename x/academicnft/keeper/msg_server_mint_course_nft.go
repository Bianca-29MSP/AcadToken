package keeper

import (
    "context"
    "time"
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "cosmossdk.io/errors"
)

func (k msgServer) MintCourseNft(goCtx context.Context, msg *types.MsgMintCourseNft) (*types.MsgMintCourseNftResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    
    // Validate if course exists in Curriculum module
    courseContent, found := k.curriculumKeeper.GetCourseContent(ctx, msg.CourseId)
    if !found {
        return nil, sdkerrors.Wrapf(types.ErrCourseNotFound, "course %s not found", msg.CourseId)
    }
    
    // Validate if institution exists
    _, found = k.curriculumKeeper.GetInstitution(ctx, msg.Institution)
    if !found {
        return nil, sdkerrors.Wrapf(types.ErrInstitutionNotFound, "institution %s not found", msg.Institution)
    }
    
    // Verify if course belongs to the specified institution
    if courseContent.Institution != msg.Institution {
        return nil, sdkerrors.Wrapf(types.ErrInvalidInstitution, 
            "course %s does not belong to institution %s", msg.CourseId, msg.Institution)
    }
    
    // Verify content hash to ensure authenticity
    if courseContent.ContentHash != msg.ContentHash {
        return nil, sdkerrors.Wrapf(types.ErrInvalidContentHash, 
            "content hash does not match curriculum records")
    }

    // Create unique NFT ID
    nftId := msg.CourseId + "_" + msg.Institution
    
    // Check if NFT already exists
    _, found = k.GetCourseNft(ctx, nftId)
    if found {
        return nil, types.ErrNFTAlreadyExists
    }
    
    // Create the course NFT
    var courseNft = types.CourseNft{
        NftId: nftId,
        Creator: msg.Creator,
        Owner: msg.Creator, // Initially the creator is the owner
        CourseId: msg.CourseId,
        Institution: msg.Institution,
        Title: msg.Title,
        Code: msg.Code,
        WorkloadHours: msg.WorkloadHours,
        Credits: msg.Credits,
        Description: msg.Description,
        Objectives: msg.Objectives,
        TopicUnits: msg.TopicUnits,
        Methodologies: msg.Methodologies,
        EvaluationMethods: msg.EvaluationMethods,
        BibliographyBasic: msg.BibliographyBasic,
        BibliographyComplementary: msg.BibliographyComplementary,
        Keywords: msg.Keywords,
        ContentHash: msg.ContentHash,
        CreatedAt: int32(time.Now().Unix()),
        ApprovedEquivalences: []string{},
    }
    
    // Store the NFT
    k.SetCourseNft(ctx, courseNft)
    
    // Add to owner index for efficient queries
    k.SetNFTByOwnerIndex(ctx, courseNft)
    
    // Emit NFT creation event
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
    
    return &types.MsgMintCourseNftResponse{}, nil
}