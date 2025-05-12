package keeper

import (
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNFTByOwnerIndex stores an index for quick retrieval of NFTs by owner
func (k Keeper) SetNFTByOwnerIndex(ctx sdk.Context, courseNft types.CourseNft) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    ownerStore := prefix.NewStore(store, KeyPrefix(types.NFTsByOwnerPrefix))
    // Key will be: owner/nftId
    key := []byte(courseNft.Owner + "/" + courseNft.NftId)
    // Value will be the NFT ID
    ownerStore.Set(key, []byte(courseNft.NftId))
}

// DeleteNFTByOwnerIndex removes an NFT from the owner index
func (k Keeper) DeleteNFTByOwnerIndex(ctx sdk.Context, owner string, nftId string) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    ownerStore := prefix.NewStore(store, KeyPrefix(types.NFTsByOwnerPrefix))
    // Key will be: owner/nftId
    key := []byte(owner + "/" + nftId)
    // Remove the entry
    ownerStore.Delete(key)
}

// KeyPrefix helper function to get a key prefix
func KeyPrefix(p string) []byte {
    return []byte(p)
}