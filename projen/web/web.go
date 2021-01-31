package web

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/deps"
	"github.com/projen/projen-go/projen/tasks"
	"github.com/projen/projen-go/projen/vscode"
	"github.com/projen/projen-go/projen/github"
)

// Class interface
type NextComponentIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Experimental.
// Struct proxy
type NextComponent struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (n *NextComponent) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewNextComponent(project projen.NodeProjectIface, options NextComponentOptionsIface) NextComponentIface {
	_init_.Initialize()
	self := NextComponent{}
	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (n *NextComponent) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextComponent) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextComponent) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// NextComponentOptionsIface is the public interface for the custom type NextComponentOptions
// Experimental.
type NextComponentOptionsIface interface {
	GetTailwind() bool
	GetTypescript() bool
}

// Experimental.
// Struct proxy
type NextComponentOptions struct {
	// Setup Tailwind as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind bool `json:"tailwind"`
	// Whether to apply options specific for TypeScript Next.js projects.
	// Experimental.
	Typescript bool `json:"typescript"`
}

func (n *NextComponentOptions) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		n,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextComponentOptions) GetTypescript() bool {
	var returns bool
	_jsii_.Get(
		n,
		"typescript",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// NextJsCommonProjectOptionsIface is the public interface for the custom type NextJsCommonProjectOptions
// Experimental.
type NextJsCommonProjectOptionsIface interface {
	GetAssetsdir() string
	GetTailwind() bool
}

// Experimental.
// Struct proxy
type NextJsCommonProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir string `json:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind bool `json:"tailwind"`
}

func (n *NextJsCommonProjectOptions) GetAssetsdir() string {
	var returns string
	_jsii_.Get(
		n,
		"assetsdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsCommonProjectOptions) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		n,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type NextJsProjectIface interface {
	GetComponents() []projen.ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []projen.FileBaseIface
	GetGitignore() projen.IgnoreFileIface
	GetLogger() projen.LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() projen.ProjectType
	GetRoot() projen.ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() projen.GitpodIface
	GetJsiiFqn() string
	GetParent() projen.ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackage() projen.NodePackageIface
	GetPackageManager() projen.NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() projen.JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() projen.IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetAssetsdir() string
	GetSrcdir() string
	GetTailwind() bool
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) projen.FileBaseIface
	TryFindJsonFile(filePath string) projen.JsonFileIface
	TryFindObjectFile(filePath string) projen.ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// Next.js project without TypeScript.
// Experimental.
// Struct proxy
type NextJsProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []projen.ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []projen.FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore projen.IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger projen.LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root projen.ProjectIface `json:"root"`
	// Experimental.
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Experimental.
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Experimental.
	Gitpod projen.GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package projen.NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest projen.JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore projen.IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// The directory in which app assets reside.
	// Experimental.
	Assetsdir string `json:"assetsdir"`
	// The directory in which source files reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// Setup Tailwind as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind bool `json:"tailwind"`
}

func (n *NextJsProject) GetComponents() []projen.ComponentIface {
	var returns []projen.ComponentIface
	_jsii_.Get(
		n,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ComponentIface)(nil)).Elem(): reflect.TypeOf((*projen.Component)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetFiles() []projen.FileBaseIface {
	var returns []projen.FileBaseIface
	_jsii_.Get(
		n,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetGitignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		n,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetLogger() projen.LoggerIface {
	var returns projen.LoggerIface
	_jsii_.Get(
		n,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerIface)(nil)).Elem(): reflect.TypeOf((*projen.Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetName() string {
	var returns string
	_jsii_.Get(
		n,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		n,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		n,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetRoot() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		n,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		n,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		n,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetGitpod() projen.GitpodIface {
	var returns projen.GitpodIface
	_jsii_.Get(
		n,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.GitpodIface)(nil)).Elem(): reflect.TypeOf((*projen.Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		n,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		n,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		n,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetPackage() projen.NodePackageIface {
	var returns projen.NodePackageIface
	_jsii_.Get(
		n,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageIface)(nil)).Elem(): reflect.TypeOf((*projen.NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		n,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		n,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		n,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetJest() projen.JestIface {
	var returns projen.JestIface
	_jsii_.Get(
		n,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestIface)(nil)).Elem(): reflect.TypeOf((*projen.Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetNpmignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		n,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		n,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		n,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetAssetsdir() string {
	var returns string
	_jsii_.Get(
		n,
		"assetsdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		n,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		n,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewNextJsProject(options NextJsProjectOptionsIface) NextJsProjectIface {
	_init_.Initialize()
	self := NextJsProject{}
	_jsii_.Create(
		"projen.web.NextJsProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func NextJsProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.web.NextJsProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		n,
		"addTask",
		[]interface{}{name, props},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) TryFindFile(filePath string) projen.FileBaseIface {
	var returns projen.FileBaseIface
	_jsii_.Invoke(
		n,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) TryFindJsonFile(filePath string) projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Invoke(
		n,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) TryFindObjectFile(filePath string) projen.ObjectFileIface {
	var returns projen.ObjectFileIface
	_jsii_.Invoke(
		n,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*projen.ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		n,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// NextJsProjectOptionsIface is the public interface for the custom type NextJsProjectOptions
// Experimental.
type NextJsProjectOptionsIface interface {
	GetAssetsdir() string
	GetTailwind() bool
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() projen.LoggerOptionsIface
	GetOutdir() string
	GetParent() projen.ProjectIface
	GetProjectType() projen.ProjectType
	GetReadme() projen.SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() projen.NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackageManager() projen.NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() projen.PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() projen.JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() projen.SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetSampleCode() bool
	GetSrcdir() string
}

// Experimental.
// Struct proxy
type NextJsProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir string `json:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind bool `json:"tailwind"`
	// This is the name of your project.
	// Experimental.
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging projen.LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme projen.SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
	// Runtime dependencies of this module.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
	// Build dependencies for this module.
	// 
	// These dependencies will only be
	// available in your build environment but will not be fetched when this
	// module is consumed.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess projen.NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions projen.PeerDependencyOptionsIface `json:"peerDependencyOptions"`
	// Peer dependencies for this module.
	// 
	// Dependencies listed here are required to
	// be installed (and satisfied) by the _consumer_ of this library. Using peer
	// dependencies allows you to ensure that only a single module of a certain
	// library exists in the `node_modules` tree of your consumers.
	// 
	// Note that prior to npm@7, peer dependencies are _not_ automatically
	// installed, which means that adding peer dependencies to a library will be a
	// breaking change for your customers.
	// 
	// Unless `peerDependencyOptions.pinnedDevDependency` is disabled (it is
	// enabled by default), projen will automatically add a dev dependency with a
	// pinned version for each peer dependency. This will ensure that you build &
	// test your module against the lowest peer version required.
	// Experimental.
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions projen.JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	// 
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	// 
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	// 
	// To create a personal access token see https://github.com/settings/tokens
	// Experimental.
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion projen.SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Generate one-time sample in `pages/` and `public/` if there are no files there.
	// Experimental.
	SampleCode bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir string `json:"srcdir"`
}

func (n *NextJsProjectOptions) GetAssetsdir() string {
	var returns string
	_jsii_.Get(
		n,
		"assetsdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		n,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		n,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		n,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		n,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		n,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		n,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetLogging() projen.LoggerOptionsIface {
	var returns projen.LoggerOptionsIface
	_jsii_.Get(
		n,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		n,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		n,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetReadme() projen.SampleReadmePropsIface {
	var returns projen.SampleReadmePropsIface
	_jsii_.Get(
		n,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*projen.SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		n,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		n,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		n,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		n,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		n,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		n,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		n,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		n,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		n,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmAccess() projen.NpmAccess {
	var returns projen.NpmAccess
	_jsii_.Get(
		n,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(): reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		n,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetPeerDependencyOptions() projen.PeerDependencyOptionsIface {
	var returns projen.PeerDependencyOptionsIface
	_jsii_.Get(
		n,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		n,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		n,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		n,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		n,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		n,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		n,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		n,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		n,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		n,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		n,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		n,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		n,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		n,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetJestOptions() projen.JestOptionsIface {
	var returns projen.JestOptionsIface
	_jsii_.Get(
		n,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		n,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		n,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		n,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		n,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		n,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		n,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		n,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		n,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		n,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetProjenVersion() projen.SemverIface {
	var returns projen.SemverIface
	_jsii_.Get(
		n,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SemverIface)(nil)).Elem(): reflect.TypeOf((*projen.Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		n,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		n,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		n,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		n,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		n,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		n,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		n,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsProjectOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		n,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type NextJsTypeDefIface interface {
	GetProject() projen.ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_ projen.IResolverIface) string
}

// Experimental.
// Struct proxy
type NextJsTypeDef struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (n *NextJsTypeDef) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeDef) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		n,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDef) GetPath() string {
	var returns string
	_jsii_.Get(
		n,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDef) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		n,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDef) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		n,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewNextJsTypeDef(project NextJsTypeScriptProjectIface, filePath string, options NextJsTypeDefOptionsIface) NextJsTypeDefIface {
	_init_.Initialize()
	self := NextJsTypeDef{}
	_jsii_.Create(
		"projen.web.NextJsTypeDef",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (n *NextJsTypeDef) SetExecutable(val bool) {
	_jsii_.Set(
		n,
		"executable",
		val,
	)
}

func (n *NextJsTypeDef) SetReadonly(val bool) {
	_jsii_.Set(
		n,
		"readonly",
		val,
	)
}

func NextJsTypeDef_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.web.NextJsTypeDef",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDef) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeDef) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeDef) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeDef) SynthesizeContent(_ projen.IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		n,
		"synthesizeContent",
		[]interface{}{_},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// NextJsTypeDefOptionsIface is the public interface for the custom type NextJsTypeDefOptions
// Experimental.
type NextJsTypeDefOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
}

// Experimental.
// Struct proxy
type NextJsTypeDefOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (n *NextJsTypeDefOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		n,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDefOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		n,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDefOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		n,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeDefOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		n,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type NextJsTypeScriptProjectIface interface {
	GetComponents() []projen.ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []projen.FileBaseIface
	GetGitignore() projen.IgnoreFileIface
	GetLogger() projen.LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() projen.ProjectType
	GetRoot() projen.ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() projen.GitpodIface
	GetJsiiFqn() string
	GetParent() projen.ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackage() projen.NodePackageIface
	GetPackageManager() projen.NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() projen.JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() projen.IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() projen.EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() projen.TypescriptConfigIface
	GetAssetsdir() string
	GetNextJsTypeDef() NextJsTypeDefIface
	GetTailwind() bool
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) projen.FileBaseIface
	TryFindJsonFile(filePath string) projen.JsonFileIface
	TryFindObjectFile(filePath string) projen.ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// Next.js project with TypeScript.
// Experimental.
// Struct proxy
type NextJsTypeScriptProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []projen.ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []projen.FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore projen.IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger projen.LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root projen.ProjectIface `json:"root"`
	// Experimental.
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Experimental.
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Experimental.
	Gitpod projen.GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package projen.NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest projen.JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore projen.IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which source files reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint projen.EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig projen.TypescriptConfigIface `json:"tsconfig"`
	// The directory in which app assets reside.
	// Experimental.
	Assetsdir string `json:"assetsdir"`
	// TypeScript definition file included that ensures Next.js types are picked up by the TypeScript compiler.
	// See: https://nextjs.org/docs/basic-features/typescript
	//
	// Experimental.
	NextJsTypeDef NextJsTypeDefIface `json:"nextJsTypeDef"`
	// Setup Tailwind as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind bool `json:"tailwind"`
}

func (n *NextJsTypeScriptProject) GetComponents() []projen.ComponentIface {
	var returns []projen.ComponentIface
	_jsii_.Get(
		n,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ComponentIface)(nil)).Elem(): reflect.TypeOf((*projen.Component)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetFiles() []projen.FileBaseIface {
	var returns []projen.FileBaseIface
	_jsii_.Get(
		n,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetGitignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		n,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetLogger() projen.LoggerIface {
	var returns projen.LoggerIface
	_jsii_.Get(
		n,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerIface)(nil)).Elem(): reflect.TypeOf((*projen.Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetName() string {
	var returns string
	_jsii_.Get(
		n,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		n,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		n,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetRoot() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		n,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		n,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		n,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetGitpod() projen.GitpodIface {
	var returns projen.GitpodIface
	_jsii_.Get(
		n,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.GitpodIface)(nil)).Elem(): reflect.TypeOf((*projen.Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		n,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		n,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		n,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetPackage() projen.NodePackageIface {
	var returns projen.NodePackageIface
	_jsii_.Get(
		n,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageIface)(nil)).Elem(): reflect.TypeOf((*projen.NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		n,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		n,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		n,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetJest() projen.JestIface {
	var returns projen.JestIface
	_jsii_.Get(
		n,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestIface)(nil)).Elem(): reflect.TypeOf((*projen.Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetNpmignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		n,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		n,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		n,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		n,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetLibdir() string {
	var returns string
	_jsii_.Get(
		n,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		n,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetTestdir() string {
	var returns string
	_jsii_.Get(
		n,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		n,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetEslint() projen.EslintIface {
	var returns projen.EslintIface
	_jsii_.Get(
		n,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.EslintIface)(nil)).Elem(): reflect.TypeOf((*projen.Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetTsconfig() projen.TypescriptConfigIface {
	var returns projen.TypescriptConfigIface
	_jsii_.Get(
		n,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*projen.TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetAssetsdir() string {
	var returns string
	_jsii_.Get(
		n,
		"assetsdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetNextJsTypeDef() NextJsTypeDefIface {
	var returns NextJsTypeDefIface
	_jsii_.Get(
		n,
		"nextJsTypeDef",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NextJsTypeDefIface)(nil)).Elem(): reflect.TypeOf((*NextJsTypeDef)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		n,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewNextJsTypeScriptProject(options NextJsTypeScriptProjectOptionsIface) NextJsTypeScriptProjectIface {
	_init_.Initialize()
	self := NextJsTypeScriptProject{}
	_jsii_.Create(
		"projen.web.NextJsTypeScriptProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func NextJsTypeScriptProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.web.NextJsTypeScriptProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		n,
		"addTask",
		[]interface{}{name, props},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) TryFindFile(filePath string) projen.FileBaseIface {
	var returns projen.FileBaseIface
	_jsii_.Invoke(
		n,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) TryFindJsonFile(filePath string) projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Invoke(
		n,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) TryFindObjectFile(filePath string) projen.ObjectFileIface {
	var returns projen.ObjectFileIface
	_jsii_.Invoke(
		n,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*projen.ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NextJsTypeScriptProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		n,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// NextJsTypeScriptProjectOptionsIface is the public interface for the custom type NextJsTypeScriptProjectOptions
// Experimental.
type NextJsTypeScriptProjectOptionsIface interface {
	GetAssetsdir() string
	GetTailwind() bool
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() projen.LoggerOptionsIface
	GetOutdir() string
	GetParent() projen.ProjectIface
	GetProjectType() projen.ProjectType
	GetReadme() projen.SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() projen.NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackageManager() projen.NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() projen.PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() projen.JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() projen.SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetCompileBeforeTest() bool
	GetDisableTsconfig() bool
	GetDocgen() bool
	GetDocsDirectory() string
	GetEntrypointTypes() string
	GetEslint() bool
	GetEslintOptions() projen.EslintOptionsIface
	GetLibdir() string
	GetPackage() bool
	GetSampleCode() bool
	GetSrcdir() string
	GetTestdir() string
	GetTsconfig() projen.TypescriptConfigOptionsIface
	GetTypescriptVersion() string
}

// Experimental.
// Struct proxy
type NextJsTypeScriptProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir string `json:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind bool `json:"tailwind"`
	// This is the name of your project.
	// Experimental.
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging projen.LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme projen.SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
	// Runtime dependencies of this module.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
	// Build dependencies for this module.
	// 
	// These dependencies will only be
	// available in your build environment but will not be fetched when this
	// module is consumed.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess projen.NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions projen.PeerDependencyOptionsIface `json:"peerDependencyOptions"`
	// Peer dependencies for this module.
	// 
	// Dependencies listed here are required to
	// be installed (and satisfied) by the _consumer_ of this library. Using peer
	// dependencies allows you to ensure that only a single module of a certain
	// library exists in the `node_modules` tree of your consumers.
	// 
	// Note that prior to npm@7, peer dependencies are _not_ automatically
	// installed, which means that adding peer dependencies to a library will be a
	// breaking change for your customers.
	// 
	// Unless `peerDependencyOptions.pinnedDevDependency` is disabled (it is
	// enabled by default), projen will automatically add a dev dependency with a
	// pinned version for each peer dependency. This will ensure that you build &
	// test your module against the lowest peer version required.
	// Experimental.
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions projen.JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	// 
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	// 
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	// 
	// To create a personal access token see https://github.com/settings/tokens
	// Experimental.
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion projen.SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Compile the code before running tests.
	// Experimental.
	CompileBeforeTest bool `json:"compileBeforeTest"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes string `json:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions projen.EslintOptionsIface `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir string `json:"libdir"`
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Experimental.
	Package bool `json:"package"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	// 
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir string `json:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig projen.TypescriptConfigOptionsIface `json:"tsconfig"`
	// TypeScript version to use.
	// Experimental.
	TypescriptVersion string `json:"typescriptVersion"`
}

func (n *NextJsTypeScriptProjectOptions) GetAssetsdir() string {
	var returns string
	_jsii_.Get(
		n,
		"assetsdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		n,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		n,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		n,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		n,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		n,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		n,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetLogging() projen.LoggerOptionsIface {
	var returns projen.LoggerOptionsIface
	_jsii_.Get(
		n,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		n,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		n,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		n,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetReadme() projen.SampleReadmePropsIface {
	var returns projen.SampleReadmePropsIface
	_jsii_.Get(
		n,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*projen.SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		n,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		n,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		n,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		n,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		n,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		n,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		n,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		n,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		n,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmAccess() projen.NpmAccess {
	var returns projen.NpmAccess
	_jsii_.Get(
		n,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(): reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		n,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPeerDependencyOptions() projen.PeerDependencyOptionsIface {
	var returns projen.PeerDependencyOptionsIface
	_jsii_.Get(
		n,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		n,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		n,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		n,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		n,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		n,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		n,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		n,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		n,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		n,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		n,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		n,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		n,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		n,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetJestOptions() projen.JestOptionsIface {
	var returns projen.JestOptionsIface
	_jsii_.Get(
		n,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		n,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		n,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		n,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		n,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		n,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		n,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		n,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		n,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		n,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetProjenVersion() projen.SemverIface {
	var returns projen.SemverIface
	_jsii_.Get(
		n,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SemverIface)(nil)).Elem(): reflect.TypeOf((*projen.Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		n,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		n,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		n,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		n,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		n,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		n,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetCompileBeforeTest() bool {
	var returns bool
	_jsii_.Get(
		n,
		"compileBeforeTest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDisableTsconfig() bool {
	var returns bool
	_jsii_.Get(
		n,
		"disableTsconfig",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		n,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		n,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetEntrypointTypes() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypointTypes",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		n,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetEslintOptions() projen.EslintOptionsIface {
	var returns projen.EslintOptionsIface
	_jsii_.Get(
		n,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetLibdir() string {
	var returns string
	_jsii_.Get(
		n,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetPackage() bool {
	var returns bool
	_jsii_.Get(
		n,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		n,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		n,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetTestdir() string {
	var returns string
	_jsii_.Get(
		n,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetTsconfig() projen.TypescriptConfigOptionsIface {
	var returns projen.TypescriptConfigOptionsIface
	_jsii_.Get(
		n,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.TypescriptConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.TypescriptConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NextJsTypeScriptProjectOptions) GetTypescriptVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"typescriptVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type PostCssIface interface {
	GetFile() projen.JsonFileIface
	GetFileName() string
	GetTailwind() TailwindConfigIface
}

// Declares a PostCSS dependency with a default config file.
// Experimental.
// Struct proxy
type PostCss struct {
	// Experimental.
	File projen.JsonFileIface `json:"file"`
	// Experimental.
	FileName string `json:"fileName"`
	// Experimental.
	Tailwind TailwindConfigIface `json:"tailwind"`
}

func (p *PostCss) GetFile() projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Get(
		p,
		"file",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PostCss) GetFileName() string {
	var returns string
	_jsii_.Get(
		p,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PostCss) GetTailwind() TailwindConfigIface {
	var returns TailwindConfigIface
	_jsii_.Get(
		p,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TailwindConfigIface)(nil)).Elem(): reflect.TypeOf((*TailwindConfig)(nil)).Elem(),
		},
	)
	return returns
}


func NewPostCss(project projen.NodeProjectIface, options PostCssOptionsIface) PostCssIface {
	_init_.Initialize()
	self := PostCss{}
	_jsii_.Create(
		"projen.web.PostCss",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

// PostCssOptionsIface is the public interface for the custom type PostCssOptions
// Experimental.
type PostCssOptionsIface interface {
	GetFileName() string
	GetTailwind() bool
	GetTailwindOptions() TailwindConfigOptionsIface
}

// Experimental.
// Struct proxy
type PostCssOptions struct {
	// Experimental.
	FileName string `json:"fileName"`
	// Install Tailwind CSS as a PostCSS plugin.
	// Experimental.
	Tailwind bool `json:"tailwind"`
	// Tailwind CSS options.
	// Experimental.
	TailwindOptions TailwindConfigOptionsIface `json:"tailwindOptions"`
}

func (p *PostCssOptions) GetFileName() string {
	var returns string
	_jsii_.Get(
		p,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PostCssOptions) GetTailwind() bool {
	var returns bool
	_jsii_.Get(
		p,
		"tailwind",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PostCssOptions) GetTailwindOptions() TailwindConfigOptionsIface {
	var returns TailwindConfigOptionsIface
	_jsii_.Get(
		p,
		"tailwindOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TailwindConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*TailwindConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type ReactComponentIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Experimental.
// Struct proxy
type ReactComponent struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (r *ReactComponent) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewReactComponent(project projen.NodeProjectIface, options ReactComponentOptionsIface) ReactComponentIface {
	_init_.Initialize()
	self := ReactComponent{}
	_jsii_.Create(
		"projen.web.ReactComponent",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (r *ReactComponent) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactComponent) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactComponent) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ReactComponentOptionsIface is the public interface for the custom type ReactComponentOptions
// Experimental.
type ReactComponentOptionsIface interface {
	GetTypescript() bool
}

// Experimental.
// Struct proxy
type ReactComponentOptions struct {
	// Whether to apply options specific for TypeScript React projects.
	// Experimental.
	Typescript bool `json:"typescript"`
}

func (r *ReactComponentOptions) GetTypescript() bool {
	var returns bool
	_jsii_.Get(
		r,
		"typescript",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ReactProjectIface interface {
	GetComponents() []projen.ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []projen.FileBaseIface
	GetGitignore() projen.IgnoreFileIface
	GetLogger() projen.LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() projen.ProjectType
	GetRoot() projen.ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() projen.GitpodIface
	GetJsiiFqn() string
	GetParent() projen.ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackage() projen.NodePackageIface
	GetPackageManager() projen.NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() projen.JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() projen.IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetSrcdir() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) projen.FileBaseIface
	TryFindJsonFile(filePath string) projen.JsonFileIface
	TryFindObjectFile(filePath string) projen.ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// React project without TypeScript.
// Experimental.
// Struct proxy
type ReactProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []projen.ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []projen.FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore projen.IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger projen.LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root projen.ProjectIface `json:"root"`
	// Experimental.
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Experimental.
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Experimental.
	Gitpod projen.GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package projen.NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest projen.JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore projen.IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// The directory in which source files reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
}

func (r *ReactProject) GetComponents() []projen.ComponentIface {
	var returns []projen.ComponentIface
	_jsii_.Get(
		r,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ComponentIface)(nil)).Elem(): reflect.TypeOf((*projen.Component)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		r,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetFiles() []projen.FileBaseIface {
	var returns []projen.FileBaseIface
	_jsii_.Get(
		r,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetGitignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		r,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetLogger() projen.LoggerIface {
	var returns projen.LoggerIface
	_jsii_.Get(
		r,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerIface)(nil)).Elem(): reflect.TypeOf((*projen.Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetName() string {
	var returns string
	_jsii_.Get(
		r,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		r,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		r,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetRoot() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		r,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		r,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		r,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetGitpod() projen.GitpodIface {
	var returns projen.GitpodIface
	_jsii_.Get(
		r,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.GitpodIface)(nil)).Elem(): reflect.TypeOf((*projen.Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		r,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		r,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		r,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		r,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		r,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		r,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		r,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		r,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		r,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		r,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetPackage() projen.NodePackageIface {
	var returns projen.NodePackageIface
	_jsii_.Get(
		r,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageIface)(nil)).Elem(): reflect.TypeOf((*projen.NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		r,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		r,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		r,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		r,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetJest() projen.JestIface {
	var returns projen.JestIface
	_jsii_.Get(
		r,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestIface)(nil)).Elem(): reflect.TypeOf((*projen.Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetNpmignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		r,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		r,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		r,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		r,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewReactProject(options ReactProjectOptionsIface) ReactProjectIface {
	_init_.Initialize()
	self := ReactProject{}
	_jsii_.Create(
		"projen.web.ReactProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ReactProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.web.ReactProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		r,
		"addTask",
		[]interface{}{name, props},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) TryFindFile(filePath string) projen.FileBaseIface {
	var returns projen.FileBaseIface
	_jsii_.Invoke(
		r,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) TryFindJsonFile(filePath string) projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Invoke(
		r,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) TryFindObjectFile(filePath string) projen.ObjectFileIface {
	var returns projen.ObjectFileIface
	_jsii_.Invoke(
		r,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*projen.ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		r,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		r,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ReactProjectOptionsIface is the public interface for the custom type ReactProjectOptions
// Experimental.
type ReactProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() projen.LoggerOptionsIface
	GetOutdir() string
	GetParent() projen.ProjectIface
	GetProjectType() projen.ProjectType
	GetReadme() projen.SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() projen.NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackageManager() projen.NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() projen.PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() projen.JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() projen.SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetSampleCode() bool
	GetSrcdir() string
}

// Experimental.
// Struct proxy
type ReactProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging projen.LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme projen.SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
	// Runtime dependencies of this module.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
	// Build dependencies for this module.
	// 
	// These dependencies will only be
	// available in your build environment but will not be fetched when this
	// module is consumed.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess projen.NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions projen.PeerDependencyOptionsIface `json:"peerDependencyOptions"`
	// Peer dependencies for this module.
	// 
	// Dependencies listed here are required to
	// be installed (and satisfied) by the _consumer_ of this library. Using peer
	// dependencies allows you to ensure that only a single module of a certain
	// library exists in the `node_modules` tree of your consumers.
	// 
	// Note that prior to npm@7, peer dependencies are _not_ automatically
	// installed, which means that adding peer dependencies to a library will be a
	// breaking change for your customers.
	// 
	// Unless `peerDependencyOptions.pinnedDevDependency` is disabled (it is
	// enabled by default), projen will automatically add a dev dependency with a
	// pinned version for each peer dependency. This will ensure that you build &
	// test your module against the lowest peer version required.
	// Experimental.
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions projen.JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	// 
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	// 
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	// 
	// To create a personal access token see https://github.com/settings/tokens
	// Experimental.
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion projen.SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Generate one-time sample in `src/` and `public/` if there are no files there.
	// Experimental.
	SampleCode bool `json:"sampleCode"`
	// Source directory.
	// Experimental.
	Srcdir string `json:"srcdir"`
}

func (r *ReactProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		r,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		r,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		r,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		r,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		r,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetLogging() projen.LoggerOptionsIface {
	var returns projen.LoggerOptionsIface
	_jsii_.Get(
		r,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		r,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		r,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetReadme() projen.SampleReadmePropsIface {
	var returns projen.SampleReadmePropsIface
	_jsii_.Get(
		r,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*projen.SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		r,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		r,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		r,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		r,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		r,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		r,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		r,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		r,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		r,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		r,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		r,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		r,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		r,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmAccess() projen.NpmAccess {
	var returns projen.NpmAccess
	_jsii_.Get(
		r,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(): reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		r,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		r,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		r,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		r,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		r,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		r,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetPeerDependencyOptions() projen.PeerDependencyOptionsIface {
	var returns projen.PeerDependencyOptionsIface
	_jsii_.Get(
		r,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		r,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		r,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		r,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		r,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		r,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		r,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		r,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		r,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		r,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		r,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		r,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		r,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		r,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		r,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		r,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetJestOptions() projen.JestOptionsIface {
	var returns projen.JestOptionsIface
	_jsii_.Get(
		r,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		r,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		r,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		r,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		r,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		r,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		r,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		r,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		r,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		r,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetProjenVersion() projen.SemverIface {
	var returns projen.SemverIface
	_jsii_.Get(
		r,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SemverIface)(nil)).Elem(): reflect.TypeOf((*projen.Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		r,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		r,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		r,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		r,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		r,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		r,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		r,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		r,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		r,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		r,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		r,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactProjectOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		r,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ReactTypeDefIface interface {
	GetProject() projen.ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_ projen.IResolverIface) string
}

// Experimental.
// Struct proxy
type ReactTypeDef struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (r *ReactTypeDef) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeDef) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		r,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDef) GetPath() string {
	var returns string
	_jsii_.Get(
		r,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDef) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		r,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDef) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		r,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewReactTypeDef(project ReactTypeScriptProjectIface, filePath string, options ReactTypeDefOptionsIface) ReactTypeDefIface {
	_init_.Initialize()
	self := ReactTypeDef{}
	_jsii_.Create(
		"projen.web.ReactTypeDef",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (r *ReactTypeDef) SetExecutable(val bool) {
	_jsii_.Set(
		r,
		"executable",
		val,
	)
}

func (r *ReactTypeDef) SetReadonly(val bool) {
	_jsii_.Set(
		r,
		"readonly",
		val,
	)
}

func ReactTypeDef_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.web.ReactTypeDef",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDef) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeDef) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeDef) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeDef) SynthesizeContent(_ projen.IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		r,
		"synthesizeContent",
		[]interface{}{_},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// ReactTypeDefOptionsIface is the public interface for the custom type ReactTypeDefOptions
// Experimental.
type ReactTypeDefOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
}

// Experimental.
// Struct proxy
type ReactTypeDefOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (r *ReactTypeDefOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		r,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDefOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		r,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDefOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		r,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeDefOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		r,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ReactTypeScriptProjectIface interface {
	GetComponents() []projen.ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []projen.FileBaseIface
	GetGitignore() projen.IgnoreFileIface
	GetLogger() projen.LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() projen.ProjectType
	GetRoot() projen.ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() projen.GitpodIface
	GetJsiiFqn() string
	GetParent() projen.ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackage() projen.NodePackageIface
	GetPackageManager() projen.NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() projen.JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() projen.IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() projen.EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() projen.TypescriptConfigIface
	GetReactTypeDef() ReactTypeDefIface
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) projen.FileBaseIface
	TryFindJsonFile(filePath string) projen.JsonFileIface
	TryFindObjectFile(filePath string) projen.ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// React project with TypeScript.
// Experimental.
// Struct proxy
type ReactTypeScriptProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []projen.ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []projen.FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore projen.IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger projen.LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root projen.ProjectIface `json:"root"`
	// Experimental.
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Experimental.
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Experimental.
	Gitpod projen.GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package projen.NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest projen.JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore projen.IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which source files reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint projen.EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig projen.TypescriptConfigIface `json:"tsconfig"`
	// TypeScript definition file included that ensures React types are picked up by the TypeScript compiler.
	// Experimental.
	ReactTypeDef ReactTypeDefIface `json:"reactTypeDef"`
}

func (r *ReactTypeScriptProject) GetComponents() []projen.ComponentIface {
	var returns []projen.ComponentIface
	_jsii_.Get(
		r,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ComponentIface)(nil)).Elem(): reflect.TypeOf((*projen.Component)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		r,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetFiles() []projen.FileBaseIface {
	var returns []projen.FileBaseIface
	_jsii_.Get(
		r,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetGitignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		r,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetLogger() projen.LoggerIface {
	var returns projen.LoggerIface
	_jsii_.Get(
		r,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerIface)(nil)).Elem(): reflect.TypeOf((*projen.Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetName() string {
	var returns string
	_jsii_.Get(
		r,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		r,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		r,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetRoot() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		r,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		r,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		r,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetGitpod() projen.GitpodIface {
	var returns projen.GitpodIface
	_jsii_.Get(
		r,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.GitpodIface)(nil)).Elem(): reflect.TypeOf((*projen.Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		r,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		r,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		r,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		r,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		r,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		r,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		r,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		r,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		r,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		r,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetPackage() projen.NodePackageIface {
	var returns projen.NodePackageIface
	_jsii_.Get(
		r,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageIface)(nil)).Elem(): reflect.TypeOf((*projen.NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		r,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		r,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		r,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		r,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetJest() projen.JestIface {
	var returns projen.JestIface
	_jsii_.Get(
		r,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestIface)(nil)).Elem(): reflect.TypeOf((*projen.Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetNpmignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		r,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		r,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		r,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		r,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetLibdir() string {
	var returns string
	_jsii_.Get(
		r,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		r,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetTestdir() string {
	var returns string
	_jsii_.Get(
		r,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		r,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetEslint() projen.EslintIface {
	var returns projen.EslintIface
	_jsii_.Get(
		r,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.EslintIface)(nil)).Elem(): reflect.TypeOf((*projen.Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		r,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetTsconfig() projen.TypescriptConfigIface {
	var returns projen.TypescriptConfigIface
	_jsii_.Get(
		r,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*projen.TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) GetReactTypeDef() ReactTypeDefIface {
	var returns ReactTypeDefIface
	_jsii_.Get(
		r,
		"reactTypeDef",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ReactTypeDefIface)(nil)).Elem(): reflect.TypeOf((*ReactTypeDef)(nil)).Elem(),
		},
	)
	return returns
}


func NewReactTypeScriptProject(options ReactTypeScriptProjectOptionsIface) ReactTypeScriptProjectIface {
	_init_.Initialize()
	self := ReactTypeScriptProject{}
	_jsii_.Create(
		"projen.web.ReactTypeScriptProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ReactTypeScriptProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.web.ReactTypeScriptProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		r,
		"addTask",
		[]interface{}{name, props},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) TryFindFile(filePath string) projen.FileBaseIface {
	var returns projen.FileBaseIface
	_jsii_.Invoke(
		r,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) TryFindJsonFile(filePath string) projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Invoke(
		r,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) TryFindObjectFile(filePath string) projen.ObjectFileIface {
	var returns projen.ObjectFileIface
	_jsii_.Invoke(
		r,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*projen.ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		r,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *ReactTypeScriptProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		r,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ReactTypeScriptProjectOptionsIface is the public interface for the custom type ReactTypeScriptProjectOptions
// Experimental.
type ReactTypeScriptProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() projen.LoggerOptionsIface
	GetOutdir() string
	GetParent() projen.ProjectIface
	GetProjectType() projen.ProjectType
	GetReadme() projen.SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() projen.NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() projen.NpmTaskExecution
	GetPackageManager() projen.NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() projen.PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() projen.JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() projen.SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetCompileBeforeTest() bool
	GetDisableTsconfig() bool
	GetDocgen() bool
	GetDocsDirectory() string
	GetEntrypointTypes() string
	GetEslint() bool
	GetEslintOptions() projen.EslintOptionsIface
	GetLibdir() string
	GetPackage() bool
	GetSampleCode() bool
	GetSrcdir() string
	GetTestdir() string
	GetTsconfig() projen.TypescriptConfigOptionsIface
	GetTypescriptVersion() string
}

// Experimental.
// Struct proxy
type ReactTypeScriptProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging projen.LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme projen.SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
	// Runtime dependencies of this module.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
	// Build dependencies for this module.
	// 
	// These dependencies will only be
	// available in your build environment but will not be fetched when this
	// module is consumed.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess projen.NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution projen.NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager projen.NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions projen.PeerDependencyOptionsIface `json:"peerDependencyOptions"`
	// Peer dependencies for this module.
	// 
	// Dependencies listed here are required to
	// be installed (and satisfied) by the _consumer_ of this library. Using peer
	// dependencies allows you to ensure that only a single module of a certain
	// library exists in the `node_modules` tree of your consumers.
	// 
	// Note that prior to npm@7, peer dependencies are _not_ automatically
	// installed, which means that adding peer dependencies to a library will be a
	// breaking change for your customers.
	// 
	// Unless `peerDependencyOptions.pinnedDevDependency` is disabled (it is
	// enabled by default), projen will automatically add a dev dependency with a
	// pinned version for each peer dependency. This will ensure that you build &
	// test your module against the lowest peer version required.
	// Experimental.
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions projen.JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	// 
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	// 
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	// 
	// To create a personal access token see https://github.com/settings/tokens
	// Experimental.
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion projen.SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Compile the code before running tests.
	// Experimental.
	CompileBeforeTest bool `json:"compileBeforeTest"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes string `json:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions projen.EslintOptionsIface `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir string `json:"libdir"`
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Experimental.
	Package bool `json:"package"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	// 
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir string `json:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig projen.TypescriptConfigOptionsIface `json:"tsconfig"`
	// TypeScript version to use.
	// Experimental.
	TypescriptVersion string `json:"typescriptVersion"`
}

func (r *ReactTypeScriptProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		r,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		r,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		r,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		r,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		r,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetLogging() projen.LoggerOptionsIface {
	var returns projen.LoggerOptionsIface
	_jsii_.Get(
		r,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		r,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		r,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		r,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetReadme() projen.SampleReadmePropsIface {
	var returns projen.SampleReadmePropsIface
	_jsii_.Get(
		r,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*projen.SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		r,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		r,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		r,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		r,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		r,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		r,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		r,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		r,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		r,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		r,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		r,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		r,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		r,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmAccess() projen.NpmAccess {
	var returns projen.NpmAccess
	_jsii_.Get(
		r,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(): reflect.TypeOf((*projen.NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		r,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		r,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		r,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmTaskExecution() projen.NpmTaskExecution {
	var returns projen.NpmTaskExecution
	_jsii_.Get(
		r,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*projen.NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPackageManager() projen.NodePackageManager {
	var returns projen.NodePackageManager
	_jsii_.Get(
		r,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(): reflect.TypeOf((*projen.NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		r,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPeerDependencyOptions() projen.PeerDependencyOptionsIface {
	var returns projen.PeerDependencyOptionsIface
	_jsii_.Get(
		r,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		r,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		r,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		r,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		r,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		r,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		r,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		r,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		r,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		r,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		r,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		r,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		r,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		r,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		r,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		r,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		r,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetJestOptions() projen.JestOptionsIface {
	var returns projen.JestOptionsIface
	_jsii_.Get(
		r,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		r,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		r,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		r,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		r,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		r,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		r,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		r,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		r,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		r,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetProjenVersion() projen.SemverIface {
	var returns projen.SemverIface
	_jsii_.Get(
		r,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SemverIface)(nil)).Elem(): reflect.TypeOf((*projen.Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		r,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		r,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		r,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		r,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		r,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		r,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		r,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		r,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		r,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		r,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		r,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetCompileBeforeTest() bool {
	var returns bool
	_jsii_.Get(
		r,
		"compileBeforeTest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDisableTsconfig() bool {
	var returns bool
	_jsii_.Get(
		r,
		"disableTsconfig",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		r,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		r,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetEntrypointTypes() string {
	var returns string
	_jsii_.Get(
		r,
		"entrypointTypes",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		r,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetEslintOptions() projen.EslintOptionsIface {
	var returns projen.EslintOptionsIface
	_jsii_.Get(
		r,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetLibdir() string {
	var returns string
	_jsii_.Get(
		r,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetPackage() bool {
	var returns bool
	_jsii_.Get(
		r,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		r,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		r,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetTestdir() string {
	var returns string
	_jsii_.Get(
		r,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetTsconfig() projen.TypescriptConfigOptionsIface {
	var returns projen.TypescriptConfigOptionsIface
	_jsii_.Get(
		r,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.TypescriptConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.TypescriptConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (r *ReactTypeScriptProjectOptions) GetTypescriptVersion() string {
	var returns string
	_jsii_.Get(
		r,
		"typescriptVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type TailwindConfigIface interface {
	GetFile() projen.JsonFileIface
	GetFileName() string
}

// Declares a Tailwind CSS configuration file.
// 
// There are multiple ways to add Tailwind CSS in your node project - see:
// https://tailwindcss.com/docs/installation
// See: PostCss
//
// Experimental.
// Struct proxy
type TailwindConfig struct {
	// Experimental.
	File projen.JsonFileIface `json:"file"`
	// Experimental.
	FileName string `json:"fileName"`
}

func (t *TailwindConfig) GetFile() projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Get(
		t,
		"file",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TailwindConfig) GetFileName() string {
	var returns string
	_jsii_.Get(
		t,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewTailwindConfig(project projen.NodeProjectIface, options TailwindConfigOptionsIface) TailwindConfigIface {
	_init_.Initialize()
	self := TailwindConfig{}
	_jsii_.Create(
		"projen.web.TailwindConfig",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

// TailwindConfigOptionsIface is the public interface for the custom type TailwindConfigOptions
// Experimental.
type TailwindConfigOptionsIface interface {
	GetFileName() string
}

// Experimental.
// Struct proxy
type TailwindConfigOptions struct {
	// Experimental.
	FileName string `json:"fileName"`
}

func (t *TailwindConfigOptions) GetFileName() string {
	var returns string
	_jsii_.Get(
		t,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


