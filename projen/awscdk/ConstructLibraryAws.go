package awscdk

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

// Deprecated: use `AwsCdkConstructLibrary`.
type ConstructLibraryAws interface {
	AwsCdkConstructLibrary
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies() *bool
	// The build output directory.
	//
	// An npm tarball will be created under the `js`
	// subdirectory. For example, if this is set to `dist` (the default), the npm
	// tarball will be placed under `dist/js/boom-boom-1.2.3.tg`.
	// Deprecated: use `AwsCdkConstructLibrary`.
	ArtifactsDirectory() *string
	// The location of the npm tarball after build (`${artifactsDirectory}/js`).
	// Deprecated: use `AwsCdkConstructLibrary`.
	ArtifactsJavascriptDirectory() *string
	// Auto approve set up for this project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AutoApprove() github.AutoApprove
	// Component that sets up mergify for merging approved pull requests.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AutoMerge() github.AutoMerge
	// Deprecated: use `AwsCdkConstructLibrary`.
	BuildTask() projen.Task
	// The PR build GitHub workflow.
	//
	// `undefined` if `buildWorkflow` is disabled.
	// Deprecated: use `AwsCdkConstructLibrary`.
	BuildWorkflow() build.BuildWorkflow
	// The job ID of the build workflow.
	// Deprecated: use `AwsCdkConstructLibrary`.
	BuildWorkflowJobId() *string
	// Deprecated: use `AwsCdkConstructLibrary`.
	Bundler() javascript.Bundler
	// Deprecated: use `AwsCdkConstructLibrary`.
	CdkDeps() AwsCdkDeps
	// The target CDK version for this library.
	// Deprecated: use `AwsCdkConstructLibrary`.
	CdkVersion() *string
	// Whether to commit the managed files by default.
	// Deprecated: use `AwsCdkConstructLibrary`.
	CommitGenerated() *bool
	// Deprecated: use `AwsCdkConstructLibrary`.
	CompileTask() projen.Task
	// Returns all the components within this project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Components() *[]projen.Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Deprecated: use `AwsCdkConstructLibrary`.
	DefaultTask() projen.Task
	// Project dependencies.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Deps() projen.Dependencies
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated: use `AwsCdkConstructLibrary`.
	DevContainer() vscode.DevContainer
	// Deprecated: use `AwsCdkConstructLibrary`.
	Docgen() *bool
	// Deprecated: use `AwsCdkConstructLibrary`.
	DocsDirectory() *string
	// Whether or not the project is being ejected.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Ejected() *bool
	// Deprecated: use `package.entrypoint`
	Entrypoint() *string
	// Deprecated: use `AwsCdkConstructLibrary`.
	Eslint() javascript.Eslint
	// All files in this project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Files() *[]projen.FileBase
	// The .gitattributes file for this repository.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Gitattributes() projen.GitAttributesFile
	// Access all github components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Github() github.GitHub
	// .gitignore.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Gitignore() projen.IgnoreFile
	// Access for Gitpod.
	//
	// This will be `undefined` if gitpod boolean is false.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Gitpod() projen.Gitpod
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Deprecated: use `AwsCdkConstructLibrary`.
	InitProject() *projen.InitProject
	// The Jest configuration (if enabled).
	// Deprecated: use `AwsCdkConstructLibrary`.
	Jest() javascript.Jest
	// The directory in which compiled .js files reside.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Libdir() *string
	// Logging utilities.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Logger() projen.Logger
	// Deprecated: use `package.addField(x, y)`
	Manifest() interface{}
	// Maximum node version required by this package.
	// Deprecated: use `AwsCdkConstructLibrary`.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Deprecated: use `AwsCdkConstructLibrary`.
	MinNodeVersion() *string
	// Project name.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Name() *string
	// Deprecated: use `AwsCdkConstructLibrary`.
	NodeVersion() *string
	// The .npmignore file.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Npmignore() projen.IgnoreFile
	// The .npmrc file.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Npmrc() javascript.NpmConfig
	// Absolute output directory of this project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Outdir() *string
	// API for managing the node package.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Package() javascript.NodePackage
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager() javascript.NodePackageManager
	// Deprecated: use `AwsCdkConstructLibrary`.
	PackageTask() projen.Task
	// A parent project.
	//
	// If undefined, this is the root project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Parent() projen.Project
	// Deprecated: use `AwsCdkConstructLibrary`.
	PostCompileTask() projen.Task
	// Deprecated: use `AwsCdkConstructLibrary`.
	PreCompileTask() projen.Task
	// Deprecated: use `AwsCdkConstructLibrary`.
	Prettier() javascript.Prettier
	// Manages the build process of the project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	ProjectBuild() projen.ProjectBuild
	// Deprecated: use `AwsCdkConstructLibrary`.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Deprecated: use `AwsCdkConstructLibrary`.
	ProjenCommand() *string
	// Package publisher.
	//
	// This will be `undefined` if the project does not have a
	// release workflow.
	// Deprecated: use `release.publisher`.
	Publisher() release.Publisher
	// Release management.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Release() release.Release
	// The root project.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Root() projen.Project
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Deprecated: use `AwsCdkConstructLibrary`.
	RunScriptCommand() *string
	// The directory in which the .ts sources reside.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Srcdir() *string
	// Project tasks.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Tasks() projen.Tasks
	// The directory in which tests reside.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Testdir() *string
	// Deprecated: use `AwsCdkConstructLibrary`.
	TestTask() projen.Task
	// Deprecated: use `AwsCdkConstructLibrary`.
	Tsconfig() javascript.TypescriptConfig
	// A typescript configuration file which covers all files (sources, tests, projen).
	// Deprecated: use `AwsCdkConstructLibrary`.
	TsconfigDev() javascript.TypescriptConfig
	// Deprecated: use `AwsCdkConstructLibrary`.
	TsconfigEslint() javascript.TypescriptConfig
	// The upgrade workflow.
	// Deprecated: use `AwsCdkConstructLibrary`.
	UpgradeWorkflow() javascript.UpgradeDependencies
	// Deprecated: use `cdkVersion`.
	Version() *string
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Vscode() vscode.VsCode
	// The "watch" task.
	// Deprecated: use `AwsCdkConstructLibrary`.
	WatchTask() projen.Task
	// Deprecated: use `AwsCdkConstructLibrary`.
	WorkflowBootstrapSteps() *[]*workflows.JobStep
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddBins(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddBundledDeps(deps ...*string)
	// Adds dependencies to AWS CDK modules.
	//
	// Since this is a library project, dependencies will be added as peer dependencies.
	// Deprecated: Not supported in v2. For v1, use `project.cdkDeps.addV1Dependencies()`
	AddCdkDependencies(deps ...*string)
	// Adds AWS CDK modules as dev dependencies.
	// Deprecated: Not supported in v2. For v1, use `project.cdkDeps.addV1DevDependencies()`
	AddCdkTestDependencies(deps ...*string)
	// DEPRECATED.
	// Deprecated: use `project.compileTask.exec()`
	AddCompileCommand(commands ...*string)
	// Defines normal dependencies.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddDeps(deps ...*string)
	// Defines development/test dependencies.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddDevDeps(deps ...*string)
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddExcludeFromCleanup(globs ...*string)
	// Directly set fields in `package.json`.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddFields(fields *map[string]interface{})
	// Adds a .gitignore pattern.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddGitIgnore(pattern *string)
	// Adds keywords to package.json (deduplicated).
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddKeywords(keywords ...*string)
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddPackageIgnore(pattern *string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddPeerDeps(deps ...*string)
	// Replaces the contents of multiple npm package.json scripts.
	// Deprecated: use `AwsCdkConstructLibrary`.
	AddScripts(scripts *map[string]*string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Deprecated: use `AwsCdkConstructLibrary`.
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
	// Deprecated: use `AwsCdkConstructLibrary`.
	AnnotateGenerated(glob *string)
	// Indicates if a script by the name name is defined.
	// Deprecated: Use `project.tasks.tryFind(name)`
	HasScript(name *string) *bool
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Deprecated: use `AwsCdkConstructLibrary`.
	PostSynthesize()
	// Called before all components are synthesized.
	// Deprecated: use `AwsCdkConstructLibrary`.
	PreSynthesize()
	// Removes the npm script (always successful).
	// Deprecated: use `AwsCdkConstructLibrary`.
	RemoveScript(name *string)
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Deprecated: use `AwsCdkConstructLibrary`.
	RemoveTask(name *string) projen.Task
	// Returns the set of workflow steps which should be executed to bootstrap a workflow.
	//
	// Returns: Job steps.
	// Deprecated: use `AwsCdkConstructLibrary`.
	RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep
	// Returns the shell command to execute in order to run a task.
	//
	// This will
	// typically be `npx projen TASK`.
	// Deprecated: use `AwsCdkConstructLibrary`.
	RunTaskCommand(task projen.Task) *string
	// Replaces the contents of an npm package.json script.
	// Deprecated: use `AwsCdkConstructLibrary`.
	SetScript(name *string, command *string)
	// Synthesize all project files into `outdir`.
	//
	// 1. Call "this.preSynthesize()"
	// 2. Delete all generated files
	// 3. Synthesize all sub-projects
	// 4. Synthesize all components of this project
	// 5. Call "postSynthesize()" for all components of this project
	// 6. Call "this.postSynthesize()"
	// Deprecated: use `AwsCdkConstructLibrary`.
	Synth()
	// Finds a file at the specified relative path within this project and all its subprojects.
	//
	// Returns: a `FileBase` or undefined if there is no file in that path.
	// Deprecated: use `AwsCdkConstructLibrary`.
	TryFindFile(filePath *string) projen.FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) projen.JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Deprecated: use `AwsCdkConstructLibrary`.
	TryFindObjectFile(filePath *string) projen.ObjectFile
	// Finds a file at the specified relative path within this project and removes it.
	//
	// Returns: a `FileBase` if the file was found and removed, or undefined if
	// the file was not found.
	// Deprecated: use `AwsCdkConstructLibrary`.
	TryRemoveFile(filePath *string) projen.FileBase
}

// The jsii proxy struct for ConstructLibraryAws
type jsiiProxy_ConstructLibraryAws struct {
	jsiiProxy_AwsCdkConstructLibrary
}

func (j *jsiiProxy_ConstructLibraryAws) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) CdkDeps() AwsCdkDeps {
	var returns AwsCdkDeps
	_jsii_.Get(
		j,
		"cdkDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Npmrc() javascript.NpmConfig {
	var returns javascript.NpmConfig
	_jsii_.Get(
		j,
		"npmrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Prettier() javascript.Prettier {
	var returns javascript.Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryAws) WorkflowBootstrapSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"workflowBootstrapSteps",
		&returns,
	)
	return returns
}


// Deprecated: use `AwsCdkConstructLibrary`.
func NewConstructLibraryAws(options *AwsCdkConstructLibraryOptions) ConstructLibraryAws {
	_init_.Initialize()

	if err := validateNewConstructLibraryAwsParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConstructLibraryAws{}

	_jsii_.Create(
		"projen.awscdk.ConstructLibraryAws",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Deprecated: use `AwsCdkConstructLibrary`.
func NewConstructLibraryAws_Override(c ConstructLibraryAws, options *AwsCdkConstructLibraryOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.ConstructLibraryAws",
		[]interface{}{options},
		c,
	)
}

func ConstructLibraryAws_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.awscdk.ConstructLibraryAws",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) AddBins(bins *map[string]*string) {
	if err := c.validateAddBinsParameters(bins); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBins",
		[]interface{}{bins},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addBundledDeps",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddCdkDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addCdkDependencies",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddCdkTestDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addCdkTestDependencies",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addCompileCommand",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDeps",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDevDeps",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addExcludeFromCleanup",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddFields(fields *map[string]interface{}) {
	if err := c.validateAddFieldsParameters(fields); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addFields",
		[]interface{}{fields},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddGitIgnore(pattern *string) {
	if err := c.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addKeywords",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddPackageIgnore(pattern *string) {
	if err := c.validateAddPackageIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addPeerDeps",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddScripts(scripts *map[string]*string) {
	if err := c.validateAddScriptsParameters(scripts); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addScripts",
		[]interface{}{scripts},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	if err := c.validateAddTaskParameters(name, props); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		c,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addTestCommand",
		args,
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddTip(message *string) {
	if err := c.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addTip",
		[]interface{}{message},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AnnotateGenerated(glob *string) {
	if err := c.validateAnnotateGeneratedParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) HasScript(name *string) *bool {
	if err := c.validateHasScriptParameters(name); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConstructLibraryAws) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConstructLibraryAws) RemoveScript(name *string) {
	if err := c.validateRemoveScriptParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeScript",
		[]interface{}{name},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) RemoveTask(name *string) projen.Task {
	if err := c.validateRemoveTaskParameters(name); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		c,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep {
	if err := c.validateRenderWorkflowSetupParameters(options); err != nil {
		panic(err)
	}
	var returns *[]*workflows.JobStep

	_jsii_.Invoke(
		c,
		"renderWorkflowSetup",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) RunTaskCommand(task projen.Task) *string {
	if err := c.validateRunTaskCommandParameters(task); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) SetScript(name *string, command *string) {
	if err := c.validateSetScriptParameters(name, command); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"setScript",
		[]interface{}{name, command},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) Synth() {
	_jsii_.InvokeVoid(
		c,
		"synth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConstructLibraryAws) TryFindFile(filePath *string) projen.FileBase {
	if err := c.validateTryFindFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		c,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) TryFindJsonFile(filePath *string) projen.JsonFile {
	if err := c.validateTryFindJsonFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.JsonFile

	_jsii_.Invoke(
		c,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) TryFindObjectFile(filePath *string) projen.ObjectFile {
	if err := c.validateTryFindObjectFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.ObjectFile

	_jsii_.Invoke(
		c,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConstructLibraryAws) TryRemoveFile(filePath *string) projen.FileBase {
	if err := c.validateTryRemoveFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		c,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

