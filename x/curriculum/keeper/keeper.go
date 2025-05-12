package keeper

import (
    "fmt"
    "cosmossdk.io/core/store"
    "cosmossdk.io/log"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
)

type (
    Keeper struct {
        cdc codec.BinaryCodec
        storeService store.KVStoreService
        logger log.Logger
        // the address capable of executing a MsgUpdateParams message
        authority string
        bankKeeper types.BankKeeper
        // Reference to AcademicNFTKeeper for cross-module functionality
        academicnftKeeper types.AcademicNFTKeeper
    }
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeService store.KVStoreService,
    logger log.Logger,
    authority string,
    bankKeeper types.BankKeeper,
    academicnftKeeper types.AcademicNFTKeeper,
) Keeper {
    if _, err := sdk.AccAddressFromBech32(authority); err != nil {
        panic(fmt.Sprintf("invalid authority address: %s", authority))
    }
    
    return Keeper{
        cdc: cdc,
        storeService: storeService,
        authority: authority,
        logger: logger,
        bankKeeper: bankKeeper,
        academicnftKeeper: academicnftKeeper,
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

// SetAcademicNFTKeeper updates the AcademicNFTKeeper reference after initialization
// This resolves the circular dependency between modules
func (k *Keeper) SetAcademicNFTKeeper(academicnftKeeper types.AcademicNFTKeeper) {
    k.academicnftKeeper = academicnftKeeper
}

// Adicione estes métodos à estrutura Keeper no arquivo keeper.go do módulo curriculum:

// IsGraduationEligible verifica se um estudante é elegível para graduação
func (k Keeper) IsGraduationEligible(ctx sdk.Context, student string, institution string, program string) bool {
    // Implemente a lógica real aqui
    // Por exemplo, verifique se o estudante completou todos os requisitos para o programa
    
    // Código provisório para permitir compilação
    return true
}

// CheckPrerequisites verifica se um estudante atende aos pré-requisitos para um curso
func (k Keeper) CheckPrerequisites(ctx sdk.Context, studentAddr string, courseId string) bool {
    // Implemente a lógica real aqui
    // Por exemplo, verifique se o estudante já completou os cursos que são pré-requisitos
    
    // Código provisório para permitir compilação
    return true
}