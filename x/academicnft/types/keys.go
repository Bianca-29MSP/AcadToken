package types

// Storage key prefixes
const (
    // ModuleName defines the module name
    ModuleName = "academicnft"
    // StoreKey defines the primary store key
    StoreKey = ModuleName
    // RouterKey defines the message router key
    RouterKey = ModuleName
    // QuerierRoute defines the querier route
    QuerierRoute = ModuleName
    // ParamsKey is the key for module parameters
    ParamsKey = "Params"
    // NFTsByOwnerPrefix is the prefix for NFTs by owner index
    NFTsByOwnerPrefix = "nfts-by-owner-"
)

// Event prefixes
const (
    // EventTypeMintCourseNFT is emitted when a Course NFT is created
    EventTypeMintCourseNFT = "mint_course_nft"
    // EventTypeTransferCourseNFT is emitted when a Course NFT is transferred
    EventTypeTransferCourseNFT = "transfer_course_nft"
    // EventTypeApproveEquivalence is emitted when an equivalence is approved
    EventTypeApproveEquivalence = "approve_equivalence"
)

// Event attributes
const (
    AttributeKeyNFTID = "nft_id"
    AttributeKeyNFTID1 = "nft_id_1"
    AttributeKeyNFTID2 = "nft_id_2"
    AttributeKeyCourseID = "course_id"
    AttributeKeyInstitution = "institution"
    AttributeKeyContentHash = "content_hash"
    AttributeKeyFrom = "from"
    AttributeKeyTo = "to"
    AttributeKeyRecipient = "recipient"
    AttributeKeyScore = "score"
    AttributeKeyJustification = "justification"
    AttributeKeyEquivalentCourseID = "equivalent_course_id"
)

// KeyPrefix returns the key prefix for a specific type
func KeyPrefix(p string) []byte {
    return []byte(p)
}

// NFTByOwnerKey returns the key for a specific NFT owned by an address
func NFTByOwnerKey(owner string, nftId string) []byte {
    return append(NFTByOwnerPrefix(owner), []byte("/"+nftId)...)
}

// NFTByOwnerPrefix returns the prefix for all NFTs of an owner
func NFTByOwnerPrefix(owner string) []byte {
    return append(KeyPrefix(NFTsByOwnerPrefix), []byte(owner)...)
}