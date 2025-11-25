package python

import (
	"github.com/projen/projen-go/projen/python/uvconfig"
)

// Options for UV project.
// Experimental.
type UvOptions struct {
	// Path to the python executable to use.
	// Default: "python".
	//
	// Experimental.
	PythonExec *string `field:"optional" json:"pythonExec" yaml:"pythonExec"`
	// Declares any Python level dependencies that must be installed in order to run the project’s build system successfully.
	// Default: - no build system.
	//
	// Experimental.
	BuildSystem *BuildSystem `field:"optional" json:"buildSystem" yaml:"buildSystem"`
	// The project's basic metadata configuration.
	// Experimental.
	Project *PyprojectTomlProject `field:"optional" json:"project" yaml:"project"`
	// The configuration and metadata for uv.
	// Experimental.
	Uv *uvconfig.UvConfiguration `field:"optional" json:"uv" yaml:"uv"`
}

