package javascript


// eslint rules override.
// Experimental.
type EslintOverride struct {
	// Files or file patterns on which to apply the override.
	// Experimental.
	Files *[]*string `field:"required" json:"files" yaml:"files"`
	// Pattern(s) to exclude from this override.
	//
	// If a file matches any of the excluded patterns, the configuration won’t apply.
	// Experimental.
	ExcludedFiles *[]*string `field:"optional" json:"excludedFiles" yaml:"excludedFiles"`
	// Config(s) to extend in this override.
	// Experimental.
	Extends *[]*string `field:"optional" json:"extends" yaml:"extends"`
	// The overridden parser.
	// Experimental.
	Parser *string `field:"optional" json:"parser" yaml:"parser"`
	// `plugins` override.
	// Experimental.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// The overridden rules.
	// Experimental.
	Rules *map[string]interface{} `field:"optional" json:"rules" yaml:"rules"`
}

