package workflows


// Project card options.
// Experimental.
type ProjectCardOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

