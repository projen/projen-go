package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/typescript"
	"github.com/projen/projen-go/projen/vscode"
	"github.com/projen/projen-go/projen/web/internal"
)

// Experimental.
type NextComponent interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for NextComponent
type jsiiProxy_NextComponent struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NextComponent) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNextComponent(project javascript.NodeProject, options *NextComponentOptions) NextComponent {
	_init_.Initialize()

	j := jsiiProxy_NextComponent{}

	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNextComponent_Override(n NextComponent, project javascript.NodeProject, options *NextComponentOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		n,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (n *jsiiProxy_NextComponent) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (n *jsiiProxy_NextComponent) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (n *jsiiProxy_NextComponent) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type NextComponentOptions struct {
	// Setup Tailwind as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind *bool `json:"tailwind"`
	// Whether to apply options specific for TypeScript Next.js projects.
	// Experimental.
	Typescript *bool `json:"typescript"`
}

// Experimental.
type NextJsCommonProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir *string `json:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind *bool `json:"tailwind"`
}

// Next.js project without TypeScript.
// Experimental.
type NextJsProject interface {
	javascript.NodeProject
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	ArtifactsDirectory() *string
	Assetsdir() *string
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
	Entrypoint() *string
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	InstallWorkflowSteps() *[]*workflows.JobStep
	Jest() javascript.Jest
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
	Tailwind() *bool
	Tasks() projen.Tasks
	TestTask() projen.Task
	UpgradeWorkflow() javascript.UpgradeDependencies
	Vscode() vscode.VsCode
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

// The jsii proxy struct for NextJsProject
type jsiiProxy_NextJsProject struct {
	internal.Type__javascriptNodeProject
}

func (j *jsiiProxy_NextJsProject) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Assetsdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Tailwind() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"tailwind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewNextJsProject(options *NextJsProjectOptions) NextJsProject {
	_init_.Initialize()

	j := jsiiProxy_NextJsProject{}

	_jsii_.Create(
		"projen.web.NextJsProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewNextJsProject_Override(n NextJsProject, options *NextJsProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.NextJsProject",
		[]interface{}{options},
		n,
	)
}

func NextJsProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.web.NextJsProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

// Experimental.
func (n *jsiiProxy_NextJsProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		n,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addBundledDeps",
		args,
	)
}

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
func (n *jsiiProxy_NextJsProject) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addCompileCommand",
		args,
	)
}

// Defines normal dependencies.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDeps",
		args,
	)
}

// Defines development/test dependencies.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDevDeps",
		args,
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addExcludeFromCleanup",
		args,
	)
}

// Directly set fields in `package.json`.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
func (n *jsiiProxy_NextJsProject) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addKeywords",
		args,
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NextJsProject) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addPeerDeps",
		args,
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
func (n *jsiiProxy_NextJsProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		n,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
func (n *jsiiProxy_NextJsProject) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addTestCommand",
		args,
	)
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (n *jsiiProxy_NextJsProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NextJsProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		n,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
func (n *jsiiProxy_NextJsProject) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NextJsProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (n *jsiiProxy_NextJsProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (n *jsiiProxy_NextJsProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (n *jsiiProxy_NextJsProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NextJsProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Replaces the contents of an npm package.json script.
// Experimental.
func (n *jsiiProxy_NextJsProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NextJsProject) Synth() {
	_jsii_.InvokeVoid(
		n,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (n *jsiiProxy_NextJsProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		n,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (n *jsiiProxy_NextJsProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		n,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (n *jsiiProxy_NextJsProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		n,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type NextJsProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir *string `json:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind *bool `json:"tailwind"`
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
	// Generate one-time sample in `pages/` and `public/` if there are no files there.
	// Experimental.
	SampleCode *bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir *string `json:"srcdir"`
}

// Next.js project with TypeScript.
// Experimental.
type NextJsTypeScriptProject interface {
	typescript.TypeScriptAppProject
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	ArtifactsDirectory() *string
	Assetsdir() *string
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
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Publisher() release.Publisher
	Release() release.Release
	Root() projen.Project
	RunScriptCommand() *string
	Srcdir() *string
	Tailwind() *bool
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

// The jsii proxy struct for NextJsTypeScriptProject
type jsiiProxy_NextJsTypeScriptProject struct {
	internal.Type__typescriptTypeScriptAppProject
}

func (j *jsiiProxy_NextJsTypeScriptProject) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Assetsdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Tailwind() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"tailwind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextJsTypeScriptProject) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewNextJsTypeScriptProject(options *NextJsTypeScriptProjectOptions) NextJsTypeScriptProject {
	_init_.Initialize()

	j := jsiiProxy_NextJsTypeScriptProject{}

	_jsii_.Create(
		"projen.web.NextJsTypeScriptProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewNextJsTypeScriptProject_Override(n NextJsTypeScriptProject, options *NextJsTypeScriptProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.NextJsTypeScriptProject",
		[]interface{}{options},
		n,
	)
}

func NextJsTypeScriptProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.web.NextJsTypeScriptProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		n,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addBundledDeps",
		args,
	)
}

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
func (n *jsiiProxy_NextJsTypeScriptProject) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addCompileCommand",
		args,
	)
}

