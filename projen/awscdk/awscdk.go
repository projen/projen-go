package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/build"
	"github.com/projen/projen-go/projen/cdk"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/java"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/python"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/typescript"
	"github.com/projen/projen-go/projen/vscode"
)

// Which approval is required when deploying CDK apps.
// Experimental.
type ApprovalLevel string

const (
	// Approval is never required.
	// Experimental.
	ApprovalLevel_NEVER ApprovalLevel = "NEVER"
	// Requires approval on any IAM or security-group-related change.
	// Experimental.
	ApprovalLevel_ANY_CHANGE ApprovalLevel = "ANY_CHANGE"
	// Requires approval when IAM statements or traffic rules are added;
	//
	// removals don't require approval.
	// Experimental.
	ApprovalLevel_BROADENING ApprovalLevel = "BROADENING"
)

// Discovers and creates integration tests and lambdas from code in the project's source and test trees.
// Experimental.
type AutoDiscover interface {
	projen.Component
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

// The jsii proxy struct for AutoDiscover
type jsiiProxy_AutoDiscover struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoDiscover(project projen.Project, options *AutoDiscoverOptions) AutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_AutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.AutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoDiscover_Override(a AutoDiscover, project projen.Project, options *AutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AutoDiscover",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Common options for auto discovering project subcomponents.
// Experimental.
type AutoDiscoverCommonOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
}

// Options for `AutoDiscover`.
// Experimental.
type AutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Project source tree (relative to project output directory).
	// Experimental.
	Srcdir *string `field:"required" json:"srcdir" yaml:"srcdir"`
	// Options for AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
	// Options for lambda extensions.
	// Experimental.
	LambdaExtensionOptions *LambdaExtensionCommonOptions `field:"optional" json:"lambdaExtensionOptions" yaml:"lambdaExtensionOptions"`
	// Test source tree.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
	// Options for integration tests.
	// Experimental.
	IntegrationTestOptions *IntegrationTestCommonOptions `field:"optional" json:"integrationTestOptions" yaml:"integrationTestOptions"`
	// Auto-discover edge lambda functions.
	// Experimental.
	EdgeLambdaAutoDiscover *bool `field:"optional" json:"edgeLambdaAutoDiscover" yaml:"edgeLambdaAutoDiscover"`
	// Auto-discover integration tests.
	// Experimental.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
	// Auto-discover lambda functions.
	// Experimental.
	LambdaAutoDiscover *bool `field:"optional" json:"lambdaAutoDiscover" yaml:"lambdaAutoDiscover"`
	// Auto-discover lambda extensions.
	// Experimental.
	LambdaExtensionAutoDiscover *bool `field:"optional" json:"lambdaExtensionAutoDiscover" yaml:"lambdaExtensionAutoDiscover"`
}

