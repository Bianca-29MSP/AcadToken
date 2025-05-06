package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetCourseEquivalence set a specific courseEquivalence in the store from its index
func (k Keeper) SetCourseEquivalence(ctx context.Context, courseEquivalence types.CourseEquivalence) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseEquivalenceKeyPrefix))
	b := k.cdc.MustMarshal(&courseEquivalence)
	store.Set(types.CourseEquivalenceKey(
		courseEquivalence.Index,
	), b)
}

// GetCourseEquivalence returns a courseEquivalence from its index
func (k Keeper) GetCourseEquivalence(
	ctx context.Context,
	index string,

) (val types.CourseEquivalence, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseEquivalenceKeyPrefix))

	b := store.Get(types.CourseEquivalenceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCourseEquivalence removes a courseEquivalence from the store
func (k Keeper) RemoveCourseEquivalence(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseEquivalenceKeyPrefix))
	store.Delete(types.CourseEquivalenceKey(
		index,
	))
}

// GetAllCourseEquivalence returns all courseEquivalence
func (k Keeper) GetAllCourseEquivalence(ctx context.Context) (list []types.CourseEquivalence) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseEquivalenceKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CourseEquivalence
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
