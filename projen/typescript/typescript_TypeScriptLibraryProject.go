package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/vscode"
)

// Deprecated: use `TypeScriptProject`.
type TypeScriptLibraryProject interface {
	TypeScriptProject
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies() *bool
	// The build output directory.
	//
	// An npm tarball will be created under the `js`
	// subdirectory. For example, if this is set to `dist` (the default), the npm
	// tarball will be placed under `dist/js/boom-boom-1.2.3.tg`.
	// Deprecated: use `TypeScriptProject`.
	ArtifactsDirectory() *string
	// The location of the npm tarball after build (`${artifactsDirectory}/js`).
	// Deprecated: use `TypeScriptProject`.
	ArtifactsJavascriptDirectory() *string
	// Auto approve set up for this project.
	// Deprecated: use `TypeScriptProject`.
	AutoApprove() github.AutoApprove
	// Component that sets up mergify for merging approved pull requests.
	// Deprecated: use `TypeScriptProject`.
	AutoMerge() github.AutoMerge
	// Deprecated: use `TypeScriptProject`.
	BuildTask() projen.Task
	// The PR build GitHub workflow.
	//
	// `undefined` if `buildWorkflow` is disabled.
	// Deprecated: use `TypeScriptProject`.
	BuildWorkflow() build.BuildWorkflow
	// The job ID of the build workflow.
	// Deprecated: use `TypeScriptProject`.
	BuildWorkflowJobId() *string
	// Deprecated: use `TypeScriptProject`.
	Bundler() javascript.Bundler
	// Whether to commit the managed files by default.
	// Deprecated: use `TypeScriptProject`.
	CommitGenerated() *bool
	// Deprecated: use `TypeScriptProject`.
	CompileTask() projen.Task
	// Returns all the components within this project.
	// Deprecated: use `TypeScriptProject`.
	Components() *[]projen.Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Deprecated: use `TypeScriptProject`.
	DefaultTask() projen.Task
	// Project dependencies.
	// Deprecated: use `TypeScriptProject`.
	Deps() projen.Dependencies
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated: use `TypeScriptProject`.
	DevContainer() vscode.DevContainer
	// Deprecated: use `TypeScriptProject`.
	Docgen() *bool
	// Deprecated: use `TypeScriptProject`.
	DocsDirectory() *string
	// Whether or not the project is being ejected.
	// Deprecated: use `TypeScriptProject`.
	Ejected() *bool
	// Deprecated: use `package.entrypoint`
	Entrypoint() *string
	// Deprecated: use `TypeScriptProject`.
	Eslint() javascript.Eslint
	// All files in this project.
	// Deprecated: use `TypeScriptProject`.
	Files() *[]projen.FileBase
	// The .gitattributes file for this repository.
	// Deprecated: use `TypeScriptProject`.
	Gitattributes() projen.GitAttributesFile
	// Access all github components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated: use `TypeScriptProject`.
	Github() github.GitHub
	// .gitignore.
	// Deprecated: use `TypeScriptProject`.
	Gitignore() projen.IgnoreFile
	// Access for Gitpod.
	//
	// This will be `undefined` if gitpod boolean is false.
	// Deprecated: use `TypeScriptProject`.
	Gitpod() projen.Gitpod
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Deprecated: use `TypeScriptProject`.
	InitProject() *projen.InitProject
	// The Jest configuration (if enabled).
	// Deprecated: use `TypeScriptProject`.
	Jest() javascript.Jest
	// The directory in which compiled .js files reside.
	// Deprecated: use `TypeScriptProject`.
	Libdir() *string
	// Logging utilities.
	// Deprecated: use `TypeScriptProject`.
	Logger() projen.Logger
	// Deprecated: use `package.addField(x, y)`
	Manifest() interface{}
	// Maximum node version required by this pacakge.
	// Deprecated: use `TypeScriptProject`.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Deprecated: use `TypeScriptProject`.
	MinNodeVersion() *string
	// Project name.
	// Deprecated: use `TypeScriptProject`.
	Name() *string
	// The .npmignore file.
	// Deprecated: use `TypeScriptProject`.
	Npmignore() projen.IgnoreFile
	// Absolute output directory of this project.
	// Deprecated: use `TypeScriptProject`.
	Outdir() *string
	// API for managing the node package.
	// Deprecated: use `TypeScriptProject`.
	Package() javascript.NodePackage
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager() javascript.NodePackageManager
	// Deprecated: use `TypeScriptProject`.
	PackageTask() projen.Task
	// A parent project.
	//
	// If undefined, this is the root project.
	// Deprecated: use `TypeScriptProject`.
	Parent() projen.Project
	// Deprecated: use `TypeScriptProject`.
	PostCompileTask() projen.Task
	// Deprecated: use `TypeScriptProject`.
	PreCompileTask() projen.Task
	// Deprecated: use `TypeScriptProject`.
	Prettier() javascript.Prettier
	// Manages the build process of the project.
	// Deprecated: use `TypeScriptProject`.
	ProjectBuild() projen.ProjectBuild
	// Deprecated: use `TypeScriptProject`.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Deprecated: use `TypeScriptProject`.
	ProjenCommand() *string
	// Package publisher.
	//
	// This will be `undefined` if the project does not have a
	// release workflow.
	// Deprecated: use `release.publisher`.
	Publisher() release.Publisher
	// Release management.
	// Deprecated: use `TypeScriptProject`.
	Release() release.Release
	// The root project.
	// Deprecated: use `TypeScriptProject`.
	Root() projen.Project
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Deprecated: use `TypeScriptProject`.
	RunScriptCommand() *string
	// The directory in which the .ts sources reside.
	// Deprecated: use `TypeScriptProject`.
	Srcdir() *string
	// Project tasks.
	// Deprecated: use `TypeScriptProject`.
	Tasks() projen.Tasks
	// The directory in which tests reside.
	// Deprecated: use `TypeScriptProject`.
	Testdir() *string
	// Deprecated: use `TypeScriptProject`.
	TestTask() projen.Task
	// Deprecated: use `TypeScriptProject`.
	Tsconfig() javascript.TypescriptConfig
	// A typescript configuration file which covers all files (sources, tests, projen).
	// Deprecated: use `TypeScriptProject`.
	TsconfigDev() javascript.TypescriptConfig
	// Deprecated: use `TypeScriptProject`.
	TsconfigEslint() javascript.TypescriptConfig
	// The upgrade workflow.
	// Deprecated: use `TypeScriptProject`.
	UpgradeWorkflow() javascript.UpgradeDependencies
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated: use `TypeScriptProject`.
	Vscode() vscode.VsCode
	// The "watch" task.
	// Deprecated: use `TypeScriptProject`.
	WatchTask() projen.Task
	// Deprecated: use `TypeScriptProject`.
	AddBins(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Deprecated: use `TypeScriptProject`.
	AddBundledDeps(deps ...*string)
	// DEPRECATED.
	// Deprecated: use `project.compileTask.exec()`
	AddCompileCommand(commands ...*string)
	// Defines normal dependencies.
	// Deprecated: use `TypeScriptProject`.
	AddDeps(deps ...*string)
	// Defines development/test dependencies.
	// Deprecated: use `TypeScriptProject`.
	AddDevDeps(deps ...*string)
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Deprecated: use `TypeScriptProject`.
	AddExcludeFromCleanup(globs ...*string)
	// Directly set fields in `package.json`.
	// Deprecated: use `TypeScriptProject`.
	AddFields(fields *map[string]interface{})
	// Adds a .gitignore pattern.
	// Deprecated: use `TypeScriptProject`.
	AddGitIgnore(pattern *string)
	// Adds keywords to package.json (deduplicated).
	// Deprecated: use `TypeScriptProject`.
	AddKeywords(keywords ...*string)
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Deprecated: use `TypeScriptProject`.
	AddPackageIgnore(pattern *string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Deprecated: use `TypeScriptProject`.
	AddPeerDeps(deps ...*string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Deprecated: use `TypeScriptProject`.
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	// DEPRECATED.
	// Deprecated: use `project.testTask.exec()`
	AddTestCommand(commands ...*string)
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
	// Deprecated: use `TypeScriptProject`.
	AnnotateGenerated(glob *string)
	// Indicates if a script by the name name is defined.
	// Deprecated: use `TypeScriptProject`.
	HasScript(name *string) *bool
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Deprecated: use `TypeScriptProject`.
	PostSynthesize()
	// Called before all components are synthesized.
	// Deprecated: use `TypeScriptProject`.
	PreSynthesize()
	// Removes the npm script (always successful).
	// Deprecated: use `TypeScriptProject`.
	RemoveScript(name *string)
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Deprecated: use `TypeScriptProject`.
	RemoveTask(name *string) projen.Task
	// Returns the set of workflow steps which should be executed to bootstrap a workflow.
	//
	// Returns: Job steps.
	// Deprecated: use `TypeScriptProject`.
	RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep
	// Returns the shell command to execute in order to run a task.
	//
	// This will
	// typically be `npx projen TASK`.
	// Deprecated: use `TypeScriptProject`.
	RunTaskCommand(task projen.Task) *string
	// Replaces the contents of an npm package.json script.
	// Deprecated: use `TypeScriptProject`.
	SetScript(name *string, command *string)
	// Synthesize all project files into `outdir`.
	//
	// 1. Call "this.preSynthesize()"
	// 2. Delete all generated files
	// 3. Synthesize all sub-projects
	// 4. Synthesize all components of this project
	// 5. Call "postSynthesize()" for all components of this project
	// 6. Call "this.postSynthesize()"
	// Deprecated: use `TypeScriptProject`.
	Synth()
	// Finds a file at the specified relative path within this project and all its subprojects.
	//
	// Returns: a `FileBase` or undefined if there is no file in that path.
	// Deprecated: use `TypeScriptProject`.
	TryFindFile(filePath *string) projen.FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) projen.JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Deprecated: use `TypeScriptProject`.
	TryFindObjectFile(filePath *string) projen.ObjectFile
	// Finds a file at the specified relative path within this project and removes it.
	//
	// Returns: a `FileBase` if the file was found and removed, or undefined if
	// the file was not found.
	// Deprecated: use `TypeScriptProject`.
	TryRemoveFile(filePath *string) projen.FileBase
}

// The jsii proxy struct for TypeScriptLibraryProject
type jsiiProxy_TypeScriptLibraryProject struct {
	jsiiProxy_TypeScriptProject
}

func (j *jsiiProxy_TypeScriptLibraryProject) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Prettier() javascript.Prettier {
	var returns javascript.Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypeScriptLibraryProject) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Deprecated: use `TypeScriptProject`.
func NewTypeScriptLibraryProject(options *TypeScriptProjectOptions) TypeScriptLibraryProject {
	_init_.Initialize()

	j := jsiiProxy_TypeScriptLibraryProject{}

	_jsii_.Create(
		"projen.typescript.TypeScriptLibraryProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Deprecated: use `TypeScriptProject`.
func NewTypeScriptLibraryProject_Override(t TypeScriptLibraryProject, options *TypeScriptProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.typescript.TypeScriptLibraryProject",
		[]interface{}{options},
		t,
	)
}

func TypeScriptLibraryProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.typescript.TypeScriptLibraryProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		t,
		"addBins",
		[]interface{}{bins},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addBundledDeps",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addCompileCommand",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDeps",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDevDeps",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addExcludeFromCleanup",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addFields",
		[]interface{}{fields},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addKeywords",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addPeerDeps",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		t,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addTestCommand",
		args,
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		t,
		"addTip",
		[]interface{}{message},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		t,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		t,
		"removeScript",
		[]interface{}{name},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		t,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep

	_jsii_.Invoke(
		t,
		"renderWorkflowSetup",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		t,
		"setScript",
		[]interface{}{name, command},
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) Synth() {
	_jsii_.InvokeVoid(
		t,
		"synth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypeScriptLibraryProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		t,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		t,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		t,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypeScriptLibraryProject) TryRemoveFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		t,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

