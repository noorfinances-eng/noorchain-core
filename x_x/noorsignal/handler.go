package noorsignal

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/noorfinances-eng/noorchain-core/x/noorsignal/keeper"
	noorsignaltypes "github.com/noorfinances-eng/noorchain-core/x/noorsignal/types"
)

// NewHandler crée le handler legacy pour les Msg du module x/noorsignal.
//
// Objectif : rendre MsgCreateSignal exécutable via le router Cosmos SDK
// (chemin legacy), sans dépendre d'un MsgServer protobuf encore inexistant.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch m := msg.(type) {

		case *noorsignaltypes.MsgCreateSignal:
			// 1) Validation de base (adresses, type, date).
			if err := m.ValidateBasic(); err != nil {
				return nil, err
			}

			// 2) Appel du pipeline interne PoSS.
			participantReward, curatorReward, err := k.ProcessSignalInternal(
				ctx,
				m.Participant,
				m.Curator,
				m.SignalType,
				m.Date,
			)
			if err != nil {
				return nil, err
			}

			// 3) Construire la réponse logique (rewards théoriques).
			// Pour l’instant on ne sérialise pas la réponse dans Result.Data,
			// mais on émet des events pour indexers / explorateurs.
			_ = noorsignaltypes.MsgCreateSignalResponse{
				ParticipantReward: participantReward,
				CuratorReward:     curatorReward,
			}

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"noorsignal_create_signal",
					sdk.NewAttribute("participant", m.Participant),
					sdk.NewAttribute("curator", m.Curator),
					sdk.NewAttribute("signal_type", string(m.SignalType)),
					sdk.NewAttribute("participant_reward", participantReward.String()),
					sdk.NewAttribute("curator_reward", curatorReward.String()),
				),
			)

			// On retourne un Result minimal : les events sont déjà
			// pris en charge via le Context par Tendermint/CometBFT.
			return &sdk.Result{}, nil

		default:
			return nil, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %T",
				noorsignaltypes.ModuleName, msg,
			)
		}
	}
}
