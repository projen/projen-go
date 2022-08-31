package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manages a virtual environment through the Python venv module.
// Experimental.
type Venv interface {
	projen.Component
	IPythonEnv
	// Experimental.
	Project() projen.Project
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Venv
type jsiiProxy_Venv struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonEnv
}

func (j *jsiiProxy_Venv) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVenv(project projen.Project, options *VenvOptions) Venv {
	_init_.Initialize()

	if err := validateNewVenvParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Venv{}

	_jsii_.Create(
		"projen.python.Venv",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewVenv_Override(v Venv, project projen.Project, options *VenvOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Venv",
		[]interface{}{project, options},
		v,
	)
}

func (v *jsiiProxy_Venv) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		v,
		"setupEnvironment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

