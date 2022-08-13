package workflows


// A strategy creates a build matrix for your jobs.
//
// You can define different
// variations to run each job in.
// Experimental.
type JobStrategy struct {
	// When set to true, GitHub cancels all in-progress jobs if any matrix job fails.
	//
	// Default: true.
	// Experimental.
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	// You can define a matrix of different job configurations.
	//
	// A matrix allows
	// you to create multiple jobs by performing variable substitution in a
	// single job definition. For example, you can use a matrix to create jobs
	// for more than one supported version of a programming language, operating
	// system, or tool. A matrix reuses the job's configuration and creates a
	// job for each matrix you configure.
	//
	// A job matrix can generate a maximum of 256 jobs per workflow run. This
	// limit also applies to self-hosted runners.
	// Experimental.
	Matrix *JobMatrix `field:"optional" json:"matrix" yaml:"matrix"`
	// The maximum number of jobs that can run simultaneously when using a matrix job strategy.
	//
	// By default, GitHub will maximize the number of jobs
	// run in parallel depending on the available runners on GitHub-hosted
	// virtual machines.
	// Experimental.
	MaxParallel *float64 `field:"optional" json:"maxParallel" yaml:"maxParallel"`
}

