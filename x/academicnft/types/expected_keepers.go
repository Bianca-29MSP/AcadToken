package types

import (
    "context"
    sdk "github.com/cosmos/cosmos-sdk/types"
    curriculumtypes "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
    GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
    // Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
    SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
    // Methods imported from bank should be defined here
}

// ADICIONE ESTA NOVA INTERFACE:
// CurriculumKeeper defines the expected interface for the Curriculum module.
type CurriculumKeeper interface {
    // Defina aqui os métodos do Curriculum que o AcademicNFT precisa usar
    GetCourseContent(ctx context.Context, courseId string) (curriculumtypes.CourseContent, bool)
    SetCourseContent(ctx context.Context, content curriculumtypes.CourseContent)
    GetEquivalence(ctx context.Context, courseA, courseB string) (curriculumtypes.CourseEquivalence, bool)
    IsGraduationEligible(ctx context.Context, student, institution, program string) bool
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
    Get(context.Context, []byte, interface{})
    Set(context.Context, []byte, interface{})
}