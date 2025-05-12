package types

import (
    "context"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sharedtypes "github.com/Bianca-29MSP/AcademicToken/x/shared/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
    GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
    // Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
    SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
    // Methods imported from bank should be defined here
}

// CurriculumKeeper defines the expected interface for the Curriculum module.
type CurriculumKeeper interface {
    GetCourseContent(ctx sdk.Context, courseId string) (sharedtypes.CourseContent, bool)
    GetInstitution(ctx sdk.Context, institutionId string) (sharedtypes.Institution, bool)
    IsGraduationEligible(ctx sdk.Context, student, institution, program string) bool
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
    Get(context.Context, []byte, interface{})
    Set(context.Context, []byte, interface{})
}