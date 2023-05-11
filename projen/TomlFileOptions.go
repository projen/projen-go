package projen


// Options for `TomlFile`.
// Experimental.
type TomlFileOptions struct {
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
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `field:"optional" json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
}

