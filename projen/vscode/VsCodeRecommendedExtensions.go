package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// VS Code Workspace recommended extensions Source: https://code.visualstudio.com/docs/editor/extension-marketplace#_workspace-recommended-extensions.
// Experimental.
type VsCodeRecommendedExtensions interface {
	projen.Component
	// Experimental.
	File() projen.JsonFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Adds a list of VS Code extensions as recommendations for this workspace.
	// Experimental.
	AddRecommendations(extensions ...*string)
	// Marks a list of VS Code extensions as unwanted recommendations for this workspace.
	//
	// VS Code should not be recommend these extensions for users of this workspace.
	// Experimental.
	AddUnwantedRecommendations(extensions ...*string)
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

// The jsii proxy struct for VsCodeRecommendedExtensions
type jsiiProxy_VsCodeRecommendedExtensions struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_VsCodeRecommendedExtensions) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCodeRecommendedExtensions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VsCodeRecommendedExtensions) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVsCodeRecommendedExtensions(vscode VsCode) VsCodeRecommendedExtensions {
	_init_.Initialize()

	if err := validateNewVsCodeRecommendedExtensionsParameters(vscode); err != nil {
		panic(err)
	}
	j := jsiiProxy_VsCodeRecommendedExtensions{}

	_jsii_.Create(
		"projen.vscode.VsCodeRecommendedExtensions",
		[]interface{}{vscode},
		&j,
	)

	return &j
}

// Experimental.
func NewVsCodeRecommendedExtensions_Override(v VsCodeRecommendedExtensions, vscode VsCode) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.vscode.VsCodeRecommendedExtensions",
		[]interface{}{vscode},
		v,
	)
}

// Test whether the given construct is a component.
// Experimental.
func VsCodeRecommendedExtensions_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsCodeRecommendedExtensions_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.VsCodeRecommendedExtensions",
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
func VsCodeRecommendedExtensions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVsCodeRecommendedExtensions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.VsCodeRecommendedExtensions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) AddRecommendations(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addRecommendations",
		args,
	)
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) AddUnwantedRecommendations(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addUnwantedRecommendations",
		args,
	)
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

