package gitlab


// Specification for a secret provided by a HashiCorp Vault.
// See: https://www.vaultproject.io/
//
// Experimental.
type VaultConfig struct {
	// Experimental.
	Engine *Engine `field:"required" json:"engine" yaml:"engine"`
	// Experimental.
	Field *string `field:"required" json:"field" yaml:"field"`
	// Path to the secret.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

