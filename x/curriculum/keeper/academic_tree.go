package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetAcademicTree set a specific academic tree in the store from its index
func (k Keeper) SetAcademicTree(ctx context.Context, academicTree types.AcademicTree) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AcademicTreeKeyPrefix))
	b := k.cdc.MustMarshal(&academicTree)
	store.Set(types.AcademicTreeKey(
		academicTree.Student,
	), b)
}

// GetAcademicTreeByStudent returns an academic tree from its index
func (k Keeper) GetAcademicTreeByStudent(
	ctx context.Context,
	student string,
) (val types.AcademicTree, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AcademicTreeKeyPrefix))
	b := store.Get(types.AcademicTreeKey(
		student,
	))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAcademicTree removes an academic tree from the store
func (k Keeper) RemoveAcademicTree(
	ctx context.Context,
	student string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AcademicTreeKeyPrefix))
	store.Delete(types.AcademicTreeKey(
		student,
	))
}

// GetAllAcademicTree returns all academic trees
func (k Keeper) GetAllAcademicTree(ctx context.Context) (list []types.AcademicTree) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AcademicTreeKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.AcademicTree
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}