// Defines normal dependencies.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDeps",
		args,
	)
}

// Defines development/test dependencies.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDevDeps",
		args,
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addExcludeFromCleanup",
		args,
	)
}

// Directly set fields in `package.json`.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addKeywords",
		args,
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NextJsTypeScriptProject) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addPeerDeps",
		args,
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		n,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
func (n *jsiiProxy_NextJsTypeScriptProject) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addTestCommand",
		args,
	)
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (n *jsiiProxy_NextJsTypeScriptProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NextJsTypeScriptProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		n,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NextJsTypeScriptProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NextJsTypeScriptProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Replaces the contents of an npm package.json script.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NextJsTypeScriptProject) Synth() {
	_jsii_.InvokeVoid(
		n,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		n,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (n *jsiiProxy_NextJsTypeScriptProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		n,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (n *jsiiProxy_NextJsTypeScriptProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		n,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type NextJsTypeScriptProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir *string `json:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind *bool `json:"tailwind"`
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
}

// Declares a PostCSS dependency with a default config file.
// Experimental.
type PostCss interface {
	File() projen.JsonFile
	FileName() *string
	Tailwind() TailwindConfig
}

// The jsii proxy struct for PostCss
type jsiiProxy_PostCss struct {
	_ byte // padding
}

func (j *jsiiProxy_PostCss) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostCss) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostCss) Tailwind() TailwindConfig {
	var returns TailwindConfig
	_jsii_.Get(
		j,
		"tailwind",
		&returns,
	)
	return returns
}


// Experimental.
func NewPostCss(project javascript.NodeProject, options *PostCssOptions) PostCss {
	_init_.Initialize()

	j := jsiiProxy_PostCss{}

	_jsii_.Create(
		"projen.web.PostCss",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPostCss_Override(p PostCss, project javascript.NodeProject, options *PostCssOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.PostCss",
		[]interface{}{project, options},
		p,
	)
}

// Experimental.
type PostCssOptions struct {
	// Experimental.
	FileName *string `json:"fileName"`
	// Install Tailwind CSS as a PostCSS plugin.
	// Experimental.
	Tailwind *bool `json:"tailwind"`
	// Tailwind CSS options.
	// Experimental.
	TailwindOptions *TailwindConfigOptions `json:"tailwindOptions"`
}

// Experimental.
type ReactComponent interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for ReactComponent
type jsiiProxy_ReactComponent struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_ReactComponent) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewReactComponent(project javascript.NodeProject, options *ReactComponentOptions) ReactComponent {
	_init_.Initialize()

	j := jsiiProxy_ReactComponent{}

	_jsii_.Create(
		"projen.web.ReactComponent",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewReactComponent_Override(r ReactComponent, project javascript.NodeProject, options *ReactComponentOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.ReactComponent",
		[]interface{}{project, options},
		r,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (r *jsiiProxy_ReactComponent) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (r *jsiiProxy_ReactComponent) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (r *jsiiProxy_ReactComponent) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type ReactComponentOptions struct {
	// Rewire webpack configuration.
	//
	// Use this property to override webpack configuration properties provided
	// by create-react-app, without needing to eject.
	//
	// This property will create a `config-overrides.js` file in your root directory,
	// which will contain the desired rewiring code.
	//
	// To **override** the configuration, you can provide simple key value pairs.
	// Keys take the form of js code directives that traverse to the desired property.
	// Values should be JSON serializable objects.
	//
	// For example, the following config:
	//
	// ```json
	// rewire: { "module.unknownContextCritical": false }
	// ```
	//
	// Will translate to the following `config-overrides.js` file:
	//
	// ```js
	// module.exports = function override(config, env) {
	//    config.module.unknownContextCritical = false;
	// }
	// ```
	// See: https://github.com/timarney/react-app-rewired
	//
	// Experimental.
	Rewire *map[string]interface{} `json:"rewire"`
	// Whether to apply options specific for TypeScript React projects.
	// Experimental.
	Typescript *bool `json:"typescript"`
}

// React project without TypeScript.
// Experimental.
type ReactProject interface {
	javascript.NodeProject
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	ArtifactsDirectory() *string
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
	Entrypoint() *string
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	InstallWorkflowSteps() *[]*workflows.JobStep
	Jest() javascript.Jest
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
	TestTask() projen.Task
	UpgradeWorkflow() javascript.UpgradeDependencies
	Vscode() vscode.VsCode
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

// The jsii proxy struct for ReactProject
type jsiiProxy_ReactProject struct {
	internal.Type__javascriptNodeProject
}

func (j *jsiiProxy_ReactProject) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewReactProject(options *ReactProjectOptions) ReactProject {
	_init_.Initialize()

	j := jsiiProxy_ReactProject{}

	_jsii_.Create(
		"projen.web.ReactProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewReactProject_Override(r ReactProject, options *ReactProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.ReactProject",
		[]interface{}{options},
		r,
	)
}

func ReactProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.web.ReactProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

// Experimental.
func (r *jsiiProxy_ReactProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		r,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
func (r *jsiiProxy_ReactProject) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addBundledDeps",
		args,
	)
}

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
func (r *jsiiProxy_ReactProject) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addCompileCommand",
		args,
	)
}

// Defines normal dependencies.
// Experimental.
func (r *jsiiProxy_ReactProject) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDeps",
		args,
	)
}