// AWS CDK construct library project.
//
// A multi-language (jsii) construct library which vends constructs designed to
// use within the AWS CDK with a friendly workflow and automatic publishing to
// the construct catalog.
// Experimental.
type AwsCdkConstructLibrary interface {
	cdk.ConstructLibrary
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies() *bool
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
	CdkDeps() AwsCdkDeps
	// The target CDK version for this library.
	// Experimental.
	CdkVersion() *string
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
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion() *string
	// Project name.
	// Experimental.
	Name() *string
	// The .npmignore file.
	// Experimental.
	Npmignore() projen.IgnoreFile
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
	// Deprecated: use `cdkVersion`.
	Version() *string
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Vscode() vscode.VsCode
	// The "watch" task.
	// Experimental.
	WatchTask() projen.Task
	// Experimental.
	AddBins(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Experimental.
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
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Experimental.
	AddPackageIgnore(pattern *string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Experimental.
	AddPeerDeps(deps ...*string)
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
	// Experimental.
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
	// 3. Synthesize all sub-projects
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

// The jsii proxy struct for AwsCdkConstructLibrary
type jsiiProxy_AwsCdkConstructLibrary struct {
	internal.Type__cdkConstructLibrary
}

func (j *jsiiProxy_AwsCdkConstructLibrary) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) CdkDeps() AwsCdkDeps {
	var returns AwsCdkDeps
	_jsii_.Get(
		j,
		"cdkDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Prettier() javascript.Prettier {
	var returns javascript.Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkConstructLibrary) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkConstructLibrary(options *AwsCdkConstructLibraryOptions) AwsCdkConstructLibrary {
	_init_.Initialize()

	j := jsiiProxy_AwsCdkConstructLibrary{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkConstructLibrary",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkConstructLibrary_Override(a AwsCdkConstructLibrary, options *AwsCdkConstructLibraryOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkConstructLibrary",
		[]interface{}{options},
		a,
	)
}

func AwsCdkConstructLibrary_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.awscdk.AwsCdkConstructLibrary",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		a,
		"addBins",
		[]interface{}{bins},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addBundledDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddCdkDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addCdkDependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddCdkTestDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addCdkTestDependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addCompileCommand",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDevDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddExcludeFromCleanup(globs ...*string) {
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

func (a *jsiiProxy_AwsCdkConstructLibrary) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addFields",
		[]interface{}{fields},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addKeywords",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addPeerDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		a,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addTestCommand",
		args,
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AddTip(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addTip",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		a,
		"removeScript",
		[]interface{}{name},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		a,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep

	_jsii_.Invoke(
		a,
		"renderWorkflowSetup",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		a,
		"setScript",
		[]interface{}{name, command},
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkConstructLibrary) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		a,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		a,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkConstructLibrary) TryRemoveFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Options for `AwsCdkConstructLibrary`.
// Experimental.
type AwsCdkConstructLibraryOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `field:"optional" json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Experimental.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Experimental.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Experimental.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `field:"optional" json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `field:"optional" json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `field:"optional" json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `field:"optional" json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `field:"optional" json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `field:"optional" json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `field:"optional" json:"bugsUrl" yaml:"bugsUrl"`
	// List of dependencies to bundle into this module.
	//
	// These modules will be
	// added both to the `dependencies` section and `bundledDependencies` section of
	// your `package.json`.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps *[]*string `field:"optional" json:"bundledDeps" yaml:"bundledDeps"`
	// Options for npm packages using AWS CodeArtifact.
	//
	// This is required if publishing packages to, or installing scoped packages from AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *javascript.CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'express', 'lodash', 'foo@^2' ]
	//
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// Example:
	//   [ 'typescript', '@types/express' ]
	//
	// Experimental.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `field:"optional" json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `field:"optional" json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `field:"optional" json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess javascript.NpmAccess `field:"optional" json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `field:"optional" json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `field:"optional" json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `field:"optional" json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager javascript.NodePackageManager `field:"optional" json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *javascript.PeerDependencyOptions `field:"optional" json:"peerDependencyOptions" yaml:"peerDependencyOptions"`
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
	PeerDeps *[]*string `field:"optional" json:"peerDeps" yaml:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `field:"optional" json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// Options for privately hosted scoped packages.
	// Experimental.
	ScopedPackagesOptions *[]*javascript.ScopedPackagesOptions `field:"optional" json:"scopedPackagesOptions" yaml:"scopedPackagesOptions"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `field:"optional" json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `field:"optional" json:"stability" yaml:"stability"`
	// Version requirement of `publib` which is used to publish modules to npm.
	// Experimental.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Experimental.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Experimental.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Experimental.
	PublishDryRun *bool `field:"optional" json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Experimental.
	ReleaseBranches *map[string]*release.BranchOptions `field:"optional" json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `field:"optional" json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Experimental.
	ReleaseFailureIssue *bool `field:"optional" json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Experimental.
	ReleaseFailureIssueLabel *string `field:"optional" json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `field:"optional" json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	ReleaseTagPrefix *string `field:"optional" json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Experimental.
	ReleaseTrigger release.ReleaseTrigger `field:"optional" json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Experimental.
	ReleaseWorkflowName *string `field:"optional" json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `field:"optional" json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Experimental.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch *string `field:"required" json:"defaultReleaseBranch" yaml:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Experimental.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveUpgrades *bool `field:"optional" json:"autoApproveUpgrades" yaml:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow *bool `field:"optional" json:"buildWorkflow" yaml:"buildWorkflow"`
	// Build workflow triggers.
	// Experimental.
	BuildWorkflowTriggers *workflows.Triggers `field:"optional" json:"buildWorkflowTriggers" yaml:"buildWorkflowTriggers"`
	// Options for `Bundler`.
	// Experimental.
	BundlerOptions *javascript.BundlerOptions `field:"optional" json:"bundlerOptions" yaml:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov *bool `field:"optional" json:"codeCov" yaml:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret *string `field:"optional" json:"codeCovTokenSecret" yaml:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner *string `field:"optional" json:"copyrightOwner" yaml:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod *string `field:"optional" json:"copyrightPeriod" yaml:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Experimental.
	Dependabot *bool `field:"optional" json:"dependabot" yaml:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions *github.DependabotOptions `field:"optional" json:"dependabotOptions" yaml:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Experimental.
	DepsUpgrade *bool `field:"optional" json:"depsUpgrade" yaml:"depsUpgrade"`
	// Options for `UpgradeDependencies`.
	// Experimental.
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `field:"optional" json:"depsUpgradeOptions" yaml:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore *[]*string `field:"optional" json:"gitignore" yaml:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest *bool `field:"optional" json:"jest" yaml:"jest"`
	// Jest options.
	// Experimental.
	JestOptions *javascript.JestOptions `field:"optional" json:"jestOptions" yaml:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `field:"optional" json:"npmignore" yaml:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled *bool `field:"optional" json:"npmignoreEnabled" yaml:"npmignoreEnabled"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Experimental.
	Package *bool `field:"optional" json:"package" yaml:"package"`
	// Setup prettier.
	// Experimental.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Prettier options.
	// Experimental.
	PrettierOptions *javascript.PrettierOptions `field:"optional" json:"prettierOptions" yaml:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency *bool `field:"optional" json:"projenDevDependency" yaml:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Experimental.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options for .projenrc.js.
	// Experimental.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate *bool `field:"optional" json:"pullRequestTemplate" yaml:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents *[]*string `field:"optional" json:"pullRequestTemplateContents" yaml:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Experimental.
	Release *bool `field:"optional" json:"release" yaml:"release"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm *bool `field:"optional" json:"releaseToNpm" yaml:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `field:"optional" json:"releaseWorkflow" yaml:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps *[]*workflows.JobStep `field:"optional" json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Experimental.
	WorkflowGitIdentity *github.GitIdentity `field:"optional" json:"workflowGitIdentity" yaml:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig *bool `field:"optional" json:"disableTsconfig" yaml:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen *bool `field:"optional" json:"docgen" yaml:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory *string `field:"optional" json:"docsDirectory" yaml:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes *string `field:"optional" json:"entrypointTypes" yaml:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint *bool `field:"optional" json:"eslint" yaml:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions *javascript.EslintOptions `field:"optional" json:"eslintOptions" yaml:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir *string `field:"optional" json:"libdir" yaml:"libdir"`
	// Use TypeScript for your projenrc file (`.projenrc.ts`).
	// Experimental.
	ProjenrcTs *bool `field:"optional" json:"projenrcTs" yaml:"projenrcTs"`
	// Options for .projenrc.ts.
	// Experimental.
	ProjenrcTsOptions *typescript.ProjenrcOptions `field:"optional" json:"projenrcTsOptions" yaml:"projenrcTsOptions"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode *bool `field:"optional" json:"sampleCode" yaml:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir *string `field:"optional" json:"srcdir" yaml:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	//
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Custom tsconfig options for the development tsconfig.json file (used for testing).
	// Experimental.
	TsconfigDev *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfigDev" yaml:"tsconfigDev"`
	// The name of the development tsconfig.json file.
	// Experimental.
	TsconfigDevFile *string `field:"optional" json:"tsconfigDevFile" yaml:"tsconfigDevFile"`
	// TypeScript version to use.
	//
	// NOTE: Typescript is not semantically versioned and should remain on the
	// same minor, so we recommend using a `~` dependency (e.g. `~1.2.3`).
	// Experimental.
	TypescriptVersion *string `field:"optional" json:"typescriptVersion" yaml:"typescriptVersion"`
	// The name of the library author.
	// Experimental.
	Author *string `field:"required" json:"author" yaml:"author"`
	// Email or URL of the library author.
	// Experimental.
	AuthorAddress *string `field:"required" json:"authorAddress" yaml:"authorAddress"`
	// Git repository URL.
	// Experimental.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	//
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Experimental.
	Compat *bool `field:"optional" json:"compat" yaml:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Experimental.
	CompatIgnore *string `field:"optional" json:"compatIgnore" yaml:"compatIgnore"`
	// File path for generated docs.
	// Experimental.
	DocgenFilePath *string `field:"optional" json:"docgenFilePath" yaml:"docgenFilePath"`
	// Deprecated: use `publishToNuget`.
	Dotnet *cdk.JsiiDotNetTarget `field:"optional" json:"dotnet" yaml:"dotnet"`
	// Accepts a list of glob patterns.
	//
	// Files matching any of those patterns will be excluded from the TypeScript compiler input.
	//
	// By default, jsii will include all *.ts files (except .d.ts files) in the TypeScript compiler input.
	// This can be problematic for example when the package's build or test procedure generates .ts files
	// that cannot be compiled with jsii's compiler settings.
	// Experimental.
	ExcludeTypescript *[]*string `field:"optional" json:"excludeTypescript" yaml:"excludeTypescript"`
	// Publish Go bindings to a git repository.
	// Experimental.
	PublishToGo *cdk.JsiiGoTarget `field:"optional" json:"publishToGo" yaml:"publishToGo"`
	// Publish to maven.
	// Experimental.
	PublishToMaven *cdk.JsiiJavaTarget `field:"optional" json:"publishToMaven" yaml:"publishToMaven"`
	// Publish to NuGet.
	// Experimental.
	PublishToNuget *cdk.JsiiDotNetTarget `field:"optional" json:"publishToNuget" yaml:"publishToNuget"`
	// Publish to pypi.
	// Experimental.
	PublishToPypi *cdk.JsiiPythonTarget `field:"optional" json:"publishToPypi" yaml:"publishToPypi"`
	// Deprecated: use `publishToPyPi`.
	Python *cdk.JsiiPythonTarget `field:"optional" json:"python" yaml:"python"`
	// Experimental.
	Rootdir *string `field:"optional" json:"rootdir" yaml:"rootdir"`
	// Libraries will be picked up by the construct catalog when they are published to npm as jsii modules and will be published under:.
	//
	// https://awscdk.io/packages/[@SCOPE/]PACKAGE@VERSION
	//
	// The catalog will also post a tweet to https://twitter.com/awscdkio with the
	// package name, description and the above link. You can disable these tweets
	// through `{ announce: false }`.
	//
	// You can also add a Twitter handle through `{ twitter: 'xx' }` which will be
	// mentioned in the tweet.
	// See: https://github.com/construct-catalog/catalog
	//
	// Experimental.
	Catalog *cdk.Catalog `field:"optional" json:"catalog" yaml:"catalog"`
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// Automatically adds an `cloudfront.experimental.EdgeFunction` for each `.edge-lambda.ts` handler in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Experimental.
	EdgeLambdaAutoDiscover *bool `field:"optional" json:"edgeLambdaAutoDiscover" yaml:"edgeLambdaAutoDiscover"`
	// Automatically discovers and creates integration tests for each `.integ.ts` file in under your test directory.
	// Experimental.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
	// Automatically adds an `aws_lambda.Function` for each `.lambda.ts` handler in your source tree. If this is disabled, you either need to explicitly call `aws_lambda.Function.autoDiscover()` or define a `new aws_lambda.Function()` for each handler.
	// Experimental.
	LambdaAutoDiscover *bool `field:"optional" json:"lambdaAutoDiscover" yaml:"lambdaAutoDiscover"`
	// Automatically adds an `awscdk.LambdaExtension` for each `.lambda-extension.ts` entrypoint in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Experimental.
	LambdaExtensionAutoDiscover *bool `field:"optional" json:"lambdaExtensionAutoDiscover" yaml:"lambdaExtensionAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
}

// Manages dependencies on the AWS CDK.
// Experimental.
type AwsCdkDeps interface {
	projen.Component
	// Whether CDK dependencies are added as normal dependencies (and peer dependencies).
	// Deprecated: Not used for CDK 2.x
	CdkDependenciesAsDeps() *bool
	// The major version of the AWS CDK (e.g. 1, 2, ...).
	// Experimental.
	CdkMajorVersion() *float64
	// The minimum version of the AWS CDK (e.g. `2.0.0`).
	// Experimental.
	CdkMinimumVersion() *string
	// The dependency requirement for AWS CDK (e.g. `^2.0.0`).
	// Experimental.
	CdkVersion() *string
	// Experimental.
	Project() projen.Project
	// Adds dependencies to AWS CDK modules.
	//
	// The type of dependency is determined by the `dependencyType` option.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1Dependencies(deps ...*string)
	// Adds AWS CDK modules as dev dependencies.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1DevDependencies(deps ...*string)
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *AwsCdkPackageNames
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

// The jsii proxy struct for AwsCdkDeps
type jsiiProxy_AwsCdkDeps struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AwsCdkDeps) CdkDependenciesAsDeps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkDependenciesAsDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) CdkMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdkMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) CdkMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkDeps_Override(a AwsCdkDeps, project projen.Project, options *AwsCdkDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkDeps",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AwsCdkDeps) AddV1Dependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1Dependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDeps) AddV1DevDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1DevDependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDeps) PackageNames() *AwsCdkPackageNames {
	var returns *AwsCdkPackageNames

	_jsii_.Invoke(
		a,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDeps) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDeps) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDeps) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `AwsCdkDeps`.
// Experimental.
type AwsCdkDepsCommonOptions struct {
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
}

