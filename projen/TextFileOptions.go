package projen


// Options for `TextFile`.
// Experimental.
type TextFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Default: true.
	//
	// Experimental.
	Committed *bool `field:"optional" json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Default: true.
	//
	// Experimental.
	EditGitignore *bool `field:"optional" json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Default: false.
	//
	// Experimental.
	Executable *bool `field:"optional" json:"executable" yaml:"executable"`
	// Adds the projen marker to the file.
	// Default: - marker will be included as long as the project is not ejected.
	//
	// Experimental.
	Marker *bool `field:"optional" json:"marker" yaml:"marker"`
	// Whether the generated file should be readonly.
	// Default: true.
	//
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
	// The contents of the text file.
	//
	// You can use `addLine()` to append lines.
	// Default: [] empty file.
	//
	// Experimental.
	Lines *[]*string `field:"optional" json:"lines" yaml:"lines"`
}

