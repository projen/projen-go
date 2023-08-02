package web


// Deprecated: No longer used.
type ReactTypeDefOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Default: true.
	//
	// Deprecated: No longer used.
	Committed *bool `field:"optional" json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Default: true.
	//
	// Deprecated: No longer used.
	EditGitignore *bool `field:"optional" json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Default: false.
	//
	// Deprecated: No longer used.
	Executable *bool `field:"optional" json:"executable" yaml:"executable"`
	// Adds the projen marker to the file.
	// Default: - marker will be included as long as the project is not ejected.
	//
	// Deprecated: No longer used.
	Marker *bool `field:"optional" json:"marker" yaml:"marker"`
	// Whether the generated file should be readonly.
	// Default: true.
	//
	// Deprecated: No longer used.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

