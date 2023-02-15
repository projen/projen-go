package circleci


// Execution steps for Job.
// See: https://circleci.com/docs/2.0/configuration-reference/#steps
//
// Experimental.
type StepRun struct {
	// Experimental.
	Run *Run `field:"optional" json:"run" yaml:"run"`
}

