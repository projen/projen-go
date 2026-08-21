package projen


// Options for selecting the GitHub Runner that a job runs on.
// Experimental.
type RunsOnOptions struct {
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
}

