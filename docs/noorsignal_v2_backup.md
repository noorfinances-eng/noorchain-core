# Archive NOORSIGNAL V2 (PoSS avancé)

Ceci est une archive de l’ancien module `x/noorsignal` (version PoSS avancée avec argent réel, BankKeeper, Reserve PoSS, etc.).
x/noorsignal/README.md
# NOORCHAIN — PoSS Module (`x/noorsignal`)

This module implements the core of **Proof of Signal Social (PoSS)** for NOORCHAIN.

> Status: early technical skeleton — types, keeper and config are present,
> but no reward or halving logic is implemented yet.

---

## 1. Purpose

The `noorsignal` module is responsible for:

- recording **social signals** emitted by participants
- tracking and managing **curators** (social validators)
- storing and applying the **global PoSS configuration**:
  - base reward
  - 70% / 30% split (participant / curator)
  - daily limits
  - enable/disable flag

The actual **minting logic, reward distribution and halving** will be implemented
in later phases and wired to the global NUR supply / PoSS rules.

---

## 2. Current Code Structure

### 2.1 Types

Located in: `x/noorsignal/types/types.go`

- `Signal`  
  Represents a social signal emitted by a participant:
  - `Id uint64`
  - `Participant sdk.AccAddress`
  - `Curator sdk.AccAddress`
  - `Weight uint32` (1x, 2x, 5x… encoded as integers)
  - `Time time.Time`
  - `Metadata string` (content hash, external ID, etc.)

- `Curator`  
  Represents a social validator:
  - `Address sdk.AccAddress`
  - `Level string` (e.g. “bronze”, “silver”, “gold”)
  - `TotalSignalsValidated uint64`
  - `Active bool`

- `PossConfig`  
  Global PoSS configuration:
  - `BaseReward uint64`
  - `ParticipantShare uint32` (e.g. 70)
  - `CuratorShare uint32` (e.g. 30)
  - `MaxSignalsPerDay uint32`
  - `Enabled bool`

- `DefaultPossConfig()`  
  Returns a default configuration coherent with NOORCHAIN’s model:
  - 70% participant / 30% curator
  - module enabled
  - symbolic `BaseReward` and daily limits.

> Note: these Go structs are the conceptual model.  
> Final protobuf definitions (`.proto`) will be added later.

---

### 2.2 Keys and Store Layout

Located in: `x/noorsignal/types/keys.go`

Defines:

- `ModuleName = "noorsignal"`
- `StoreKey`, `RouterKey`, `QuerierRoute` (standard Cosmos conventions)

Store prefixes:

- `KeyPrefixSignals = []byte{0x01}`
- `KeyPrefixCurators = []byte{0x02}`
- `KeyPrefixConfig = []byte{0x03}`

Helper functions:

- `GetSignalStore(parent prefix.Store) prefix.Store`
- `GetCuratorStore(parent prefix.Store) prefix.Store`
- `GetConfigStore(parent prefix.Store) prefix.Store`

These functions are used by the keeper to obtain scoped KVStores.

---

### 2.3 Keeper

Located in: `x/noorsignal/keeper/keeper.go`

The `Keeper` struct currently holds:

- `storeKey storetypes.StoreKey`
- `cdc codec.Codec`

Constructor:

- `NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey)`

Internal helpers:

- `getStore(ctx sdk.Context) sdk.KVStore`
- `signalStore(ctx sdk.Context) prefix.Store`
- `curatorStore(ctx sdk.Context) prefix.Store`
- `configStore(ctx sdk.Context) prefix.Store`

Config management:

- `SetConfig(ctx sdk.Context, cfg PossConfig)`  
  Encodes and stores the global PoSS configuration.

- `GetConfig(ctx sdk.Context) (PossConfig, bool)`  
  Reads and decodes the configuration; returns `(config, found)`.

- `InitDefaultConfig(ctx sdk.Context)`  
  If no configuration is present, stores `DefaultPossConfig()`.  
  If a configuration already exists, it is preserved.

No signal or curator logic is implemented yet (no writes/reads for them).

---

## 3. Planned Extensions

The following elements are planned but not yet implemented:

1. **Protobuf Definitions (`.proto`)**
   - `noorsignal/signal.proto`
   - `noorsignal/curator.proto`
   - `noorsignal/config.proto`
   - `noorsignal/tx.proto` for Msg definitions

2. **Message Types (Tx)**
   Examples:
   - `MsgSubmitSignal`
   - `MsgValidateSignal`
   - `MsgUpdatePossConfig` (governance or curator-limited)

3. **Genesis Handling**
   - `InitGenesis` and `ExportGenesis` functions
   - default PoSS config included in genesis
   - optional pre-registered curators

4. **Reward Logic**
   - calculation of PoSS rewards per signal
   - enforcement of halving rules (linked to global NUR supply)
   - daily limits per participant
   - enforcement of 70% / 30% split

5. **Integration with Other Modules**
   - `BankKeeper` for NUR transfers
   - potentially `StakingKeeper` or `GovKeeper` for curator governance
   - interaction with global supply management for PoSS

---

## 4. Development Notes

- The current implementation is intentionally minimal and safe:
  - no minting
  - no reward distribution
  - no halving logic
- This skeleton allows:
  - early review of the design
  - step-by-step implementation of PoSS features
  - future audit of the module in isolation

This README should be updated as soon as:

- protobuf files are added,
- messages are implemented,
- genesis and reward logic are wired.

- x/noorsignal/genesis.go
- package noorsignal

