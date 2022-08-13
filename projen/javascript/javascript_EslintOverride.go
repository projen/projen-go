package javascript


// eslint rules override.
// Experimental.
type EslintOverride struct {
	// Files or file patterns on which to apply the override.
	// Experimental.
	Files *[]*string `field:"required" json:"files" yaml:"files"`
	// The overridden parser.
	// Experimental.
	Parser *string `field:"optional" json:"parser" yaml:"parser"`
	// The overriden rules.
	// Experimental.
	Rules *map[string]interface{} `field:"optional" json:"rules" yaml:"rules"`
}

