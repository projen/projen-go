package gitlab


// Explains what the global variable is used for, what the acceptable values are.
// See: https://docs.gitlab.com/ee/ci/yaml/#variables
//
// Experimental.
type VariableConfig struct {
	// Define a global variable that is prefilled when running a pipeline manually.
	//
	// Must be used with value.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The variable value.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

