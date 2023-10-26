package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// Experimental.
type VsCode interface {
	projen.Component
	// Experimental.
	Extensions() VsCodeRecommendedExtensions
	// Experimental.
	LaunchConfiguration() VsCodeLaunchConfig
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Experimental.
	Settings() VsCodeSettings
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

// The jsii proxy struct for VsCode
type jsiiProxy_VsCode struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_VsCode) Extensions() VsCodeRecommendedExtensions {
	var returns VsCodeRecommendedExtensions
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCode) LaunchConfiguration() VsCodeLaunchConfig {
	var returns VsCodeLaunchConfig
	_jsii_.Get(
		j,
		"launchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCode) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCode) Settings() VsCodeSettings {
	var returns VsCodeSettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func NewVsCode(project projen.Project) VsCode {
	_init_.Initialize()

	if err := validateNewVsCodeParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_VsCode{}

	_jsii_.Create(
		"projen.vscode.VsCode",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewVsCode_Override(v VsCode, project projen.Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.vscode.VsCode",
		[]interface{}{project},
		v,
	)
}

// Test whether the given construct is a component.
// Experimental.
func VsCode_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsCode_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.VsCode",
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
func VsCode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsCode_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.VsCode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsCode) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCode) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCode) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

