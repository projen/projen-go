package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/java/internal"
	"github.com/projen/projen-go/projen/vscode"
)

// Java project.
// Experimental.
type JavaProject interface {
	github.GitHubProject
	// Auto approve set up for this project.
	// Deprecated.
	AutoApprove() github.AutoApprove
	// Experimental.
	BuildTask() projen.Task
	// Whether to commit the managed files by default.
	// Experimental.
	CommitGenerated() *bool
	// Compile component.
	// Experimental.
	Compile() MavenCompile
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
	Junit() Junit
	// Logging utilities.
	// Experimental.
	Logger() projen.Logger
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
	// Packaging component.
	// Experimental.
	Packaging() MavenPackaging
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() projen.Project
	// API for managing `pom.xml`.
	// Experimental.
	Pom() Pom
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
	Projenrc() Projenrc
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
	AddPlugin(spec *string, options *PluginOptions) *projen.Dependency
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

// The jsii proxy struct for JavaProject
type jsiiProxy_JavaProject struct {
	internal.Type__githubGitHubProject
}

func (j *jsiiProxy_JavaProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Compile() MavenCompile {
	var returns MavenCompile
	_jsii_.Get(
		j,
		"compile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Distdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Junit() Junit {
	var returns Junit
	_jsii_.Get(
		j,
		"junit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Packaging() MavenPackaging {
	var returns MavenPackaging
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Pom() Pom {
	var returns Pom
	_jsii_.Get(
		j,
		"pom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Projenrc() Projenrc {
	var returns Projenrc
	_jsii_.Get(
		j,
		"projenrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Subprojects() *[]projen.Project {
	var returns *[]projen.Project
	_jsii_.Get(
		j,
		"subprojects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewJavaProject(options *JavaProjectOptions) JavaProject {
	_init_.Initialize()

	if err := validateNewJavaProjectParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaProject{}

	_jsii_.Create(
		"projen.java.JavaProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewJavaProject_Override(j JavaProject, options *JavaProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.JavaProject",
		[]interface{}{options},
		j,
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
func JavaProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJavaProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.JavaProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a project.
// Experimental.
func JavaProject_IsProject(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJavaProject_IsProjectParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.JavaProject",
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
func JavaProject_Of(construct constructs.IConstruct) projen.Project {
	_init_.Initialize()

	if err := validateJavaProject_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns projen.Project

	_jsii_.StaticInvoke(
		"projen.java.JavaProject",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func JavaProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.java.JavaProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaProject) AddDependency(spec *string) {
	if err := j.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addDependency",
		[]interface{}{spec},
	)
}

func (j *jsiiProxy_JavaProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		j,
		"addExcludeFromCleanup",
		args,
	)
}

func (j *jsiiProxy_JavaProject) AddGitIgnore(pattern *string) {
	if err := j.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_JavaProject) AddPackageIgnore(_pattern *string) {
	if err := j.validateAddPackageIgnoreParameters(_pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (j *jsiiProxy_JavaProject) AddPlugin(spec *string, options *PluginOptions) *projen.Dependency {
	if err := j.validateAddPluginParameters(spec, options); err != nil {
		panic(err)
	}
	var returns *projen.Dependency

	_jsii_.Invoke(
		j,
		"addPlugin",
		[]interface{}{spec, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	if err := j.validateAddTaskParameters(name, props); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		j,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) AddTestDependency(spec *string) {
	if err := j.validateAddTestDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addTestDependency",
		[]interface{}{spec},
	)
}

func (j *jsiiProxy_JavaProject) AddTip(message *string) {
	if err := j.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addTip",
		[]interface{}{message},
	)
}

func (j *jsiiProxy_JavaProject) AnnotateGenerated(glob *string) {
	if err := j.validateAnnotateGeneratedParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (j *jsiiProxy_JavaProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaProject) RemoveTask(name *string) projen.Task {
	if err := j.validateRemoveTaskParameters(name); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		j,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) RunTaskCommand(task projen.Task) *string {
	if err := j.validateRunTaskCommandParameters(task); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) Synth() {
	_jsii_.InvokeVoid(
		j,
		"synth",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JavaProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) TryFindFile(filePath *string) projen.FileBase {
	if err := j.validateTryFindFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		j,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	if err := j.validateTryFindJsonFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.JsonFile

	_jsii_.Invoke(
		j,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	if err := j.validateTryFindObjectFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.ObjectFile

	_jsii_.Invoke(
		j,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaProject) TryRemoveFile(filePath *string) projen.FileBase {
	if err := j.validateTryRemoveFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		j,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

