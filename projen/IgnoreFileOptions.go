package projen


// Experimental.
type IgnoreFileOptions struct {
	// Filter out comment lines?
	// Default: true.
	//
	// Experimental.
	FilterCommentLines *bool `field:"optional" json:"filterCommentLines" yaml:"filterCommentLines"`
	// Filter out blank/empty lines?
	// Default: true.
	//
	// Experimental.
	FilterEmptyLines *bool `field:"optional" json:"filterEmptyLines" yaml:"filterEmptyLines"`
	// Patterns to add to the ignore file.
	// Default: [].
	//
	// Experimental.
	IgnorePatterns *[]*string `field:"optional" json:"ignorePatterns" yaml:"ignorePatterns"`
}

