package workflows


// Run settings for a job.
// Experimental.
type RunSettings struct {
	// Which shell to use for running the step.
	//
	// Example:
	//   "bash"
	//
	// Experimental.
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// Working directory to use when running the step.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

