package javascript


// Experimental.
type EslintCommandOptions struct {
	// Extra flag arguments to pass to eslint command.
	//
	// Each element is passed to eslint as a single argument, exactly as given: no
	// shell parses these, so a flag and its value need separate elements
	// (`["--rulesdir", "my rules"]`, not `["--rulesdir 'my rules'"]`) and values
	// must not be quoted.
	//
	// Example:
	//   ["--cache", "--max-warnings=0"]
	//
	// Experimental.
	ExtraArgs *[]*string `field:"optional" json:"extraArgs" yaml:"extraArgs"`
	// Whether to fix eslint issues when running the eslint task.
	// Default: true.
	//
	// Experimental.
	Fix *bool `field:"optional" json:"fix" yaml:"fix"`
}

