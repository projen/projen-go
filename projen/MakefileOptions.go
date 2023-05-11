package projen


// Options for Makefiles.
// Experimental.
type MakefileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `field:"optional" json:"committed" yaml:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `field:"optional" json:"editGitignore" yaml:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `field:"optional" json:"executable" yaml:"executable"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker *bool `field:"optional" json:"marker" yaml:"marker"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
	// List of targets to build when Make is invoked without specifying any targets.
	// Experimental.
	All *[]*string `field:"optional" json:"all" yaml:"all"`
	// Rules to include in the Makefile.
	// Experimental.
	Rules *[]*Rule `field:"optional" json:"rules" yaml:"rules"`
}