// Defines development/test dependencies.
// Experimental.
func (r *jsiiProxy_ReactProject) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDevDeps",
		args,
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
func (r *jsiiProxy_ReactProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addExcludeFromCleanup",
		args,
	)
}

// Directly set fields in `package.json`.
// Experimental.
func (r *jsiiProxy_ReactProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (r *jsiiProxy_ReactProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		r,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
func (r *jsiiProxy_ReactProject) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addKeywords",
		args,
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (r *jsiiProxy_ReactProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_ReactProject) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addPeerDeps",
		args,
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
func (r *jsiiProxy_ReactProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		r,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
func (r *jsiiProxy_ReactProject) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addTestCommand",
		args,
	)
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (r *jsiiProxy_ReactProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_ReactProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		r,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
func (r *jsiiProxy_ReactProject) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ReactProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (r *jsiiProxy_ReactProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (r *jsiiProxy_ReactProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		r,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (r *jsiiProxy_ReactProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ReactProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Replaces the contents of an npm package.json script.
// Experimental.
func (r *jsiiProxy_ReactProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_ReactProject) Synth() {
	_jsii_.InvokeVoid(
		r,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (r *jsiiProxy_ReactProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		r,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (r *jsiiProxy_ReactProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		r,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (r *jsiiProxy_ReactProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		r,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type ReactProjectOptions struct {
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
	// Rewire webpack configuration.
	//
	// Use this property to override webpack configuration properties provided
	// by create-react-app, without needing to eject.
	//
	// This property will create a `config-overrides.js` file in your root directory,
	// which will contain the desired rewiring code.
	//
	// To **override** the configuration, you can provide simple key value pairs.
	// Keys take the form of js code directives that traverse to the desired property.
	// Values should be JSON serializable objects.
	//
	// For example, the following config:
	//
	// ```json
	// rewire: { "module.unknownContextCritical": false }
	// ```
	//
	// Will translate to the following `config-overrides.js` file:
	//
	// ```js
	// module.exports = function override(config, env) {
	//    config.module.unknownContextCritical = false;
	// }
	// ```
	// See: https://github.com/timarney/react-app-rewired
	//
	// Experimental.
	Rewire *map[string]interface{} `json:"rewire"`
	// Generate one-time sample in `src/` and `public/` if there are no files there.
	// Experimental.
	SampleCode *bool `json:"sampleCode"`
	// Source directory.
	// Experimental.
	Srcdir *string `json:"srcdir"`
}

// Experimental.
type ReactRewireOptions struct {
	// Rewire webpack configuration.
	//
	// Use this property to override webpack configuration properties provided
	// by create-react-app, without needing to eject.
	//
	// This property will create a `config-overrides.js` file in your root directory,
	// which will contain the desired rewiring code.
	//
	// To **override** the configuration, you can provide simple key value pairs.
	// Keys take the form of js code directives that traverse to the desired property.
	// Values should be JSON serializable objects.
	//
	// For example, the following config:
	//
	// ```json
	// rewire: { "module.unknownContextCritical": false }
	// ```
	//
	// Will translate to the following `config-overrides.js` file:
	//
	// ```js
	// module.exports = function override(config, env) {
	//    config.module.unknownContextCritical = false;
	// }
	// ```
	// See: https://github.com/timarney/react-app-rewired
	//
	// Experimental.
	Rewire *map[string]interface{} `json:"rewire"`
}

// Experimental.
type ReactTypeDef interface {
	projen.FileBase
	AbsolutePath() *string
	Executable() *bool
	SetExecutable(val *bool)
	Path() *string
	Project() projen.Project
	Readonly() *bool
	SetReadonly(val *bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_arg projen.IResolver) *string
}

// The jsii proxy struct for ReactTypeDef
type jsiiProxy_ReactTypeDef struct {
	internal.Type__projenFileBase
}

func (j *jsiiProxy_ReactTypeDef) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewReactTypeDef(project ReactTypeScriptProject, filePath *string, options *ReactTypeDefOptions) ReactTypeDef {
	_init_.Initialize()

	j := jsiiProxy_ReactTypeDef{}

	_jsii_.Create(
		"projen.web.ReactTypeDef",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewReactTypeDef_Override(r ReactTypeDef, project ReactTypeScriptProject, filePath *string, options *ReactTypeDefOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.ReactTypeDef",
		[]interface{}{project, filePath, options},
		r,
	)
}

func (j *jsiiProxy_ReactTypeDef) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_ReactTypeDef) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func ReactTypeDef_PROJEN_MARKER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.web.ReactTypeDef",
		"PROJEN_MARKER",
		&returns,
	)
	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (r *jsiiProxy_ReactTypeDef) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (r *jsiiProxy_ReactTypeDef) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (r *jsiiProxy_ReactTypeDef) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (r *jsiiProxy_ReactTypeDef) SynthesizeContent(_arg projen.IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// Experimental.
type ReactTypeDefOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	//
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed *bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore *bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable *bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `json:"readonly"`
}

// React project with TypeScript.
// Experimental.
type ReactTypeScriptProject interface {
	typescript.TypeScriptAppProject
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	ArtifactsDirectory() *string
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
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Publisher() release.Publisher
	ReactTypeDef() ReactTypeDef
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

// The jsii proxy struct for ReactTypeScriptProject
type jsiiProxy_ReactTypeScriptProject struct {
	internal.Type__typescriptTypeScriptAppProject
}

func (j *jsiiProxy_ReactTypeScriptProject) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Bundler() javascript.Bundler {
	var returns javascript.Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Docgen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) DocsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Eslint() javascript.Eslint {
	var returns javascript.Eslint
	_jsii_.Get(
		j,
		"eslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Jest() javascript.Jest {
	var returns javascript.Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Libdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Package() javascript.NodePackage {
	var returns javascript.NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) PackageManager() javascript.NodePackageManager {
	var returns javascript.NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) ReactTypeDef() ReactTypeDef {
	var returns ReactTypeDef
	_jsii_.Get(
		j,
		"reactTypeDef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Srcdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) TsconfigDev() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigDev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) TsconfigEslint() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfigEslint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) UpgradeWorkflow() javascript.UpgradeDependencies {
	var returns javascript.UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeScriptProject) WatchTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewReactTypeScriptProject(options *ReactTypeScriptProjectOptions) ReactTypeScriptProject {
	_init_.Initialize()

	j := jsiiProxy_ReactTypeScriptProject{}

	_jsii_.Create(
		"projen.web.ReactTypeScriptProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewReactTypeScriptProject_Override(r ReactTypeScriptProject, options *ReactTypeScriptProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.ReactTypeScriptProject",
		[]interface{}{options},
		r,
	)
}

func ReactTypeScriptProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.web.ReactTypeScriptProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		r,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addBundledDeps",
		args,
	)
}

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
func (r *jsiiProxy_ReactTypeScriptProject) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addCompileCommand",
		args,
	)
}

// Defines normal dependencies.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDeps",
		args,
	)
}

// Defines development/test dependencies.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDevDeps",
		args,
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addExcludeFromCleanup",
		args,
	)
}

// Directly set fields in `package.json`.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		r,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addKeywords",
		args,
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_ReactTypeScriptProject) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addPeerDeps",
		args,
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		r,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
func (r *jsiiProxy_ReactTypeScriptProject) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addTestCommand",
		args,
	)
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (r *jsiiProxy_ReactTypeScriptProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_ReactTypeScriptProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		r,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ReactTypeScriptProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		r,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ReactTypeScriptProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Replaces the contents of an npm package.json script.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_ReactTypeScriptProject) Synth() {
	_jsii_.InvokeVoid(
		r,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		r,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (r *jsiiProxy_ReactTypeScriptProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		r,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (r *jsiiProxy_ReactTypeScriptProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		r,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type ReactTypeScriptProjectOptions struct {
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
	// Rewire webpack configuration.
	//
	// Use this property to override webpack configuration properties provided
	// by create-react-app, without needing to eject.
	//
	// This property will create a `config-overrides.js` file in your root directory,
	// which will contain the desired rewiring code.
	//
	// To **override** the configuration, you can provide simple key value pairs.
	// Keys take the form of js code directives that traverse to the desired property.
	// Values should be JSON serializable objects.
	//
	// For example, the following config:
	//
	// ```json
	// rewire: { "module.unknownContextCritical": false }
	// ```
	//
	// Will translate to the following `config-overrides.js` file:
	//
	// ```js
	// module.exports = function override(config, env) {
	//    config.module.unknownContextCritical = false;
	// }
	// ```
	// See: https://github.com/timarney/react-app-rewired
	//
	// Experimental.
	Rewire *map[string]interface{} `json:"rewire"`
}

// Declares a Tailwind CSS configuration file.
//
// There are multiple ways to add Tailwind CSS in your node project - see:
// https://tailwindcss.com/docs/installation
// See: PostCss
//
// Experimental.
type TailwindConfig interface {
	File() projen.JsonFile
	FileName() *string
}

// The jsii proxy struct for TailwindConfig
type jsiiProxy_TailwindConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_TailwindConfig) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TailwindConfig) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTailwindConfig(project javascript.NodeProject, options *TailwindConfigOptions) TailwindConfig {
	_init_.Initialize()

	j := jsiiProxy_TailwindConfig{}

	_jsii_.Create(
		"projen.web.TailwindConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTailwindConfig_Override(t TailwindConfig, project javascript.NodeProject, options *TailwindConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.TailwindConfig",
		[]interface{}{project, options},
		t,
	)
}

// Experimental.
type TailwindConfigOptions struct {
	// Experimental.
	FileName *string `json:"fileName"`
}

