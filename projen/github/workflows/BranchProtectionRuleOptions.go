package workflows


// Branch Protection Rule options.
// Experimental.
type BranchProtectionRuleOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

