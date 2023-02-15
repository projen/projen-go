package workflows


// Registry package options.
// Experimental.
type RegistryPackageOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

