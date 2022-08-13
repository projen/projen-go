package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// Experimental.
type VsCode interface {
	projen.Component
	// Experimental.
	LaunchConfiguration() VsCodeLaunchConfig
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
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for VsCode
type jsiiProxy_VsCode struct {
	internal.Type__projenComponent
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

func (j *jsiiProxy_VsCode) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVsCode(project projen.Project) VsCode {
	_init_.Initialize()

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

