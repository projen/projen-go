package vscode

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/tasks"
)

// Class interface
type DevContainerIface interface {
	projen.IDevEnvironmentIface
	GetProject() projen.ProjectIface
	GetConfig() interface{}
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddDockerImage(image projen.DevEnvironmentDockerImageIface)
	AddPorts(ports string)
	AddTasks(tasks tasks.TaskIface)
	AddVscodeExtensions(extensions string)
}

// A development environment running VSCode in a container;
// 
// used by GitHub
// codespaces.
// Experimental.
// Struct proxy
type DevContainer struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// Direct access to the devcontainer configuration (escape hatch).
	// Experimental.
	Config interface{} `json:"config"`
}

func (d *DevContainer) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		d,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevContainer) GetConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		d,
		"config",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewDevContainer(project projen.ProjectIface, options DevContainerOptionsIface) DevContainerIface {
	_init_.Initialize()
	self := DevContainer{}
	_jsii_.Create(
		"projen.vscode.DevContainer",
		[]interface{}{project, options},
		[]_jsii_.FQN{"projen.IDevEnvironment"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (d *DevContainer) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DevContainer) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DevContainer) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DevContainer) AddDockerImage(image projen.DevEnvironmentDockerImageIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addDockerImage",
		[]interface{}{image},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DevContainer) AddPorts(ports string) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addPorts",
		[]interface{}{ports},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DevContainer) AddTasks(tasks tasks.TaskIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addTasks",
		[]interface{}{tasks},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DevContainer) AddVscodeExtensions(extensions string) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addVscodeExtensions",
		[]interface{}{extensions},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// DevContainerOptionsIface is the public interface for the custom type DevContainerOptions
// Experimental.
type DevContainerOptionsIface interface {
	GetDockerImage() projen.DevEnvironmentDockerImageIface
	GetPorts() []string
	GetTasks() []tasks.TaskIface
	GetVscodeExtensions() []string
}

