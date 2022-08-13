package circleci


// Specifies which triggers will cause this workflow to be executed.
//
// Default behavior is to trigger the workflow when pushing to a branch.
// See: https://circleci.com/docs/2.0/configuration-reference/#triggers
//
// Experimental.
type Triggers struct {
	// Experimental.
	Schedule *Schedule `field:"optional" json:"schedule" yaml:"schedule"`
}

