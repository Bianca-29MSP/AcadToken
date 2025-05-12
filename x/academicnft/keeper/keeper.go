package keeper

import (
    "fmt"
    "cosmossdk.io/core/store"
    "cosmossdk.io/log"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
)

type (
    Keeper struct {
        cdc codec.BinaryCodec
        storeService store.KVStoreService
        logger log.Logger
        // the address capable of executing a MsgUpdateParams message
        authority string
        // Reference to CurriculumKeeper for cross-module functionality
        curriculumKeeper types.CurriculumKeeper
    }
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeService store.KVStoreService,
    logger log.Logger,
    authority string,
    curriculumKeeper types.CurriculumKeeper,
) Keeper {
    if _, err := sdk.AccAddressFromBech32(authority); err != nil {
        panic(fmt.Sprintf("invalid authority address: %s", authority))
    }
    
    return Keeper{
        cdc: cdc,
        storeService: storeService,
        authority: authority,
        logger: logger,
        curriculumKeeper: curriculumKeeper,
    }
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
    return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
    return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}