package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// MsgServer est le point d'entrée pour les transactions (Msg)
// du module PoSS (noorsignal).
type MsgServer struct {
	Keeper
}

// NewMsgServer construit un MsgServer à partir d'un Keeper PoSS.
func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{Keeper: k}
}

// SubmitSignal gère la réception d'un MsgSubmitSignal
// (émission d'un nouveau signal social PoSS).
//
// Étapes :
// - validation du poids
// - récupération de la config PoSS (MaxSignalsPerDay, etc.)
// - calcul du "jour" (dayBucket) à partir du BlockTime
// - vérification de la limite quotidienne pour le participant
// - création et stockage du signal
// - incrément du compteur quotidien.
func (s MsgServer) SubmitSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgSubmitSignal,
) (*sdk.Result, error) {
	// 1) Récupérer le sdk.Context à partir du context gRPC.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 2) Valider le poids du signal.
	if msg.Weight == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be >= 1")
	}
	if msg.Weight > 100 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "weight must be <= 100")
	}

	// 3) Convertir l'adresse du participant (Bech32 -> sdk.AccAddress).
	participantAddr, err := msg.GetParticipantAddress()
	if err != nil {
		return nil, err
	}

	// 4) Récupérer la configuration PoSS.
	cfg, found := s.Keeper.GetConfig(ctx)
	if !found {
		// Si aucune config n'est présente, on utilise simplement
		// la config par défaut (sans l'enregistrer forcément).
		cfg = noorsignaltypes.DefaultPossConfig()
	}

	// 5) Vérifier la limite quotidienne si elle est active (> 0).
	var dayBucket uint64
	if cfg.MaxSignalsPerDay > 0 {
		// Calcul du "jour" sous forme d'entier : timestampUnix / 86400.
		ts := ctx.BlockTime().Unix()
		if ts < 0 {
			ts = 0
		}
		dayBucket = uint64(ts) / 86400

		current := s.getDailySignalCount(ctx, participantAddr, dayBucket)
		if current >= cfg.MaxSignalsPerDay {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "daily signal limit reached")
		}
	}

	// 6) Construire un Signal de base (sans curator pour l'instant).
	signal := noorsignaltypes.Signal{
		Participant: participantAddr,
		Curator:     nil,
		Weight:      msg.Weight,
		Time:        ctx.BlockTime(),
		Metadata:    msg.Metadata,
	}

	// 7) Créer et stocker le signal via le Keeper.
	_ = s.Keeper.CreateSignal(ctx, signal)

	// 8) Incrémenter le compteur quotidien si une limite est active.
	if cfg.MaxSignalsPerDay > 0 {
		s.incrementDailySignalCount(ctx, participantAddr, dayBucket)
	}

	// 9) Retourner un sdk.Result simple (sans events pour l'instant).
	return &sdk.Result{}, nil
}

// ValidateSignal gère la réception d'un MsgValidateSignal
// (validation d'un signal existant par un curator).
func (s MsgServer) ValidateSignal(
	goCtx context.Context,
	msg *noorsignaltypes.MsgValidateSignal,
) (*sdk.Result, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1) Convertir l'adresse du curator (Bech32 -> sdk.AccAddress).
	curatorAddr, err := msg.GetCuratorAddress()
	if err != nil {
		return nil, err
	}

	// 2) Récupérer le signal à valider.
	signal, found := s.Keeper.GetSignal(ctx, msg.SignalId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "signal not found")
	}

	// 3) Vérifier qu'il n'est pas déjà validé.
	if signal.Curator != nil && len(signal.Curator) > 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "signal already validated")
	}

	// 4) Associer le curator au signal.
	signal.Curator = curatorAddr

	// 5) Mettre à jour le signal dans le store.
	s.Keeper.SetSignal(ctx, signal)

	// TODO (plus tard) :
	// - vérifier que le curator est autorisé
	// - calculer et distribuer les récompenses 70% / 30%
	// - émettre des events.

	return &sdk.Result{}, nil
}
