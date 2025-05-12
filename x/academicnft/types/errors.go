package types

import (
    sdkerrors "cosmossdk.io/errors"
)

// Error codes for the AcademicNFT module
var (
    ErrNFTNotFound         = sdkerrors.Register(ModuleName, 1, "NFT not found")
    ErrNotOwner            = sdkerrors.Register(ModuleName, 2, "not the owner of NFT")
    ErrNFTAlreadyExists    = sdkerrors.Register(ModuleName, 3, "NFT already exists")
    ErrInvalidInput        = sdkerrors.Register(ModuleName, 4, "invalid input")
    ErrUnauthorized        = sdkerrors.Register(ModuleName, 5, "unauthorized")
    ErrInvalidSigner       = sdkerrors.Register(ModuleName, 6, "invalid signer")
    ErrCourseNotFound      = sdkerrors.Register(ModuleName, 7, "course not found")
    ErrInstitutionNotFound = sdkerrors.Register(ModuleName, 8, "institution not found")
    ErrInvalidInstitution  = sdkerrors.Register(ModuleName, 9, "invalid institution")
    ErrInvalidContentHash  = sdkerrors.Register(ModuleName, 10, "invalid content hash")
)