// Manages dependencies on the AWS CDK for Java projects.
// Experimental.
type AwsCdkDepsJava interface {
	AwsCdkDeps
	// Whether CDK dependencies are added as normal dependencies (and peer dependencies).
	// Deprecated: Not used for CDK 2.x
	CdkDependenciesAsDeps() *bool
	// The major version of the AWS CDK (e.g. 1, 2, ...).
	// Experimental.
	CdkMajorVersion() *float64
	// The minimum version of the AWS CDK (e.g. `2.0.0`).
	// Experimental.
	CdkMinimumVersion() *string
	// The dependency requirement for AWS CDK (e.g. `^2.0.0`).
	// Experimental.
	CdkVersion() *string
	// Experimental.
	Project() projen.Project
	// Adds dependencies to AWS CDK modules.
	//
	// The type of dependency is determined by the `dependencyType` option.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1Dependencies(deps ...*string)
	// Adds AWS CDK modules as dev dependencies.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1DevDependencies(deps ...*string)
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *AwsCdkPackageNames
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

// The jsii proxy struct for AwsCdkDepsJava
type jsiiProxy_AwsCdkDepsJava struct {
	jsiiProxy_AwsCdkDeps
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkDependenciesAsDeps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkDependenciesAsDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdkMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkDepsJava(project projen.Project, options *AwsCdkDepsOptions) AwsCdkDepsJava {
	_init_.Initialize()

	j := jsiiProxy_AwsCdkDepsJava{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsJava",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkDepsJava_Override(a AwsCdkDepsJava, project projen.Project, options *AwsCdkDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsJava",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) AddV1Dependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1Dependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) AddV1DevDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1DevDependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) PackageNames() *AwsCdkPackageNames {
	var returns *AwsCdkPackageNames

	_jsii_.Invoke(
		a,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDepsJava) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Manages dependencies on the AWS CDK for Node.js projects.
// Experimental.
type AwsCdkDepsJs interface {
	AwsCdkDeps
	// Whether CDK dependencies are added as normal dependencies (and peer dependencies).
	// Deprecated: Not used for CDK 2.x
	CdkDependenciesAsDeps() *bool
	// The major version of the AWS CDK (e.g. 1, 2, ...).
	// Experimental.
	CdkMajorVersion() *float64
	// The minimum version of the AWS CDK (e.g. `2.0.0`).
	// Experimental.
	CdkMinimumVersion() *string
	// The dependency requirement for AWS CDK (e.g. `^2.0.0`).
	// Experimental.
	CdkVersion() *string
	// Experimental.
	Project() projen.Project
	// Adds dependencies to AWS CDK modules.
	//
	// The type of dependency is determined by the `dependencyType` option.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1Dependencies(deps ...*string)
	// Adds AWS CDK modules as dev dependencies.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1DevDependencies(deps ...*string)
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *AwsCdkPackageNames
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

// The jsii proxy struct for AwsCdkDepsJs
type jsiiProxy_AwsCdkDepsJs struct {
	jsiiProxy_AwsCdkDeps
}

func (j *jsiiProxy_AwsCdkDepsJs) CdkDependenciesAsDeps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkDependenciesAsDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJs) CdkMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdkMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJs) CdkMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJs) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJs) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkDepsJs(project projen.Project, options *AwsCdkDepsOptions) AwsCdkDepsJs {
	_init_.Initialize()

	j := jsiiProxy_AwsCdkDepsJs{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsJs",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkDepsJs_Override(a AwsCdkDepsJs, project projen.Project, options *AwsCdkDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsJs",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AwsCdkDepsJs) AddV1Dependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1Dependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDepsJs) AddV1DevDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1DevDependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDepsJs) PackageNames() *AwsCdkPackageNames {
	var returns *AwsCdkPackageNames

	_jsii_.Invoke(
		a,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDepsJs) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJs) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJs) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type AwsCdkDepsOptions struct {
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// The type of dependency to use for runtime AWS CDK and `constructs` modules.
	//
	// For libraries, use peer dependencies and for apps use runtime dependencies.
	// Experimental.
	DependencyType projen.DependencyType `field:"required" json:"dependencyType" yaml:"dependencyType"`
}

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
	// 3. Synthesize all sub-projects
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
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddPlugin(spec *string, options *java.PluginOptions) *projen.Dependency {
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
	_jsii_.InvokeVoid(
		a,
		"addTestDependency",
		[]interface{}{spec},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AddTip(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addTip",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_AwsCdkJavaApp) AnnotateGenerated(glob *string) {
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
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type AwsCdkJavaAppOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `field:"optional" json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Experimental.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Experimental.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Experimental.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// The artifactId is generally the name that the project is known by.
	//
	// Although
	// the groupId is important, people within the group will rarely mention the
	// groupId in discussion (they are often all be the same ID, such as the
	// MojoHaus project groupId: org.codehaus.mojo). It, along with the groupId,
	// creates a key that separates this project from every other project in the
	// world (at least, it should :) ). Along with the groupId, the artifactId
	// fully defines the artifact's living quarters within the repository. In the
	// case of the above project, my-project lives in
	// $M2_REPO/org/codehaus/mojo/my-project.
	// Experimental.
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// This is generally unique amongst an organization or a project.
	//
	// For example,
	// all core Maven artifacts do (well, should) live under the groupId
	// org.apache.maven. Group ID's do not necessarily use the dot notation, for
	// example, the junit project. Note that the dot-notated groupId does not have
	// to correspond to the package structure that the project contains. It is,
	// however, a good practice to follow. When stored within a repository, the
	// group acts much like the Java packaging structure does in an operating
	// system. The dots are replaced by OS specific directory separators (such as
	// '/' in Unix) which becomes a relative directory structure from the base
	// repository. In the example given, the org.codehaus.mojo group lives within
	// the directory $M2_REPO/org/codehaus/mojo.
	// Experimental.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// This is the last piece of the naming puzzle.
	//
	// groupId:artifactId denotes a
	// single project but they cannot delineate which incarnation of that project
	// we are talking about. Do we want the junit:junit of 2018 (version 4.12), or
	// of 2007 (version 3.8.2)? In short: code changes, those changes should be
	// versioned, and this element keeps those versions in line. It is also used
	// within an artifact's repository to separate versions from each other.
	// my-project version 1.0 files live in the directory structure
	// $M2_REPO/org/codehaus/mojo/my-project/1.0.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Project packaging format.
	// Experimental.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Compile options.
	// Experimental.
	CompileOptions *java.MavenCompileOptions `field:"optional" json:"compileOptions" yaml:"compileOptions"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// Final artifact output directory.
	// Experimental.
	Distdir *string `field:"optional" json:"distdir" yaml:"distdir"`
	// Include junit tests.
	// Experimental.
	Junit *bool `field:"optional" json:"junit" yaml:"junit"`
	// junit options.
	// Experimental.
	JunitOptions *java.JunitOptions `field:"optional" json:"junitOptions" yaml:"junitOptions"`
	// Packaging options.
	// Experimental.
	PackagingOptions *java.MavenPackagingOptions `field:"optional" json:"packagingOptions" yaml:"packagingOptions"`
	// Use projenrc in java.
	//
	// This will install `projen` as a java dependency and will add a `synth` task which
	// will compile & execute `main()` from `src/main/java/projenrc.java`.
	// Experimental.
	ProjenrcJava *bool `field:"optional" json:"projenrcJava" yaml:"projenrcJava"`
	// Options related to projenrc in java.
	// Experimental.
	ProjenrcJavaOptions *java.ProjenrcOptions `field:"optional" json:"projenrcJavaOptions" yaml:"projenrcJavaOptions"`
	// List of test dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addTestDependency()`.
	// Experimental.
	TestDeps *[]*string `field:"optional" json:"testDeps" yaml:"testDeps"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample *bool `field:"optional" json:"sample" yaml:"sample"`
	// The java package to use for the code sample.
	// Experimental.
	SampleJavaPackage *string `field:"optional" json:"sampleJavaPackage" yaml:"sampleJavaPackage"`
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// The name of the Java class with the static `main()` method.
	//
	// This method
	// should call `app.synth()` on the CDK app.
	// Experimental.
	MainClass *string `field:"required" json:"mainClass" yaml:"mainClass"`
}

// Language-specific AWS CDK package names.
// Experimental.
type AwsCdkPackageNames struct {
	// Fully qualified name of the assertions library package.
	// Experimental.
	Assertions *string `field:"required" json:"assertions" yaml:"assertions"`
	// Fully qualified name of the constructs library package.
	// Experimental.
	Constructs *string `field:"required" json:"constructs" yaml:"constructs"`
	// Fully qualified name of the core framework package for CDKv1.
	// Experimental.
	CoreV1 *string `field:"required" json:"coreV1" yaml:"coreV1"`
	// Fully qualified name of the core framework package for CDKv2.
	// Experimental.
	CoreV2 *string `field:"required" json:"coreV2" yaml:"coreV2"`
	// Fully qualified name of the assert library package Can be empty as it's only really available for javascript projects.
	// Experimental.
	Assert *string `field:"optional" json:"assert" yaml:"assert"`
}

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
	// Project tasks.
	// Experimental.
	Tasks() projen.Tasks
	// The directory in which the python tests reside.
	// Experimental.
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
	// 3. Synthesize all sub-projects
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

func (j *jsiiProxy_AwsCdkPythonApp) SetPytest(val python.Pytest) {
	_jsii_.Set(
		j,
		"pytest",
		val,
	)
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
	_jsii_.InvokeVoid(
		a,
		"addDependency",
		[]interface{}{spec},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddDevDependency(spec *string) {
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
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AddTask(name *string, props *projen.TaskOptions) projen.Task {
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
	_jsii_.InvokeVoid(
		a,
		"addTip",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_AwsCdkPythonApp) AnnotateGenerated(glob *string) {
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

func (a *jsiiProxy_AwsCdkPythonApp) TryFindFile(filePath *string) projen.FileBase {
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
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Options for `AwsCdkPythonApp`.
// Experimental.
type AwsCdkPythonAppOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `field:"optional" json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Experimental.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Experimental.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Experimental.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `field:"required" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `field:"required" json:"authorName" yaml:"authorName"`
	// Version of the package.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *python.PoetryPyprojectOptionsWithoutDeps `field:"optional" json:"poetryOptions" yaml:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `field:"optional" json:"setupConfig" yaml:"setupConfig"`
	// Name of the python package as used in imports and filenames.
	//
	// Must only consist of alphanumeric characters and underscores.
	// Experimental.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// List of dev dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDevDependency()`.
	// Experimental.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Use pip with a requirements.txt file to track project dependencies.
	// Experimental.
	Pip *bool `field:"optional" json:"pip" yaml:"pip"`
	// Use poetry to manage your project dependencies, virtual environment, and (optional) packaging/publishing.
	// Experimental.
	Poetry *bool `field:"optional" json:"poetry" yaml:"poetry"`
	// Use projenrc in javascript.
	//
	// This will install `projen` as a JavaScript dependency and add a `synth`
	// task which will run `.projenrc.js`.
	// Experimental.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options related to projenrc in JavaScript.
	// Experimental.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Use projenrc in Python.
	//
	// This will install `projen` as a Python dependency and add a `synth`
	// task which will run `.projenrc.py`.
	// Experimental.
	ProjenrcPython *bool `field:"optional" json:"projenrcPython" yaml:"projenrcPython"`
	// Options related to projenrc in python.
	// Experimental.
	ProjenrcPythonOptions *python.ProjenrcOptions `field:"optional" json:"projenrcPythonOptions" yaml:"projenrcPythonOptions"`
	// Include pytest tests.
	// Experimental.
	Pytest *bool `field:"optional" json:"pytest" yaml:"pytest"`
	// pytest options.
	// Experimental.
	PytestOptions *python.PytestOptions `field:"optional" json:"pytestOptions" yaml:"pytestOptions"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample *bool `field:"optional" json:"sample" yaml:"sample"`
	// Use setuptools with a setup.py script for packaging and publishing.
	// Experimental.
	Setuptools *bool `field:"optional" json:"setuptools" yaml:"setuptools"`
	// Use venv to manage a virtual environment for installing dependencies inside.
	// Experimental.
	Venv *bool `field:"optional" json:"venv" yaml:"venv"`
	// Venv options.
	// Experimental.
	VenvOptions *python.VenvOptions `field:"optional" json:"venvOptions" yaml:"venvOptions"`
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// The CDK app's entrypoint (relative to the source directory, which is "src" by default).
	// Experimental.
	AppEntrypoint *string `field:"optional" json:"appEntrypoint" yaml:"appEntrypoint"`
	// Python sources directory.
	// Experimental.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
}

// AWS CDK app in TypeScript.
// Experimental.
type AwsCdkTypeScriptApp interface {
	typescript.TypeScriptAppProject
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies() *bool
	// The CDK app entrypoint.
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
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion() *string
	// Project name.
	// Experimental.
	Name() *string
	// The .npmignore file.
	// Experimental.
	Npmignore() projen.IgnoreFile
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
	AddBins(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Experimental.
	AddBundledDeps(deps ...*string)
	// Adds an AWS CDK module dependencies.
	// Experimental.
	AddCdkDependency(modules ...*string)
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
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Experimental.
	AddPackageIgnore(pattern *string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Experimental.
	AddPeerDeps(deps ...*string)
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
	// Experimental.
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
	// 3. Synthesize all sub-projects
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

// The jsii proxy struct for AwsCdkTypeScriptApp
type jsiiProxy_AwsCdkTypeScriptApp struct {
	internal.Type__typescriptTypeScriptAppProject
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) AppEntrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) CdkConfig() CdkConfig {
	var returns CdkConfig
	_jsii_.Get(
		j,
		"cdkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) CdkDeps() AwsCdkDeps {
	var returns AwsCdkDeps
	_jsii_.Get(
		j,
		"cdkDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) CdkTasks() CdkTasks {
	var returns CdkTasks
	_jsii_.Get(
		j,
		"cdkTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Prettier() javascript.Prettier {
	var returns javascript.Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkTypeScriptApp) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkTypeScriptApp(options *AwsCdkTypeScriptAppOptions) AwsCdkTypeScriptApp {
	_init_.Initialize()

	j := jsiiProxy_AwsCdkTypeScriptApp{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkTypeScriptApp",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkTypeScriptApp_Override(a AwsCdkTypeScriptApp, options *AwsCdkTypeScriptAppOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkTypeScriptApp",
		[]interface{}{options},
		a,
	)
}

func AwsCdkTypeScriptApp_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.awscdk.AwsCdkTypeScriptApp",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		a,
		"addBins",
		[]interface{}{bins},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addBundledDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddCdkDependency(modules ...*string) {
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

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addCompileCommand",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDevDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddExcludeFromCleanup(globs ...*string) {
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

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addFields",
		[]interface{}{fields},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addKeywords",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addPeerDeps",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		a,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addTestCommand",
		args,
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AddTip(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addTip",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		a,
		"removeScript",
		[]interface{}{name},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		a,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) RenderWorkflowSetup(options *javascript.RenderWorkflowSetupOptions) *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep

	_jsii_.Invoke(
		a,
		"renderWorkflowSetup",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		a,
		"setScript",
		[]interface{}{name, command},
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		a,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		a,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkTypeScriptApp) TryRemoveFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		a,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type AwsCdkTypeScriptAppOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `field:"optional" json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Experimental.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Experimental.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Experimental.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `field:"optional" json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `field:"optional" json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `field:"optional" json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `field:"optional" json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `field:"optional" json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `field:"optional" json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `field:"optional" json:"bugsUrl" yaml:"bugsUrl"`
	// List of dependencies to bundle into this module.
	//
	// These modules will be
	// added both to the `dependencies` section and `bundledDependencies` section of
	// your `package.json`.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps *[]*string `field:"optional" json:"bundledDeps" yaml:"bundledDeps"`
	// Options for npm packages using AWS CodeArtifact.
	//
	// This is required if publishing packages to, or installing scoped packages from AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *javascript.CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'express', 'lodash', 'foo@^2' ]
	//
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// Example:
	//   [ 'typescript', '@types/express' ]
	//
	// Experimental.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `field:"optional" json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `field:"optional" json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `field:"optional" json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess javascript.NpmAccess `field:"optional" json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `field:"optional" json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `field:"optional" json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `field:"optional" json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager javascript.NodePackageManager `field:"optional" json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *javascript.PeerDependencyOptions `field:"optional" json:"peerDependencyOptions" yaml:"peerDependencyOptions"`
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
	PeerDeps *[]*string `field:"optional" json:"peerDeps" yaml:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `field:"optional" json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// Options for privately hosted scoped packages.
	// Experimental.
	ScopedPackagesOptions *[]*javascript.ScopedPackagesOptions `field:"optional" json:"scopedPackagesOptions" yaml:"scopedPackagesOptions"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `field:"optional" json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `field:"optional" json:"stability" yaml:"stability"`
	// Version requirement of `publib` which is used to publish modules to npm.
	// Experimental.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Experimental.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Experimental.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Experimental.
	PublishDryRun *bool `field:"optional" json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Experimental.
	ReleaseBranches *map[string]*release.BranchOptions `field:"optional" json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `field:"optional" json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Experimental.
	ReleaseFailureIssue *bool `field:"optional" json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Experimental.
	ReleaseFailureIssueLabel *string `field:"optional" json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `field:"optional" json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	ReleaseTagPrefix *string `field:"optional" json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Experimental.
	ReleaseTrigger release.ReleaseTrigger `field:"optional" json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Experimental.
	ReleaseWorkflowName *string `field:"optional" json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `field:"optional" json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Experimental.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch *string `field:"required" json:"defaultReleaseBranch" yaml:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Experimental.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveUpgrades *bool `field:"optional" json:"autoApproveUpgrades" yaml:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow *bool `field:"optional" json:"buildWorkflow" yaml:"buildWorkflow"`
	// Build workflow triggers.
	// Experimental.
	BuildWorkflowTriggers *workflows.Triggers `field:"optional" json:"buildWorkflowTriggers" yaml:"buildWorkflowTriggers"`
	// Options for `Bundler`.
	// Experimental.
	BundlerOptions *javascript.BundlerOptions `field:"optional" json:"bundlerOptions" yaml:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov *bool `field:"optional" json:"codeCov" yaml:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret *string `field:"optional" json:"codeCovTokenSecret" yaml:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner *string `field:"optional" json:"copyrightOwner" yaml:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod *string `field:"optional" json:"copyrightPeriod" yaml:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Experimental.
	Dependabot *bool `field:"optional" json:"dependabot" yaml:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions *github.DependabotOptions `field:"optional" json:"dependabotOptions" yaml:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Experimental.
	DepsUpgrade *bool `field:"optional" json:"depsUpgrade" yaml:"depsUpgrade"`
	// Options for `UpgradeDependencies`.
	// Experimental.
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `field:"optional" json:"depsUpgradeOptions" yaml:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore *[]*string `field:"optional" json:"gitignore" yaml:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest *bool `field:"optional" json:"jest" yaml:"jest"`
	// Jest options.
	// Experimental.
	JestOptions *javascript.JestOptions `field:"optional" json:"jestOptions" yaml:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `field:"optional" json:"npmignore" yaml:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled *bool `field:"optional" json:"npmignoreEnabled" yaml:"npmignoreEnabled"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Experimental.
	Package *bool `field:"optional" json:"package" yaml:"package"`
	// Setup prettier.
	// Experimental.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Prettier options.
	// Experimental.
	PrettierOptions *javascript.PrettierOptions `field:"optional" json:"prettierOptions" yaml:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency *bool `field:"optional" json:"projenDevDependency" yaml:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Experimental.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options for .projenrc.js.
	// Experimental.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate *bool `field:"optional" json:"pullRequestTemplate" yaml:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents *[]*string `field:"optional" json:"pullRequestTemplateContents" yaml:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Experimental.
	Release *bool `field:"optional" json:"release" yaml:"release"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm *bool `field:"optional" json:"releaseToNpm" yaml:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `field:"optional" json:"releaseWorkflow" yaml:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps *[]*workflows.JobStep `field:"optional" json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Experimental.
	WorkflowGitIdentity *github.GitIdentity `field:"optional" json:"workflowGitIdentity" yaml:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig *bool `field:"optional" json:"disableTsconfig" yaml:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen *bool `field:"optional" json:"docgen" yaml:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory *string `field:"optional" json:"docsDirectory" yaml:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes *string `field:"optional" json:"entrypointTypes" yaml:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint *bool `field:"optional" json:"eslint" yaml:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions *javascript.EslintOptions `field:"optional" json:"eslintOptions" yaml:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir *string `field:"optional" json:"libdir" yaml:"libdir"`
	// Use TypeScript for your projenrc file (`.projenrc.ts`).
	// Experimental.
	ProjenrcTs *bool `field:"optional" json:"projenrcTs" yaml:"projenrcTs"`
	// Options for .projenrc.ts.
	// Experimental.
	ProjenrcTsOptions *typescript.ProjenrcOptions `field:"optional" json:"projenrcTsOptions" yaml:"projenrcTsOptions"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode *bool `field:"optional" json:"sampleCode" yaml:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir *string `field:"optional" json:"srcdir" yaml:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	//
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Custom tsconfig options for the development tsconfig.json file (used for testing).
	// Experimental.
	TsconfigDev *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfigDev" yaml:"tsconfigDev"`
	// The name of the development tsconfig.json file.
	// Experimental.
	TsconfigDevFile *string `field:"optional" json:"tsconfigDevFile" yaml:"tsconfigDevFile"`
	// TypeScript version to use.
	//
	// NOTE: Typescript is not semantically versioned and should remain on the
	// same minor, so we recommend using a `~` dependency (e.g. `~1.2.3`).
	// Experimental.
	TypescriptVersion *string `field:"optional" json:"typescriptVersion" yaml:"typescriptVersion"`
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// The CDK app's entrypoint (relative to the source directory, which is "src" by default).
	// Experimental.
	AppEntrypoint *string `field:"optional" json:"appEntrypoint" yaml:"appEntrypoint"`
	// Automatically adds an `cloudfront.experimental.EdgeFunction` for each `.edge-lambda.ts` handler in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Experimental.
	EdgeLambdaAutoDiscover *bool `field:"optional" json:"edgeLambdaAutoDiscover" yaml:"edgeLambdaAutoDiscover"`
	// Automatically discovers and creates integration tests for each `.integ.ts` file in under your test directory.
	// Experimental.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
	// Automatically adds an `awscdk.LambdaFunction` for each `.lambda.ts` handler in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Experimental.
	LambdaAutoDiscover *bool `field:"optional" json:"lambdaAutoDiscover" yaml:"lambdaAutoDiscover"`
	// Automatically adds an `awscdk.LambdaExtension` for each `.lambda-extension.ts` entrypoint in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Experimental.
	LambdaExtensionAutoDiscover *bool `field:"optional" json:"lambdaExtensionAutoDiscover" yaml:"lambdaExtensionAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
}

// Represents cdk.json file.
// Experimental.
type CdkConfig interface {
	projen.Component
	// Name of the cdk.out directory.
	// Experimental.
	Cdkout() *string
	// List of glob patterns to be excluded by CDK.
	// Experimental.
	Exclude() *[]*string
	// List of glob patterns to be included by CDK.
	// Experimental.
	Include() *[]*string
	// Represents the JSON file.
	// Experimental.
	Json() projen.JsonFile
	// Experimental.
	Project() projen.Project
	// Add excludes to `cdk.json`.
	// Experimental.
	AddExcludes(patterns ...*string)
	// Add includes to `cdk.json`.
	// Experimental.
	AddIncludes(patterns ...*string)
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

// The jsii proxy struct for CdkConfig
type jsiiProxy_CdkConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_CdkConfig) Cdkout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Json() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkConfig(project projen.Project, options *CdkConfigOptions) CdkConfig {
	_init_.Initialize()

	j := jsiiProxy_CdkConfig{}

	_jsii_.Create(
		"projen.awscdk.CdkConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkConfig_Override(c CdkConfig, project projen.Project, options *CdkConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.CdkConfig",
		[]interface{}{project, options},
		c,
	)
}

func (c *jsiiProxy_CdkConfig) AddExcludes(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addExcludes",
		args,
	)
}

func (c *jsiiProxy_CdkConfig) AddIncludes(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addIncludes",
		args,
	)
}

func (c *jsiiProxy_CdkConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

// Common options for `cdk.json`.
// Experimental.
type CdkConfigCommonOptions struct {
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
}

// Options for `CdkJson`.
// Experimental.
type CdkConfigOptions struct {
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
	// The command line to execute in order to synthesize the CDK application (language specific).
	// Experimental.
	App *string `field:"required" json:"app" yaml:"app"`
}

// Adds standard AWS CDK tasks to your project.
// Experimental.
type CdkTasks interface {
	projen.Component
	// Deploys your app.
	// Experimental.
	Deploy() projen.Task
	// Destroys all the stacks.
	// Experimental.
	Destroy() projen.Task
	// Diff against production.
	// Experimental.
	Diff() projen.Task
	// Experimental.
	Project() projen.Project
	// Synthesizes your app.
	// Experimental.
	Synth() projen.Task
	// Synthesizes your app and suppresses stdout.
	// Experimental.
	SynthSilent() projen.Task
	// Watch task.
	// Experimental.
	Watch() projen.Task
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

// The jsii proxy struct for CdkTasks
type jsiiProxy_CdkTasks struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_CdkTasks) Deploy() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"deploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Destroy() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"destroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Diff() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"diff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Synth() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"synth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) SynthSilent() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"synthSilent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Watch() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watch",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkTasks(project projen.Project) CdkTasks {
	_init_.Initialize()

	j := jsiiProxy_CdkTasks{}

	_jsii_.Create(
		"projen.awscdk.CdkTasks",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkTasks_Override(c CdkTasks, project projen.Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.CdkTasks",
		[]interface{}{project},
		c,
	)
}

func (c *jsiiProxy_CdkTasks) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkTasks) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkTasks) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

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
	// Maximum node version required by this pacakge.
	// Deprecated: use `AwsCdkConstructLibrary`.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Deprecated: use `AwsCdkConstructLibrary`.
	MinNodeVersion() *string
	// Project name.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Name() *string
	// The .npmignore file.
	// Deprecated: use `AwsCdkConstructLibrary`.
	Npmignore() projen.IgnoreFile
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
	// Deprecated: use `AwsCdkConstructLibrary`.
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

func (j *jsiiProxy_ConstructLibraryAws) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
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


// Deprecated: use `AwsCdkConstructLibrary`.
func NewConstructLibraryAws(options *AwsCdkConstructLibraryOptions) ConstructLibraryAws {
	_init_.Initialize()

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
	_jsii_.InvokeVoid(
		c,
		"addFields",
		[]interface{}{fields},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AddGitIgnore(pattern *string) {
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

func (c *jsiiProxy_ConstructLibraryAws) AddTask(name *string, props *projen.TaskOptions) projen.Task {
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
	_jsii_.InvokeVoid(
		c,
		"addTip",
		[]interface{}{message},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		c,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) HasScript(name *string) *bool {
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
	_jsii_.InvokeVoid(
		c,
		"removeScript",
		[]interface{}{name},
	)
}

func (c *jsiiProxy_ConstructLibraryAws) RemoveTask(name *string) projen.Task {
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
	var returns projen.FileBase

	_jsii_.Invoke(
		c,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Deprecated: use `AwsCdkConstructLibraryOptions`.
type ConstructLibraryAwsOptions struct {
	// This is the name of your project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Logging *projen.LoggerOptions `field:"optional" json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJsonOptions *projen.ProjenrcOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AllowLibraryDependencies *bool `field:"optional" json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Author's Organization.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorOrganization *bool `field:"optional" json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorUrl *string `field:"optional" json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoDetectBin *bool `field:"optional" json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Bin *map[string]*string `field:"optional" json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BugsEmail *string `field:"optional" json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BugsUrl *string `field:"optional" json:"bugsUrl" yaml:"bugsUrl"`
	// List of dependencies to bundle into this module.
	//
	// These modules will be
	// added both to the `dependencies` section and `bundledDependencies` section of
	// your `package.json`.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BundledDeps *[]*string `field:"optional" json:"bundledDeps" yaml:"bundledDeps"`
	// Options for npm packages using AWS CodeArtifact.
	//
	// This is required if publishing packages to, or installing scoped packages from AWS CodeArtifact.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CodeArtifactOptions *javascript.CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'express', 'lodash', 'foo@^2' ]
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// Example:
	//   [ 'typescript', '@types/express' ]
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Licensed *bool `field:"optional" json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MaxNodeVersion *string `field:"optional" json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MinNodeVersion *string `field:"optional" json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmAccess javascript.NpmAccess `field:"optional" json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `field:"optional" json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmRegistryUrl *string `field:"optional" json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmTokenSecret *string `field:"optional" json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PackageManager javascript.NodePackageManager `field:"optional" json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PeerDependencyOptions *javascript.PeerDependencyOptions `field:"optional" json:"peerDependencyOptions" yaml:"peerDependencyOptions"`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PeerDeps *[]*string `field:"optional" json:"peerDeps" yaml:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	RepositoryDirectory *string `field:"optional" json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// Options for privately hosted scoped packages.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ScopedPackagesOptions *[]*javascript.ScopedPackagesOptions `field:"optional" json:"scopedPackagesOptions" yaml:"scopedPackagesOptions"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Scripts *map[string]*string `field:"optional" json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Stability *string `field:"optional" json:"stability" yaml:"stability"`
	// Version requirement of `publib` which is used to publish modules to npm.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishDryRun *bool `field:"optional" json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseBranches *map[string]*release.BranchOptions `field:"optional" json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `field:"optional" json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseFailureIssue *bool `field:"optional" json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseFailureIssueLabel *string `field:"optional" json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `field:"optional" json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseTagPrefix *string `field:"optional" json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseTrigger release.ReleaseTrigger `field:"optional" json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseWorkflowName *string `field:"optional" json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `field:"optional" json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// The name of the main release branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DefaultReleaseBranch *string `field:"required" json:"defaultReleaseBranch" yaml:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoApproveUpgrades *bool `field:"optional" json:"autoApproveUpgrades" yaml:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BuildWorkflow *bool `field:"optional" json:"buildWorkflow" yaml:"buildWorkflow"`
	// Build workflow triggers.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BuildWorkflowTriggers *workflows.Triggers `field:"optional" json:"buildWorkflowTriggers" yaml:"buildWorkflowTriggers"`
	// Options for `Bundler`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BundlerOptions *javascript.BundlerOptions `field:"optional" json:"bundlerOptions" yaml:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CodeCov *bool `field:"optional" json:"codeCov" yaml:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CodeCovTokenSecret *string `field:"optional" json:"codeCovTokenSecret" yaml:"codeCovTokenSecret"`
	// License copyright owner.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CopyrightOwner *string `field:"optional" json:"copyrightOwner" yaml:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CopyrightPeriod *string `field:"optional" json:"copyrightPeriod" yaml:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Dependabot *bool `field:"optional" json:"dependabot" yaml:"dependabot"`
	// Options for dependabot.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DependabotOptions *github.DependabotOptions `field:"optional" json:"dependabotOptions" yaml:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DepsUpgrade *bool `field:"optional" json:"depsUpgrade" yaml:"depsUpgrade"`
	// Options for `UpgradeDependencies`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `field:"optional" json:"depsUpgradeOptions" yaml:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Gitignore *[]*string `field:"optional" json:"gitignore" yaml:"gitignore"`
	// Setup jest unit tests.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Jest *bool `field:"optional" json:"jest" yaml:"jest"`
	// Jest options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	JestOptions *javascript.JestOptions `field:"optional" json:"jestOptions" yaml:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `field:"optional" json:"npmignore" yaml:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmignoreEnabled *bool `field:"optional" json:"npmignoreEnabled" yaml:"npmignoreEnabled"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Package *bool `field:"optional" json:"package" yaml:"package"`
	// Setup prettier.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Prettier options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PrettierOptions *javascript.PrettierOptions `field:"optional" json:"prettierOptions" yaml:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenDevDependency *bool `field:"optional" json:"projenDevDependency" yaml:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options for .projenrc.js.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Version of projen to install.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Include a GitHub pull request template.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PullRequestTemplate *bool `field:"optional" json:"pullRequestTemplate" yaml:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PullRequestTemplateContents *[]*string `field:"optional" json:"pullRequestTemplateContents" yaml:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Release *bool `field:"optional" json:"release" yaml:"release"`
	// Automatically release to npm when new versions are introduced.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseToNpm *bool `field:"optional" json:"releaseToNpm" yaml:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `field:"optional" json:"releaseWorkflow" yaml:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowBootstrapSteps *[]*workflows.JobStep `field:"optional" json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowGitIdentity *github.GitIdentity `field:"optional" json:"workflowGitIdentity" yaml:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DisableTsconfig *bool `field:"optional" json:"disableTsconfig" yaml:"disableTsconfig"`
	// Docgen by Typedoc.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Docgen *bool `field:"optional" json:"docgen" yaml:"docgen"`
	// Docs directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DocsDirectory *string `field:"optional" json:"docsDirectory" yaml:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	EntrypointTypes *string `field:"optional" json:"entrypointTypes" yaml:"entrypointTypes"`
	// Setup eslint.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Eslint *bool `field:"optional" json:"eslint" yaml:"eslint"`
	// Eslint options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	EslintOptions *javascript.EslintOptions `field:"optional" json:"eslintOptions" yaml:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Libdir *string `field:"optional" json:"libdir" yaml:"libdir"`
	// Use TypeScript for your projenrc file (`.projenrc.ts`).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcTs *bool `field:"optional" json:"projenrcTs" yaml:"projenrcTs"`
	// Options for .projenrc.ts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcTsOptions *typescript.ProjenrcOptions `field:"optional" json:"projenrcTsOptions" yaml:"projenrcTsOptions"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	SampleCode *bool `field:"optional" json:"sampleCode" yaml:"sampleCode"`
	// Typescript sources directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Srcdir *string `field:"optional" json:"srcdir" yaml:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	//
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
	// Custom TSConfig.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Tsconfig *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Custom tsconfig options for the development tsconfig.json file (used for testing).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	TsconfigDev *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfigDev" yaml:"tsconfigDev"`
	// The name of the development tsconfig.json file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	TsconfigDevFile *string `field:"optional" json:"tsconfigDevFile" yaml:"tsconfigDevFile"`
	// TypeScript version to use.
	//
	// NOTE: Typescript is not semantically versioned and should remain on the
	// same minor, so we recommend using a `~` dependency (e.g. `~1.2.3`).
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	TypescriptVersion *string `field:"optional" json:"typescriptVersion" yaml:"typescriptVersion"`
	// The name of the library author.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Author *string `field:"required" json:"author" yaml:"author"`
	// Email or URL of the library author.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorAddress *string `field:"required" json:"authorAddress" yaml:"authorAddress"`
	// Git repository URL.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	//
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Compat *bool `field:"optional" json:"compat" yaml:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CompatIgnore *string `field:"optional" json:"compatIgnore" yaml:"compatIgnore"`
	// File path for generated docs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DocgenFilePath *string `field:"optional" json:"docgenFilePath" yaml:"docgenFilePath"`
	// Deprecated: use `publishToNuget`.
	Dotnet *cdk.JsiiDotNetTarget `field:"optional" json:"dotnet" yaml:"dotnet"`
	// Accepts a list of glob patterns.
	//
	// Files matching any of those patterns will be excluded from the TypeScript compiler input.
	//
	// By default, jsii will include all *.ts files (except .d.ts files) in the TypeScript compiler input.
	// This can be problematic for example when the package's build or test procedure generates .ts files
	// that cannot be compiled with jsii's compiler settings.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ExcludeTypescript *[]*string `field:"optional" json:"excludeTypescript" yaml:"excludeTypescript"`
	// Publish Go bindings to a git repository.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToGo *cdk.JsiiGoTarget `field:"optional" json:"publishToGo" yaml:"publishToGo"`
	// Publish to maven.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToMaven *cdk.JsiiJavaTarget `field:"optional" json:"publishToMaven" yaml:"publishToMaven"`
	// Publish to NuGet.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToNuget *cdk.JsiiDotNetTarget `field:"optional" json:"publishToNuget" yaml:"publishToNuget"`
	// Publish to pypi.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToPypi *cdk.JsiiPythonTarget `field:"optional" json:"publishToPypi" yaml:"publishToPypi"`
	// Deprecated: use `publishToPyPi`.
	Python *cdk.JsiiPythonTarget `field:"optional" json:"python" yaml:"python"`
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Rootdir *string `field:"optional" json:"rootdir" yaml:"rootdir"`
	// Libraries will be picked up by the construct catalog when they are published to npm as jsii modules and will be published under:.
	//
	// https://awscdk.io/packages/[@SCOPE/]PACKAGE@VERSION
	//
	// The catalog will also post a tweet to https://twitter.com/awscdkio with the
	// package name, description and the above link. You can disable these tweets
	// through `{ announce: false }`.
	//
	// You can also add a Twitter handle through `{ twitter: 'xx' }` which will be
	// mentioned in the tweet.
	// See: https://github.com/construct-catalog/catalog
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Catalog *cdk.Catalog `field:"optional" json:"catalog" yaml:"catalog"`
	// Minimum version of the AWS CDK to depend on.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the @aws-cdk/assert library?
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// Automatically adds an `cloudfront.experimental.EdgeFunction` for each `.edge-lambda.ts` handler in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	EdgeLambdaAutoDiscover *bool `field:"optional" json:"edgeLambdaAutoDiscover" yaml:"edgeLambdaAutoDiscover"`
	// Automatically discovers and creates integration tests for each `.integ.ts` file in under your test directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
	// Automatically adds an `aws_lambda.Function` for each `.lambda.ts` handler in your source tree. If this is disabled, you either need to explicitly call `aws_lambda.Function.autoDiscover()` or define a `new aws_lambda.Function()` for each handler.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	LambdaAutoDiscover *bool `field:"optional" json:"lambdaAutoDiscover" yaml:"lambdaAutoDiscover"`
	// Automatically adds an `awscdk.LambdaExtension` for each `.lambda-extension.ts` entrypoint in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	LambdaExtensionAutoDiscover *bool `field:"optional" json:"lambdaExtensionAutoDiscover" yaml:"lambdaExtensionAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
}

// Creates edge lambdas from entry points discovered in the project's source tree.
// Experimental.
type EdgeLambdaAutoDiscover interface {
	cdk.AutoDiscoverBase
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
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

// The jsii proxy struct for EdgeLambdaAutoDiscover
type jsiiProxy_EdgeLambdaAutoDiscover struct {
	internal.Type__cdkAutoDiscoverBase
}

func (j *jsiiProxy_EdgeLambdaAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeLambdaAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewEdgeLambdaAutoDiscover(project projen.Project, options *EdgeLambdaAutoDiscoverOptions) EdgeLambdaAutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_EdgeLambdaAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.EdgeLambdaAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewEdgeLambdaAutoDiscover_Override(e EdgeLambdaAutoDiscover, project projen.Project, options *EdgeLambdaAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.EdgeLambdaAutoDiscover",
		[]interface{}{project, options},
		e,
	)
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"postSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"preSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `EdgeLambdaAutoDiscover`.
// Experimental.
type EdgeLambdaAutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Project source tree (relative to project output directory).
	// Experimental.
	Srcdir *string `field:"required" json:"srcdir" yaml:"srcdir"`
	// Options for AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
}

// Cloud integration tests.
// Experimental.
type IntegrationTest interface {
	cdk.IntegrationTestBase
	// Synthesizes the integration test and compares against a local copy (runs during build).
	// Experimental.
	AssertTask() projen.Task
	// Deploy the integration test and update the snapshot upon success.
	// Experimental.
	DeployTask() projen.Task
	// Destroy the integration test resources.
	// Experimental.
	DestroyTask() projen.Task
	// Integration test name.
	// Experimental.
	Name() *string
	// Experimental.
	Project() projen.Project
	// Snapshot output directory.
	// Experimental.
	SnapshotDir() *string
	// Just update snapshot (without deployment).
	// Experimental.
	SnapshotTask() projen.Task
	// Temporary directory for each integration test.
	// Experimental.
	TmpDir() *string
	// The watch task.
	// Experimental.
	WatchTask() projen.Task
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

// The jsii proxy struct for IntegrationTest
type jsiiProxy_IntegrationTest struct {
	internal.Type__cdkIntegrationTestBase
}

func (j *jsiiProxy_IntegrationTest) AssertTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"assertTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) DeployTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"deployTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) DestroyTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"destroyTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) SnapshotDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) SnapshotTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"snapshotTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) TmpDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tmpDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTest) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTest(project projen.Project, options *IntegrationTestOptions) IntegrationTest {
	_init_.Initialize()

	j := jsiiProxy_IntegrationTest{}

	_jsii_.Create(
		"projen.awscdk.IntegrationTest",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTest_Override(i IntegrationTest, project projen.Project, options *IntegrationTestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.IntegrationTest",
		[]interface{}{project, options},
		i,
	)
}

func (i *jsiiProxy_IntegrationTest) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTest) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTest) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

