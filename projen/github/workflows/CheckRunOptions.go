package workflows


// Check run options.
// Experimental.
type CheckRunOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

