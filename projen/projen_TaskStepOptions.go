// CDK for software projects
package projen


// Options for task steps.
// Experimental.
type TaskStepOptions struct {
	// The working directory for this step.
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// Step name.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

