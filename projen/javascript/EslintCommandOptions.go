package javascript


// Experimental.
type EslintCommandOptions struct {
	// Extra flag arguments to pass to eslint command.
	// Experimental.
	ExtraArgs *[]*string `field:"optional" json:"extraArgs" yaml:"extraArgs"`
	// Whether to fix eslint issues when running the eslint task.
	// Default: true.
	//
	// Experimental.
	Fix *bool `field:"optional" json:"fix" yaml:"fix"`
}

