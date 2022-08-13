package gitlab


// Filtering options for when a job will run.
// Experimental.
type Filter struct {
	// Filter job creation based on files that were modified in a git push.
	// Experimental.
	Changes *[]*string `field:"optional" json:"changes" yaml:"changes"`
	// Filter job based on if Kubernetes integration is active.
	// Experimental.
	Kubernetes KubernetesEnum `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Control when to add jobs to a pipeline based on branch names or pipeline types.
	// Experimental.
	Refs *[]*string `field:"optional" json:"refs" yaml:"refs"`
	// Filter job by checking comparing values of environment variables.
	//
	// Read more about variable expressions: https://docs.gitlab.com/ee/ci/variables/README.html#variables-expressions
	// Experimental.
	Variables *[]*string `field:"optional" json:"variables" yaml:"variables"`
}

