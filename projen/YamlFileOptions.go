package projen


// Options for `JsonFile`.
// Experimental.
type YamlFileOptions struct {
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
	// The object that will be serialized. You can modify the object's contents before synthesis.
	//
	// Serialization of the object is similar to JSON.stringify with few enhancements:
	// - values that are functions will be called during synthesis and the result will be serialized - this allow to have lazy values.
	// - `Set` will be converted to array
	// - `Map` will be converted to a plain object ({ key: value, ... }})
	// - `RegExp` without flags will be converted to string representation of the source.
	// Default: {} an empty object (use `file.obj` to mutate).
	//
	// Experimental.
	Obj interface{} `field:"optional" json:"obj" yaml:"obj"`
	// Omits empty objects and arrays.
	// Default: false.
	//
	// Experimental.
	OmitEmpty *bool `field:"optional" json:"omitEmpty" yaml:"omitEmpty"`
	// Maximum line width (set to 0 to disable folding).
	// Default: - 0.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}

