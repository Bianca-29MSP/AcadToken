package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
)

// ParamsKey é a chave para armazenar os parâmetros no store
var ParamsKey = []byte("Params")

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(ParamsKey)
	if bz == nil {
		return params
	}
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(ParamsKey, bz)
	return nil
}