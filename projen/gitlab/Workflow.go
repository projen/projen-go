package gitlab


// Used to control pipeline behavior.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflow
//
// Experimental.
type Workflow struct {
	// You can use name to define a name for pipelines.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Used to control whether or not a whole pipeline is created.
	// Experimental.
	Rules *[]*WorkflowRule `field:"optional" json:"rules" yaml:"rules"`
}

