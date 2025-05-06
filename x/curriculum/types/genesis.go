package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InstitutionList:       []Institution{},
		CourseTokenList:       []CourseToken{},
		CourseEquivalenceList: []CourseEquivalence{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in institution
	institutionIndexMap := make(map[string]struct{})

	for _, elem := range gs.InstitutionList {
		index := string(InstitutionKey(elem.Index))
		if _, ok := institutionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for institution")
		}
		institutionIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in courseToken
	courseTokenIndexMap := make(map[string]struct{})

	for _, elem := range gs.CourseTokenList {
		index := string(CourseTokenKey(elem.Index))
		if _, ok := courseTokenIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for courseToken")
		}
		courseTokenIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in courseEquivalence
	courseEquivalenceIndexMap := make(map[string]struct{})

	for _, elem := range gs.CourseEquivalenceList {
		index := string(CourseEquivalenceKey(elem.Index))
		if _, ok := courseEquivalenceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for courseEquivalence")
		}
		courseEquivalenceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
