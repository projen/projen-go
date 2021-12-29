package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build"
	"github.com/projen/projen-go/projen/cdk"
	"github.com/projen/projen-go/projen/cdktf/internal"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/typescript"
	"github.com/projen/projen-go/projen/vscode"
)

// CDKTF construct library project.
//
// A multi-language (jsii) construct library which vends constructs designed to
// use within the CDK for Terraform (CDKTF), with a friendly workflow and
// automatic publishing to the construct catalog.
// Experimental.
type ConstructLibraryCdktf interface {
	cdk.ConstructLibrary
	AllowLibraryDependencies() *bool
	ArtifactsDirectory() *string
	ArtifactsJavascriptDirectory() *string
	AutoApprove() github.AutoApprove
	AutoMerge() github.AutoMerge
	BuildTask() projen.Task
	BuildWorkflow() build.BuildWorkflow
	BuildWorkflowJobId() *string
	Bundler() javascript.Bundler
	CompileTask() projen.Task
	Components() *[]projen.Component
	DefaultTask() projen.Task
	Deps() projen.Dependencies
	DevContainer() vscode.DevContainer
	Docgen() *bool
	DocsDirectory() *string
	Entrypoint() *string
	Eslint() javascript.Eslint
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	InstallWorkflowSteps() *[]*workflows.JobStep
	Jest() javascript.Jest
	Libdir() *string
	Logger() projen.Logger
	Manifest() interface{}
	MaxNodeVersion() *string
	MinNodeVersion() *string
	Name() *string
	Npmignore() projen.IgnoreFile
	Outdir() *string
	Package() javascript.NodePackage
	PackageManager() javascript.NodePackageManager
	PackageTask() projen.Task
	Parent() projen.Project
	PostCompileTask() projen.Task
	PreCompileTask() projen.Task
	Prettier() javascript.Prettier
	PrettierIgnore() projen.IgnoreFile
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Publisher() release.Publisher
	Release() release.Release
	Root() projen.Project
	RunScriptCommand() *string
	Srcdir() *string
	Tasks() projen.Tasks
	Testdir() *string
	TestTask() projen.Task
	Tsconfig() javascript.TypescriptConfig
	TsconfigDev() javascript.TypescriptConfig
	TsconfigEslint() javascript.TypescriptConfig
	UpgradeWorkflow() javascript.UpgradeDependencies
	Vscode() vscode.VsCode
	WatchTask() projen.Task
	AddBins(bins *map[string]*string)
	AddBundledDeps(deps ...*string)
	AddCompileCommand(commands ...*string)
	AddDeps(deps ...*string)
	AddDevDeps(deps ...*string)
	AddExcludeFromCleanup(globs ...*string)
	AddFields(fields *map[string]interface{})
	AddGitIgnore(pattern *string)
	AddKeywords(keywords ...*string)
	AddPackageIgnore(pattern *string)
	AddPeerDeps(deps ...*string)
	AddPrettierIgnore(pattern *string)
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	AddTestCommand(commands ...*string)
	AddTip(message *string)
	AnnotateGenerated(glob *string)
	HasScript(name *string) *bool
	PostSynthesize()
	PreSynthesize()
	RemoveScript(name *string)
	RemoveTask(name *string) projen.Task
	RunTaskCommand(task projen.Task) *string
	SetScript(name *string, command *string)
	Synth()
	TryFindFile(filePath *string) projen.FileBase
	TryFindJsonFile(filePath *string) projen.JsonFile
	TryFindObjectFile(filePath *string) projen.ObjectFile
}

// The jsii proxy struct for ConstructLibraryCdktf
type jsiiProxy_ConstructLibraryCdktf struct {
	internal.Type__cdkConstructLibrary
}

func (j *jsiiProxy_ConstructLibraryCdktf) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Prettier() javascript.Prettier {
	var returns javascript.Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) PrettierIgnore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"prettierIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConstructLibraryCdktf) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewConstructLibraryCdktf(options *ConstructLibraryCdktfOptions) ConstructLibraryCdktf {
	_init_.Initialize()

	j := jsiiProxy_ConstructLibraryCdktf{}

	_jsii_.Create(
		"projen.cdktf.ConstructLibraryCdktf",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewConstructLibraryCdktf_Override(c ConstructLibraryCdktf, options *ConstructLibraryCdktfOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdktf.ConstructLibraryCdktf",
		[]interface{}{options},
		c,
	)
}

