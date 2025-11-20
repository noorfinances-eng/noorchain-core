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

