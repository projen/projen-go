package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/cdk"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/java"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/typescript"
	"github.com/projen/projen-go/projen/vscode"
)

// Which approval is required when deploying CDK apps.
// Experimental.
type ApprovalLevel string

const (
	ApprovalLevel_NEVER ApprovalLevel = "NEVER"
	ApprovalLevel_ANY_CHANGE ApprovalLevel = "ANY_CHANGE"
	ApprovalLevel_BROADENING ApprovalLevel = "BROADENING"
)

// Automatically creates a `LambdaFunction` for all `.lambda.ts` files under the source directory of the project.
// Experimental.
type AutoDiscover interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (a *jsiiProxy_AutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (a *jsiiProxy_AutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `AutoDiscover`.
// Experimental.
type AutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `json:"cdkDeps"`
	// Project source tree (relative to project output directory).
	// Experimental.
	Srcdir *string `json:"srcdir"`
	// Test source tree.
	// Experimental.
	Testdir *string `json:"testdir"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `json:"tsconfigPath"`
	// Options for auto-discovery of AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `json:"lambdaOptions"`
}

// AWS CDK construct library project.
//
// A multi-language (jsii) construct library which vends constructs designed to
// use within the AWS CDK with a friendly workflow and automatic publishing to
// the construct catalog.
// Experimental.
type AwsCdkConstructLibrary interface {
	cdk.ConstructLibrary
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	AutoApprove() github.AutoApprove
	AutoMerge() github.AutoMerge
	BuildTask() projen.Task
	BuildWorkflow() github.TaskWorkflow
	BuildWorkflowJobId() *string
	Bundler() javascript.Bundler
	CdkDeps() AwsCdkDeps
	CdkVersion() *string
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
	Version() *string
	Vscode() vscode.VsCode
	WatchTask() projen.Task
	AddBins(bins *map[string]*string)
	AddBundledDeps(deps ...*string)
	AddCdkDependencies(deps ...*string)
	AddCdkTestDependencies(deps ...*string)
	AddCompileCommand(commands ...*string)
	AddDeps(deps ...*string)
	AddDevDeps(deps ...*string)
	AddExcludeFromCleanup(globs ...*string)
	AddFields(fields *map[string]interface{})
	AddGitIgnore(pattern *string)
	AddKeywords(keywords ...*string)
	AddPackageIgnore(pattern *string)
	AddPeerDeps(deps ...*string)
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

func (j *jsiiProxy_AwsCdkConstructLibrary) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
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

func (j *jsiiProxy_AwsCdkConstructLibrary) BuildWorkflow() github.TaskWorkflow {
	var returns github.TaskWorkflow
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

func (j *jsiiProxy_AwsCdkConstructLibrary) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
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

// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		a,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
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

// Adds dependencies to AWS CDK modules.
//
// Since this is a library project, dependencies will be added as peer dependencies.
// Deprecated: Not supported in v2. For v1, use `project.cdkDeps.addV1Dependencies()`
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

// Adds AWS CDK modules as dev dependencies.
// Deprecated: Not supported in v2. For v1, use `project.cdkDeps.addV1DevDependencies()`
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

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
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

// Defines normal dependencies.
// Experimental.
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

// Defines development/test dependencies.
// Experimental.
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

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
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

// Directly set fields in `package.json`.
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
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

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
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

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
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

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
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

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (a *jsiiProxy_AwsCdkConstructLibrary) AddTip(message *string) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AwsCdkConstructLibrary) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
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

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		a,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
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

// Returns the shell command to execute in order to run a task.
//
// This will
// typically be `npx projen TASK`.
// Experimental.
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

// Replaces the contents of an npm package.json script.
// Experimental.
func (a *jsiiProxy_AwsCdkConstructLibrary) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AwsCdkConstructLibrary) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
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

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
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

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
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

// Options for `AwsCdkConstructLibrary`.
// Experimental.
type AwsCdkConstructLibraryOptions struct {
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
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper *bool `json:"antitamper"`
	// A directory which will contain artifacts to be published to npm.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory"`
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
	// Experimental.
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
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Experimental.
	Package *bool `json:"package"`
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
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Deprecated: The
	CdkAssert *bool `json:"cdkAssert"`
	// Install the @aws-cdk/assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'
	// Experimental.
	CdkAssertions *bool `json:"cdkAssertions"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Deprecated: For CDK 2.x use 'peerDeps' instead
	CdkDependencies *[]*string `json:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `json:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' instead
	CdkTestDependencies *[]*string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `json:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `json:"constructsVersion"`
	// Automatically adds an `aws_lambda.Function` for each `.lambda.ts` handler in your source tree. If this is disabled, you either need to explicitly call `aws_lambda.Function.autoDiscover()` or define a `new aws_lambda.Function()` for each handler.
	// Experimental.
	LambdaAutoDiscover *bool `json:"lambdaAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `json:"lambdaOptions"`
}

// Manages dependencies on the AWS CDK.
// Experimental.
type AwsCdkDeps interface {
	projen.Component
	CdkDependenciesAsDeps() *bool
	CdkMajorVersion() *float64
	CdkMinimumVersion() *string
	CdkVersion() *string
	Project() projen.Project
	AddV1Dependencies(deps ...*string)
	AddV1DevDependencies(deps ...*string)
	PostSynthesize()
	PreSynthesize()
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
func NewAwsCdkDeps(project projen.Project, options *AwsCdkDepsOptions) AwsCdkDeps {
	_init_.Initialize()

	j := jsiiProxy_AwsCdkDeps{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkDeps",
		[]interface{}{project, options},
		&j,
	)

	return &j
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

// Adds dependencies to AWS CDK modules.
//
// The type of dependency is determined by the `dependencyType` option.
//
// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
// `project.addDeps()` as appropriate.
// Experimental.
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

// Adds AWS CDK modules as dev dependencies.
//
// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
// `project.addDeps()` as appropriate.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AwsCdkDeps) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (a *jsiiProxy_AwsCdkDeps) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	CdkVersion *string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Deprecated: The
	CdkAssert *bool `json:"cdkAssert"`
	// Install the @aws-cdk/assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'
	// Experimental.
	CdkAssertions *bool `json:"cdkAssertions"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Deprecated: For CDK 2.x use 'peerDeps' instead
	CdkDependencies *[]*string `json:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `json:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' instead
	CdkTestDependencies *[]*string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `json:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `json:"constructsVersion"`
}

// Experimental.
type AwsCdkDepsOptions struct {
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Deprecated: The
	CdkAssert *bool `json:"cdkAssert"`
	// Install the @aws-cdk/assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'
	// Experimental.
	CdkAssertions *bool `json:"cdkAssertions"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Deprecated: For CDK 2.x use 'peerDeps' instead
	CdkDependencies *[]*string `json:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `json:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' instead
	CdkTestDependencies *[]*string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `json:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `json:"constructsVersion"`
	// The type of dependency to use for runtime AWS CDK and `constructs` modules.
	//
	// For libraries, use peer dependencies and for apps use runtime dependencies.
	// Experimental.
	DependencyType projen.DependencyType `json:"dependencyType"`
}

// AWS CDK app in Java.
// Experimental.
type AwsCdkJavaApp interface {
	java.JavaProject
	AutoApprove() github.AutoApprove
	BuildTask() projen.Task
	CdkConfig() CdkConfig
	CdkTasks() CdkTasks
	CdkVersion() *string
	Compile() java.MavenCompile
	CompileTask() projen.Task
	Components() *[]projen.Component
	DefaultTask() projen.Task
	Deps() projen.Dependencies
	DevContainer() vscode.DevContainer
	Distdir() *string
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	Junit() java.Junit
	Logger() projen.Logger
	MainClass() *string
	MainClassName() *string
	MainPackage() *string
	Name() *string
	Outdir() *string
	PackageTask() projen.Task
	Packaging() java.MavenPackaging
	Parent() projen.Project
	Pom() java.Pom
	PostCompileTask() projen.Task
	PreCompileTask() projen.Task
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Projenrc() java.Projenrc
	Root() projen.Project
	Tasks() projen.Tasks
	TestTask() projen.Task
	Vscode() vscode.VsCode
	AddCdkDependency(modules ...*string)
	AddDependency(spec *string)
	AddExcludeFromCleanup(globs ...*string)
	AddGitIgnore(pattern *string)
	AddPackageIgnore(_pattern *string)
	AddPlugin(spec *string, options *java.PluginOptions) *projen.Dependency
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	AddTestDependency(spec *string)
	AddTip(message *string)
	AnnotateGenerated(glob *string)
	PostSynthesize()
	PreSynthesize()
	RemoveTask(name *string) projen.Task
	RunTaskCommand(task projen.Task) *string
	Synth()
	TryFindFile(filePath *string) projen.FileBase
	TryFindJsonFile(filePath *string) projen.JsonFile
	TryFindObjectFile(filePath *string) projen.ObjectFile
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

func (j *jsiiProxy_AwsCdkJavaApp) CdkTasks() CdkTasks {
	var returns CdkTasks
	_jsii_.Get(
		j,
		"cdkTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkJavaApp) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
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

// Adds an AWS CDK module dependencies.
// Experimental.
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

// Adds a runtime dependency.
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		a,
		"addDependency",
		[]interface{}{spec},
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
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

// Adds a .gitignore pattern.
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

// Adds a build plugin to the pom.
//
// The plug in is also added as a BUILD dep to the project.
// Experimental.
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

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
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

// Adds a test dependency.
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) AddTestDependency(spec *string) {
	_jsii_.InvokeVoid(
		a,
		"addTestDependency",
		[]interface{}{spec},
	)
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (a *jsiiProxy_AwsCdkJavaApp) AddTip(message *string) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AwsCdkJavaApp) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
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

// Returns the shell command to execute in order to run a task.
//
// By default, this is `npx projen@<version> <task>`
// Experimental.
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

// Synthesize all project files into `outdir`.
//
// 1. Call "this.preSynthesize()"
// 2. Delete all generated files
// 3. Synthesize all sub-projects
// 4. Synthesize all components of this project
// 5. Call "postSynthesize()" for all components of this project
// 6. Call "this.postSynthesize()"
// Experimental.
func (a *jsiiProxy_AwsCdkJavaApp) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
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

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
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

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
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

// Experimental.
type AwsCdkJavaAppOptions struct {
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
	ArtifactId *string `json:"artifactId"`
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
	GroupId *string `json:"groupId"`
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
	Version *string `json:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description *string `json:"description"`
	// Project packaging format.
	// Experimental.
	Packaging *string `json:"packaging"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url *string `json:"url"`
	// Compile options.
	// Experimental.
	CompileOptions *java.MavenCompileOptions `json:"compileOptions"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `json:"deps"`
	// Final artifact output directory.
	// Experimental.
	Distdir *string `json:"distdir"`
	// Include junit tests.
	// Experimental.
	Junit *bool `json:"junit"`
	// junit options.
	// Experimental.
	JunitOptions *java.JunitOptions `json:"junitOptions"`
	// Packaging options.
	// Experimental.
	PackagingOptions *java.MavenPackagingOptions `json:"packagingOptions"`
	// Use projenrc in java.
	//
	// This will install `projen` as a java dependency and will add a `synth` task which
	// will compile & execute `main()` from `src/main/java/projenrc.java`.
	// Experimental.
	ProjenrcJava *bool `json:"projenrcJava"`
	// Options related to projenrc in java.
	// Experimental.
	ProjenrcJavaOptions *java.ProjenrcOptions `json:"projenrcJavaOptions"`
	// List of test dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addTestDependency()`.
	// Experimental.
	TestDeps *[]*string `json:"testDeps"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample *bool `json:"sample"`
	// The java package to use for the code sample.
	// Experimental.
	SampleJavaPackage *string `json:"sampleJavaPackage"`
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `json:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `json:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]*string `json:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `json:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `json:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `json:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `json:"watchIncludes"`
	// AWS CDK version to use (you can use semantic versioning).
	// Experimental.
	CdkVersion *string `json:"cdkVersion"`
	// The name of the Java class with the static `main()` method.
	//
	// This method
	// should call `app.synth()` on the CDK app.
	// Experimental.
	MainClass *string `json:"mainClass"`
	// Which AWS CDK modules this app uses.
	//
	// The `core` module is included by
	// default and you can add additional modules here by stating only the package
	// name (e.g. `aws-lambda`).
	// Experimental.
	CdkDependencies *[]*string `json:"cdkDependencies"`
}

// AWS CDK app in TypeScript.
// Experimental.
type AwsCdkTypeScriptApp interface {
	typescript.TypeScriptAppProject
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	AppEntrypoint() *string
	AutoApprove() github.AutoApprove
	AutoMerge() github.AutoMerge
	BuildTask() projen.Task
	BuildWorkflow() github.TaskWorkflow
	BuildWorkflowJobId() *string
	Bundler() javascript.Bundler
	CdkConfig() CdkConfig
	CdkDeps() AwsCdkDeps
	CdkTasks() CdkTasks
	CdkVersion() *string
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
	Vscode() vscode.VsCode
	WatchTask() projen.Task
	AddBins(bins *map[string]*string)
	AddBundledDeps(deps ...*string)
	AddCdkDependency(modules ...*string)
	AddCompileCommand(commands ...*string)
	AddDeps(deps ...*string)
	AddDevDeps(deps ...*string)
	AddExcludeFromCleanup(globs ...*string)
	AddFields(fields *map[string]interface{})
	AddGitIgnore(pattern *string)
	AddKeywords(keywords ...*string)
	AddPackageIgnore(pattern *string)
	AddPeerDeps(deps ...*string)
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

func (j *jsiiProxy_AwsCdkTypeScriptApp) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
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

func (j *jsiiProxy_AwsCdkTypeScriptApp) BuildWorkflow() github.TaskWorkflow {
	var returns github.TaskWorkflow
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

func (j *jsiiProxy_AwsCdkTypeScriptApp) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
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

// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		a,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
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

// Adds an AWS CDK module dependencies.
// Experimental.
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

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
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

// Defines normal dependencies.
// Experimental.
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

// Defines development/test dependencies.
// Experimental.
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

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
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

// Directly set fields in `package.json`.
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
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

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		a,
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

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
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

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
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

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (a *jsiiProxy_AwsCdkTypeScriptApp) AddTip(message *string) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AwsCdkTypeScriptApp) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		a,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
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

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		a,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
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

// Returns the shell command to execute in order to run a task.
//
// This will
// typically be `npx projen TASK`.
// Experimental.
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

// Replaces the contents of an npm package.json script.
// Experimental.
func (a *jsiiProxy_AwsCdkTypeScriptApp) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AwsCdkTypeScriptApp) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
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

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
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

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
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

// Experimental.
type AwsCdkTypeScriptAppOptions struct {
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
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper *bool `json:"antitamper"`
	// A directory which will contain artifacts to be published to npm.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory"`
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
	// Experimental.
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
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Experimental.
	Package *bool `json:"package"`
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
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Experimental.
	BuildCommand *string `json:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `json:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]*string `json:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `json:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `json:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `json:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `json:"watchIncludes"`
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Deprecated: The
	CdkAssert *bool `json:"cdkAssert"`
	// Install the @aws-cdk/assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'
	// Experimental.
	CdkAssertions *bool `json:"cdkAssertions"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Deprecated: For CDK 2.x use 'peerDeps' instead
	CdkDependencies *[]*string `json:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `json:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' instead
	CdkTestDependencies *[]*string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `json:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `json:"constructsVersion"`
	// The CDK app's entrypoint (relative to the source directory, which is "src" by default).
	// Experimental.
	AppEntrypoint *string `json:"appEntrypoint"`
	// Automatically adds an `awscdk.LambdaFunction` for each `.lambda.ts` handler in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Experimental.
	LambdaAutoDiscover *bool `json:"lambdaAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `json:"lambdaOptions"`
}

// Represents cdk.json file.
// Experimental.
type CdkConfig interface {
	projen.Component
	Cdkout() *string
	Json() projen.JsonFile
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (c *jsiiProxy_CdkConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (c *jsiiProxy_CdkConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	BuildCommand *string `json:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `json:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]*string `json:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `json:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `json:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `json:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `json:"watchIncludes"`
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
	BuildCommand *string `json:"buildCommand"`
	// cdk.out directory.
	// Experimental.
	Cdkout *string `json:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context *map[string]*string `json:"context"`
	// Include all feature flags in cdk.json.
	// Experimental.
	FeatureFlags *bool `json:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval ApprovalLevel `json:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Experimental.
	WatchExcludes *[]*string `json:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Experimental.
	WatchIncludes *[]*string `json:"watchIncludes"`
	// The command line to execute in order to synthesize the CDK application (language specific).
	// Experimental.
	App *string `json:"app"`
}

// Adds standard AWS CDK tasks to your project.
// Experimental.
type CdkTasks interface {
	projen.Component
	Deploy() projen.Task
	Destroy() projen.Task
	Diff() projen.Task
	Project() projen.Project
	Synth() projen.Task
	Watch() projen.Task
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (c *jsiiProxy_CdkTasks) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (c *jsiiProxy_CdkTasks) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (c *jsiiProxy_CdkTasks) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

// Deprecated: use `AwsCdkConstructLibrary`
type ConstructLibraryAws interface {
	AwsCdkConstructLibrary
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	AutoApprove() github.AutoApprove
	AutoMerge() github.AutoMerge
	BuildTask() projen.Task
	BuildWorkflow() github.TaskWorkflow
	BuildWorkflowJobId() *string
	Bundler() javascript.Bundler
	CdkDeps() AwsCdkDeps
	CdkVersion() *string
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
	Version() *string
	Vscode() vscode.VsCode
	WatchTask() projen.Task
	AddBins(bins *map[string]*string)
	AddBundledDeps(deps ...*string)
	AddCdkDependencies(deps ...*string)
	AddCdkTestDependencies(deps ...*string)
	AddCompileCommand(commands ...*string)
	AddDeps(deps ...*string)
	AddDevDeps(deps ...*string)
	AddExcludeFromCleanup(globs ...*string)
	AddFields(fields *map[string]interface{})
	AddGitIgnore(pattern *string)
	AddKeywords(keywords ...*string)
	AddPackageIgnore(pattern *string)
	AddPeerDeps(deps ...*string)
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

func (j *jsiiProxy_ConstructLibraryAws) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
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

func (j *jsiiProxy_ConstructLibraryAws) BuildWorkflow() github.TaskWorkflow {
	var returns github.TaskWorkflow
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

func (j *jsiiProxy_ConstructLibraryAws) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
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


// Deprecated: use `AwsCdkConstructLibrary`
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

// Deprecated: use `AwsCdkConstructLibrary`
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

// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) AddBins(bins *map[string]*string) {
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
// Deprecated: use `AwsCdkConstructLibrary`
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

// Adds dependencies to AWS CDK modules.
//
// Since this is a library project, dependencies will be added as peer dependencies.
// Deprecated: Not supported in v2. For v1, use `project.cdkDeps.addV1Dependencies()`
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

// Adds AWS CDK modules as dev dependencies.
// Deprecated: Not supported in v2. For v1, use `project.cdkDeps.addV1DevDependencies()`
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

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
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

// Defines normal dependencies.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Defines development/test dependencies.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Directly set fields in `package.json`.
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		c,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Deprecated: use `AwsCdkConstructLibrary`
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

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) AddPackageIgnore(pattern *string) {
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
// Deprecated: use `AwsCdkConstructLibrary`
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

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Deprecated: use `AwsCdkConstructLibrary`
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

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
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

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (c *jsiiProxy_ConstructLibraryAws) AddTip(message *string) {
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
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		c,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		c,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Returns the shell command to execute in order to run a task.
//
// This will
// typically be `npx projen TASK`.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Replaces the contents of an npm package.json script.
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) SetScript(name *string, command *string) {
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
// Deprecated: use `AwsCdkConstructLibrary`
func (c *jsiiProxy_ConstructLibraryAws) Synth() {
	_jsii_.InvokeVoid(
		c,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Deprecated: use `AwsCdkConstructLibrary`
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

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
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

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Deprecated: use `AwsCdkConstructLibrary`
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

// Deprecated: use `AwsCdkConstructLibraryOptions`
type ConstructLibraryAwsOptions struct {
	// This is the name of your project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Name *string `json:"name"`
	// Configure logging options such as verbosity.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Logging *projen.LoggerOptions `json:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Outdir *string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Parent projen.Project `json:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenCommand *string `json:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenrcJson *bool `json:"projenrcJson"`
	// Options for .projenrc.json.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Clobber *bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DevContainer *bool `json:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Github *bool `json:"github"`
	// Options for GitHub integration.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	GithubOptions *github.GitHubOptions `json:"githubOptions"`
	// Add a Gitpod development environment.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Readme *projen.SampleReadmeProps `json:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Stale *bool `json:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	StaleOptions *github.StaleOptions `json:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Vscode *bool `json:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AllowLibraryDependencies *bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorName *string `json:"authorName"`
	// Author's Organization.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorOrganization *bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorUrl *string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AutoDetectBin *bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Bin *map[string]*string `json:"bin"`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	BundledDeps *[]*string `json:"bundledDeps"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Deps *[]*string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DevDeps *[]*string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Entrypoint *string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Homepage *string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Keywords *[]*string `json:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	License *string `json:"license"`
	// Indicates if a license should be added.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Licensed *bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MaxNodeVersion *string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MinNodeVersion *string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmAccess javascript.NpmAccess `json:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry *string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmRegistryUrl *string `json:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmTokenSecret *string `json:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PackageManager javascript.NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PackageName *string `json:"packageName"`
	// Options for `peerDeps`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PeerDeps *[]*string `json:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Repository *string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	RepositoryDirectory *string `json:"repositoryDirectory"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Scripts *map[string]*string `json:"scripts"`
	// Package's Stability.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Stability *string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Antitamper *bool `json:"antitamper"`
	// A directory which will contain artifacts to be published to npm.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ArtifactsDirectory *string `json:"artifactsDirectory"`
	// Version requirement of `jsii-release` which is used to publish modules to npm.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	JsiiReleaseVersion *string `json:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MajorVersion *float64 `json:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmDistTag *string `json:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Prerelease *string `json:"prerelease"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishTasks *bool `json:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseBranches *map[string]*release.BranchOptions `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `json:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseFailureIssue *bool `json:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseTagPrefix *string `json:"releaseTagPrefix"`
	// The release trigger to use.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseTrigger release.ReleaseTrigger `json:"releaseTrigger"`
	// The name of the default release workflow.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseWorkflowName *string `json:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `json:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	VersionrcOptions *map[string]interface{} `json:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowContainerImage *string `json:"workflowContainerImage"`
	// Github Runner selection labels.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowRunsOn *[]*string `json:"workflowRunsOn"`
	// The name of the main release branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DefaultReleaseBranch *string `json:"defaultReleaseBranch"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AutoApproveProjenUpgrades *bool `json:"autoApproveProjenUpgrades"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AutoApproveUpgrades *bool `json:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	BuildWorkflow *bool `json:"buildWorkflow"`
	// Options for `Bundler`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	BundlerOptions *javascript.BundlerOptions `json:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CodeCov *bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CodeCovTokenSecret *string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CopyrightOwner *string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CopyrightPeriod *string `json:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Dependabot *bool `json:"dependabot"`
	// Options for dependabot.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DependabotOptions *github.DependabotOptions `json:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DepsUpgrade *bool `json:"depsUpgrade"`
	// Options for depsUpgrade.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `json:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Gitignore *[]*string `json:"gitignore"`
	// Setup jest unit tests.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Jest *bool `json:"jest"`
	// Jest options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	JestOptions *javascript.JestOptions `json:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MutableBuild *bool `json:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmignoreEnabled *bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenDevDependency *bool `json:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenrcJs *bool `json:"projenrcJs"`
	// Options for .projenrc.js.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenrcJsOptions *javascript.ProjenrcOptions `json:"projenrcJsOptions"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `autoApproveProjenUpgrades`.
	ProjenUpgradeAutoMerge *bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenUpgradeSecret *string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenVersion *string `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PullRequestTemplate *bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PullRequestTemplateContents *[]*string `json:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Release *bool `json:"release"`
	// Automatically release to npm when new versions are introduced.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseToNpm *bool `json:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowBootstrapSteps *[]interface{} `json:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowGitIdentity *github.GitIdentity `json:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowNodeVersion *string `json:"workflowNodeVersion"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DisableTsconfig *bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Docgen *bool `json:"docgen"`
	// Docs directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DocsDirectory *string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	EntrypointTypes *string `json:"entrypointTypes"`
	// Setup eslint.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Eslint *bool `json:"eslint"`
	// Eslint options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	EslintOptions *javascript.EslintOptions `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Libdir *string `json:"libdir"`
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Package *bool `json:"package"`
	// Use TypeScript for your projenrc file (`.projenrc.ts`).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenrcTs *bool `json:"projenrcTs"`
	// Options for .projenrc.ts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenrcTsOptions *typescript.ProjenrcOptions `json:"projenrcTsOptions"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	SampleCode *bool `json:"sampleCode"`
	// Typescript sources directory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Srcdir *string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	//
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Testdir *string `json:"testdir"`
	// Custom TSConfig.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Tsconfig *javascript.TypescriptConfigOptions `json:"tsconfig"`
	// Custom tsconfig options for the development tsconfig.json file (used for testing).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	TsconfigDev *javascript.TypescriptConfigOptions `json:"tsconfigDev"`
	// The name of the development tsconfig.json file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	TsconfigDevFile *string `json:"tsconfigDevFile"`
	// TypeScript version to use.
	//
	// NOTE: Typescript is not semantically versioned and should remain on the
	// same minor, so we recommend using a `~` dependency (e.g. `~1.2.3`).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	TypescriptVersion *string `json:"typescriptVersion"`
	// The name of the library author.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Author *string `json:"author"`
	// Email or URL of the library author.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorAddress *string `json:"authorAddress"`
	// Git repository URL.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	RepositoryUrl *string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	//
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Compat *bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ExcludeTypescript *[]*string `json:"excludeTypescript"`
	// Publish Go bindings to a git repository.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToGo *cdk.JsiiGoTarget `json:"publishToGo"`
	// Publish to maven.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToMaven *cdk.JsiiJavaTarget `json:"publishToMaven"`
	// Publish to NuGet.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToNuget *cdk.JsiiDotNetTarget `json:"publishToNuget"`
	// Publish to pypi.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToPypi *cdk.JsiiPythonTarget `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python *cdk.JsiiPythonTarget `json:"python"`
	// Deprecated: use `AwsCdkConstructLibraryOptions`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Catalog *cdk.Catalog `json:"catalog"`
	// Minimum version of the AWS CDK to depend on.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkVersion *string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Deprecated: The
	CdkAssert *bool `json:"cdkAssert"`
	// Install the @aws-cdk/assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkAssertions *bool `json:"cdkAssertions"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Deprecated: For CDK 2.x use 'peerDeps' instead
	CdkDependencies *[]*string `json:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `json:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' instead
	CdkTestDependencies *[]*string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkVersionPinning *bool `json:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ConstructsVersion *string `json:"constructsVersion"`
	// Automatically adds an `aws_lambda.Function` for each `.lambda.ts` handler in your source tree. If this is disabled, you either need to explicitly call `aws_lambda.Function.autoDiscover()` or define a `new aws_lambda.Function()` for each handler.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	LambdaAutoDiscover *bool `json:"lambdaAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	LambdaOptions *LambdaFunctionCommonOptions `json:"lambdaOptions"`
}

