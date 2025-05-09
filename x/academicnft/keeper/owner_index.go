package keeper

import (
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNFTByOwnerIndex armazena um índice para rápida recuperação de NFTs por proprietário
func (k Keeper) SetNFTByOwnerIndex(ctx sdk.Context, courseNft types.CourseNft) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    ownerStore := prefix.NewStore(store, []byte(types.NFTsByOwnerPrefix))
    
    // A chave será: owner/nftId
    key := []byte(courseNft.Owner + "/" + courseNft.NftId)
    
    // O valor será o ID do NFT
    ownerStore.Set(key, []byte(courseNft.NftId))
}

// DeleteNFTByOwnerIndex remove um NFT do índice de proprietário
func (k Keeper) DeleteNFTByOwnerIndex(ctx sdk.Context, owner string, nftId string) {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    ownerStore := prefix.NewStore(store, []byte(types.NFTsByOwnerPrefix))
    
    // A chave será: owner/nftId
    key := []byte(owner + "/" + nftId)
    
    // Remove a entrada
    ownerStore.Delete(key)
}

// GetNFTsByOwner retorna todos os IDs de NFTs pertencentes a um proprietário
func (k Keeper) GetNFTsByOwner(ctx sdk.Context, owner string) []string {
    store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    ownerStore := prefix.NewStore(store, []byte(types.NFTsByOwnerPrefix))
    
    // Prefixo para o proprietário específico
    ownerPrefix := []byte(owner + "/")
    nftByOwnerStore := prefix.NewStore(ownerStore, ownerPrefix)
    
    var nftIds []string
    iterator := nftByOwnerStore.Iterator(nil, nil)
    defer iterator.Close()
    
    for ; iterator.Valid(); iterator.Next() {
        nftIds = append(nftIds, string(iterator.Value()))
    }
    
    return nftIds
}