func ConstructLibraryCdktf_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.cdktf.ConstructLibraryCdktf",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		c,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddBundledDeps(deps ...*string) {
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

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
func (c *jsiiProxy_ConstructLibraryCdktf) AddCompileCommand(commands ...*string) {
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

// Defines normal dependencies.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddDeps(deps ...*string) {
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

// Defines development/test dependencies.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddDevDeps(deps ...*string) {
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

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddExcludeFromCleanup(globs ...*string) {
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

// Directly set fields in `package.json`.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		c,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddKeywords(keywords ...*string) {
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

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		c,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

// Defines peer dependencies.
//
// When adding peer dependencies, a devDependency will also be added on the
// pinned version of the declared peer. This will ensure that you are testing
// your code against the minimum version required from your consumers.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddPeerDeps(deps ...*string) {
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

// Defines Prettier ignore Patterns these patterns will be added to the file .prettierignore.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddPrettierIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrettierIgnore",
		[]interface{}{pattern},
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		c,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
func (c *jsiiProxy_ConstructLibraryCdktf) AddTestCommand(commands ...*string) {
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

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (c *jsiiProxy_ConstructLibraryCdktf) AddTip(message *string) {
	_jsii_.InvokeVoid(
		c,
		"addTip",
		[]interface{}{message},
	)
}

// Marks the provided file(s) as being generated.
//
// This is achieved using the
// github-linguist attributes. Generated files do not count against the
// repository statistics and language breakdown.
// See: https://github.com/github/linguist/blob/master/docs/overrides.md
//
// Deprecated.
func (c *jsiiProxy_ConstructLibraryCdktf) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		c,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		c,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		c,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the shell command to execute in order to run a task.
//
// This will
// typically be `npx projen TASK`.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Replaces the contents of an npm package.json script.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		c,
		"setScript",
		[]interface{}{name, command},
	)
}

// Synthesize all project files into `outdir`.
//
// 1. Call "this.preSynthesize()"
// 2. Delete all generated files
// 3. Synthesize all sub-projects
// 4. Synthesize all components of this project
// 5. Call "postSynthesize()" for all components of this project
// 6. Call "this.postSynthesize()"
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) Synth() {
	_jsii_.InvokeVoid(
		c,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		c,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (c *jsiiProxy_ConstructLibraryCdktf) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		c,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (c *jsiiProxy_ConstructLibraryCdktf) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		c,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type ConstructLibraryCdktfOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `json:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level
	ProjectType projen.ProjectType `json:"projectType"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Experimental.
	ProjenTokenSecret *string `json:"projenTokenSecret"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `json:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `json:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `json:"bugsUrl"`
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
	BundledDeps *[]*string `json:"bundledDeps"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *javascript.CodeArtifactOptions `json:"codeArtifactOptions"`
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
	Deps *[]*string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `json:"description"`
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
	DevDeps *[]*string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `json:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess javascript.NpmAccess `json:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry *string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `json:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `json:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager javascript.NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *javascript.PeerDependencyOptions `json:"peerDependencyOptions"`
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
	PeerDeps *[]*string `json:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `json:"repositoryDirectory"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `json:"stability"`
	// Version requirement of `jsii-release` which is used to publish modules to npm.
	// Experimental.
	JsiiReleaseVersion *string `json:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Experimental.
	MajorVersion *float64 `json:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Experimental.
	NpmDistTag *string `json:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Experimental.
	Prerelease *string `json:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Experimental.
	PublishDryRun *bool `json:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `json:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Experimental.
	ReleaseBranches *map[string]*release.BranchOptions `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `json:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Experimental.
	ReleaseFailureIssue *bool `json:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Experimental.
	ReleaseFailureIssueLabel *string `json:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `json:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	ReleaseTagPrefix *string `json:"releaseTagPrefix"`
	// The release trigger to use.
	// Experimental.
	ReleaseTrigger release.ReleaseTrigger `json:"releaseTrigger"`
	// The name of the default release workflow.
	// Experimental.
	ReleaseWorkflowName *string `json:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `json:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Experimental.
	VersionrcOptions *map[string]interface{} `json:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage *string `json:"workflowContainerImage"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `json:"workflowRunsOn"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch *string `json:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveProjenUpgrades *bool `json:"autoApproveProjenUpgrades"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveUpgrades *bool `json:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow *bool `json:"buildWorkflow"`
	// Options for `Bundler`.
	// Experimental.
	BundlerOptions *javascript.BundlerOptions `json:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov *bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret *string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner *string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod *string `json:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Experimental.
	Dependabot *bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions *github.DependabotOptions `json:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Experimental.
	DepsUpgrade *bool `json:"depsUpgrade"`
	// Options for depsUpgrade.
	// Experimental.
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `json:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore *[]*string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest *bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions *javascript.JestOptions `json:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `json:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled *bool `json:"npmignoreEnabled"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Experimental.
	Package *bool `json:"package"`
	// Setup prettier.
	// Experimental.
	Prettier *bool `json:"prettier"`
	// Defines an .prettierIgnore file.
	// Experimental.
	PrettierIgnoreEnabled *bool `json:"prettierIgnoreEnabled"`
	// Prettier options.
	// Experimental.
	PrettierOptions *javascript.PrettierOptions `json:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency *bool `json:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Experimental.
	ProjenrcJs *bool `json:"projenrcJs"`
	// Options for .projenrc.js.
	// Experimental.
	ProjenrcJsOptions *javascript.ProjenrcOptions `json:"projenrcJsOptions"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `autoApproveProjenUpgrades`.
	ProjenUpgradeAutoMerge *bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule *[]*string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	//
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	//
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	//
	// To create a personal access token see https://github.com/settings/tokens
	// Deprecated: use `githubTokenSecret` instead.
	ProjenUpgradeSecret *string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion *string `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate *bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents *[]*string `json:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Experimental.
	Release *bool `json:"release"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm *bool `json:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps *[]interface{} `json:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Experimental.
	WorkflowGitIdentity *github.GitIdentity `json:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion *string `json:"workflowNodeVersion"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig *bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen *bool `json:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory *string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes *string `json:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint *bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions *javascript.EslintOptions `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir *string `json:"libdir"`
	// Use TypeScript for your projenrc file (`.projenrc.ts`).
	// Experimental.
	ProjenrcTs *bool `json:"projenrcTs"`
	// Options for .projenrc.ts.
	// Experimental.
	ProjenrcTsOptions *typescript.ProjenrcOptions `json:"projenrcTsOptions"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode *bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir *string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	//
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir *string `json:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig *javascript.TypescriptConfigOptions `json:"tsconfig"`
	// Custom tsconfig options for the development tsconfig.json file (used for testing).
	// Experimental.
	TsconfigDev *javascript.TypescriptConfigOptions `json:"tsconfigDev"`
	// The name of the development tsconfig.json file.
	// Experimental.
	TsconfigDevFile *string `json:"tsconfigDevFile"`
	// TypeScript version to use.
	//
	// NOTE: Typescript is not semantically versioned and should remain on the
	// same minor, so we recommend using a `~` dependency (e.g. `~1.2.3`).
	// Experimental.
	TypescriptVersion *string `json:"typescriptVersion"`
	// The name of the library author.
	// Experimental.
	Author *string `json:"author"`
	// Email or URL of the library author.
	// Experimental.
	AuthorAddress *string `json:"authorAddress"`
	// Git repository URL.
	// Experimental.
	RepositoryUrl *string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	//
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Experimental.
	Compat *bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Experimental.
	CompatIgnore *string `json:"compatIgnore"`
	// Deprecated: use `publishToNuget`
	Dotnet *cdk.JsiiDotNetTarget `json:"dotnet"`
	// Accepts a list of glob patterns.
	//
	// Files matching any of those patterns will be excluded from the TypeScript compiler input.
	//
	// By default, jsii will include all *.ts files (except .d.ts files) in the TypeScript compiler input.
	// This can be problematic for example when the package's build or test procedure generates .ts files
	// that cannot be compiled with jsii's compiler settings.
	// Experimental.
	ExcludeTypescript *[]*string `json:"excludeTypescript"`
	// Publish Go bindings to a git repository.
	// Experimental.
	PublishToGo *cdk.JsiiGoTarget `json:"publishToGo"`
	// Publish to maven.
	// Experimental.
	PublishToMaven *cdk.JsiiJavaTarget `json:"publishToMaven"`
	// Publish to NuGet.
	// Experimental.
	PublishToNuget *cdk.JsiiDotNetTarget `json:"publishToNuget"`
	// Publish to pypi.
	// Experimental.
	PublishToPypi *cdk.JsiiPythonTarget `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python *cdk.JsiiPythonTarget `json:"python"`
	// Experimental.
	Rootdir *string `json:"rootdir"`
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
	Catalog *cdk.Catalog `json:"catalog"`
	// Minimum target version this library is tested against.
	// Experimental.
	CdktfVersion *string `json:"cdktfVersion"`
}

