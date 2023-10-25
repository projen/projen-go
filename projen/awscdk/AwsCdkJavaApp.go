package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/java"
	"github.com/projen/projen-go/projen/vscode"
)

// AWS CDK app in Java.
// Experimental.
type AwsCdkJavaApp interface {
	java.JavaProject
	// Auto approve set up for this project.
	// Deprecated.
	AutoApprove() github.AutoApprove
	// Experimental.
	BuildTask() projen.Task
	// The `cdk.json` file.
	// Experimental.
	CdkConfig() CdkConfig
	// CDK dependency management helper class.
	// Experimental.
	CdkDeps() AwsCdkDeps
	// CDK tasks.
	// Experimental.
	CdkTasks() CdkTasks
	// Whether to commit the managed files by default.
	// Experimental.
	CommitGenerated() *bool
	// Compile component.
	// Experimental.
	Compile() java.MavenCompile
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
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated.
	DevContainer() vscode.DevContainer
	// Maven artifact output directory.
	// Experimental.
	Distdir() *string
	// Whether or not the project is being ejected.
	// Experimental.
	Ejected() *bool
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
	// JUnit component.
	// Experimental.
	Junit() java.Junit
	// Logging utilities.
	// Experimental.
	Logger() projen.Logger
	// The full name of the main class of the java app (package.Class).
	// Experimental.
	MainClass() *string
	// The name of the Java class with the static `main()` method.
	// Experimental.
	MainClassName() *string
	// The name of the Java package that includes the main class.
	// Experimental.
	MainPackage() *string
	// Project name.
	// Experimental.
	Name() *string
	// Absolute output directory of this project.
	// Experimental.
	Outdir() *string
	// Experimental.
	PackageTask() projen.Task
	// Packaging component.
	// Experimental.
	Packaging() java.MavenPackaging
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() projen.Project
	// API for managing `pom.xml`.
	// Experimental.
	Pom() java.Pom
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
	// Projenrc component.
	// Experimental.
	Projenrc() java.Projenrc
	// The root project.
	// Experimental.
	Root() projen.Project
	// Returns all the subprojects within this project.
	// Experimental.
	Subprojects() *[]projen.Project
	// Project tasks.
	// Experimental.
	Tasks() projen.Tasks
	// Experimental.
	TestTask() projen.Task
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Vscode() vscode.VsCode
	// Adds an AWS CDK module dependencies.
	// Deprecated: In CDK 2.x all modules are available by default. Alpha modules should be added using the standard 'deps'
	AddCdkDependency(modules ...*string)
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
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
	// Adds a build plugin to the pom.
	//
	// The plug in is also added as a BUILD dep to the project.
	// Experimental.
	AddPlugin(spec *string, options *java.PluginOptions) *projen.Dependency
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Experimental.
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	// Adds a test dependency.
	// Experimental.
	AddTestDependency(spec *string)
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

// The jsii proxy struct for AwsCdkJavaApp
type jsiiProxy_AwsCdkJavaApp struct {
	internal.Type__javaJavaProject
}

func (j *jsiiProxy_AwsCdkJavaApp) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) CdkConfig() CdkConfig {
	var returns CdkConfig
	_jsii_.Get(
		j,
		"cdkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) CdkDeps() AwsCdkDeps {
	var returns AwsCdkDeps
	_jsii_.Get(
		j,
		"cdkDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) CdkTasks() CdkTasks {
	var returns CdkTasks
	_jsii_.Get(
		j,
		"cdkTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Compile() java.MavenCompile {
	var returns java.MavenCompile
	_jsii_.Get(
		j,
		"compile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Distdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Junit() java.Junit {
	var returns java.Junit
	_jsii_.Get(
		j,
		"junit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) MainClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) MainClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) MainPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Packaging() java.MavenPackaging {
	var returns java.MavenPackaging
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Pom() java.Pom {
	var returns java.Pom
	_jsii_.Get(
		j,
		"pom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Projenrc() java.Projenrc {
	var returns java.Projenrc
	_jsii_.Get(
		j,
		"projenrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Subprojects() *[]projen.Project {
	var returns *[]projen.Project
	_jsii_.Get(
		j,
		"subprojects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkJavaApp(options *AwsCdkJavaAppOptions) AwsCdkJavaApp {
	_init_.Initialize()

	if err := validateNewAwsCdkJavaAppParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCdkJavaApp{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkJavaApp",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkJavaApp_Override(a AwsCdkJavaApp, options *AwsCdkJavaAppOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkJavaApp",
		[]interface{}{options},
		a,
	)
}

func AwsCdkJavaApp_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.awscdk.AwsCdkJavaApp",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCdkJavaApp) AddCdkDependency(modules ...*string) {
	args := []interface{}{}
	for _, a := range modules {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addCdkDependency",
		args,
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddDependency(spec *string) {
	if err := a.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addDependency",
		[]interface{}{spec},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddExcludeFromCleanup(globs ...*string) {
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

func (a *jsiiProxy_AwsCdkJavaApp) AddGitIgnore(pattern *string) {
	if err := a.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddPackageIgnore(_pattern *string) {
	if err := a.validateAddPackageIgnoreParameters(_pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddPlugin(spec *string, options *java.PluginOptions) *projen.Dependency {
	if err := a.validateAddPluginParameters(spec, options); err != nil {
		panic(err)
	}
	var returns *projen.Dependency

	_jsii_.Invoke(
		a,
		"addPlugin",
		[]interface{}{spec, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkJavaApp) AddTask(name *string, props *projen.TaskOptions) projen.Task {
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

func (a *jsiiProxy_AwsCdkJavaApp) AddTestDependency(spec *string) {
	if err := a.validateAddTestDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addTestDependency",
		[]interface{}{spec},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddTip(message *string) {
	if err := a.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addTip",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AnnotateGenerated(glob *string) {
	if err := a.validateAnnotateGeneratedParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) RemoveTask(name *string) projen.Task {
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

func (a *jsiiProxy_AwsCdkJavaApp) RunTaskCommand(task projen.Task) *string {
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

func (a *jsiiProxy_AwsCdkJavaApp) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) TryFindFile(filePath *string) projen.FileBase {
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

func (a *jsiiProxy_AwsCdkJavaApp) TryFindJsonFile(filePath *string) projen.JsonFile {
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

func (a *jsiiProxy_AwsCdkJavaApp) TryFindObjectFile(filePath *string) projen.ObjectFile {
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

func (a *jsiiProxy_AwsCdkJavaApp) TryRemoveFile(filePath *string) projen.FileBase {
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