import (
	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// InitGenesis initialise le module PoSS (noorsignal) à partir du GenesisState.
func InitGenesis(
	ctx sdk.Context,
	k noorsignalkeeper.Keeper,
	genState noorsignaltypes.GenesisState,
) []abci.ValidatorUpdate {
	// 1) Config PoSS : on applique simplement la config fournie
	// (si elle est "zero value", ce sera la responsabilité de DefaultGenesis).
	k.SetConfig(ctx, genState.Config)

	// 2) Curators
	for _, c := range genState.Curators {
		k.SetCurator(ctx, c)
	}

	// Pas de validators spécifiques à renvoyer pour PoSS.
	return []abci.ValidatorUpdate{}
}

// ExportGenesis exporte l'état courant du module PoSS vers un GenesisState.
func ExportGenesis(
	ctx sdk.Context,
	k noorsignalkeeper.Keeper,
) noorsignaltypes.GenesisState {
	cfg, found := k.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// Pour l’instant, on n'exporte pas les signaux ni les compteurs journaliers
	// et on laisse la liste de curators vide (ou à améliorer plus tard).
	return noorsignaltypes.GenesisState{
		Config:   cfg,
		Curators: []noorsignaltypes.Curator{},
	}
}
x/noorsignal/module.go
package noorsignal

import (
	"context"
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	noorsignalkeeper "github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ------------------------------------------------------------
// AppModuleBasic : partie "statique" du module (sans Keeper)
// ------------------------------------------------------------

type AppModuleBasic struct{}

var _ module.AppModuleBasic = AppModuleBasic{}

func (AppModuleBasic) Name() string { return noorsignaltypes.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// Pour l’instant, on ne déclare pas encore d’interfaces spécifiques.
func (AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {}

// DefaultGenesis renvoie l'état de genèse par défaut du module.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	// Ici on ne dépend PAS de types.DefaultGenesis (qui n'existe pas encore).
	gs := noorsignaltypes.GenesisState{
		Config:   noorsignaltypes.DefaultPossConfig(),
		Curators: []noorsignaltypes.Curator{},
	}

	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ValidateGenesis valide l'état de genèse (version minimale).
func (AppModuleBasic) ValidateGenesis(
	_ codec.JSONCodec,
	_ client.TxEncodingConfig,
	bz json.RawMessage,
) error {
	if len(bz) == 0 {
		return nil
	}

	var gs noorsignaltypes.GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		return err
	}

	// Version minimale : on ne fait pas encore de validation complexe.
	return nil
}

// RegisterGRPCGatewayRoutes : pas encore de routes spécifiques.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(
	_ client.Context,
	_ *runtime.ServeMux,
) {
}

// GetTxCmd : pas encore de commandes CLI Tx spécifiques.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

// GetQueryCmd : pas encore de commandes CLI Query spécifiques.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

// ------------------------------------------------------------
// AppModule : partie "avec Keeper" (logique métier)
// ------------------------------------------------------------

type AppModule struct {
	AppModuleBasic
	keeper noorsignalkeeper.Keeper
}

// On s'assure qu'on implémente bien module.AppModule
var _ module.AppModule = AppModule{}

// IsAppModule est une méthode "marqueur" demandée par certains SDK récents.
func (AppModule) IsAppModule() {}

// IsOnePerModuleType est un autre marqueur requis par l'interface module.AppModule
// dans Cosmos SDK v0.50+. Il ne fait rien, il sert juste à indiquer
// qu'il n'y a qu'un module de ce type par app.
func (AppModule) IsOnePerModuleType() {}

// NewAppModule construit le module PoSS avec son Keeper.
func NewAppModule(k noorsignalkeeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// RegisterServices : pour l’instant, on ne branche PAS encore Msg/Query gRPC
// car les fichiers proto / pb.go ne sont pas finalisés.
func (am AppModule) RegisterServices(_ module.Configurator) {
	// On ajoutera plus tard :
	// - RegisterMsgServer
	// - RegisterQueryServer
}

// InitGenesis initialise l'état de genèse du module.
func (am AppModule) InitGenesis(
	ctx context.Context,
	_ codec.JSONCodec,
	data json.RawMessage,
) []abci.ValidatorUpdate {
	var gs noorsignaltypes.GenesisState
	if len(data) == 0 {
		gs = noorsignaltypes.GenesisState{
			Config:   noorsignaltypes.DefaultPossConfig(),
			Curators: []noorsignaltypes.Curator{},
		}
	} else {
		if err := json.Unmarshal(data, &gs); err != nil {
			panic(err)
		}
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return InitGenesis(sdkCtx, am.keeper, gs)
}

// ExportGenesis exporte l'état de genèse courant.
func (am AppModule) ExportGenesis(
	ctx context.Context,
	_ codec.JSONCodec,
) json.RawMessage {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	gs := ExportGenesis(sdkCtx, am.keeper)

	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

// ConsensusVersion permet d’indiquer une version du module (pour migrations futures).
func (am AppModule) ConsensusVersion() uint64 { return 1 }

x/noorsignal/keeper/keeper.go
package keeper

import (
	"encoding/binary"
	"encoding/json"

	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// Keeper est le gestionnaire principal du module PoSS (noorsignal) pour NOORCHAIN.
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec // conservé si on veut l'utiliser plus tard
}

// NewKeeper construit un nouveau Keeper PoSS pour NOORCHAIN.
func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// getStore retourne le KVStore brut du module à partir du contexte.
func (k Keeper) getStore(ctx sdk.Context) storetypes.KVStore {
	return ctx.KVStore(k.storeKey)
}

// ---------- Stores préfixés ----------

func (k Keeper) signalStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetSignalStore(parent)
}

func (k Keeper) curatorStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetCuratorStore(parent)
}

func (k Keeper) configStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetConfigStore(parent)
}

func (k Keeper) dailyCounterStore(ctx sdk.Context) prefix.Store {
	parent := k.getStore(ctx)
	return noorsignaltypes.GetDailyCounterStore(parent)
}

// ---------------------------
// Gestion de la configuration
// ---------------------------

func (k Keeper) SetConfig(ctx sdk.Context, cfg noorsignaltypes.PossConfig) {
	store := k.configStore(ctx)
	bz, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	store.Set([]byte("config"), bz)
}

func (k Keeper) GetConfig(ctx sdk.Context) (noorsignaltypes.PossConfig, bool) {
	store := k.configStore(ctx)

	bz := store.Get([]byte("config"))
	if bz == nil {
		return noorsignaltypes.PossConfig{}, false
	}

	var cfg noorsignaltypes.PossConfig
	if err := json.Unmarshal(bz, &cfg); err != nil {
		panic(err)
	}
	return cfg, true
}

func (k Keeper) InitDefaultConfig(ctx sdk.Context) {
	_, found := k.GetConfig(ctx)
	if found {
		return
	}

	defaultCfg := noorsignaltypes.DefaultPossConfig()
	k.SetConfig(ctx, defaultCfg)
}

// ---------------------------------
// Calcul des récompenses PoSS (aide)
// ---------------------------------

func (k Keeper) ComputeSignalRewardsFromConfig(
	ctx sdk.Context,
	weight uint32,
	era uint32,
) (total uint64, participant uint64, curator uint64, found bool) {
	cfg, ok := k.GetConfig(ctx)
	if !ok {
		return 0, 0, 0, false
	}

	total, participant, curator = noorsignaltypes.ComputeSignalRewards(cfg, weight, era)
	return total, participant, curator, true
}

func (k Keeper) ComputeSignalRewardsCurrentEra(
	ctx sdk.Context,
	weight uint32,
) (total uint64, participant uint64, curator uint64, found bool) {
	cfg, ok := k.GetConfig(ctx)
	if !ok {
		return 0, 0, 0, false
	}

	era := cfg.EraIndex
	total, participant, curator = noorsignaltypes.ComputeSignalRewards(cfg, weight, era)
	return total, participant, curator, true
}

// -------------------------------------
// Gestion des identifiants et des signaux
// -------------------------------------

func (k Keeper) getNextSignalID(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)

	bz := store.Get(noorsignaltypes.KeyNextSignalID)
	if bz == nil {
		return 1
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setNextSignalID(ctx sdk.Context, nextID uint64) {
	store := k.getStore(ctx)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, nextID)
	store.Set(noorsignaltypes.KeyNextSignalID, bz)
}

func (k Keeper) CreateSignal(ctx sdk.Context, sig noorsignaltypes.Signal) noorsignaltypes.Signal {
	nextID := k.getNextSignalID(ctx)
	sig.Id = nextID

	sstore := k.signalStore(ctx)
	key := noorsignaltypes.SignalKey(sig.Id)

	bz, err := json.Marshal(sig)
	if err != nil {
		panic(err)
	}
	sstore.Set(key, bz)

	k.setNextSignalID(ctx, nextID+1)
	return sig
}

func (k Keeper) SetSignal(ctx sdk.Context, sig noorsignaltypes.Signal) {
	sstore := k.signalStore(ctx)
	key := noorsignaltypes.SignalKey(sig.Id)

	bz, err := json.Marshal(sig)
	if err != nil {
		panic(err)
	}
	sstore.Set(key, bz)
}

func (k Keeper) GetSignal(ctx sdk.Context, id uint64) (noorsignaltypes.Signal, bool) {
	sstore := k.signalStore(ctx)
	key := noorsignaltypes.SignalKey(id)

	bz := sstore.Get(key)
	if bz == nil {
		return noorsignaltypes.Signal{}, false
	}

	var sig noorsignaltypes.Signal
	if err := json.Unmarshal(bz, &sig); err != nil {
		panic(err)
	}
	return sig, true
}

// -------------------------------------
// Compteurs quotidiens PoSS (limites)
// -------------------------------------

func (k Keeper) getDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	store := k.dailyCounterStore(ctx)
	key := noorsignaltypes.DailyCounterKey(addr, dayBucket)

	bz := store.Get(key)
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint32(bz)
}

func (k Keeper) setDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
	count uint32,
) {
	store := k.dailyCounterStore(ctx)
	key := noorsignaltypes.DailyCounterKey(addr, dayBucket)

	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, count)
	store.Set(key, bz)
}

