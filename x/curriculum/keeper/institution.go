package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetInstitution set a specific institution in the store from its index
func (k Keeper) SetInstitution(ctx context.Context, institution types.Institution) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InstitutionKeyPrefix))
	b := k.cdc.MustMarshal(&institution)
	store.Set(types.InstitutionKey(
		institution.Index,
	), b)
}

// GetInstitution returns a institution from its index
func (k Keeper) GetInstitution(
	ctx context.Context,
	index string,

) (val types.Institution, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InstitutionKeyPrefix))

	b := store.Get(types.InstitutionKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInstitution removes a institution from the store
func (k Keeper) RemoveInstitution(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InstitutionKeyPrefix))
	store.Delete(types.InstitutionKey(
		index,
	))
}

// GetAllInstitution returns all institution
func (k Keeper) GetAllInstitution(ctx context.Context) (list []types.Institution) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.InstitutionKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Institution
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
