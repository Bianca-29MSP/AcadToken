package app

import (
	"context"
	"io"

	_ "cosmossdk.io/api/cosmos/tx/config/v1" // import for side-effects
	clienthelpers "cosmossdk.io/client/v2/helpers"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	_ "cosmossdk.io/x/circuit" // import for side-effects
	circuitkeeper "cosmossdk.io/x/circuit/keeper"
	_ "cosmossdk.io/x/evidence" // import for side-effects
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	_ "cosmossdk.io/x/feegrant/module" // import for side-effects
	nftkeeper "cosmossdk.io/x/nft/keeper"
	_ "cosmossdk.io/x/nft/module" // import for side-effects
	_ "cosmossdk.io/x/upgrade"    // import for side-effects
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	abci "github.com/cometbft/cometbft/abci/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/vesting" // import for side-effects
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/authz/module" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/bank"         // import for side-effects
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/consensus" // import for side-effects
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/crisis" // import for side-effects
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/distribution" // import for side-effects
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/group/module" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/mint"         // import for side-effects
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/params" // import for side-effects
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	_ "github.com/cosmos/cosmos-sdk/x/slashing" // import for side-effects
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/staking" // import for side-effects
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	_ "github.com/cosmos/ibc-go/modules/capability" // import for side-effects
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	_ "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts" // import for side-effects
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	_ "github.com/cosmos/ibc-go/v8/modules/apps/29-fee" // import for side-effects
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	academictokenmodulekeeper "github.com/Bianca-29MSP/AcademicToken/x/academictoken/keeper"
	curriculummodulekeeper "github.com/Bianca-29MSP/AcademicToken/x/curriculum/keeper"
	academicnftmodulekeeper "github.com/Bianca-29MSP/AcademicToken/x/academicnft/keeper"
	
	// Import module types for intermodule communication
	academicnfttypes "github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	curriculumtypes "github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sharedtypes "github.com/Bianca-29MSP/AcademicToken/x/shared/types"
	academictokentypes "github.com/Bianca-29MSP/AcademicToken/x/academictoken/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport

	"github.com/Bianca-29MSP/AcademicToken/docs"
)

const (
	// Name is the name of the application.
	Name = "AcademicToken"
	// AccountAddressPrefix is the prefix for accounts addresses.
	AccountAddressPrefix = "acad"
	// ChainCoinType is the coin type of the chain.
	ChainCoinType = 118
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper

	SlashingKeeper       slashingkeeper.Keeper
	MintKeeper           mintkeeper.Keeper
	GovKeeper            *govkeeper.Keeper
	CrisisKeeper         *crisiskeeper.Keeper
	UpgradeKeeper        *upgradekeeper.Keeper
	ParamsKeeper         paramskeeper.Keeper
	AuthzKeeper          authzkeeper.Keeper
	EvidenceKeeper       evidencekeeper.Keeper
	FeeGrantKeeper       feegrantkeeper.Keeper
	GroupKeeper          groupkeeper.Keeper
	NFTKeeper            nftkeeper.Keeper
	CircuitBreakerKeeper circuitkeeper.Keeper

	// IBC
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	CapabilityKeeper    *capabilitykeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper

	// Scoped IBC
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedKeepers             map[string]capabilitykeeper.ScopedKeeper

	AcademictokenKeeper academictokenmodulekeeper.Keeper
	CurriculumKeeper    curriculummodulekeeper.Keeper
	AcademicnftKeeper   academicnftmodulekeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// simulation manager
	sm *module.SimulationManager
}

// KVStoreAdapter adapta KVStoreKey para KVStoreService
type KVStoreAdapter struct {
	storeKey *storetypes.KVStoreKey
}

func NewKVStoreAdapter(storeKey *storetypes.KVStoreKey) *KVStoreAdapter {
	return &KVStoreAdapter{
		storeKey: storeKey,
	}
}

// OpenKVStore implementa a interface KVStoreService para o Cosmos SDK v0.50.13
func (a *KVStoreAdapter) OpenKVStore(ctx context.Context) store.KVStore {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return NewKVStoreWrapper(sdkCtx.KVStore(a.storeKey))
}

// KVStoreWrapper adapta o KVStore para ter a assinatura correta dos métodos
type KVStoreWrapper struct {
	kvStore storetypes.KVStore
}

func NewKVStoreWrapper(kvStore storetypes.KVStore) *KVStoreWrapper {
	return &KVStoreWrapper{kvStore: kvStore}
}

