// CDK for software projects
package projen


// A Make rule.
// Experimental.
type Rule struct {
	// Files to be created or updated by this rule.
	//
	// If the rule is phony then instead this represents the command's name(s).
	// Experimental.
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// Marks whether the target is phony.
	// Experimental.
	Phony *bool `field:"optional" json:"phony" yaml:"phony"`
	// Files that are used as inputs to create a target.
	// Experimental.
	Prerequisites *[]*string `field:"optional" json:"prerequisites" yaml:"prerequisites"`
	// Commands that are run (using prerequisites as inputs) to create a target.
	// Experimental.
	Recipe *[]*string `field:"optional" json:"recipe" yaml:"recipe"`
}

