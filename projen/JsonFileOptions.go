package projen


// Options for `JsonFile`.
// Experimental.
type JsonFileOptions struct {
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
	// The object that will be serialized.
	//
	// You can modify the object's contents
	// before synthesis.
	// Default: {} an empty object (use `file.obj` to mutate).
	//
	// Experimental.
	Obj interface{} `field:"optional" json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Default: false.
	//
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
	// Allow the use of comments in this file.
	// Default: - false for .json files, true for .json5 and .jsonc files
	//
	// Experimental.
	AllowComments *bool `field:"optional" json:"allowComments" yaml:"allowComments"`
	// Adds a newline at the end of the file.
	// Default: true.
	//
	// Experimental.
	Newline *bool `field:"optional" json:"newline" yaml:"newline"`
}