func (w *KVStoreWrapper) Get(key []byte) ([]byte, error) {
	return w.kvStore.Get(key), nil
}

func (w *KVStoreWrapper) Has(key []byte) (bool, error) {
	return w.kvStore.Has(key), nil
}

func (w *KVStoreWrapper) Set(key, value []byte) error {
	w.kvStore.Set(key, value)
	return nil
}

func (w *KVStoreWrapper) Delete(key []byte) error {
	w.kvStore.Delete(key)
	return nil
}

func (w *KVStoreWrapper) Iterator(start, end []byte) (store.Iterator, error) {
	return w.kvStore.Iterator(start, end), nil
}

func (w *KVStoreWrapper) ReverseIterator(start, end []byte) (store.Iterator, error) {
	return w.kvStore.ReverseIterator(start, end), nil
}

// AcademicNFTAdapter adapts the AcademicNFTKeeper to the interface expected by CurriculumKeeper
type AcademicNFTAdapter struct {
	keeper academicnftmodulekeeper.Keeper
}

// NewAcademicNFTAdapter creates a new adapter for the AcademicNFTKeeper
func NewAcademicNFTAdapter(k academicnftmodulekeeper.Keeper) *AcademicNFTAdapter {
	return &AcademicNFTAdapter{keeper: k}
}

// GetCourseNft adapts the GetCourseNft method to the interface expected by CurriculumKeeper
func (a *AcademicNFTAdapter) GetCourseNft(ctx sdk.Context, nftId string) (sharedtypes.CourseNft, bool) {
	nft, found := a.keeper.GetCourseNft(ctx, nftId)
	if !found {
		return sharedtypes.CourseNft{}, false
	}
	
	// Convert the type
	sharedNft := sharedtypes.CourseNft{
		NftId:                   nft.NftId,
		Creator:                 nft.Creator,
		Owner:                   nft.Owner,
		CourseId:                nft.CourseId,
		Institution:             nft.Institution,
		Title:                   nft.Title,
		Code:                    nft.Code,
		WorkloadHours:           nft.WorkloadHours,
		Credits:                 nft.Credits,
		Description:             nft.Description,
		Objectives:              nft.Objectives,
		TopicUnits:              nft.TopicUnits,
		Methodologies:           nft.Methodologies,
		EvaluationMethods:       nft.EvaluationMethods,
		BibliographyBasic:       nft.BibliographyBasic,
		BibliographyComplementary: nft.BibliographyComplementary,
		Keywords:                nft.Keywords,
		ContentHash:             nft.ContentHash,
		CreatedAt:               nft.CreatedAt,
		ApprovedEquivalences:    nft.ApprovedEquivalences,
	}
	
	return sharedNft, true
}

// HasCourseNFT adapts the HasCourseNFT method to the interface expected by CurriculumKeeper
func (a *AcademicNFTAdapter) HasCourseNFT(ctx sdk.Context, owner string, courseId string) bool {
	return a.keeper.HasCourseNFT(ctx, owner, courseId)
}

// GetNFTsByOwner adapts the GetNFTsByOwner method to the interface expected by CurriculumKeeper
func (a *AcademicNFTAdapter) GetNFTsByOwner(ctx sdk.Context, owner string) []sharedtypes.CourseNft {
	nfts := a.keeper.GetNFTsByOwner(ctx, owner)
	sharedNfts := make([]sharedtypes.CourseNft, len(nfts))
	
	for i, nft := range nfts {
		sharedNfts[i] = sharedtypes.CourseNft{
			NftId:                   nft.NftId,
			Creator:                 nft.Creator,
			Owner:                   nft.Owner,
			CourseId:                nft.CourseId,
			Institution:             nft.Institution,
			Title:                   nft.Title,
			Code:                    nft.Code,
			WorkloadHours:           nft.WorkloadHours,
			Credits:                 nft.Credits,
			Description:             nft.Description,
			Objectives:              nft.Objectives,
			TopicUnits:              nft.TopicUnits,
			Methodologies:           nft.Methodologies,
			EvaluationMethods:       nft.EvaluationMethods,
			BibliographyBasic:       nft.BibliographyBasic,
			BibliographyComplementary: nft.BibliographyComplementary,
			Keywords:                nft.Keywords,
			ContentHash:             nft.ContentHash,
			CreatedAt:               nft.CreatedAt,
			ApprovedEquivalences:    nft.ApprovedEquivalences,
		}
	}
	
	return sharedNfts
}

// CurriculumAdapter adapts the CurriculumKeeper to the interface expected by AcademicNFTKeeper
type CurriculumAdapter struct {
	keeper curriculummodulekeeper.Keeper
}

