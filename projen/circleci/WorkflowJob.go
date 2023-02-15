package circleci


// A Job is part of Workflow.
//
// A Job can be created with {@link Job} or it can be provided by the orb.
// See: https://circleci.com/docs/2.0/configuration-reference/#jobs-in-workflow
//
// Experimental.
type WorkflowJob struct {
	// name of dynamic key *.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The name of the context(s).
	//
	// The initial default name is org-global. Each context name must be unique.
	// Experimental.
	Context *[]*string `field:"optional" json:"context" yaml:"context"`
	// Job Filters can have the key branches or tags.
	// Experimental.
	Filters *Filter `field:"optional" json:"filters" yaml:"filters"`
	// Experimental.
	Matrix *Matrix `field:"optional" json:"matrix" yaml:"matrix"`
	// A replacement for the job name.
	//
	// Useful when calling a job multiple times.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Parameters passed to job when referencing a job from orb.
	// Experimental.
	OrbParameters *map[string]interface{} `field:"optional" json:"orbParameters" yaml:"orbParameters"`
	// A list of jobs that must succeed for the job to start.
	// Experimental.
	Requires *[]*string `field:"optional" json:"requires" yaml:"requires"`
	// A job may have a type of approval indicating it must be manually approved before downstream jobs may proceed.
	// Experimental.
	Type JobType `field:"optional" json:"type" yaml:"type"`
}

