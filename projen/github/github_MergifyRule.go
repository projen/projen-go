package github


// Experimental.
type MergifyRule struct {
	// A dictionary made of Actions that will be executed on the matching pull requests.
	// See: https://docs.mergify.io/actions/#actions
	//
	// Experimental.
	Actions *map[string]interface{} `field:"required" json:"actions" yaml:"actions"`
	// A list of Conditions string that must match against the pull request for the rule to be applied.
	// See: https://docs.mergify.io/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The name of the rule.
	//
	// This is not used by the engine directly,
	// but is used when reporting information about a rule.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