func (k Keeper) incrementDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	current := k.getDailySignalCount(ctx, addr, dayBucket)
	next := current + 1
	k.setDailySignalCount(ctx, addr, dayBucket, next)
	return next
}

// Helpers exportés pour le MsgServer (BankKeeper à venir).
func (k Keeper) GetDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	return k.getDailySignalCount(ctx, addr, dayBucket)
}

func (k Keeper) IncrementDailySignalCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
	dayBucket uint64,
) uint32 {
	return k.incrementDailySignalCount(ctx, addr, dayBucket)
}

// -------------------------------------
// Gestion des Curators PoSS
// -------------------------------------

func (k Keeper) curatorKey(addr sdk.AccAddress) []byte {
	return addr.Bytes()
}

func (k Keeper) SetCurator(ctx sdk.Context, curator noorsignaltypes.Curator) {
	store := k.curatorStore(ctx)
	key := k.curatorKey(curator.Address)

	bz, err := json.Marshal(curator)
	if err != nil {
		panic(err)
	}
	store.Set(key, bz)
}

func (k Keeper) GetCurator(ctx sdk.Context, addr sdk.AccAddress) (noorsignaltypes.Curator, bool) {
	store := k.curatorStore(ctx)
	key := k.curatorKey(addr)

	bz := store.Get(key)
	if bz == nil {
		return noorsignaltypes.Curator{}, false
	}

	var curator noorsignaltypes.Curator
	if err := json.Unmarshal(bz, &curator); err != nil {
		panic(err)
	}
	return curator, true
}

func (k Keeper) IsActiveCurator(ctx sdk.Context, addr sdk.AccAddress) bool {
	curator, found := k.GetCurator(ctx, addr)
	if !found {
		return false
	}
	return curator.Active
}

// IncrementCuratorValidatedCount augmente de 1 le nombre total de signaux
// validés par ce Curator, si le Curator existe.
func (k Keeper) IncrementCuratorValidatedCount(
	ctx sdk.Context,
	addr sdk.AccAddress,
) {
	curator, found := k.GetCurator(ctx, addr)
	if !found {
		return
	}

	curator.TotalSignalsValidated++
	k.SetCurator(ctx, curator)
}
x/noorsignal/keeper/msg_server.go
package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
type MsgServer struct {
	Keeper Keeper
}

// NewMsgServer construit un MsgServer basé uniquement sur le Keeper PoSS.
//
// NOTE : Dans cette V1, on ne branche PAS encore de BankKeeper.
//        Les transferts réels de NUR seront ajoutés plus tard.
func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{
		Keeper: k,
	}
}

// ---------------------------------------------------------
// SubmitSignal : un participant émet un signal PoSS
// ---------------------------------------------------------
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Validation basique (poids, adresse…)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	// 2) Récupérer la config PoSS actuelle (ou défaut)
	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// 3) Vérifier la limite quotidienne si activée
	var dayBucket uint64
	if cfg.MaxSignalsPerDay > 0 {
		ts := ctx.BlockTime().Unix()
		if ts < 0 {
			ts = 0
		}
		dayBucket = uint64(ts) / 86400

		current := s.Keeper.GetDailySignalCount(ctx, participantAddr, dayBucket)
		if current >= cfg.MaxSignalsPerDay {
			return nil, errors.New("daily signal limit reached")
		}
	}

	// 4) Construire le signal (sans rewards, sans curator)
	signal := noorsignaltypes.Signal{
		Participant:       participantAddr,
		Curator:           nil,
		Weight:            msg.Weight,
		Time:              ctx.BlockTime(),
		Metadata:          msg.Metadata,
		TotalReward:       0,
		RewardParticipant: 0,
		RewardCurator:     0,
	}

	// 5) Enregistrer le signal (ID auto-incrémenté)
	signal = s.Keeper.CreateSignal(ctx, signal)

	// 6) Incrémenter le compteur quotidien
	if cfg.MaxSignalsPerDay > 0 {
		s.Keeper.IncrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	// 7) Émettre un event poss.signal_submitted
	ctx.EventManager().EmitEvent(
		noorsignaltypes.NewSignalSubmittedEvent(signal, ctx.BlockHeight()),
	)

	return &sdk.Result{}, nil
}

// ---------------------------------------------------------
// ValidateSignal : un curator valide un signal
// (V1 : pas encore de BankKeeper.SendCoins)
// ---------------------------------------------------------
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Validation basique
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	curatorAddr, err := msg.GetCuratorAddress()
	if err != nil {
		return nil, err
	}

	// 2) Vérifier que le curator est actif
	if !s.Keeper.IsActiveCurator(ctx, curatorAddr) {
		return nil, errors.New("curator not active or not authorized")
	}

	// 3) Charger le signal
	signal, found := s.Keeper.GetSignal(ctx, msg.SignalId)
	if !found {
		return nil, errors.New("signal not found")
	}

	// 4) Vérifier qu'il n'est pas déjà validé
	if signal.Curator != nil && len(signal.Curator) > 0 {
		return nil, errors.New("signal already validated")
	}

	// 5) Calcul des rewards PoSS (local, sans transfert réel)
	total, part, cur, ok := s.Keeper.ComputeSignalRewardsCurrentEra(ctx, signal.Weight)
	if !ok {
		total = 0
		part = 0
		cur = 0
	}

	// 6) Mettre à jour le signal
	signal.Curator = curatorAddr
	signal.TotalReward = total
	signal.RewardParticipant = part
	signal.RewardCurator = cur

	s.Keeper.SetSignal(ctx, signal)

	// 7) Incrémenter le compteur de validations du curator
	s.Keeper.IncrementCuratorValidatedCount(ctx, curatorAddr)

	// 8) Émettre un event poss.signal_validated
	ctx.EventManager().EmitEvent(
		noorsignaltypes.NewSignalValidatedEvent(signal, ctx.BlockHeight()),
	)

	// 9) FUTUR : ici on utilisera BankKeeper.SendCoins pour
	// distribuer les rewards depuis la réserve PoSS.

	return &sdk.Result{}, nil
}

