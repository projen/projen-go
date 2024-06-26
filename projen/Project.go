package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen/internal"
)

// Base project.
// Experimental.
type Project interface {
	constructs.Construct
	// Experimental.
	BuildTask() Task
	// Whether to commit the managed files by default.
	// Experimental.
	CommitGenerated() *bool
	// Experimental.
	CompileTask() Task
	// Returns all the components within this project.
	// Experimental.
	Components() *[]Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Experimental.
	DefaultTask() Task
	// Project dependencies.
	// Experimental.
	Deps() Dependencies
	// Whether or not the project is being ejected.
	// Experimental.
	Ejected() *bool
	// All files in this project.
	// Experimental.
	Files() *[]FileBase
	// The .gitattributes file for this repository.
	// Experimental.
	Gitattributes() GitAttributesFile
	// .gitignore.
	// Experimental.
	Gitignore() IgnoreFile
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Experimental.
	InitProject() *InitProject
	// Logging utilities.
	// Experimental.
	Logger() Logger
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
	PackageTask() Task
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() Project
	// Experimental.
	PostCompileTask() Task
	// Experimental.
	PreCompileTask() Task
	// Manages the build process of the project.
	// Experimental.
	ProjectBuild() ProjectBuild
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand() *string
	// The root project.
	// Experimental.
	Root() Project
	// Returns all the subprojects within this project.
	// Experimental.
	Subprojects() *[]Project
	// Project tasks.
	// Experimental.
	Tasks() Tasks
	// Experimental.
	TestTask() Task
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
	AddTask(name *string, props *TaskOptions) Task
	// Prints a "tip" message during synthesis.
	// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
	AddTip(message *string)
	// Consider a set of files as "generated".
	//
	// This method is implemented by
	// derived classes and used for example, to add git attributes to tell GitHub
	// that certain files are generated.
	// Experimental.
	AnnotateGenerated(_glob *string)
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
	RemoveTask(name *string) Task
	// Returns the shell command to execute in order to run a task.
	//
	// By default, this is `npx projen@<version> <task>`.
	// Experimental.
	RunTaskCommand(task Task) *string
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
	TryFindFile(filePath *string) FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Experimental.
	TryFindObjectFile(filePath *string) ObjectFile
	// Finds a file at the specified relative path within this project and removes it.
	//
	// Returns: a `FileBase` if the file was found and removed, or undefined if
	// the file was not found.
	// Experimental.
	TryRemoveFile(filePath *string) FileBase
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Project) BuildTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) CompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Components() *[]Component {
	var returns *[]Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DefaultTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Deps() Dependencies {
	var returns Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Files() *[]FileBase {
	var returns *[]FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Gitattributes() GitAttributesFile {
	var returns GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Gitignore() IgnoreFile {
	var returns IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) InitProject() *InitProject {
	var returns *InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Logger() Logger {
	var returns Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PackageTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Parent() Project {
	var returns Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PostCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PreCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectBuild() ProjectBuild {
	var returns ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Root() Project {
	var returns Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Subprojects() *[]Project {
	var returns *[]Project
	_jsii_.Get(
		j,
		"subprojects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Tasks() Tasks {
	var returns Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TestTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewProject(options *ProjectOptions) Project {
	_init_.Initialize()

	if err := validateNewProjectParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Project{}

	_jsii_.Create(
		"projen.Project",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewProject_Override(p Project, options *ProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Project",
		[]interface{}{options},
		p,
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
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a project.
// Experimental.
func Project_IsProject(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsProjectParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Project",
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
func Project_Of(construct constructs.IConstruct) Project {
	_init_.Initialize()

	if err := validateProject_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns Project

	_jsii_.StaticInvoke(
		"projen.Project",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Project_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.Project",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Project) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addExcludeFromCleanup",
		args,
	)
}

func (p *jsiiProxy_Project) AddGitIgnore(pattern *string) {
	if err := p.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (p *jsiiProxy_Project) AddPackageIgnore(_pattern *string) {
	if err := p.validateAddPackageIgnoreParameters(_pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (p *jsiiProxy_Project) AddTask(name *string, props *TaskOptions) Task {
	if err := p.validateAddTaskParameters(name, props); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		p,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) AddTip(message *string) {
	if err := p.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addTip",
		[]interface{}{message},
	)
}

func (p *jsiiProxy_Project) AnnotateGenerated(_glob *string) {
	if err := p.validateAnnotateGeneratedParameters(_glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"annotateGenerated",
		[]interface{}{_glob},
	)
}

func (p *jsiiProxy_Project) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) RemoveTask(name *string) Task {
	if err := p.validateRemoveTaskParameters(name); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		p,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) RunTaskCommand(task Task) *string {
	if err := p.validateRunTaskCommandParameters(task); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) Synth() {
	_jsii_.InvokeVoid(
		p,
		"synth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) TryFindFile(filePath *string) FileBase {
	if err := p.validateTryFindFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns FileBase

	_jsii_.Invoke(
		p,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) TryFindJsonFile(filePath *string) JsonFile {
	if err := p.validateTryFindJsonFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns JsonFile

	_jsii_.Invoke(
		p,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) TryFindObjectFile(filePath *string) ObjectFile {
	if err := p.validateTryFindObjectFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns ObjectFile

	_jsii_.Invoke(
		p,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) TryRemoveFile(filePath *string) FileBase {
	if err := p.validateTryRemoveFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns FileBase

	_jsii_.Invoke(
		p,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

