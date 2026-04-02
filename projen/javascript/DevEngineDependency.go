package javascript


// A dependency entry for the `devEngines` field in `package.json`.
// Experimental.
type DevEngineDependency struct {
	// The name of the dependency.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// What action to take if validation fails.
	// Default: "error".
	//
	// Experimental.
	OnFail *string `field:"optional" json:"onFail" yaml:"onFail"`
	// The version range for the dependency.
	// Default: "*".
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

