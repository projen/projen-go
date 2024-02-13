package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

// Experimental.
type IPythonDeps interface {
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Installs dependencies (called during post-synthesis).
	// Experimental.
	InstallDependencies()
	// A task that installs and updates dependencies.
	// Experimental.
	InstallCiTask() projen.Task
}

// The jsii proxy for IPythonDeps
type jsiiProxy_IPythonDeps struct {
	_ byte // padding
}

func (i *jsiiProxy_IPythonDeps) AddDependency(spec *string) {
	if err := i.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDependency",
		[]interface{}{spec},
	)
}

func (i *jsiiProxy_IPythonDeps) AddDevDependency(spec *string) {
	if err := i.validateAddDevDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (i *jsiiProxy_IPythonDeps) InstallDependencies() {
	_jsii_.InvokeVoid(
		i,
		"installDependencies",
		nil, // no parameters
	)
}

func (j *jsiiProxy_IPythonDeps) InstallCiTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installCiTask",
		&returns,
	)
	return returns
}

