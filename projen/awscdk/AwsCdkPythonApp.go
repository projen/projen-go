package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/python"
	"github.com/projen/projen-go/projen/vscode"
)

// AWS CDK app in Python.
// Experimental.
type AwsCdkPythonApp interface {
	python.PythonProject
	// The CDK app entrypoint.
	// Experimental.
	AppEntrypoint() *string
	// Auto approve set up for this project.
	// Deprecated.
	AutoApprove() github.AutoApprove
	// Experimental.
	BuildTask() projen.Task
	// cdk.json configuration.
	// Experimental.
	CdkConfig() CdkConfig
	// Experimental.
	CdkDeps() AwsCdkDeps
	// Common CDK tasks.
	// Experimental.
	CdkTasks() CdkTasks
	// The CDK version this app is using.
	// Experimental.
	CdkVersion() *string
	// Whether to commit the managed files by default.
	// Experimental.
	CommitGenerated() *bool
	// Experimental.
	CompileTask() projen.Task
	// Returns all the components within this project.
	// Experimental.
	Components() *[]projen.Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Experimental.
	DefaultTask() projen.Task
	// Project dependencies.
	// Experimental.
	Deps() projen.Dependencies
	// API for managing dependencies.
	// Experimental.
	DepsManager() python.IPythonDeps
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated.
	DevContainer() vscode.DevContainer
	// Whether or not the project is being ejected.
	// Experimental.
	Ejected() *bool
	// API for mangaging the Python runtime environment.
	// Experimental.
	EnvManager() python.IPythonEnv
	// All files in this project.
	// Experimental.
	Files() *[]projen.FileBase
	// The .gitattributes file for this repository.
	// Experimental.
	Gitattributes() projen.GitAttributesFile
	// Access all github components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Github() github.GitHub
	// .gitignore.
	// Experimental.
	Gitignore() projen.IgnoreFile
	// Access for Gitpod.
	//
	// This will be `undefined` if gitpod boolean is false.
	// Deprecated.
	Gitpod() projen.Gitpod
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Experimental.
	InitProject() *projen.InitProject
	// Logging utilities.
	// Experimental.
	Logger() projen.Logger
	// Python module name (the project name, with any hyphens or periods replaced with underscores).
	// Experimental.
	ModuleName() *string
	// Project name.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Absolute output directory of this project.
	// Experimental.
	Outdir() *string
	// Experimental.
	PackageTask() projen.Task
	// API for managing packaging the project as a library.
	//
	// Only applies when the `projectType` is LIB.
	// Experimental.
	PackagingManager() python.IPythonPackaging
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() projen.Project
	// Experimental.
	PostCompileTask() projen.Task
	// Experimental.
	PreCompileTask() projen.Task
	// Manages the build process of the project.
	// Experimental.
	ProjectBuild() projen.ProjectBuild
	// Deprecated.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand() *string
	// Pytest component.
	// Experimental.
	Pytest() python.Pytest
	// Experimental.
	SetPytest(val python.Pytest)
	// The root project.
	// Experimental.
	Root() projen.Project
	// The directory in which the python sample tests reside.
	// Experimental.
	SampleTestdir() *string
	// Returns all the subprojects within this project.
	// Experimental.
	Subprojects() *[]projen.Project
	// Project tasks.
	// Experimental.
	Tasks() projen.Tasks
	// The directory in which the python tests reside.
	// Deprecated: Use `sampleTestdir` instead.
	Testdir() *string
	// Experimental.
	TestTask() projen.Task
	// Version of the package for distribution (should follow semver).
	// Experimental.
	Version() *string
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Vscode() vscode.VsCode
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Experimental.
	AddExcludeFromCleanup(globs ...*string)
	// Adds a .gitignore pattern.
	// Experimental.
	AddGitIgnore(pattern *string)
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Experimental.
	AddPackageIgnore(_pattern *string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Experimental.
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	// Prints a "tip" message during synthesis.
	// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
	AddTip(message *string)
	// Marks the provided file(s) as being generated.
	//
	// This is achieved using the
	// github-linguist attributes. Generated files do not count against the
	// repository statistics and language breakdown.
	// See: https://github.com/github/linguist/blob/master/docs/overrides.md
	//
	// Deprecated.
	AnnotateGenerated(glob *string)
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before all components are synthesized.
	// Experimental.
	PreSynthesize()
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Experimental.
	RemoveTask(name *string) projen.Task
	// Returns the shell command to execute in order to run a task.
	//
	// By default, this is `npx projen@<version> <task>`.
	// Experimental.
	RunTaskCommand(task projen.Task) *string
	// Synthesize all project files into `outdir`.
	//
	// 1. Call "this.preSynthesize()"
	// 2. Delete all generated files
	// 3. Synthesize all subprojects
	// 4. Synthesize all components of this project
	// 5. Call "postSynthesize()" for all components of this project
	// 6. Call "this.postSynthesize()"
	// Experimental.
	Synth()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Finds a file at the specified relative path within this project and all its subprojects.
	//
	// Returns: a `FileBase` or undefined if there is no file in that path.
	// Experimental.
	TryFindFile(filePath *string) projen.FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) projen.JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Experimental.
	TryFindObjectFile(filePath *string) projen.ObjectFile
	// Finds a file at the specified relative path within this project and removes it.
	//
	// Returns: a `FileBase` if the file was found and removed, or undefined if
	// the file was not found.
	// Experimental.
	TryRemoveFile(filePath *string) projen.FileBase
}