// Constructor options for the DevContainer component.
// 
// The default docker image used for GitHub Codespaces is defined here:
// See: https://github.com/microsoft/vscode-dev-containers/tree/master/containers/codespaces-linux
//
// Experimental.
// Struct proxy
type DevContainerOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage projen.DevEnvironmentDockerImageIface `json:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports []string `json:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks []tasks.TaskIface `json:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions []string `json:"vscodeExtensions"`
}

func (d *DevContainerOptions) GetDockerImage() projen.DevEnvironmentDockerImageIface {
	var returns projen.DevEnvironmentDockerImageIface
	_jsii_.Get(
		d,
		"dockerImage",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.DevEnvironmentDockerImageIface)(nil)).Elem(): reflect.TypeOf((*projen.DevEnvironmentDockerImage)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevContainerOptions) GetPorts() []string {
	var returns []string
	_jsii_.Get(
		d,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevContainerOptions) GetTasks() []tasks.TaskIface {
	var returns []tasks.TaskIface
	_jsii_.Get(
		d,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevContainerOptions) GetVscodeExtensions() []string {
	var returns []string
	_jsii_.Get(
		d,
		"vscodeExtensions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Controls the visibility of the VSCode Debug Console panel during a debugging session Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type InternalConsoleOptions string

const (
	InternalConsoleOptionsNeverOpen InternalConsoleOptions = "NEVER_OPEN"
	InternalConsoleOptionsOpenOnFirstSessionStart InternalConsoleOptions = "OPEN_ON_FIRST_SESSION_START"
	InternalConsoleOptionsOpenOnSessionStart InternalConsoleOptions = "OPEN_ON_SESSION_START"
)

// PresentationIface is the public interface for the custom type Presentation
// Experimental.
type PresentationIface interface {
	GetGroup() string
	GetHidden() bool
	GetOrder() float64
}

// VSCode launch configuration Presentation interface "using the order, group, and hidden attributes in the presentation object you can sort, group, and hide configurations and compounds in the Debug configuration dropdown and in the Debug quick pick." Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
// Struct proxy
type Presentation struct {
	// Experimental.
	Group string `json:"group"`
	// Experimental.
	Hidden bool `json:"hidden"`
	// Experimental.
	Order float64 `json:"order"`
}

func (p *Presentation) GetGroup() string {
	var returns string
	_jsii_.Get(
		p,
		"group",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Presentation) GetHidden() bool {
	var returns bool
	_jsii_.Get(
		p,
		"hidden",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Presentation) GetOrder() float64 {
	var returns float64
	_jsii_.Get(
		p,
		"order",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ServerReadyActionIface is the public interface for the custom type ServerReadyAction
// Experimental.
type ServerReadyActionIface interface {
	GetAction() string
	GetPattern() string
	GetUriFormat() string
}

// VSCode launch configuration ServerReadyAction interface "if you want to open a URL in a web browser whenever the program under debugging outputs a specific message to the debug console or integrated terminal." Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
// Struct proxy
type ServerReadyAction struct {
	// Experimental.
	Action string `json:"action"`
	// Experimental.
	Pattern string `json:"pattern"`
	// Experimental.
	UriFormat string `json:"uriFormat"`
}

func (s *ServerReadyAction) GetAction() string {
	var returns string
	_jsii_.Get(
		s,
		"action",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServerReadyAction) GetPattern() string {
	var returns string
	_jsii_.Get(
		s,
		"pattern",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServerReadyAction) GetUriFormat() string {
	var returns string
	_jsii_.Get(
		s,
		"uriFormat",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type VsCodeIface interface {
	GetProject() projen.ProjectIface
	GetLaunchConfiguration() VsCodeLaunchConfigIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Experimental.
// Struct proxy
type VsCode struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// Experimental.
	LaunchConfiguration VsCodeLaunchConfigIface `json:"launchConfiguration"`
}

func (v *VsCode) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		v,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCode) GetLaunchConfiguration() VsCodeLaunchConfigIface {
	var returns VsCodeLaunchConfigIface
	_jsii_.Get(
		v,
		"launchConfiguration",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VsCodeLaunchConfigIface)(nil)).Elem(): reflect.TypeOf((*VsCodeLaunchConfig)(nil)).Elem(),
		},
	)
	return returns
}


func NewVsCode(project projen.ProjectIface) VsCodeIface {
	_init_.Initialize()
	self := VsCode{}
	_jsii_.Create(
		"projen.vscode.VsCode",
		[]interface{}{project},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (v *VsCode) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *VsCode) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *VsCode) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type VsCodeLaunchConfigIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddConfiguration(cfg VsCodeLaunchConfigurationEntryIface)
}

// VSCode launch configuration file (launch.json), useful for enabling in-editor debugger.
// Experimental.
// Struct proxy
type VsCodeLaunchConfig struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (v *VsCodeLaunchConfig) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		v,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewVsCodeLaunchConfig(vscode VsCodeIface) VsCodeLaunchConfigIface {
	_init_.Initialize()
	self := VsCodeLaunchConfig{}
	_jsii_.Create(
		"projen.vscode.VsCodeLaunchConfig",
		[]interface{}{vscode},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (v *VsCodeLaunchConfig) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *VsCodeLaunchConfig) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *VsCodeLaunchConfig) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *VsCodeLaunchConfig) AddConfiguration(cfg VsCodeLaunchConfigurationEntryIface) {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"addConfiguration",
		[]interface{}{cfg},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// VsCodeLaunchConfigurationEntryIface is the public interface for the custom type VsCodeLaunchConfigurationEntry
// Experimental.
type VsCodeLaunchConfigurationEntryIface interface {
	GetName() string
	GetRequest() string
	GetType() string
	GetArgs() []string
	GetDebugServer() float64
	GetInternalConsoleOptions() InternalConsoleOptions
	GetOutFiles() []string
	GetPostDebugTask() string
	GetPreLaunchTask() string
	GetPresentation() PresentationIface
	GetProgram() string
	GetRuntimeArgs() []string
	GetServerReadyAction() ServerReadyActionIface
	GetSkipFiles() []string
	GetUrl() string
	GetWebRoot() string
}

// Options for a 'VsCodeLaunchConfigurationEntry' Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
// Struct proxy
type VsCodeLaunchConfigurationEntry struct {
	// Experimental.
	Name string `json:"name"`
	// Experimental.
	Request string `json:"request"`
	// Experimental.
	Type string `json:"type"`
	// Experimental.
	Args []string `json:"args"`
	// Experimental.
	DebugServer float64 `json:"debugServer"`
	// Experimental.
	InternalConsoleOptions InternalConsoleOptions `json:"internalConsoleOptions"`
	// Experimental.
	OutFiles []string `json:"outFiles"`
	// Experimental.
	PostDebugTask string `json:"postDebugTask"`
	// Experimental.
	PreLaunchTask string `json:"preLaunchTask"`
	// Experimental.
	Presentation PresentationIface `json:"presentation"`
	// Experimental.
	Program string `json:"program"`
	// Experimental.
	RuntimeArgs []string `json:"runtimeArgs"`
	// Experimental.
	ServerReadyAction ServerReadyActionIface `json:"serverReadyAction"`
	// Experimental.
	SkipFiles []string `json:"skipFiles"`
	// Experimental.
	Url string `json:"url"`
	// Experimental.
	WebRoot string `json:"webRoot"`
}

func (v *VsCodeLaunchConfigurationEntry) GetName() string {
	var returns string
	_jsii_.Get(
		v,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetRequest() string {
	var returns string
	_jsii_.Get(
		v,
		"request",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetType() string {
	var returns string
	_jsii_.Get(
		v,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetArgs() []string {
	var returns []string
	_jsii_.Get(
		v,
		"args",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetDebugServer() float64 {
	var returns float64
	_jsii_.Get(
		v,
		"debugServer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetInternalConsoleOptions() InternalConsoleOptions {
	var returns InternalConsoleOptions
	_jsii_.Get(
		v,
		"internalConsoleOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*InternalConsoleOptions)(nil)).Elem(): reflect.TypeOf((*InternalConsoleOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetOutFiles() []string {
	var returns []string
	_jsii_.Get(
		v,
		"outFiles",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetPostDebugTask() string {
	var returns string
	_jsii_.Get(
		v,
		"postDebugTask",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetPreLaunchTask() string {
	var returns string
	_jsii_.Get(
		v,
		"preLaunchTask",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetPresentation() PresentationIface {
	var returns PresentationIface
	_jsii_.Get(
		v,
		"presentation",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PresentationIface)(nil)).Elem(): reflect.TypeOf((*Presentation)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetProgram() string {
	var returns string
	_jsii_.Get(
		v,
		"program",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetRuntimeArgs() []string {
	var returns []string
	_jsii_.Get(
		v,
		"runtimeArgs",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetServerReadyAction() ServerReadyActionIface {
	var returns ServerReadyActionIface
	_jsii_.Get(
		v,
		"serverReadyAction",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServerReadyActionIface)(nil)).Elem(): reflect.TypeOf((*ServerReadyAction)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetSkipFiles() []string {
	var returns []string
	_jsii_.Get(
		v,
		"skipFiles",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetUrl() string {
	var returns string
	_jsii_.Get(
		v,
		"url",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VsCodeLaunchConfigurationEntry) GetWebRoot() string {
	var returns string
	_jsii_.Get(
		v,
		"webRoot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


