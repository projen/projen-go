package github


// The Mergify conditional operators that can be used are: `or` and `and`.
//
// Note: The number of nested conditions is limited to 3.
// See: https://docs.mergify.io/conditions/#combining-conditions-with-operators
//
// Experimental.
type MergifyConditionalOperator struct {
	// Experimental.
	And *[]interface{} `field:"optional" json:"and" yaml:"and"`
	// Experimental.
	Or *[]interface{} `field:"optional" json:"or" yaml:"or"`
}

