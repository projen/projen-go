package gitlab


// A jobs in a previous stage whose sole completion is needed to start the current job.
// Experimental.
type Need struct {
	// Experimental.
	Job *string `field:"required" json:"job" yaml:"job"`
	// Experimental.
	Artifacts *bool `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Experimental.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	// Experimental.
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// Experimental.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

