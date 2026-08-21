package projen


// Resolved `runsOn`/`runsOnGroup` config for a job.
//
// Exactly one of the two
// fields is set.
// Experimental.
type RunsOnConfig struct {
	// Github Runner selection labels.
	// Default: - not set if `runsOnGroup` is used.
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Default: - not set if `runsOn` is used.
	//
	// Experimental.
	RunsOnGroup *GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
}

