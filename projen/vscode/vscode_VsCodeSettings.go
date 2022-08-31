package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// VS Code Workspace settings Source: https://code.visualstudio.com/docs/getstarted/settings#_workspace-settings.
// Experimental.
type VsCodeSettings interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// Adds a workspace setting.
	// Experimental.
	AddSetting(setting *string, value interface{}, language *string)
	// Adds a workspace setting.
	// Experimental.
	AddSettings(settings *map[string]interface{}, languages interface{})
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

// The jsii proxy struct for VsCodeSettings
type jsiiProxy_VsCodeSettings struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_VsCodeSettings) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVsCodeSettings(vscode VsCode) VsCodeSettings {
	_init_.Initialize()

	if err := validateNewVsCodeSettingsParameters(vscode); err != nil {
		panic(err)
	}
	j := jsiiProxy_VsCodeSettings{}

	_jsii_.Create(
		"projen.vscode.VsCodeSettings",
		[]interface{}{vscode},
		&j,
	)

	return &j
}

// Experimental.
func NewVsCodeSettings_Override(v VsCodeSettings, vscode VsCode) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.vscode.VsCodeSettings",
		[]interface{}{vscode},
		v,
	)
}

func (v *jsiiProxy_VsCodeSettings) AddSetting(setting *string, value interface{}, language *string) {
	if err := v.validateAddSettingParameters(setting, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addSetting",
		[]interface{}{setting, value, language},
	)
}

func (v *jsiiProxy_VsCodeSettings) AddSettings(settings *map[string]interface{}, languages interface{}) {
	if err := v.validateAddSettingsParameters(settings, languages); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addSettings",
		[]interface{}{settings, languages},
	)
}

func (v *jsiiProxy_VsCodeSettings) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeSettings) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VsCodeSettings) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

