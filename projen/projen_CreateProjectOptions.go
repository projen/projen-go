// CDK for software projects
package projen


// Experimental.
type CreateProjectOptions struct {
	// Directory that the project will be generated in.
	// Experimental.
	Dir *string `field:"required" json:"dir" yaml:"dir"`
	// Fully-qualified name of the project type (usually formatted as `module.ProjectType`).
	//
	// Example:
	//   `projen.TypescriptProject`
	//
	// Experimental.
	ProjectFqn *string `field:"required" json:"projectFqn" yaml:"projectFqn"`
	// Project options.
	//
	// Only JSON-like values can be passed in (strings,
	// booleans, numbers, enums, arrays, and objects that are not
	// derived from classes).
	//
	// Consult the API reference of the project type you are generating for
	// information about what fields and types are available.
	// Experimental.
	ProjectOptions *map[string]interface{} `field:"required" json:"projectOptions" yaml:"projectOptions"`
	// Should we render commented-out default options in the projenrc file?
	//
	// Does not apply to projenrc.json files.
	// Experimental.
	OptionHints InitProjectOptionHints `field:"optional" json:"optionHints" yaml:"optionHints"`
	// Should we execute post synthesis hooks?
	//
	// (usually package manager install).
	// Experimental.
	Post *bool `field:"optional" json:"post" yaml:"post"`
	// Should we call `project.synth()` or instantiate the project (could still have side-effects) and render the .projenrc file.
	// Experimental.
	Synth *bool `field:"optional" json:"synth" yaml:"synth"`
}