x/noorsignal/keeper/msg_server_admin.go
package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// ---------------------------------------------------------
// Messages admin : AddCurator / RemoveCurator / SetConfig
// ---------------------------------------------------------

// AddCurator : ajoute (ou réactive) un curator.
func (s MsgServer) AddCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgAddCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validation basique du message
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// ⚠️ V1 : on ne fait pas encore de vérification forte de "Authority"
	// (gov, adresse fondation, etc.). On vérifiera ça plus tard.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	// Créer ou réactiver le curator
	curator := noorsignaltypes.Curator{
		Address: curatorAddr,
		Active:  true,
		// TotalSignalsValidated reste à 0 par défaut
	}

	s.Keeper.SetCurator(ctx, curator)

	// V1 : pas encore d'events spécifiques, on se contente du résultat OK.
	return &sdk.Result{}, nil
}

// RemoveCurator : désactive un curator existant.
func (s MsgServer) RemoveCurator(
	goCtx context.Context,
	msg *noorsignaltypes.MsgRemoveCurator,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// V1 : pas de check avancé sur Authority, on gère plus tard.
	curatorAddr, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	curator, found := s.Keeper.GetCurator(ctx, curatorAddr)
	if !found {
		return nil, errors.New("curator not found")
	}

	curator.Active = false
	s.Keeper.SetCurator(ctx, curator)

	return &sdk.Result{}, nil
}

// SetConfig : met à jour la configuration globale PoSS.
// V1 : on se contente de s'assurer qu'une config existe, sans encore
// appliquer tous les champs du message (BaseReward, ratios, etc.).
func (s MsgServer) SetConfig(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSetConfig,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// V1 : on ne fait pas encore une logique fine d'authority.
	// On se contente de garantir qu'une config existe.
	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// TODO (plus tard) : appliquer réellement les champs de msg à cfg.

	s.Keeper.SetConfig(ctx, cfg)
	return &sdk.Result{}, nil
}

x/noorsignal/keeper/query_server.go
package keeper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// QueryServer implémente le service gRPC "Query" du module PoSS (noorsignal).
// Il s'appuie sur le Keeper interne pour accéder au store.
type QueryServer struct {
	Keeper
}

// NewQueryServer construit un QueryServer à partir d'un Keeper PoSS.
func NewQueryServer(k Keeper) QueryServer {
	return QueryServer{Keeper: k}
}

// Signal retourne un signal unique via son identifiant.
func (q QueryServer) Signal(
	goCtx context.Context,
	req *noorsignaltypes.QuerySignalRequest,
) (*noorsignaltypes.QuerySignalResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sig, found := q.Keeper.GetSignal(ctx, req.Id)
	if !found {
		return nil, errors.New("signal not found")
	}

	// On renvoie directement la struct Signal Go définie dans types.go.
	return &noorsignaltypes.QuerySignalResponse{
		Signal: &sig,
	}, nil
}

// Signals retourne une liste paginée de signaux PoSS.
func (q QueryServer) Signals(
	goCtx context.Context,
	req *noorsignaltypes.QuerySignalsRequest,
) (*noorsignaltypes.QuerySignalsResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := q.signalStore(ctx) // prefix.Store sur les signaux
	var signals []noorsignaltypes.Signal

	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(key, value []byte) error {
			var sig noorsignaltypes.Signal
			if err := json.Unmarshal(value, &sig); err != nil {
				return err
			}
			signals = append(signals, sig)
			return nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to paginate signals: %w", err)
	}

	return &noorsignaltypes.QuerySignalsResponse{
		Signals:    signals,
		Pagination: pageRes,
	}, nil
}

// Curator retourne les informations d'un Curator via son adresse.
func (q QueryServer) Curator(
	goCtx context.Context,
	req *noorsignaltypes.QueryCuratorRequest,
) (*noorsignaltypes.QueryCuratorResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}
	if req.Address == "" {
		return nil, errors.New("address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid curator address: %w", err)
	}

	curator, found := q.Keeper.GetCurator(ctx, addr)
	if !found {
		return nil, errors.New("curator not found")
	}

	return &noorsignaltypes.QueryCuratorResponse{
		Curator: &curator,
	}, nil
}

