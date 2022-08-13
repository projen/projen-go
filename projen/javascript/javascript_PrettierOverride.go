package javascript


// Experimental.
type PrettierOverride struct {
	// Include these files in this override.
	// Experimental.
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// The options to apply for this override.
	// Experimental.
	Settings *PrettierSettings `field:"required" json:"settings" yaml:"settings"`
	// Exclude these files from this override.
	// Experimental.
	ExcludeFiles interface{} `field:"optional" json:"excludeFiles" yaml:"excludeFiles"`
}

