package workflows


// Probject column options.
// Experimental.
type ProjectColumnOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