// Config retourne la configuration PoSS actuelle.
func (q QueryServer) Config(
	goCtx context.Context,
	req *noorsignaltypes.QueryConfigRequest,
) (*noorsignaltypes.QueryConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cfg, found := q.Keeper.GetConfig(ctx)
	if !found {
		// Si aucune config n'est encore enregistrée, on renvoie la config par défaut.
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	return &noorsignaltypes.QueryConfigResponse{
		Config: cfg,
	}, nil
}

// DailyCount retourne le nombre de signaux émis par une adresse pour un
// "day bucket" donné (par ex. block_time / 86400).
func (q QueryServer) DailyCount(
	goCtx context.Context,
	req *noorsignaltypes.QueryDailyCountRequest,
) (*noorsignaltypes.QueryDailyCountResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}
	if req.Address == "" {
		return nil, errors.New("address cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	count := q.Keeper.getDailySignalCount(ctx, addr, req.Day)

	return &noorsignaltypes.QueryDailyCountResponse{
		Count: count,
	}, nil
}

// signalStore retourne un prefix.Store explicitement typé ici pour le QueryServer.
// On redélègue au helper existant du Keeper.
func (q QueryServer) signalStore(ctx sdk.Context) prefix.Store {
	return q.Keeper.signalStore(ctx)
}

x/noorsignal/proto/noorsignal/events.proto
syntax = "proto3";

package noorchain.noorsignal;

option go_package = "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types";

// -----------------------------------------------------------
// Events Protobuf pour le module PoSS (noorsignal)
// -----------------------------------------------------------

// Événement émis lorsqu'un signal est soumis.
message EventSignalSubmitted {
  uint64 signal_id   = 1;
  string participant = 2;
  uint32 weight      = 3;
  string metadata    = 4;
}

// Événement émis lorsqu'un signal est validé par un curator.
message EventSignalValidated {
  uint64 signal_id         = 1;
  string participant       = 2;
  string curator           = 3;
  uint64 reward_total      = 4;
  uint64 reward_participant = 5;
  uint64 reward_curator     = 6;
}

x/noorsignal/proto/noorsignal/query.proto
syntax = "proto3";

package noorchain.noorsignal;

option go_package = "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types";

import "google/api/annotations.proto";

// -----------------------------------------------------------
// SERVICE QUERY PoSS (noorsignal)
// -----------------------------------------------------------
service Query {
  // Récupère un signal par son ID.
  rpc Signal(QuerySignalRequest) returns (QuerySignalResponse) {
    option (google.api.http).get = "/noorchain/noorsignal/signal/{id}";
  }

  // Liste tous les signaux existants.
  rpc Signals(QuerySignalsRequest) returns (QuerySignalsResponse) {
    option (google.api.http).get = "/noorchain/noorsignal/signals";
  }

  // Récupère les informations d'un curator.
  rpc Curator(QueryCuratorRequest) returns (QueryCuratorResponse) {
    option (google.api.http).get = "/noorchain/noorsignal/curator/{address}";
  }

  // Liste tous les curators.
  rpc Curators(QueryCuratorsRequest) returns (QueryCuratorsResponse) {
    option (google.api.http).get = "/noorchain/noorsignal/curators";
  }

  // Récupère la configuration PoSS.
  rpc Config(QueryConfigRequest) returns (QueryConfigResponse) {
    option (google.api.http).get = "/noorchain/noorsignal/config";
  }

  // Nombre de signaux d'un participant pour un jour donné.
  rpc DailyCount(QueryDailyCountRequest) returns (QueryDailyCountResponse) {
    option (google.api.http).get = "/noorchain/noorsignal/daily/{address}/{day}";
  }
}

// -----------------------------------------------------------
// Messages pour les requêtes / réponses
// -----------------------------------------------------------

// --- Signal ---
message QuerySignalRequest {
  uint64 id = 1;
}

message QuerySignalResponse {
  Signal signal = 1;
}

// --- Signals list ---
message QuerySignalsRequest {}

message QuerySignalsResponse {
  repeated Signal signals = 1;
}

// --- Curator ---
message QueryCuratorRequest {
  string address = 1;
}

message QueryCuratorResponse {
  Curator curator = 1;
}

// --- Curators list ---
message QueryCuratorsRequest {}

message QueryCuratorsResponse {
  repeated Curator curators = 1;
}

// --- Config ---
message QueryConfigRequest {}

message QueryConfigResponse {
  PossConfig config = 1;
}

// --- Daily counter ---
message QueryDailyCountRequest {
  string address = 1;
  uint64 day      = 2;
}

message QueryDailyCountResponse {
  uint32 count = 1;
}

x/noorsignal/proto/noorsignal.proto
syntax = "proto3";

package noorsignal;

option go_package = "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types";

import "google/protobuf/timestamp.proto";

// Signal représente un signal social PoSS émis sur NOORCHAIN.
message Signal {
  uint64 id = 1;
  bytes participant = 2;
  bytes curator = 3;
  uint32 weight = 4;
  google.protobuf.Timestamp time = 5;
  string metadata = 6;

  uint64 total_reward = 7;
  uint64 reward_participant = 8;
  uint64 reward_curator = 9;
}

// Curator représente un validateur de signaux PoSS (Curator NOOR).
message Curator {
  bytes address = 1;
  string level = 2;
  uint64 total_signals_validated = 3;
  bool active = 4;
}

// PossConfig contient la configuration globale du système PoSS.
message PossConfig {
  uint64 base_reward = 1;
  uint32 participant_share = 2;
  uint32 curator_share = 3;
  uint32 max_signals_per_day = 4;
  bool enabled = 5;
  uint32 era_index = 6;
}

// GenesisState représente l'état initial du module noorsignal.
message GenesisState {
  PossConfig config = 1;
  repeated Signal signals = 2;
  repeated Curator curators = 3;
}

x/noorsignal/proto/query.proto
syntax = "proto3";

package noorsignal;

option go_package = "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types";

// Requête pour récupérer la configuration PoSS actuelle.
message QueryConfigRequest {}

// Réponse contenant la configuration PoSS.
message QueryConfigResponse {
  PossConfig config = 1;
}

// Service Query pour le module noorsignal.
service Query {
  rpc Config(QueryConfigRequest) returns (QueryConfigResponse);
}

x/noorsignal/proto/tx.proto
syntax = "proto3";

package noorsignal;

option go_package = "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types";

// MsgSubmitSignal : un participant soumet un signal PoSS.
message MsgSubmitSignal {
  string participant = 1;
  uint32 weight = 2;
  string metadata = 3;
}

message MsgSubmitSignalResponse {}

// MsgValidateSignal : un curator valide un signal existant.
message MsgValidateSignal {
  string curator = 1;
  uint64 signal_id = 2;
}

message MsgValidateSignalResponse {}

// MsgAddCurator : ajoute ou réactive un curator avec un niveau donné.
message MsgAddCurator {
  string authority = 1;
  string curator = 2;
  string level = 3;
}

message MsgAddCuratorResponse {}

// MsgRemoveCurator : désactive un curator existant.
message MsgRemoveCurator {
  string authority = 1;
  string curator = 2;
}

message MsgRemoveCuratorResponse {}

// MsgSetConfig : met à jour la configuration PoSS (base_reward, ratios, etc.).
message MsgSetConfig {
  string authority = 1;
  string base_reward = 2;
  uint32 participant_ratio = 3;
  uint32 curator_ratio = 4;
  uint32 max_signals_per_day = 5;
  uint64 era_index = 6;
}

message MsgSetConfigResponse {}

// Service Msg pour la gestion des transactions PoSS.
service Msg {
  rpc SubmitSignal(MsgSubmitSignal) returns (MsgSubmitSignalResponse);
  rpc ValidateSignal(MsgValidateSignal) returns (MsgValidateSignalResponse);
  rpc AddCurator(MsgAddCurator) returns (MsgAddCuratorResponse);
  rpc RemoveCurator(MsgRemoveCurator) returns (MsgRemoveCuratorResponse);
  rpc SetConfig(MsgSetConfig) returns (MsgSetConfigResponse);
}

x/noorsignal/types/addresses.go
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Adresses Bech32 officielles NOORCHAIN pour le genesis 5/5/5/5/80.
//
// Chaque adresse correspond à un wallet distinct, avec une seed
// différente, comme défini dans le plan NOORCHAIN 1.0.
//
// IMPORTANT :
// - Ne pas modifier ces valeurs après le lancement du mainnet.
// - Ces adresses seront utilisées pour :
//   * le genesis mainnet
//   * la documentation publique
//   * la configuration PoSS / BankKeeper.
const (
	FoundationAddressBech32  = "noor1dwzpnw9g9p2cucj2w2dnxk2at4amaqsrvmyka0"
	DevAddressBech32         = "noor1s2gzjrec9elucycpj66d2eyyw09tjuqfcasy7k"
	StimulusAddressBech32    = "noor1qt2h4crdtngyw4yn3yy8sqraqpzh4ghx0c4l46"
	PresaleAddressBech32     = "noor1pxt0wlq8xswj0l6jm2spzcspjwjnnvaggatalz"
	PossReserveAddressBech32 = "noor1c4gg5n37ycnfvaalsaa4nl22pfzp9ujqwr0mne"
)

// mustAccAddressFromBech32 convertit une string Bech32 en sdk.AccAddress
// et panic si l'adresse est invalide. Cela doit uniquement échouer si
// le code contient une erreur, pas en production.
func mustAccAddressFromBech32(addr string) sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		panic(err)
	}
	return acc
}