// Creates integration tests from entry points discovered in the test tree.
// Experimental.
type IntegrationTestAutoDiscover interface {
	cdk.IntegrationTestAutoDiscoverBase
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
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

// The jsii proxy struct for IntegrationTestAutoDiscover
type jsiiProxy_IntegrationTestAutoDiscover struct {
	internal.Type__cdkIntegrationTestAutoDiscoverBase
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTestAutoDiscover(project projen.Project, options *IntegrationTestAutoDiscoverOptions) IntegrationTestAutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_IntegrationTestAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.IntegrationTestAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTestAutoDiscover_Override(i IntegrationTestAutoDiscover, project projen.Project, options *IntegrationTestAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.IntegrationTestAutoDiscover",
		[]interface{}{project, options},
		i,
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `IntegrationTestAutoDiscover`.
// Experimental.
type IntegrationTestAutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Test source tree.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
	// Options for integration tests.
	// Experimental.
	IntegrationTestOptions *IntegrationTestCommonOptions `field:"optional" json:"integrationTestOptions" yaml:"integrationTestOptions"`
}

// Experimental.
type IntegrationTestCommonOptions struct {
	// Destroy the test app after a successful deployment.
	//
	// If disabled, leaves the
	// app deployed in the dev account.
	// Experimental.
	DestroyAfterDeploy *bool `field:"optional" json:"destroyAfterDeploy" yaml:"destroyAfterDeploy"`
	// Enables path metadata, adding `aws:cdk:path`, with the defining construct's path, to the CloudFormation metadata for each synthesized resource.
	// Experimental.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
}

// Options for `IntegrationTest`.
// Experimental.
type IntegrationTestOptions struct {
	// Destroy the test app after a successful deployment.
	//
	// If disabled, leaves the
	// app deployed in the dev account.
	// Experimental.
	DestroyAfterDeploy *bool `field:"optional" json:"destroyAfterDeploy" yaml:"destroyAfterDeploy"`
	// Enables path metadata, adding `aws:cdk:path`, with the defining construct's path, to the CloudFormation metadata for each synthesized resource.
	// Experimental.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// A path from the project root directory to a TypeScript file which contains the integration test app.
	//
	// This is relative to the root directory of the project.
	//
	// Example:
	//   "test/subdir/foo.integ.ts"
	//
	// Experimental.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// The path of the tsconfig.json file to use when running integration test cdk apps.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Name of the integration test.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// A list of stacks within the integration test to deploy/destroy.
	// Experimental.
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
}

// Creates lambdas from entry points discovered in the project's source tree.
// Experimental.
type LambdaAutoDiscover interface {
	cdk.AutoDiscoverBase
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
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

// The jsii proxy struct for LambdaAutoDiscover
type jsiiProxy_LambdaAutoDiscover struct {
	internal.Type__cdkAutoDiscoverBase
}

func (j *jsiiProxy_LambdaAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaAutoDiscover(project projen.Project, options *LambdaAutoDiscoverOptions) LambdaAutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_LambdaAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.LambdaAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaAutoDiscover_Override(l LambdaAutoDiscover, project projen.Project, options *LambdaAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaAutoDiscover",
		[]interface{}{project, options},
		l,
	)
}

func (l *jsiiProxy_LambdaAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `LambdaAutoDiscover`.
// Experimental.
type LambdaAutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Project source tree (relative to project output directory).
	// Experimental.
	Srcdir *string `field:"required" json:"srcdir" yaml:"srcdir"`
	// Options for AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
}

// Create a Lambda Extension.
// Experimental.
type LambdaExtension interface {
	projen.Component
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

// The jsii proxy struct for LambdaExtension
type jsiiProxy_LambdaExtension struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_LambdaExtension) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaExtension(project projen.Project, options *LambdaExtensionOptions) LambdaExtension {
	_init_.Initialize()

	j := jsiiProxy_LambdaExtension{}

	_jsii_.Create(
		"projen.awscdk.LambdaExtension",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaExtension_Override(l LambdaExtension, project projen.Project, options *LambdaExtensionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaExtension",
		[]interface{}{project, options},
		l,
	)
}

func (l *jsiiProxy_LambdaExtension) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtension) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtension) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

// Creates Lambda Extensions from entrypoints discovered in the project's source tree.
// Experimental.
type LambdaExtensionAutoDiscover interface {
	cdk.AutoDiscoverBase
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
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

// The jsii proxy struct for LambdaExtensionAutoDiscover
type jsiiProxy_LambdaExtensionAutoDiscover struct {
	internal.Type__cdkAutoDiscoverBase
}

func (j *jsiiProxy_LambdaExtensionAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaExtensionAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaExtensionAutoDiscover(project projen.Project, options *LambdaExtensionAutoDiscoverOptions) LambdaExtensionAutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_LambdaExtensionAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.LambdaExtensionAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaExtensionAutoDiscover_Override(l LambdaExtensionAutoDiscover, project projen.Project, options *LambdaExtensionAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaExtensionAutoDiscover",
		[]interface{}{project, options},
		l,
	)
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `LambdaExtensionAutoDiscover`.
// Experimental.
type LambdaExtensionAutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Project source tree (relative to project output directory).
	// Experimental.
	Srcdir *string `field:"required" json:"srcdir" yaml:"srcdir"`
	// Options for lambda extensions.
	// Experimental.
	LambdaExtensionOptions *LambdaExtensionCommonOptions `field:"optional" json:"lambdaExtensionOptions" yaml:"lambdaExtensionOptions"`
}

// Common options for creating lambda extensions.
// Experimental.
type LambdaExtensionCommonOptions struct {
	// Bundling options for this AWS Lambda extension.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// The extension's compatible runtimes.
	// Experimental.
	CompatibleRuntimes *[]LambdaRuntime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
}

// Options for creating lambda extensions.
// Experimental.
type LambdaExtensionOptions struct {
	// Bundling options for this AWS Lambda extension.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// The extension's compatible runtimes.
	// Experimental.
	CompatibleRuntimes *[]LambdaRuntime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// A path from the project root directory to a TypeScript file which contains the AWS Lambda extension entrypoint (stand-alone script).
	//
	// This is relative to the root directory of the project.
	//
	// Example:
	//   "src/subdir/foo.lambda-extension.ts"
	//
	// Experimental.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// The name of the generated TypeScript source file.
	//
	// This file should also be
	// under the source tree.
	// Experimental.
	ConstructFile *string `field:"optional" json:"constructFile" yaml:"constructFile"`
	// The name of the generated `lambda.LayerVersion` subclass.
	// Experimental.
	ConstructName *string `field:"optional" json:"constructName" yaml:"constructName"`
	// Name of the extension.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Generates a pre-bundled AWS Lambda function construct from handler code.
//
// To use this, create an AWS Lambda handler file under your source tree with
// the `.lambda.ts` extension and add a `LambdaFunction` component to your
// typescript project pointing to this entrypoint.
//
// This will add a task to your "compile" step which will use `esbuild` to
// bundle the handler code into the build directory. It will also generate a
// file `src/foo-function.ts` with a custom AWS construct called `FooFunction`
// which extends `@aws-cdk/aws-lambda.Function` which is bound to the bundled
// handle through an asset.
//
// Example:
//   new LambdaFunction(myProject, {
//     srcdir: myProject.srcdir,
//     entrypoint: 'src/foo.lambda.ts',
//   });
//
// Experimental.
type LambdaFunction interface {
	projen.Component
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

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_LambdaFunction) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Defines a pre-bundled AWS Lambda function construct from handler code.
// Experimental.
func NewLambdaFunction(project projen.Project, options *LambdaFunctionOptions) LambdaFunction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"projen.awscdk.LambdaFunction",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Defines a pre-bundled AWS Lambda function construct from handler code.
// Experimental.
func NewLambdaFunction_Override(l LambdaFunction, project projen.Project, options *LambdaFunctionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaFunction",
		[]interface{}{project, options},
		l,
	)
}

func (l *jsiiProxy_LambdaFunction) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

// Common options for `LambdaFunction`.
//
// Applies to all functions in
// auto-discovery.
// Experimental.
type LambdaFunctionCommonOptions struct {
	// Whether to automatically reuse TCP connections when working with the AWS SDK for JavaScript.
	//
	// This sets the `AWS_NODEJS_CONNECTION_REUSE_ENABLED` environment variable
	// to `1`.
	//
	// Not applicable when `edgeLambda` is set to `true` because environment
	// variables are not supported in Lambda@Edge.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/node-reusing-connections.html
	//
	// Experimental.
	AwsSdkConnectionReuse *bool `field:"optional" json:"awsSdkConnectionReuse" yaml:"awsSdkConnectionReuse"`
	// Bundling options for this AWS Lambda function.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// Whether to create a `cloudfront.experimental.EdgeFunction` instead of a `lambda.Function`.
	// Experimental.
	EdgeLambda *bool `field:"optional" json:"edgeLambda" yaml:"edgeLambda"`
	// The node.js version to target.
	// Experimental.
	Runtime LambdaRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

// Options for `Function`.
// Experimental.
type LambdaFunctionOptions struct {
	// Whether to automatically reuse TCP connections when working with the AWS SDK for JavaScript.
	//
	// This sets the `AWS_NODEJS_CONNECTION_REUSE_ENABLED` environment variable
	// to `1`.
	//
	// Not applicable when `edgeLambda` is set to `true` because environment
	// variables are not supported in Lambda@Edge.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/node-reusing-connections.html
	//
	// Experimental.
	AwsSdkConnectionReuse *bool `field:"optional" json:"awsSdkConnectionReuse" yaml:"awsSdkConnectionReuse"`
	// Bundling options for this AWS Lambda function.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// Whether to create a `cloudfront.experimental.EdgeFunction` instead of a `lambda.Function`.
	// Experimental.
	EdgeLambda *bool `field:"optional" json:"edgeLambda" yaml:"edgeLambda"`
	// The node.js version to target.
	// Experimental.
	Runtime LambdaRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// A path from the project root directory to a TypeScript file which contains the AWS Lambda handler entrypoint (exports a `handler` function).
	//
	// This is relative to the root directory of the project.
	//
	// Example:
	//   "src/subdir/foo.lambda.ts"
	//
	// Experimental.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// The name of the generated TypeScript source file.
	//
	// This file should also be
	// under the source tree.
	// Experimental.
	ConstructFile *string `field:"optional" json:"constructFile" yaml:"constructFile"`
	// The name of the generated `lambda.Function` subclass.
	// Experimental.
	ConstructName *string `field:"optional" json:"constructName" yaml:"constructName"`
}

// The runtime for the AWS Lambda function.
// Experimental.
type LambdaRuntime interface {
	// Experimental.
	EsbuildPlatform() *string
	// The esbuild setting to use.
	// Experimental.
	EsbuildTarget() *string
	// The Node.js runtime to use.
	// Experimental.
	FunctionRuntime() *string
}

// The jsii proxy struct for LambdaRuntime
type jsiiProxy_LambdaRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_LambdaRuntime) EsbuildPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esbuildPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRuntime) EsbuildTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esbuildTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaRuntime) FunctionRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionRuntime",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaRuntime(functionRuntime *string, esbuildTarget *string) LambdaRuntime {
	_init_.Initialize()

	j := jsiiProxy_LambdaRuntime{}

	_jsii_.Create(
		"projen.awscdk.LambdaRuntime",
		[]interface{}{functionRuntime, esbuildTarget},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaRuntime_Override(l LambdaRuntime, functionRuntime *string, esbuildTarget *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaRuntime",
		[]interface{}{functionRuntime, esbuildTarget},
		l,
	)
}

func LambdaRuntime_NODEJS_10_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_10_X",
		&returns,
	)
	return returns
}

func LambdaRuntime_NODEJS_12_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_12_X",
		&returns,
	)
	return returns
}

func LambdaRuntime_NODEJS_14_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_14_X",
		&returns,
	)
	return returns
}

func LambdaRuntime_NODEJS_16_X() LambdaRuntime {
	_init_.Initialize()
	var returns LambdaRuntime
	_jsii_.StaticGet(
		"projen.awscdk.LambdaRuntime",
		"NODEJS_16_X",
		&returns,
	)
	return returns
}

