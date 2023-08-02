package projen


// Resolve options.
// Experimental.
type ResolveOptions struct {
	// Context arguments.
	// Default: [].
	//
	// Experimental.
	Args *[]interface{} `field:"optional" json:"args" yaml:"args"`
	// Omits empty arrays and objects.
	// Default: false.
	//
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
}

