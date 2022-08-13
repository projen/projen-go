package gitlab


// A CI/CD secret.
// Experimental.
type Secret struct {
	// Experimental.
	Vault *VaultConfig `field:"required" json:"vault" yaml:"vault"`
}

