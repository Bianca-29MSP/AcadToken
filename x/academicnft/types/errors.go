package types

import (
	sdkerrors "cosmossdk.io/errors"
)

// Códigos de erro para o módulo AcademicNFT
var (
	ErrNFTNotFound     = sdkerrors.Register(ModuleName, 1, "NFT não encontrado")
	ErrNotOwner        = sdkerrors.Register(ModuleName, 2, "não é o proprietário do NFT")
	ErrNFTAlreadyExists = sdkerrors.Register(ModuleName, 3, "NFT já existente")
	ErrInvalidInput    = sdkerrors.Register(ModuleName, 4, "entrada inválida")
	ErrUnauthorized    = sdkerrors.Register(ModuleName, 5, "não autorizado")
	ErrInvalidSigner   = sdkerrors.Register(ModuleName, 6, "assinante inválido")
)