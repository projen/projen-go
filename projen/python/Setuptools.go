package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manages packaging through setuptools with a setup.py script.
// Experimental.
type Setuptools interface {
	projen.Component
	IPythonPackaging
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
	// A task that uploads the package to the Test PyPI repository.
	// Experimental.
	PublishTestTask() projen.Task
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Setuptools
type jsiiProxy_Setuptools struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Setuptools) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) PublishTestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTestTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewSetuptools(project projen.Project, options *PythonPackagingOptions) Setuptools {
	_init_.Initialize()

	if err := validateNewSetuptoolsParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Setuptools{}

	_jsii_.Create(
		"projen.python.Setuptools",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSetuptools_Override(s Setuptools, project projen.Project, options *PythonPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Setuptools",
		[]interface{}{project, options},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Setuptools_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSetuptools_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Setuptools",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Setuptools_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSetuptools_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Setuptools",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Setuptools) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

