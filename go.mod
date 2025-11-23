module github.com/noorfinances-eng/noorchain-core

go 1.24

require (
    github.com/cometbft/cometbft v1.0.0
    github.com/cosmos/cosmos-sdk v0.50.5
)

// --- Corrections de compatibilité ---
// 1) Certaines libs demandent encore tendermint v0.34.27,
//    mais cette révision n'est plus disponible.
//    On redirige explicitement vers une version stable existante.
replace github.com/tendermint/tendermint v0.34.27 => github.com/tendermint/tendermint v0.34.24

// 2) Même problème déjà vu avec gogo/protobuf v1.3.3 : on force 1.3.2.
replace github.com/gogo/protobuf v1.3.3 => github.com/gogo/protobuf v1.3.2
