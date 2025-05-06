package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetCourseToken set a specific courseToken in the store from its index
func (k Keeper) SetCourseToken(ctx context.Context, courseToken types.CourseToken) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseTokenKeyPrefix))
	b := k.cdc.MustMarshal(&courseToken)
	store.Set(types.CourseTokenKey(
		courseToken.Index,
	), b)
}

// GetCourseToken returns a courseToken from its index
func (k Keeper) GetCourseToken(
	ctx context.Context,
	index string,

) (val types.CourseToken, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseTokenKeyPrefix))

	b := store.Get(types.CourseTokenKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCourseToken removes a courseToken from the store
func (k Keeper) RemoveCourseToken(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseTokenKeyPrefix))
	store.Delete(types.CourseTokenKey(
		index,
	))
}

// GetAllCourseToken returns all courseToken
func (k Keeper) GetAllCourseToken(ctx context.Context) (list []types.CourseToken) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CourseTokenKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CourseToken
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