// NewCurriculumAdapter creates a new adapter for the CurriculumKeeper
func NewCurriculumAdapter(k curriculummodulekeeper.Keeper) *CurriculumAdapter {
	return &CurriculumAdapter{keeper: k}
}

// GetCourseContent adapts the GetCourseContent method to the interface expected by AcademicNFTKeeper
func (a *CurriculumAdapter) GetCourseContent(ctx sdk.Context, courseId string) (sharedtypes.CourseContent, bool) {
	content, found := a.keeper.GetCourseContent(ctx, courseId)
	if !found {
		return sharedtypes.CourseContent{}, false
	}
	
	// Convert the type
	sharedContent := sharedtypes.CourseContent{
		CourseId:                  content.CourseId,
		Institution:               content.Institution,
		Title:                     content.Title,
		Code:                      content.Code,
		WorkloadHours:             content.WorkloadHours,
		Credits:                   content.Credits,
		Description:               content.Description,
		Objectives:                content.Objectives,
		TopicUnits:                content.TopicUnits,
		Methodologies:             content.Methodologies,
		EvaluationMethods:         content.EvaluationMethods,
		BibliographyBasic:         content.BibliographyBasic,
		BibliographyComplementary: content.BibliographyComplementary,
		Keywords:                  content.Keywords,
		ContentHash:               content.ContentHash,
	}
	
	return sharedContent, true
}

// GetInstitution adapts the GetInstitution method to the interface expected by AcademicNFTKeeper
func (a *CurriculumAdapter) GetInstitution(ctx sdk.Context, institutionId string) (sharedtypes.Institution, bool) {
	institution, found := a.keeper.GetInstitution(ctx, institutionId)
	if !found {
		return sharedtypes.Institution{}, false
	}
	
	// Converter para o tipo compartilhado usando a estrutura correta
	sharedInstitution := sharedtypes.Institution{
		Id:           institution.Index,
		Name:         institution.Name,
		Address:      institution.Address,
		IsAuthorized: institution.IsAuthorized, // Corrigido para usar o campo boolean correto
	}
	
	return sharedInstitution, true
}

// IsGraduationEligible adapts the IsGraduationEligible method
// Implementação temporária até que a função seja implementada no Keeper real
func (a *CurriculumAdapter) IsGraduationEligible(ctx sdk.Context, student string, institution string, program string) bool {
	// Como o método não existe no keeper real, implementamos uma versão temporária
	// Implemente aqui a lógica real quando disponível no keeper
	return true
}

// CheckPrerequisites adapts the CheckPrerequisites method
// Implementação temporária até que a função seja implementada no Keeper real
func (a *CurriculumAdapter) CheckPrerequisites(ctx sdk.Context, studentAddr string, courseId string) bool {
	// Como o método não existe no keeper real, implementamos uma versão temporária
	// Implemente aqui a lógica real quando disponível no keeper
	return true
}

func init() {
	var err error
	clienthelpers.EnvPrefix = Name
	DefaultNodeHome, err = clienthelpers.GetNodeHomeDirectory("." + Name)
	if err != nil {
		panic(err)
	}
}

// getGovProposalHandlers return the chain proposal handlers.
func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

// AppConfig returns the default app config.
func AppConfig() depinject.Config {
	return depinject.Configs(
		appConfig,
		// Alternatively, load the app config from a YAML file.
		// appconfig.LoadYAML(AppConfigYAML),
		depinject.Supply(
			// supply custom module basics
			map[string]module.AppModuleBasic{
				genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
				govtypes.ModuleName:     gov.NewAppModuleBasic(getGovProposalHandlers()),
				// this line is used by starport scaffolding # stargate/appConfig/moduleBasic
			},
		),
	)
}

