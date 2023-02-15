package gitlab


// The engine configuration for a secret.
// Experimental.
type Engine struct {
	// Name of the secrets engine.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path to the secrets engine.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

