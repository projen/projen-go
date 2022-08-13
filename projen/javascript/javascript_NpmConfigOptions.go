package javascript


// Options to configure the local NPM config.
// Experimental.
type NpmConfigOptions struct {
	// URL of the registry mirror to use.
	//
	// You can change this or add scoped registries using the addRegistry method.
	// Experimental.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
}

