package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// VSCode launch configuration file (launch.json), useful for enabling in-editor debugger.
// Experimental.
type VsCodeLaunchConfig interface {
	projen.Component
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
}

// The jsii proxy struct for VsCodeLaunchConfig
type jsiiProxy_VsCodeLaunchConfig struct {
	internal.Type__projenComponent
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

