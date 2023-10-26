package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// VSCode launch configuration file (launch.json), useful for enabling in-editor debugger.
// Experimental.
type VsCodeLaunchConfig interface {
	projen.Component
	// Experimental.
	File() projen.JsonFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Adds a VsCodeLaunchConfigurationEntry (e.g. a node.js debugger) to `.vscode/launch.json. Each configuration entry has following mandatory fields: type, request and name. See https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes for details.
	// Experimental.
	AddConfiguration(cfg *VsCodeLaunchConfigurationEntry)
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

// The jsii proxy struct for VsCodeLaunchConfig
type jsiiProxy_VsCodeLaunchConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_VsCodeLaunchConfig) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCodeLaunchConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCodeLaunchConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVsCodeLaunchConfig(vscode VsCode) VsCodeLaunchConfig {
	_init_.Initialize()

	if err := validateNewVsCodeLaunchConfigParameters(vscode); err != nil {
		panic(err)
	}
	j := jsiiProxy_VsCodeLaunchConfig{}

	_jsii_.Create(
		"projen.vscode.VsCodeLaunchConfig",
		[]interface{}{vscode},
		&j,
	)

	return &j
}

// Experimental.
func NewVsCodeLaunchConfig_Override(v VsCodeLaunchConfig, vscode VsCode) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.vscode.VsCodeLaunchConfig",
		[]interface{}{vscode},
		v,
	)
}

// Test whether the given construct is a component.
// Experimental.
func VsCodeLaunchConfig_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsCodeLaunchConfig_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.VsCodeLaunchConfig",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func VsCodeLaunchConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsCodeLaunchConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.VsCodeLaunchConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsCodeLaunchConfig) AddConfiguration(cfg *VsCodeLaunchConfigurationEntry) {
	if err := v.validateAddConfigurationParameters(cfg); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addConfiguration",
		[]interface{}{cfg},
	)
}

func (v *jsiiProxy_VsCodeLaunchConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeLaunchConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeLaunchConfig) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeLaunchConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