// The jsii proxy struct for AwsCdkPythonApp
type jsiiProxy_AwsCdkPythonApp struct {
	internal.Type__pythonPythonProject
}

func (j *jsiiProxy_AwsCdkPythonApp) AppEntrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) CdkConfig() CdkConfig {
	var returns CdkConfig
	_jsii_.Get(
		j,
		"cdkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) CdkDeps() AwsCdkDeps {
	var returns AwsCdkDeps
	_jsii_.Get(
		j,
		"cdkDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) CdkTasks() CdkTasks {
	var returns CdkTasks
	_jsii_.Get(
		j,
		"cdkTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) DepsManager() python.IPythonDeps {
	var returns python.IPythonDeps
	_jsii_.Get(
		j,
		"depsManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) EnvManager() python.IPythonEnv {
	var returns python.IPythonEnv
	_jsii_.Get(
		j,
		"envManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) ModuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) PackagingManager() python.IPythonPackaging {
	var returns python.IPythonPackaging
	_jsii_.Get(
		j,
		"packagingManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Pytest() python.Pytest {
	var returns python.Pytest
	_jsii_.Get(
		j,
		"pytest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) SampleTestdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleTestdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Subprojects() *[]projen.Project {
	var returns *[]projen.Project
	_jsii_.Get(
		j,
		"subprojects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkPythonApp) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkPythonApp(options *AwsCdkPythonAppOptions) AwsCdkPythonApp {
	_init_.Initialize()

	if err := validateNewAwsCdkPythonAppParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCdkPythonApp{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkPythonApp",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkPythonApp_Override(a AwsCdkPythonApp, options *AwsCdkPythonAppOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkPythonApp",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AwsCdkPythonApp)SetPytest(val python.Pytest) {
	_jsii_.Set(
		j,
		"pytest",
		val,
	)
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
func AwsCdkPythonApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkPythonApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkPythonApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a project.
// Experimental.
func AwsCdkPythonApp_IsProject(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkPythonApp_IsProjectParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkPythonApp",
		"isProject",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Find the closest ancestor project for given construct.
//
// When given a project, this it the project itself.
// Experimental.
func AwsCdkPythonApp_Of(construct constructs.IConstruct) projen.Project {
	_init_.Initialize()

	if err := validateAwsCdkPythonApp_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns projen.Project

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkPythonApp",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func AwsCdkPythonApp_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.awscdk.AwsCdkPythonApp",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) AddDependency(spec *string) {
	if err := a.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDependency",
		[]interface{}{spec},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddDevDependency(spec *string) {
	if err := a.validateAddDevDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addExcludeFromCleanup",
		args,
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddGitIgnore(pattern *string) {
	if err := a.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddPackageIgnore(_pattern *string) {
	if err := a.validateAddPackageIgnoreParameters(_pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	if err := a.validateAddTaskParameters(name, props); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		a,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) AddTip(message *string) {
	if err := a.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addTip",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AnnotateGenerated(glob *string) {
	if err := a.validateAnnotateGeneratedParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) RemoveTask(name *string) projen.Task {
	if err := a.validateRemoveTaskParameters(name); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		a,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) RunTaskCommand(task projen.Task) *string {
	if err := a.validateRunTaskCommandParameters(task); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) TryFindFile(filePath *string) projen.FileBase {
	if err := a.validateTryFindFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) TryFindJsonFile(filePath *string) projen.JsonFile {
	if err := a.validateTryFindJsonFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.JsonFile

	_jsii_.Invoke(
		a,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) TryFindObjectFile(filePath *string) projen.ObjectFile {
	if err := a.validateTryFindObjectFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.ObjectFile

	_jsii_.Invoke(
		a,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkPythonApp) TryRemoveFile(filePath *string) projen.FileBase {
	if err := a.validateTryRemoveFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

