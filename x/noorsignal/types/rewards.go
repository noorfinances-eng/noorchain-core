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