// Adresses AccAddress prêtes à l'emploi pour les modules qui en ont besoin
// (par exemple pour la réserve PoSS dans BankKeeper).
var (
	FoundationAccAddress  = mustAccAddressFromBech32(FoundationAddressBech32)
	DevAccAddress         = mustAccAddressFromBech32(DevAddressBech32)
	StimulusAccAddress    = mustAccAddressFromBech32(StimulusAddressBech32)
	PresaleAccAddress     = mustAccAddressFromBech32(PresaleAddressBech32)
	PossReserveAccAddress = mustAccAddressFromBech32(PossReserveAddressBech32)
)

x/noorsignal/types/convert.go
package types

// Ce fichier est volontairement minimal.
//
// À ce stade du projet NOORCHAIN, les conversions Protobuf
// (SignalProto, CuratorProto, PossConfigProto, etc.)
// NE SONT PAS encore définies.
//
// Elles seront ajoutées uniquement lorsque :
// - les fichiers .proto seront écrits,
// - les fichiers .pb.go seront générés,
// - le keeper ou le module aura besoin d’une conversion.
//
// Laisser ce fichier vide permet de compiler sans erreur.

x/noorsignal/types/events.go
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------------------------------------------------------------------------
// Noms des événements PoSS
// -----------------------------------------------------------------------------

const (
	// Signal soumis par un participant
	EventTypeSignalSubmitted = "poss.signal_submitted"

	// Signal validé par un curator
	EventTypeSignalValidated = "poss.signal_validated"

	// Curator ajouté / mis à jour
	EventTypeCuratorAdded = "poss.curator_added"

	// Curator désactivé
	EventTypeCuratorRemoved = "poss.curator_removed"

	// Configuration PoSS mise à jour
	EventTypeConfigUpdated = "poss.config_updated"
)

// -----------------------------------------------------------------------------
// Clés d'attributs pour tous les events PoSS
// -----------------------------------------------------------------------------

const (
	AttrKeySignalID          = "signal_id"
	AttrKeyParticipant       = "participant"
	AttrKeyCurator           = "curator"
	AttrKeyWeight            = "weight"
	AttrKeyTotalReward       = "total_reward"
	AttrKeyRewardParticipant = "reward_participant"
	AttrKeyRewardCurator     = "reward_curator"
	AttrKeyBlockHeight       = "block_height"

	AttrKeyLevel             = "level"
	AttrKeyAuthority         = "authority"
	AttrKeyBaseReward        = "base_reward"
	AttrKeyMaxSignalsPerDay  = "max_signals_per_day"
	AttrKeyEraIndex          = "era_index"
	AttrKeyParticipantRatio  = "participant_ratio"
	AttrKeyCuratorRatio      = "curator_ratio"
)

// -----------------------------------------------------------------------------
// Helpers pour construire les événements de signaux
// -----------------------------------------------------------------------------

// NewSignalSubmittedEvent construit un événement poss.signal_submitted
// à partir d'un Signal fraîchement créé.
func NewSignalSubmittedEvent(sig Signal, blockHeight int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalSubmitted,
		sdk.NewAttribute(AttrKeySignalID, fmt.Sprintf("%d", sig.Id)),
		sdk.NewAttribute(AttrKeyParticipant, sig.Participant.String()),
		sdk.NewAttribute(AttrKeyWeight, fmt.Sprintf("%d", sig.Weight)),
		sdk.NewAttribute(AttrKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
	)
}

// NewSignalValidatedEvent construit un événement poss.signal_validated
// à partir d'un Signal qui vient d'être validé (avec rewards).
func NewSignalValidatedEvent(sig Signal, blockHeight int64) sdk.Event {
	return sdk.NewEvent(
		EventTypeSignalValidated,
		sdk.NewAttribute(AttrKeySignalID, fmt.Sprintf("%d", sig.Id)),
		sdk.NewAttribute(AttrKeyParticipant, sig.Participant.String()),
		sdk.NewAttribute(AttrKeyCurator, sig.Curator.String()),
		sdk.NewAttribute(AttrKeyWeight, fmt.Sprintf("%d", sig.Weight)),
		sdk.NewAttribute(AttrKeyTotalReward, fmt.Sprintf("%d", sig.TotalReward)),
		sdk.NewAttribute(AttrKeyRewardParticipant, fmt.Sprintf("%d", sig.RewardParticipant)),
		sdk.NewAttribute(AttrKeyRewardCurator, fmt.Sprintf("%d", sig.RewardCurator)),
		sdk.NewAttribute(AttrKeyBlockHeight, fmt.Sprintf("%d", blockHeight)),
	)
}

x/noorsignal/types/keys.go
package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/store/prefix"
)

// Nom officiel du module
const ModuleName = "noorsignal"

// Clé du store principal
const StoreKey = ModuleName

// Router key
const RouterKey = ModuleName

// -----------------------------------------------------------------------------
// Prefixes des sous-stores
// -----------------------------------------------------------------------------

var (
	// Configuration globale PoSS
	KeyPrefixPoSSConfig = []byte{0x01}

	// Signaux
	KeyPrefixSignal = []byte{0x02}

	// Curateurs
	KeyPrefixCurator = []byte{0x03}

	// Compteurs journaliers (anti-abus)
	KeyPrefixDailyCounter = []byte{0x04}

	// Next Signal ID
	KeyNextSignalID = []byte{0x05}
)

// -----------------------------------------------------------------------------
//  Sous-stores (prefix stores)
// -----------------------------------------------------------------------------

func GetConfigStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixPoSSConfig)
}

func GetSignalStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixSignal)
}

func GetCuratorStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixCurator)
}

func GetDailyCounterStore(parent storetypes.KVStore) prefix.Store {
	return prefix.NewStore(parent, KeyPrefixDailyCounter)
}

// -----------------------------------------------------------------------------
//  Clés utilitaires
// -----------------------------------------------------------------------------

// SignalKey construit la clé d’un signal à partir de son ID.
func SignalKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// DailyCounterKey construit la clé d’un compteur journalier.
func DailyCounterKey(addr sdk.AccAddress, dayBucket uint64) []byte {
	// format : address || dayBucket (8 bytes)
	day := make([]byte, 8)
	binary.BigEndian.PutUint64(day, dayBucket)

	return append(addr.Bytes(), day...)
}

x/noorsignal/types/msg.go
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
)

//
// -----------------------------------------------------------------------------
//  MsgSubmitSignal
// -----------------------------------------------------------------------------

type MsgSubmitSignal struct {
	Participant string `json:"participant" yaml:"participant"`
	Weight      uint32 `json:"weight" yaml:"weight"`
	Metadata    string `json:"metadata" yaml:"metadata"`
}

func (m MsgSubmitSignal) Route() string { return "noorsignal" }
func (m MsgSubmitSignal) Type() string  { return "submit_signal" }

