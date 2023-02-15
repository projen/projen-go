package workflows


// Default settings for all steps in the job.
// Experimental.
type JobDefaults struct {
	// Default run settings.
	// Experimental.
	Run *RunSettings `field:"optional" json:"run" yaml:"run"`
}

