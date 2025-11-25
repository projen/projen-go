package uvconfig


// A subset of the Python Package Metadata 2.3 standard as specified in <https://packaging.python.org/specifications/core-metadata/>.
// Experimental.
type StaticMetadata struct {
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	ProvidesExtra *[]*string `field:"optional" json:"providesExtra" yaml:"providesExtra"`
	// Experimental.
	RequiresDist *[]*string `field:"optional" json:"requiresDist" yaml:"requiresDist"`
	// PEP 508-style Python requirement, e.g., `>=3.10`.
	// Experimental.
	RequiresPython *string `field:"optional" json:"requiresPython" yaml:"requiresPython"`
	// PEP 440-style package version, e.g., `1.2.3`.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

