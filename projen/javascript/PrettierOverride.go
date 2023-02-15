package javascript


// Experimental.
type PrettierOverride struct {
	// Include these files in this override.
	// Experimental.
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// The options to apply for this override.
	// Experimental.
	Options *PrettierSettings `field:"required" json:"options" yaml:"options"`
	// Exclude these files from this override.
	// Experimental.
	ExcludeFiles interface{} `field:"optional" json:"excludeFiles" yaml:"excludeFiles"`
}

