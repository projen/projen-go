package circleci


// Used for orchestrating all jobs.
//
// Each workflow consists of the workflow name as a key and a map as a value.
// A name should be unique within the current config.yml.
// The top-level keys for the Workflows configuration are version and jobs.
// See: https://circleci.com/docs/2.0/configuration-reference/#workflows
//
// Experimental.
type Workflow struct {
	// name of dynamic key *.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Experimental.
	Jobs *[]*WorkflowJob `field:"optional" json:"jobs" yaml:"jobs"`
	// Experimental.
	Triggers *[]*Triggers `field:"optional" json:"triggers" yaml:"triggers"`
	// when is too dynamic to be casted to interfaces.
	//
	// Check Docu as reference.
	// See: https://circleci.com/docs/2.0/configuration-reference/#logic-statement-examples
	//
	// Experimental.
	When interface{} `field:"optional" json:"when" yaml:"when"`
}

