package gitlab


// Used to control whether or not a whole pipeline is created.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflowrules
//
// Experimental.
type WorkflowRule struct {
	// Experimental.
	Changes *[]*string `field:"optional" json:"changes" yaml:"changes"`
	// Experimental.
	Exists *[]*string `field:"optional" json:"exists" yaml:"exists"`
	// Experimental.
	If *string `field:"optional" json:"if" yaml:"if"`
	// Experimental.
	Variables *map[string]interface{} `field:"optional" json:"variables" yaml:"variables"`
	// Experimental.
	When WorkflowWhen `field:"optional" json:"when" yaml:"when"`
}

