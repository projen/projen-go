package gitlab


// Used to control pipeline behavior.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflow
//
// Experimental.
type Workflow struct {
	// Used to control whether or not a whole pipeline is created.
	// Experimental.
	Rules *[]*WorkflowRule `field:"optional" json:"rules" yaml:"rules"`
}