func (m MsgSubmitSignal) ValidateBasic() error {
	if m.Participant == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "participant address cannot be empty")
	}
	if m.Weight == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "weight must be >= 1")
	}
	if m.Weight > 100 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "weight must be <= 100")
	}
	return nil
}

func (m MsgSubmitSignal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Participant)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgSubmitSignal) GetParticipantAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(m.Participant)
}

//
// -----------------------------------------------------------------------------
//  MsgValidateSignal
// -----------------------------------------------------------------------------

type MsgValidateSignal struct {
	Curator  string `json:"curator" yaml:"curator"`
	SignalId uint64 `json:"signal_id" yaml:"signal_id"`
}

func (m MsgValidateSignal) Route() string { return "noorsignal" }
func (m MsgValidateSignal) Type() string  { return "validate_signal" }

func (m MsgValidateSignal) ValidateBasic() error {
	if m.Curator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "curator address cannot be empty")
	}
	if m.SignalId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "signal_id must be >= 1")
	}
	return nil
}

func (m MsgValidateSignal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Curator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgValidateSignal) GetCuratorAddress() (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(m.Curator)
}

//
// -----------------------------------------------------------------------------
//  ADMIN MESSAGES
// -----------------------------------------------------------------------------
//  MsgAddCurator
//  MsgRemoveCurator
//  MsgSetConfig
// -----------------------------------------------------------------------------

// -----------------------------
// MsgAddCurator
// -----------------------------
type MsgAddCurator struct {
	Authority string `json:"authority" yaml:"authority"`
	Curator   string `json:"curator" yaml:"curator"`
	Level     string `json:"level" yaml:"level"`
}

func (m MsgAddCurator) Route() string { return "noorsignal" }
func (m MsgAddCurator) Type() string  { return "add_curator" }

func (m MsgAddCurator) ValidateBasic() error {
	if m.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}
	if m.Curator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "curator cannot be empty")
	}
	if m.Level == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "level cannot be empty")
	}
	return nil
}

func (m MsgAddCurator) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// -----------------------------
// MsgRemoveCurator
// -----------------------------
type MsgRemoveCurator struct {
	Authority string `json:"authority" yaml:"authority"`
	Curator   string `json:"curator" yaml:"curator"`
}

func (m MsgRemoveCurator) Route() string { return "noorsignal" }
func (m MsgRemoveCurator) Type() string  { return "remove_curator" }

func (m MsgRemoveCurator) ValidateBasic() error {
	if m.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}
	if m.Curator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "curator cannot be empty")
	}
	return nil
}

func (m MsgRemoveCurator) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// -----------------------------
// MsgSetConfig
// -----------------------------
type MsgSetConfig struct {
	Authority        string `json:"authority" yaml:"authority"`
	BaseReward       string `json:"base_reward" yaml:"base_reward"`
	MaxSignalsPerDay uint32 `json:"max_signals_per_day" yaml:"max_signals_per_day"`
	EraIndex         uint64 `json:"era_index" yaml:"era_index"`
	ParticipantRatio uint32 `json:"participant_ratio" yaml:"participant_ratio"`
	CuratorRatio     uint32 `json:"curator_ratio" yaml:"curator_ratio"`
}

func (m MsgSetConfig) Route() string { return "noorsignal" }
func (m MsgSetConfig) Type() string  { return "set_config" }

func (m MsgSetConfig) ValidateBasic() error {
	if m.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}

	if m.BaseReward == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "base_reward cannot be empty")
	}

	if (m.ParticipantRatio + m.CuratorRatio) != 100 {
		return errorsmod.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("participant_ratio + curator_ratio must equal 100 (got %d)",
				m.ParticipantRatio+m.CuratorRatio),
		)
	}

	return nil
}

func (m MsgSetConfig) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

x/noorsignal/types/query.go
package types

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
)

// -----------------------------------------------------------------------------
//  Query types for the PoSS module (noorsignal)
// -----------------------------------------------------------------------------
//
// NOTE:
// Ces structs sont des versions Go "conceptuelles" des messages définis
// dans proto/noorsignal/query.proto. La version finale sera générée à partir
// des fichiers .proto. Pour l'instant, elles nous permettent de réfléchir et
// de prototyper le QueryServer sans bloquer sur la génération protobuf.
// -----------------------------------------------------------------------------

// QuerySignalRequest représente une requête pour récupérer un signal
// unique via son identifiant.
type QuerySignalRequest struct {
	Id uint64 `json:"id" yaml:"id"`
}

// QuerySignalResponse contient un signal unique, si trouvé.
type QuerySignalResponse struct {
	Signal *Signal `json:"signal" yaml:"signal"`
}

// QuerySignalsRequest représente une requête pour récupérer une liste
// paginée de signaux PoSS.
type QuerySignalsRequest struct {
	Pagination *query.PageRequest `json:"pagination" yaml:"pagination"`
}

// QuerySignalsResponse contient une liste paginée de signaux.
type QuerySignalsResponse struct {
	Signals    []Signal             `json:"signals" yaml:"signals"`
	Pagination *query.PageResponse  `json:"pagination" yaml:"pagination"`
}

// QueryCuratorRequest représente une requête pour récupérer les
// informations d'un Curator PoSS via son adresse.
type QueryCuratorRequest struct {
	Address string `json:"address" yaml:"address"`
}

// QueryCuratorResponse contient les informations d'un Curator, si trouvé.
type QueryCuratorResponse struct {
	Curator *Curator `json:"curator" yaml:"curator"`
}

// QueryConfigRequest représente une requête pour obtenir la configuration
// globale PoSS actuelle.
type QueryConfigRequest struct {
	// pas de champ pour l'instant
}

// QueryConfigResponse contient la configuration PoSS actuelle.
type QueryConfigResponse struct {
	Config PossConfig `json:"config" yaml:"config"`
}

// QueryDailyCountRequest représente une requête pour connaître le nombre
// de signaux émis par une adresse donnée sur un "day bucket" donné
// (par ex. block_time / 86400).
type QueryDailyCountRequest struct {
	Address string `json:"address" yaml:"address"`
	Day     uint64 `json:"day" yaml:"day"`
}

// QueryDailyCountResponse contient le compteur de signaux quotidiens.
type QueryDailyCountResponse struct {
	Count uint32 `json:"count" yaml:"count"`
}

x/noorsignal/types/rewards.go
package types

// ComputeHalvingFactor retourne le facteur de division associé à une "ère"
// PoSS donnée.
//
// era = 0  → aucun halving (facteur 1)
// era = 1  → premier halving (facteur 2)
// era = 2  → deuxième halving (facteur 4)
// etc.
//
// Dans le modèle NOORCHAIN :
// - chaque "ère" correspond à une période de 8 ans
// - l'ère 0 couvre les 8 premières années
// - l'ère 1 couvre les années 8–16, etc.
func ComputeHalvingFactor(era uint32) uint64 {
	if era == 0 {
		return 1
	}

	// 1 << era = 2^era
	return 1 << era
}

