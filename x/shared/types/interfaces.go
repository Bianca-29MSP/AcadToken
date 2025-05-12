package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Define shared structs for common data

// CourseContent represents course information shared between modules
type CourseContent struct {
	CourseId                  string
	Institution               string
	Title                     string
	Code                      string
	WorkloadHours             uint64
	Credits                   uint64
	Description               string
	Objectives                []string
	TopicUnits                []string
	Methodologies             []string
	EvaluationMethods         []string
	BibliographyBasic         []string
	BibliographyComplementary []string
	Keywords                  []string
	ContentHash               string
}

// Institution represents an academic institution
type Institution struct {
	Id           string 
	Name         string
	Address      string
	IsAuthorized bool 
}

// CourseToken represents a token for a course
type CourseToken struct {
	Index          string   
	TokenId        string   
	Name           string   
	Code           string   
	ContentHash    string   
	Institution    string   
	CompletionDate string   
	Grade          uint64   
	Prerequisites  []string 
	Owner          string   
}

// AcademicTree represents a student's academic progress
type AcademicTree struct {
	Student          string   
	Institution      string   
	CompletedTokens  []string 
	AvailableTokens  []string 
}

// CourseEquivalence represents an equivalence between courses
type CourseEquivalence struct {
	Index              string 
	SourceTokenId      string 
	TargetInstitution  string 
	TargetCourseCode   string 
	EquivalenceStatus  uint64 
	ApprovalCount      uint64 
}

// CourseNft represents an NFT for a completed course
type CourseNft struct {
	NftId                     string
	Creator                   string
	Owner                     string
	CourseId                  string
	Institution               string
	Title                     string
	Code                      string
	WorkloadHours             uint64
	Credits                   uint64
	Description               string
	Objectives                []string
	TopicUnits                []string
	Methodologies             []string
	EvaluationMethods         []string
	BibliographyBasic         []string
	BibliographyComplementary []string
	Keywords                  []string
	ContentHash               string
	CreatedAt                 int32
	ApprovedEquivalences      []string
}

// CurriculumKeeper defines the expected interfaces for the Curriculum module
type CurriculumKeeper interface {
    GetCourseContent(ctx sdk.Context, courseId string) (CourseContent, bool)
    GetInstitution(ctx sdk.Context, institutionId string) (Institution, bool)
    IsGraduationEligible(ctx sdk.Context, student string, institution string, program string) bool
    CheckPrerequisites(ctx sdk.Context, studentAddr string, courseId string) bool
}

// AcademicNFTKeeper defines the expected interfaces for the AcademicNFT module
type AcademicNFTKeeper interface {
	GetCourseNft(ctx sdk.Context, nftId string) (CourseNft, bool)
	HasCourseNFT(ctx sdk.Context, owner string, courseId string) bool
	GetNFTsByOwner(ctx sdk.Context, owner string) []CourseNft // Adicionado método necessário
}