package keeper

import (
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCourseNft sets a specific courseNft in the store
func (k Keeper) SetCourseNft(ctx sdk.Context, courseNft types.CourseNft) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    b := k.cdc.MustMarshal(&courseNft)
    store.Set(types.CourseNftKey(courseNft.NftId), b)
}

// GetCourseNft returns a courseNft by ID
func (k Keeper) GetCourseNft(ctx sdk.Context, nftId string) (val types.CourseNft, found bool) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    b := store.Get(types.CourseNftKey(nftId))
    if b == nil {
        return val, false
    }
    k.cdc.MustUnmarshal(b, &val)
    return val, true
}

// RemoveCourseNft removes a courseNft from the store
func (k Keeper) RemoveCourseNft(ctx sdk.Context, nftId string) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store.Delete(types.CourseNftKey(nftId))
}

// GetAllCourseNft returns all courseNft
func (k Keeper) GetAllCourseNft(ctx sdk.Context) (list []types.CourseNft) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    prefixStore := prefix.NewStore(store, types.CourseNftKeyPrefix)
    iterator := prefixStore.Iterator(nil, nil)
    defer iterator.Close()
    for ; iterator.Valid(); iterator.Next() {
        var val types.CourseNft
        k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
    }
    return
}

// HasCourseNFT checks if a student has an NFT for a specific course
func (k Keeper) HasCourseNFT(ctx sdk.Context, owner string, courseId string) bool {
    nfts := k.GetNFTsByOwner(ctx, owner)
    
    for _, nft := range nfts {
        if nft.CourseId == courseId {
            return true
        }
    }
    
    return false
}

// GetNFTsByOwner returns all NFTs belonging to an owner
func (k Keeper) GetNFTsByOwner(ctx sdk.Context, owner string) []types.CourseNft {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    ownerStore := prefix.NewStore(store, []byte(types.NFTsByOwnerPrefix))
    // Prefix for the specific owner
    ownerPrefix := []byte(owner + "/")
    nftByOwnerStore := prefix.NewStore(ownerStore, ownerPrefix)
    
    var nfts []types.CourseNft
    iterator := nftByOwnerStore.Iterator(nil, nil)
    defer iterator.Close()
    
    for ; iterator.Valid(); iterator.Next() {
        // Extract NFT ID from the value
        nftId := string(iterator.Value())
        nft, found := k.GetCourseNft(ctx, nftId)
        if found {
            nfts = append(nfts, nft)
        }
    }
    
    return nfts
}