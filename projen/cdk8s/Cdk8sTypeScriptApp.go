package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build"
	"github.com/projen/projen-go/projen/cdk8s/internal"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/typescript"
	"github.com/projen/projen-go/projen/vscode"
)

// CDK8s app in TypeScript.
// Experimental.
type Cdk8sTypeScriptApp interface {
	typescript.TypeScriptAppProject
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies() *bool
	// The CDK8s app entrypoint.
	// Experimental.
	AppEntrypoint() *string
	// The build output directory.
	//
	// An npm tarball will be created under the `js`
	// subdirectory. For example, if this is set to `dist` (the default), the npm
	// tarball will be placed under `dist/js/boom-boom-1.2.3.tg`.
	// Experimental.
	ArtifactsDirectory() *string
	// The location of the npm tarball after build (`${artifactsDirectory}/js`).
	// Experimental.
	ArtifactsJavascriptDirectory() *string
	// Auto approve set up for this project.
	// Deprecated.
	AutoApprove() github.AutoApprove
	// Component that sets up mergify for merging approved pull requests.
	// Experimental.
	AutoMerge() github.AutoMerge
	// Experimental.
	BuildTask() projen.Task
	// The PR build GitHub workflow.
	//
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow() build.BuildWorkflow
	// The job ID of the build workflow.
	// Experimental.
	BuildWorkflowJobId() *string
	// Experimental.
	Bundler() javascript.Bundler
	// Experimental.
	Cdk8sDeps() Cdk8sDeps
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
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated.
	DevContainer() vscode.DevContainer
	// Experimental.
	Docgen() *bool
	// Experimental.
	DocsDirectory() *string
	// Whether or not the project is being ejected.
	// Experimental.
	Ejected() *bool
	// Deprecated: use `package.entrypoint`
	Entrypoint() *string
	// Experimental.
	Eslint() javascript.Eslint
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
	// The Jest configuration (if enabled).
	// Experimental.
	Jest() javascript.Jest
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir() *string
	// Logging utilities.
	// Experimental.
	Logger() projen.Logger
	// Deprecated: use `package.addField(x, y)`
	Manifest() interface{}
	// Maximum node version required by this package.
	// Experimental.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion() *string
	// Project name.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeVersion() *string
	// The .npmignore file.
	// Experimental.
	Npmignore() projen.IgnoreFile
	// The .npmrc file.
	// Experimental.
	Npmrc() javascript.NpmConfig
	// Absolute output directory of this project.
	// Experimental.
	Outdir() *string
	// API for managing the node package.
	// Experimental.
	Package() javascript.NodePackage
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager() javascript.NodePackageManager
	// Experimental.
	PackageTask() projen.Task
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() projen.Project
	// Experimental.
	PostCompileTask() projen.Task
	// Experimental.
	PreCompileTask() projen.Task
	// Experimental.
	Prettier() javascript.Prettier
	// Manages the build process of the project.
	// Experimental.
	ProjectBuild() projen.ProjectBuild
	// Deprecated.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand() *string
	// Package publisher.
	//
	// This will be `undefined` if the project does not have a
	// release workflow.
	// Deprecated: use `release.publisher`.
	Publisher() release.Publisher
	// Release management.
	// Experimental.
	Release() release.Release
	// The root project.
	// Experimental.
	Root() projen.Project
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand() *string
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir() *string
	// Returns all the subprojects within this project.
	// Experimental.
	Subprojects() *[]projen.Project
	// Project tasks.
	// Experimental.
	Tasks() projen.Tasks
	// The directory in which tests reside.
	// Experimental.
	Testdir() *string
	// Experimental.
	TestTask() projen.Task
	// Experimental.
	Tsconfig() javascript.TypescriptConfig
	// A typescript configuration file which covers all files (sources, tests, projen).
	// Experimental.
	TsconfigDev() javascript.TypescriptConfig
	// Experimental.
	TsconfigEslint() javascript.TypescriptConfig
	// The upgrade workflow.
	// Experimental.
	UpgradeWorkflow() javascript.UpgradeDependencies
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Vscode() vscode.VsCode
	// The "watch" task.
	// Experimental.
	WatchTask() projen.Task
	// Experimental.
	WorkflowBootstrapSteps() *[]*workflows.JobStep
	// Experimental.
	WorkflowPackageCache() *bool
	// Experimental.
	AddBins(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Experimental.
	AddBundledDeps(deps ...*string)
	// DEPRECATED.
	// Deprecated: use `project.compileTask.exec()`
	AddCompileCommand(commands ...*string)
	// Defines normal dependencies.
	// Experimental.
	AddDeps(deps ...*string)
	// Defines development/test dependencies.
	// Experimental.
	AddDevDeps(deps ...*string)
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Experimental.
	AddExcludeFromCleanup(globs ...*string)
	// Directly set fields in `package.json`.
	// Experimental.
	AddFields(fields *map[string]interface{})
	// Adds a .gitignore pattern.
	// Experimental.
	AddGitIgnore(pattern *string)
	// Adds keywords to package.json (deduplicated).
	// Experimental.
	AddKeywords(keywords ...*string)
	// Adds patterns to be ignored by npm.
	// Experimental.
	AddPackageIgnore(pattern *string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Experimental.
	AddPeerDeps(deps ...*string)
	// Replaces the contents of multiple npm package.json scripts.
	// Experimental.
	AddScripts(scripts *map[string]*string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Experimental.
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
	// Deprecated.
	AnnotateGenerated(glob *string)
	// Indicates if a script by the name name is defined.
	// Deprecated: Use `project.tasks.tryFind(name)`
	HasScript(name *string) *bool
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before all components are synthesized.
	// Experimental.
	PreSynthesize()
	// Removes the npm script (always successful).
	// Experimental.
	RemoveScript(name *string)
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Experimental.
	RemoveTask(name *string) projen.Task
	// Returns the set of workflow steps which should be executed to bootstrap a workflow.
	//
	// Returns: Job steps.
	// Experimental.
	RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep
	// Returns the shell command to execute in order to run a task.
	//
	// This will
	// typically be `npx projen TASK`.
	// Experimental.
	RunTaskCommand(task projen.Task) *string
	// Replaces the contents of an npm package.json script.
	// Experimental.
	SetScript(name *string, command *string)
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

// The jsii proxy struct for Cdk8sTypeScriptApp
type jsiiProxy_Cdk8sTypeScriptApp struct {
	internal.Type__typescriptTypeScriptAppProject
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) AppEntrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Cdk8sDeps() Cdk8sDeps {
	var returns Cdk8sDeps
	_jsii_.Get(
		j,
		"cdk8sDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Npmrc() javascript.NpmConfig {
	var returns javascript.NpmConfig
	_jsii_.Get(
		j,
		"npmrc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Prettier() javascript.Prettier {
	var returns javascript.Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Subprojects() *[]projen.Project {
	var returns *[]projen.Project
	_jsii_.Get(
		j,
		"subprojects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) WorkflowBootstrapSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"workflowBootstrapSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sTypeScriptApp) WorkflowPackageCache() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"workflowPackageCache",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdk8sTypeScriptApp(options *Cdk8sTypeScriptAppOptions) Cdk8sTypeScriptApp {
	_init_.Initialize()

	if err := validateNewCdk8sTypeScriptAppParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cdk8sTypeScriptApp{}

	_jsii_.Create(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewCdk8sTypeScriptApp_Override(c Cdk8sTypeScriptApp, options *Cdk8sTypeScriptAppOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		[]interface{}{options},
		c,
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
func Cdk8sTypeScriptApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdk8sTypeScriptApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a project.
// Experimental.
func Cdk8sTypeScriptApp_IsProject(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdk8sTypeScriptApp_IsProjectParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk8s.Cdk8sTypeScriptApp",
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
func Cdk8sTypeScriptApp_Of(construct constructs.IConstruct) projen.Project {
	_init_.Initialize()

	if err := validateCdk8sTypeScriptApp_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns projen.Project

	_jsii_.StaticInvoke(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Cdk8sTypeScriptApp_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func Cdk8sTypeScriptApp_DEFAULT_TS_JEST_TRANFORM_PATTERN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.cdk8s.Cdk8sTypeScriptApp",
		"DEFAULT_TS_JEST_TRANFORM_PATTERN",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddBins(bins *map[string]*string) {
	if err := c.validateAddBinsParameters(bins); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBins",
		[]interface{}{bins},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddBundledDeps(deps ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddCompileCommand(commands ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddDeps(deps ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddDevDeps(deps ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddExcludeFromCleanup(globs ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddFields(fields *map[string]interface{}) {
	if err := c.validateAddFieldsParameters(fields); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addFields",
		[]interface{}{fields},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddGitIgnore(pattern *string) {
	if err := c.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddKeywords(keywords ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddPackageIgnore(pattern *string) {
	if err := c.validateAddPackageIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddPeerDeps(deps ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddScripts(scripts *map[string]*string) {
	if err := c.validateAddScriptsParameters(scripts); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addScripts",
		[]interface{}{scripts},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddTask(name *string, props *projen.TaskOptions) projen.Task {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddTestCommand(commands ...*string) {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) AddTip(message *string) {
	if err := c.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addTip",
		[]interface{}{message},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) AnnotateGenerated(glob *string) {
	if err := c.validateAnnotateGeneratedParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) HasScript(name *string) *bool {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) RemoveScript(name *string) {
	if err := c.validateRemoveScriptParameters(name); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeScript",
		[]interface{}{name},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) RemoveTask(name *string) projen.Task {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) RunTaskCommand(task projen.Task) *string {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) SetScript(name *string, command *string) {
	if err := c.validateSetScriptParameters(name, command); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"setScript",
		[]interface{}{name, command},
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) Synth() {
	_jsii_.InvokeVoid(
		c,
		"synth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cdk8sTypeScriptApp) TryFindFile(filePath *string) projen.FileBase {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) TryFindJsonFile(filePath *string) projen.JsonFile {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) TryFindObjectFile(filePath *string) projen.ObjectFile {
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

func (c *jsiiProxy_Cdk8sTypeScriptApp) TryRemoveFile(filePath *string) projen.FileBase {
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

