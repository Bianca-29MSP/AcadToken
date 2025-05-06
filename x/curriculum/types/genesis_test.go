package types_test

import (
	"testing"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				InstitutionList: []types.Institution{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				CourseTokenList: []types.CourseToken{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				CourseEquivalenceList: []types.CourseEquivalence{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated institution",
			genState: &types.GenesisState{
				InstitutionList: []types.Institution{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated courseToken",
			genState: &types.GenesisState{
				CourseTokenList: []types.CourseToken{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated courseEquivalence",
			genState: &types.GenesisState{
				CourseEquivalenceList: []types.CourseEquivalence{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
