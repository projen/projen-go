// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Programmatic API for projen.
// Experimental.
type Projects interface {
}

// The jsii proxy struct for Projects
type jsiiProxy_Projects struct {
	_ byte // padding
}

// Creates a new project with defaults.
//
// This function creates the project type in-process (with in VM) and calls
// `.synth()` on it (if `options.synth` is not `false`).
//
// At the moment, it also generates a `.projenrc.js` file with the same code
// that was just executed. In the future, this will also be done by the project
// type, so we can easily support multiple languages of projenrc.
// Experimental.
func Projects_CreateProject(options *CreateProjectOptions) {
	_init_.Initialize()

	if err := validateProjects_CreateProjectParameters(options); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"projen.Projects",
		"createProject",
		[]interface{}{options},
	)
}

