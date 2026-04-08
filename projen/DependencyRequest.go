package projen


// A request for a dependency.
//
// Unlike adding a dependency directly,
// requesting a dependency will intelligently merge with existing
// dependencies of the same name and type.
// Experimental.
type DependencyRequest struct {
	// The package name.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional metadata.
	// Default: - none.
	//
	// Experimental.
	Metadata *map[string]interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Dependency type.
	//
	// If not provided, an existing dependency of any type
	// will satisfy the request. If none exists, it is added as BUILD.
	// Default: - any existing type, or DependencyType.BUILD
	//
	// Experimental.
	Type DependencyType `field:"optional" json:"type" yaml:"type"`
	// Semantic version constraint.
	// Default: - any version.
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

