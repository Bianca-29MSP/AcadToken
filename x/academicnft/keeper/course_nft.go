package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetCourseNft set a specific courseNft in the store from its index
func (k Keeper) SetCourseNft(ctx context.Context, courseNft types.CourseNft) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseNftKeyPrefix))
	b := k.cdc.MustMarshal(&courseNft)
	store.Set(types.CourseNftKey(
		courseNft.Index,
	), b)
}

// GetCourseNft returns a courseNft from its index
func (k Keeper) GetCourseNft(
	ctx context.Context,
	index string,

) (val types.CourseNft, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseNftKeyPrefix))

	b := store.Get(types.CourseNftKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCourseNft removes a courseNft from the store
func (k Keeper) RemoveCourseNft(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseNftKeyPrefix))
	store.Delete(types.CourseNftKey(
		index,
	))
}

// GetAllCourseNft returns all courseNft
func (k Keeper) GetAllCourseNft(ctx context.Context) (list []types.CourseNft) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseNftKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CourseNft
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
