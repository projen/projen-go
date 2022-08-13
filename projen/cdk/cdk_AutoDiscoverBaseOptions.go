package cdk


// Options for `AutoDiscoverBase`.
// Experimental.
type AutoDiscoverBaseOptions struct {
	// Locate files with the given extension.
	//
	// Example:
	//   ".integ.ts"
	//
	// Experimental.
	Extension *string `field:"required" json:"extension" yaml:"extension"`
	// Locate entrypoints in the given project directory.
	//
	// Example:
	//   "test"
	//
	// Experimental.
	Projectdir *string `field:"required" json:"projectdir" yaml:"projectdir"`
}