// New returns a reference to an initialized App.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) (*App, error) {
	var (
		app        = &App{ScopedKeepers: make(map[string]capabilitykeeper.ScopedKeeper)}
		appBuilder *runtime.AppBuilder

		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			AppConfig(),
			depinject.Supply(
				appOpts, // supply app options
				logger,  // supply logger
				// Supply with IBC keeper getter for the IBC modules with App Wiring.
				// The IBC Keeper cannot be passed because it has not been initiated yet.
				// Passing the getter, the app IBC Keeper will always be accessible.
				// This needs to be removed after IBC supports App Wiring.
				app.GetIBCKeeper,
				app.GetCapabilityScopedKeeper,
			),
		)
	)

	// First, inject all standard keepers except our custom keepers
	if err := depinject.Inject(appConfig,
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.StakingKeeper,
		&app.DistrKeeper,
		&app.ConsensusParamsKeeper,
		&app.SlashingKeeper,
		&app.MintKeeper,
		&app.GovKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.ParamsKeeper,
		&app.AuthzKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
		&app.NFTKeeper,
		&app.GroupKeeper,
		&app.CircuitBreakerKeeper,
		// Custom keepers are manually initialized below
	); err != nil {
		panic(err)
	}

	// add to default baseapp options
	// enable optimistic execution
	baseAppOptions = append(baseAppOptions, baseapp.SetOptimisticExecution())

	app.App = appBuilder.Build(db, traceStore, baseAppOptions...)

	// First, the main AcademicToken keeper
	app.AcademictokenKeeper = academictokenmodulekeeper.NewKeeper(
		app.appCodec,
		NewKVStoreAdapter(app.GetKey(academictokentypes.StoreKey)),
		logger,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	
	// Initialize Curriculum keeper with temporary nil AcademicNFTKeeper
	app.CurriculumKeeper = curriculummodulekeeper.NewKeeper(
		app.appCodec,
		NewKVStoreAdapter(app.GetKey(curriculumtypes.StoreKey)),
		logger,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		app.BankKeeper,
		nil, // AcademicNFTKeeper will be set after initialization
	)
	
	// Create an adapter for the CurriculumKeeper
	curriculumAdapter := NewCurriculumAdapter(app.CurriculumKeeper)
	
	// Initialize AcademicNFT keeper with Curriculum adapter
	app.AcademicnftKeeper = academicnftmodulekeeper.NewKeeper(
		app.appCodec,
		NewKVStoreAdapter(app.GetKey(academicnfttypes.StoreKey)),
		logger,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		curriculumAdapter,
	)
	
	// Create an adapter to resolve the type incompatibility
	academicNFTAdapter := NewAcademicNFTAdapter(app.AcademicnftKeeper)
	
	// Update CurriculumKeeper with AcademicNFTKeeper reference via the adapter
	app.CurriculumKeeper.SetAcademicNFTKeeper(academicNFTAdapter)

	// register streaming services
	if err := app.RegisterStreamingServices(appOpts, app.kvStoreKeys()); err != nil {
		return nil, err
	}

	/****  Module Options ****/

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// create the simulation manager and define the order of the modules for deterministic simulations
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// A custom InitChainer sets if extra pre-init-genesis logic is required.
	// This is necessary for manually registered modules that do not support app wiring.
	// Manually set the module version map as shown below.
	// The upgrade module will automatically handle de-duplication of the module version map.
	app.SetInitChainer(func(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
		if err := app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap()); err != nil {
			return nil, err
		}
		return app.App.InitChainer(ctx, req)
	})

	if err := app.Load(loadLatest); err != nil {
		return nil, err
	}

	return app, nil
}

// LegacyAmino returns App's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns App's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns App's interfaceRegistry.
func (app *App) InterfaceRegistry() codectypes.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns App's tx config.
func (app *App) TxConfig() client.TxConfig {
	return app.txConfig
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	kvStoreKey, ok := app.UnsafeFindStoreKey(storeKey).(*storetypes.KVStoreKey)
	if !ok {
		return nil
	}
	return kvStoreKey
}

// GetMemKey returns the MemoryStoreKey for the provided store key.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	key, ok := app.UnsafeFindStoreKey(storeKey).(*storetypes.MemoryStoreKey)
	if !ok {
		return nil
	}

	return key
}

// kvStoreKeys returns all the kv store keys registered inside App.
func (app *App) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// GetSubspace returns a param subspace for a given module name.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// GetIBCKeeper returns the IBC keeper.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetCapabilityScopedKeeper returns the capability scoped keeper.
func (app *App) GetCapabilityScopedKeeper(moduleName string) capabilitykeeper.ScopedKeeper {
	sk, ok := app.ScopedKeepers[moduleName]
	if !ok {
		sk = app.CapabilityKeeper.ScopeToModule(moduleName)
		app.ScopedKeepers[moduleName] = sk
	}
	return sk
}

// SimulationManager implements the SimulationApp interface.
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}

	// register app's OpenAPI routes.
	docs.RegisterOpenAPIService(Name, apiSvr.Router)
}

// These are minimal types to match the existing app_config.go content

// ModuleAccPerm represents a module account permission
type ModuleAccPerm struct {
	Account     string
	Permissions []string
}