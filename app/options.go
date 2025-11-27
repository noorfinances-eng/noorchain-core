package app

// AppOptions définit une interface minimale pour récupérer des options
// (flags, configuration, etc.) au moment de l'initialisation de NOORCHAIN.
//
// Pour la Phase 2, on reste volontairement très simple.
// Plus tard, on branchera un vrai système (Viper / Cobra) dessus.
type AppOptions interface {
	// Get retourne la valeur associée à une clé de configuration.
	// Si la clé n'existe pas, l'implémentation peut retourner nil.
	Get(key string) interface{}
}
