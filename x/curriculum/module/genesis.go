package curriculum

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/keeper"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the institution
	for _, elem := range genState.InstitutionList {
		k.SetInstitution(ctx, elem)
	}
	// Set all the courseToken
	for _, elem := range genState.CourseTokenList {
		k.SetCourseToken(ctx, elem)
	}
	// Set all the courseEquivalence
	for _, elem := range genState.CourseEquivalenceList {
		k.SetCourseEquivalence(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.InstitutionList = k.GetAllInstitution(ctx)
	genesis.CourseTokenList = k.GetAllCourseToken(ctx)
	genesis.CourseEquivalenceList = k.GetAllCourseEquivalence(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
