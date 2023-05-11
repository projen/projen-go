package projen


// Coordinates of the dependency (name and version).
// Experimental.
type DependencyCoordinates struct {
	// The package manager name of the dependency (e.g. `leftpad` for npm).
	//
	// NOTE: For package managers that use complex coordinates (like Maven), we
	// will codify it into a string somehow.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Semantic version version requirement.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

