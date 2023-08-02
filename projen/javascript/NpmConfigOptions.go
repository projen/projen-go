package javascript


// Options to configure the local NPM config.
// Experimental.
type NpmConfigOptions struct {
	// Omits empty objects and arrays.
	// Default: false.
	//
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
	// URL of the registry mirror to use.
	//
	// You can change this or add scoped registries using the addRegistry method.
	// Default: - use npmjs default registry.
	//
	// Experimental.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
}

