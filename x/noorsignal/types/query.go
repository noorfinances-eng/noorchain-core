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
