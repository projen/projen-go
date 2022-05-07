package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// A development environment running VSCode in a container;
//
// used by GitHub
// codespaces.
// Experimental.
type DevContainer interface {
	projen.Component
	projen.IDevEnvironment
	// Direct access to the devcontainer configuration (escape hatch).
	// Experimental.
	Config() interface{}
	// Experimental.
	Project() projen.Project
	// Add a custom Docker image or Dockerfile for the container.
	// Experimental.
	AddDockerImage(image projen.DevEnvironmentDockerImage)
	// Adds ports that should be exposed (forwarded) from the container.
	// Experimental.
	AddPorts(ports ...*string)
	// Adds tasks to run when the container starts.
	//
	// Tasks will be run in sequence.
	// Experimental.
	AddTasks(tasks ...projen.Task)
	// Adds a list of VSCode extensions that should be automatically installed in the container.
	// Experimental.
	AddVscodeExtensions(extensions ...*string)
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

// The jsii proxy struct for DevContainer
type jsiiProxy_DevContainer struct {
	internal.Type__projenComponent
	internal.Type__projenIDevEnvironment
}

func (j *jsiiProxy_DevContainer) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevContainer) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDevContainer(project projen.Project, options *DevContainerOptions) DevContainer {
	_init_.Initialize()

	j := jsiiProxy_DevContainer{}

	_jsii_.Create(
		"projen.vscode.DevContainer",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewDevContainer_Override(d DevContainer, project projen.Project, options *DevContainerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.vscode.DevContainer",
		[]interface{}{project, options},
		d,
	)
}

func (d *jsiiProxy_DevContainer) AddDockerImage(image projen.DevEnvironmentDockerImage) {
	_jsii_.InvokeVoid(
		d,
		"addDockerImage",
		[]interface{}{image},
	)
}

func (d *jsiiProxy_DevContainer) AddPorts(ports ...*string) {
	args := []interface{}{}
	for _, a := range ports {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addPorts",
		args,
	)
}

func (d *jsiiProxy_DevContainer) AddTasks(tasks ...projen.Task) {
	args := []interface{}{}
	for _, a := range tasks {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addTasks",
		args,
	)
}

func (d *jsiiProxy_DevContainer) AddVscodeExtensions(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addVscodeExtensions",
		args,
	)
}

func (d *jsiiProxy_DevContainer) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevContainer) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevContainer) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

// Constructor options for the DevContainer component.
//
// The default docker image used for GitHub Codespaces is defined here:.
// See: https://github.com/microsoft/vscode-dev-containers/tree/master/containers/codespaces-linux
//
// Experimental.
type DevContainerOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage projen.DevEnvironmentDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks *[]projen.Task `field:"optional" json:"tasks" yaml:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions *[]*string `field:"optional" json:"vscodeExtensions" yaml:"vscodeExtensions"`
}

// Controls the visibility of the VSCode Debug Console panel during a debugging session Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type InternalConsoleOptions string

const (
	// Experimental.
	InternalConsoleOptions_NEVER_OPEN InternalConsoleOptions = "NEVER_OPEN"
	// Experimental.
	InternalConsoleOptions_OPEN_ON_FIRST_SESSION_START InternalConsoleOptions = "OPEN_ON_FIRST_SESSION_START"
	// Experimental.
	InternalConsoleOptions_OPEN_ON_SESSION_START InternalConsoleOptions = "OPEN_ON_SESSION_START"
)

// VSCode launch configuration Presentation interface "using the order, group, and hidden attributes in the presentation object you can sort, group, and hide configurations and compounds in the Debug configuration dropdown and in the Debug quick pick." Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type Presentation struct {
	// Experimental.
	Group *string `field:"required" json:"group" yaml:"group"`
	// Experimental.
	Hidden *bool `field:"required" json:"hidden" yaml:"hidden"`
	// Experimental.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

// VSCode launch configuration ServerReadyAction interface "if you want to open a URL in a web browser whenever the program under debugging outputs a specific message to the debug console or integrated terminal." Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type ServerReadyAction struct {
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Experimental.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Experimental.
	UriFormat *string `field:"optional" json:"uriFormat" yaml:"uriFormat"`
}

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

// Options for a 'VsCodeLaunchConfigurationEntry' Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type VsCodeLaunchConfigurationEntry struct {
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	Request *string `field:"required" json:"request" yaml:"request"`
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Experimental.
	DebugServer *float64 `field:"optional" json:"debugServer" yaml:"debugServer"`
	// Experimental.
	InternalConsoleOptions InternalConsoleOptions `field:"optional" json:"internalConsoleOptions" yaml:"internalConsoleOptions"`
	// Experimental.
	OutFiles *[]*string `field:"optional" json:"outFiles" yaml:"outFiles"`
	// Experimental.
	PostDebugTask *string `field:"optional" json:"postDebugTask" yaml:"postDebugTask"`
	// Experimental.
	PreLaunchTask *string `field:"optional" json:"preLaunchTask" yaml:"preLaunchTask"`
	// Experimental.
	Presentation *Presentation `field:"optional" json:"presentation" yaml:"presentation"`
	// Experimental.
	Program *string `field:"optional" json:"program" yaml:"program"`
	// Experimental.
	RuntimeArgs *[]*string `field:"optional" json:"runtimeArgs" yaml:"runtimeArgs"`
	// Experimental.
	ServerReadyAction *ServerReadyAction `field:"optional" json:"serverReadyAction" yaml:"serverReadyAction"`
	// Experimental.
	SkipFiles *[]*string `field:"optional" json:"skipFiles" yaml:"skipFiles"`
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Experimental.
	WebRoot *string `field:"optional" json:"webRoot" yaml:"webRoot"`
}