// ComputeSignalRewards calcule la récompense totale PoSS pour un signal
// et la répartit entre participant et curator.
//
// Paramètres :
// - cfg   : configuration PoSS (BaseReward, shares, Enabled)
// - weight: poids du signal (1x, 2x, 5x, etc. sous forme d'entier)
// - era   : indice d'ère pour le halving (0 = aucune division, 1 = /2, etc.)
//
// Retourne :
// - totalReward : récompense totale (après halving) pour ce signal
// - participant : part pour le participant (70% typiquement)
// - curator     : part pour le curator (30% typiquement)
//
// Remarques :
// - si PoSS est désactivé (cfg.Enabled=false), ou BaseReward/weight=0,
//   tout le monde reçoit 0.
// - cette fonction ne gère pas les plafonds globaux, ni la supply max;
//   elle se contente de la logique locale par signal.
func ComputeSignalRewards(cfg PossConfig, weight uint32, era uint32) (totalReward uint64, participant uint64, curator uint64) {
	// Cas simples : PoSS désactivé ou paramètres nuls.
	if !cfg.Enabled || cfg.BaseReward == 0 || weight == 0 {
		return 0, 0, 0
	}

	// 1) Récompense brute = BaseReward * weight.
	base := cfg.BaseReward
	total := base * uint64(weight)

	// 2) Appliquer le halving selon l'ère.
	factor := ComputeHalvingFactor(era)
	if factor == 0 {
		// Sécurité : ne jamais diviser par 0.
		return 0, 0, 0
	}
	total = total / factor

	// 3) Appliquer la répartition 70% / 30% (ou autre selon cfg).
	//
	// On suppose que ParticipantShare + CuratorShare <= 100.
	// Si ce n'est pas le cas, le "reste" est implicitement non attribué.
	participant = total * uint64(cfg.ParticipantShare) / 100
	curator = total * uint64(cfg.CuratorShare) / 100

	return total, participant, curator
}

x/noorsignal/types/types.go
package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Signal représente un signal social PoSS émis sur NOORCHAIN.
//
// Il est d'abord créé par un participant via MsgSubmitSignal, puis
// potentiellement validé par un curator via MsgValidateSignal.
//
// Les champs de récompense (TotalReward, RewardParticipant,
// RewardCurator) sont remplis lors de la validation, sur la base de
// la configuration PoSS (BaseReward, EraIndex, 70/30, etc.).
type Signal struct {
	Id          uint64         `json:"id" protobuf:"varint,1,opt,name=id,proto3"`
	Participant sdk.AccAddress `json:"participant" protobuf:"bytes,2,opt,name=participant,proto3"`
	Curator     sdk.AccAddress `json:"curator" protobuf:"bytes,3,opt,name=curator,proto3"`
	Weight      uint32         `json:"weight" protobuf:"varint,4,opt,name=weight,proto3"`
	Time        time.Time      `json:"time" protobuf:"bytes,5,opt,name=time,proto3,stdtime"`
	Metadata    string         `json:"metadata" protobuf:"bytes,6,opt,name=metadata,proto3"`

	// Champs de récompense PoSS, remplis lors de la validation du signal.
	TotalReward       uint64 `json:"total_reward" protobuf:"varint,7,opt,name=total_reward,json=totalReward,proto3"`
	RewardParticipant uint64 `json:"reward_participant" protobuf:"varint,8,opt,name=reward_participant,json=rewardParticipant,proto3"`
	RewardCurator     uint64 `json:"reward_curator" protobuf:"varint,9,opt,name=reward_curator,json=rewardCurator,proto3"`
}

// Curator représente un validateur de signaux PoSS (Curator NOOR).
type Curator struct {
	Address              sdk.AccAddress `json:"address" protobuf:"bytes,1,opt,name=address,proto3"`
	Level                string         `json:"level" protobuf:"bytes,2,opt,name=level,proto3"`
	TotalSignalsValidated uint64        `json:"total_signals_validated" protobuf:"varint,3,opt,name=total_signals_validated,json=totalSignalsValidated,proto3"`
	Active               bool           `json:"active" protobuf:"varint,4,opt,name=active,proto3"`
}

// PossConfig contient la configuration globale du système PoSS.
//
// BaseReward : unité de base (ex: 100) pour le calcul, avant halving.
// ParticipantShare / CuratorShare : parts en pourcentage (ex: 70 / 30).
// MaxSignalsPerDay : limite de signaux par participant et par jour.
// Enabled : active ou non le système PoSS.
// EraIndex : indice d'ère pour le halving (0 = pas de division, 1 = /2, 2 = /4, etc.).
type PossConfig struct {
	BaseReward       uint64 `json:"base_reward" protobuf:"varint,1,opt,name=base_reward,json=baseReward,proto3"`
	ParticipantShare uint32 `json:"participant_share" protobuf:"varint,2,opt,name=participant_share,json=participantShare,proto3"`
	CuratorShare     uint32 `json:"curator_share" protobuf:"varint,3,opt,name=curator_share,json=curatorShare,proto3"`
	MaxSignalsPerDay uint32 `json:"max_signals_per_day" protobuf:"varint,4,opt,name=max_signals_per_day,json=maxSignalsPerDay,proto3"`
	Enabled          bool   `json:"enabled" protobuf:"varint,5,opt,name=enabled,proto3"`
	EraIndex         uint32 `json:"era_index" protobuf:"varint,6,opt,name=era_index,json=eraIndex,proto3"`
}

// DefaultPossConfig retourne une configuration PoSS par défaut.
func DefaultPossConfig() PossConfig {
	return PossConfig{
		BaseReward:       100, // valeur symbolique pour V1
		ParticipantShare: 70,
		CuratorShare:     30,
		MaxSignalsPerDay: 50, // limite quotidienne par participant
		Enabled:          true,
		EraIndex:         0,
	}
}

// GenesisState représente l'état initial du module noorsignal.
type GenesisState struct {
	Config       PossConfig `json:"config" protobuf:"bytes,1,opt,name=config,proto3"`
	Signals      []Signal   `json:"signals" protobuf:"bytes,2,rep,name=signals,proto3"`
	Curators     []Curator  `json:"curators" protobuf:"bytes,3,rep,name=curators,proto3"`

	// Comptes économiques PoSS (initialisés au genesis)
	PoSSReserveAddr  string `json:"poss_reserve_addr" protobuf:"bytes,4,opt,name=poss_reserve_addr,json=possReserveAddr,proto3"`
	FoundationAddr   string `json:"foundation_addr" protobuf:"bytes,5,opt,name=foundation_addr,json=foundationAddr,proto3"`
	DevWalletAddr    string `json:"dev_wallet_addr" protobuf:"bytes,6,opt,name=dev_wallet_addr,json=devWalletAddr,proto3"`
	StimulusAddr     string `json:"stimulus_addr" protobuf:"bytes,7,opt,name=stimulus_addr,json=stimulusAddr,proto3"`
	PreSaleAddr      string `json:"presale_addr" protobuf:"bytes,8,opt,name=presale_addr,json=presaleAddr,proto3"`
}

Archive noorsignal V2 (PoSS avancé)
