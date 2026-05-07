package workflows


// An input for a workflow_dispatch event.
// Experimental.
type WorkflowDispatchInput struct {
	// A description of the input.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The default value of the input.
	// Default: - no default.
	//
	// Experimental.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Options for the input if type is 'choice'.
	// Default: - no options.
	//
	// Experimental.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Whether the input is required.
	// Default: false.
	//
	// Experimental.
	Required *bool `field:"optional" json:"required" yaml:"required"`
	// The type of the input.
	// Default: "string".
	//
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

