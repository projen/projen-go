package workflows


// Version requirement for tools.
// Experimental.
type ToolRequirement struct {
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
}

