package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetCourseContent set a specific courseContent in the store from its index
func (k Keeper) SetCourseContent(ctx context.Context, courseContent types.CourseContent) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseContentKeyPrefix))
	b := k.cdc.MustMarshal(&courseContent)
	store.Set(types.CourseContentKey(
		courseContent.Index,
	), b)
}

// GetCourseContent returns a courseContent from its index
func (k Keeper) GetCourseContent(
	ctx context.Context,
	index string,

) (val types.CourseContent, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseContentKeyPrefix))

	b := store.Get(types.CourseContentKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCourseContent removes a courseContent from the store
func (k Keeper) RemoveCourseContent(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseContentKeyPrefix))
	store.Delete(types.CourseContentKey(
		index,
	))
}

// GetAllCourseContent returns all courseContent
func (k Keeper) GetAllCourseContent(ctx context.Context) (list []types.CourseContent) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseContentKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CourseContent
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
