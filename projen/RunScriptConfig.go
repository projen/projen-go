package projen


// The resolved configuration for running a script.
// Experimental.
type RunScriptConfig struct {
	// Dependencies required to run the script.
	// Experimental.
	Dependencies *[]*DependencyRequest `field:"required" json:"dependencies" yaml:"dependencies"`
	// The task steps to execute the script.
	// Experimental.
	Steps *[]*TaskStep `field:"required" json:"steps" yaml:"steps"`
}