// Cloud integration tests.
// Experimental.
type IntegrationTest interface {
	projen.Component
	AssertTask() projen.Task
	DeployTask() projen.Task
	DestroyTask() projen.Task
	Project() projen.Project
	SnapshotTask() projen.Task
	WatchTask() projen.Task
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for IntegrationTest
type jsiiProxy_IntegrationTest struct {
	internal.Type__projenComponent
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

func (j *jsiiProxy_IntegrationTest) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (i *jsiiProxy_IntegrationTest) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (i *jsiiProxy_IntegrationTest) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (i *jsiiProxy_IntegrationTest) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type IntegrationTestCommonOptions struct {
	// Destroy the test app after a successful deployment.
	//
	// If disabled, leaves the
	// app deployed in the dev account.
	// Experimental.
	DestroyAfterDeploy *bool `json:"destroyAfterDeploy"`
}

// Options for `IntegrationTest`.
// Experimental.
type IntegrationTestOptions struct {
	// Destroy the test app after a successful deployment.
	//
	// If disabled, leaves the
	// app deployed in the dev account.
	// Experimental.
	DestroyAfterDeploy *bool `json:"destroyAfterDeploy"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `json:"cdkDeps"`
	// A path from the project root directory to a TypeScript file which contains the integration test app.
	//
	// This is relative to the root directory of the project.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Entrypoint *string `json:"entrypoint"`
	// The path of the tsconfig.json file to use when running integration test cdk apps.
	// Experimental.
	TsconfigPath *string `json:"tsconfigPath"`
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
// TODO: EXAMPLE
//
// Experimental.
type LambdaFunction interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (l *jsiiProxy_LambdaFunction) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (l *jsiiProxy_LambdaFunction) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	// Bundling options for this AWS Lambda function.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `json:"bundlingOptions"`
	// The node.js version to target.
	// Experimental.
	Runtime LambdaRuntime `json:"runtime"`
}

// Options for `Function`.
// Experimental.
type LambdaFunctionOptions struct {
	// Bundling options for this AWS Lambda function.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `json:"bundlingOptions"`
	// The node.js version to target.
	// Experimental.
	Runtime LambdaRuntime `json:"runtime"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `json:"cdkDeps"`
	// A path from the project root directory to a TypeScript file which contains the AWS Lambda handler entrypoint (exports a `handler` function).
	//
	// This is relative to the root directory of the project.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Entrypoint *string `json:"entrypoint"`
	// The name of the generated TypeScript source file.
	//
	// This file should also be
	// under the source tree.
	// Experimental.
	ConstructFile *string `json:"constructFile"`
	// The name of the generated `lambda.Function` subclass.
	// Experimental.
	ConstructName *string `json:"constructName"`
}

// The runtime for the AWS Lambda function.
// Experimental.
type LambdaRuntime interface {
	EsbuildPlatform() *string
	EsbuildTarget() *string
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

