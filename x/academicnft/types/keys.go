package types

// Prefixos para as chaves de armazenamento
const (
    // ModuleName define o nome do módulo
    ModuleName = "academicnft"
    // StoreKey define a chave de armazenamento principal
    StoreKey = ModuleName
    // RouterKey define a mensagem de router
    RouterKey = ModuleName
    // QuerierRoute define a rota do querier
    QuerierRoute = ModuleName
    // ParamsKey é a chave para os parâmetros do módulo
    ParamsKey = "Params"

    // NFTsByOwnerPrefix é o prefixo para o índice de NFTs por proprietário
    NFTsByOwnerPrefix = "nfts-by-owner-"
)

// Prefixos para eventos
const (
    // EventTypeMintCourseNFT é o evento emitido quando um Course NFT é criado
    EventTypeMintCourseNFT = "mint_course_nft"
    // EventTypeTransferCourseNFT é o evento emitido quando um Course NFT é transferido
    EventTypeTransferCourseNFT = "transfer_course_nft"
    // EventTypeApproveEquivalence é o evento emitido quando uma equivalência é aprovada
    EventTypeApproveEquivalence = "approve_equivalence"
)

// Atributos de evento
const (
    AttributeKeyNFTID        = "nft_id"
    AttributeKeyNFTID1       = "nft_id_1"
    AttributeKeyNFTID2       = "nft_id_2"
    AttributeKeyCourseID     = "course_id"
    AttributeKeyInstitution  = "institution"
    AttributeKeyContentHash  = "content_hash"
    AttributeKeyRecipient    = "recipient"
    AttributeKeyScore        = "score"
    AttributeKeyJustification = "justification"
)

func KeyPrefix(p string) []byte {
    return []byte(p)
}