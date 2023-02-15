package circleci


// Parameters are declared by name under a job, command, or executor.
// See: https://circleci.com/docs/2.0/reusing-config#using-the-parameters-declaration
//
// Experimental.
type PipelineParameter struct {
	// The parameter type, required.
	// Experimental.
	Type PipelineParameterType `field:"required" json:"type" yaml:"type"`
	// The default value for the parameter.
	//
	// If not present, the parameter is implied to be required.
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Used to generate documentation for your orb.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

