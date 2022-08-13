package javascript


// Options for scoped packages.
// Experimental.
type ScopedPackagesOptions struct {
	// URL of the registry for scoped packages.
	// Experimental.
	RegistryUrl *string `field:"required" json:"registryUrl" yaml:"registryUrl"`
	// Scope of the packages.
	//
	// Example:
	//   "@angular"
	//
	// Experimental.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

