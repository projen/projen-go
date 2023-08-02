package projen


// Represents a project dependency.
// Experimental.
type Dependency struct {
	// The package manager name of the dependency (e.g. `leftpad` for npm).
	//
	// NOTE: For package managers that use complex coordinates (like Maven), we
	// will codify it into a string somehow.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Semantic version version requirement.
	// Default: - requirement is managed by the package manager (e.g. npm/yarn).
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Which type of dependency this is (runtime, build-time, etc).
	// Experimental.
	Type DependencyType `field:"required" json:"type" yaml:"type"`
	// Additional JSON metadata associated with the dependency (package manager specific).
	// Default: {}.
	//
	// Experimental.
	Metadata *map[string]interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

