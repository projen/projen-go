package javascript


// A reference to another TypeScript project.
// Experimental.
type TypescriptConfigReference struct {
	// Path to a directory containing `tsconfig.json`, or to a configuration file.
	//
	// Relative paths are resolved from the referencing configuration file's directory.
	// Backslashes are converted to forward slashes. Absolute paths are not rebased.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

