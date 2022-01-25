package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/java/internal"
	"github.com/projen/projen-go/projen/vscode"
)

// Java project.
// Experimental.
type JavaProject interface {
	github.GitHubProject
	AutoApprove() github.AutoApprove
	BuildTask() projen.Task
	Compile() MavenCompile
	CompileTask() projen.Task
	Components() *[]projen.Component
	DefaultTask() projen.Task
	Deps() projen.Dependencies
	DevContainer() vscode.DevContainer
	Distdir() *string
	Ejected() *bool
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	Junit() Junit
	Logger() projen.Logger
	Name() *string
	Outdir() *string
	PackageTask() projen.Task
	Packaging() MavenPackaging
	Parent() projen.Project
	Pom() Pom
	PostCompileTask() projen.Task
	PreCompileTask() projen.Task
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Projenrc() Projenrc
	Root() projen.Project
	Tasks() projen.Tasks
	TestTask() projen.Task
	Vscode() vscode.VsCode
	AddDependency(spec *string)
	AddExcludeFromCleanup(globs ...*string)
	AddGitIgnore(pattern *string)
	AddPackageIgnore(_pattern *string)
	AddPlugin(spec *string, options *PluginOptions) *projen.Dependency
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

// Adds a runtime dependency.
// Experimental.
func (j *jsiiProxy_JavaProject) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		j,
		"addDependency",
		[]interface{}{spec},
	)
}

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
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

// Adds a .gitignore pattern.
// Experimental.
func (j *jsiiProxy_JavaProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (j *jsiiProxy_JavaProject) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

// Adds a build plugin to the pom.
//
// The plug in is also added as a BUILD dep to the project.
// Experimental.
func (j *jsiiProxy_JavaProject) AddPlugin(spec *string, options *PluginOptions) *projen.Dependency {
	var returns *projen.Dependency

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JavaProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		j,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Adds a test dependency.
// Experimental.
func (j *jsiiProxy_JavaProject) AddTestDependency(spec *string) {
	_jsii_.InvokeVoid(
		j,
		"addTestDependency",
		[]interface{}{spec},
	)
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (j *jsiiProxy_JavaProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		j,
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
func (j *jsiiProxy_JavaProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		j,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (j *jsiiProxy_JavaProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (j *jsiiProxy_JavaProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
func (j *jsiiProxy_JavaProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JavaProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		j,
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
func (j *jsiiProxy_JavaProject) Synth() {
	_jsii_.InvokeVoid(
		j,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
func (j *jsiiProxy_JavaProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		j,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (j *jsiiProxy_JavaProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		j,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
func (j *jsiiProxy_JavaProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		j,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Options for `JavaProject`.
// Experimental.
type JavaProjectCommonOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level
	ProjectType projen.ProjectType `json:"projectType" yaml:"projectType"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Experimental.
	ProjenTokenSecret *string `json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode" yaml:"vscode"`
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
	ArtifactId *string `json:"artifactId" yaml:"artifactId"`
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
	GroupId *string `json:"groupId" yaml:"groupId"`
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
	Version *string `json:"version" yaml:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Project packaging format.
	// Experimental.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url *string `json:"url" yaml:"url"`
	// Compile options.
	// Experimental.
	CompileOptions *MavenCompileOptions `json:"compileOptions" yaml:"compileOptions"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `json:"deps" yaml:"deps"`
	// Final artifact output directory.
	// Experimental.
	Distdir *string `json:"distdir" yaml:"distdir"`
	// Include junit tests.
	// Experimental.
	Junit *bool `json:"junit" yaml:"junit"`
	// junit options.
	// Experimental.
	JunitOptions *JunitOptions `json:"junitOptions" yaml:"junitOptions"`
	// Packaging options.
	// Experimental.
	PackagingOptions *MavenPackagingOptions `json:"packagingOptions" yaml:"packagingOptions"`
	// Use projenrc in java.
	//
	// This will install `projen` as a java dependency and will add a `synth` task which
	// will compile & execute `main()` from `src/main/java/projenrc.java`.
	// Experimental.
	ProjenrcJava *bool `json:"projenrcJava" yaml:"projenrcJava"`
	// Options related to projenrc in java.
	// Experimental.
	ProjenrcJavaOptions *ProjenrcOptions `json:"projenrcJavaOptions" yaml:"projenrcJavaOptions"`
	// List of test dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addTestDependency()`.
	// Experimental.
	TestDeps *[]*string `json:"testDeps" yaml:"testDeps"`
}

// Options for `JavaProject`.
// Experimental.
type JavaProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level
	ProjectType projen.ProjectType `json:"projectType" yaml:"projectType"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Experimental.
	ProjenTokenSecret *string `json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode" yaml:"vscode"`
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
	ArtifactId *string `json:"artifactId" yaml:"artifactId"`
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
	GroupId *string `json:"groupId" yaml:"groupId"`
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
	Version *string `json:"version" yaml:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Project packaging format.
	// Experimental.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url *string `json:"url" yaml:"url"`
	// Compile options.
	// Experimental.
	CompileOptions *MavenCompileOptions `json:"compileOptions" yaml:"compileOptions"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps *[]*string `json:"deps" yaml:"deps"`
	// Final artifact output directory.
	// Experimental.
	Distdir *string `json:"distdir" yaml:"distdir"`
	// Include junit tests.
	// Experimental.
	Junit *bool `json:"junit" yaml:"junit"`
	// junit options.
	// Experimental.
	JunitOptions *JunitOptions `json:"junitOptions" yaml:"junitOptions"`
	// Packaging options.
	// Experimental.
	PackagingOptions *MavenPackagingOptions `json:"packagingOptions" yaml:"packagingOptions"`
	// Use projenrc in java.
	//
	// This will install `projen` as a java dependency and will add a `synth` task which
	// will compile & execute `main()` from `src/main/java/projenrc.java`.
	// Experimental.
	ProjenrcJava *bool `json:"projenrcJava" yaml:"projenrcJava"`
	// Options related to projenrc in java.
	// Experimental.
	ProjenrcJavaOptions *ProjenrcOptions `json:"projenrcJavaOptions" yaml:"projenrcJavaOptions"`
	// List of test dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addTestDependency()`.
	// Experimental.
	TestDeps *[]*string `json:"testDeps" yaml:"testDeps"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample *bool `json:"sample" yaml:"sample"`
	// The java package to use for the code sample.
	// Experimental.
	SampleJavaPackage *string `json:"sampleJavaPackage" yaml:"sampleJavaPackage"`
}

// Implements JUnit-based testing.
// Experimental.
type Junit interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Junit
type jsiiProxy_Junit struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Junit) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewJunit(project projen.Project, options *JunitOptions) Junit {
	_init_.Initialize()

	j := jsiiProxy_Junit{}

	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJunit_Override(j Junit, project projen.Project, options *JunitOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		j,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (j *jsiiProxy_Junit) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (j *jsiiProxy_Junit) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (j *jsiiProxy_Junit) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Junit`.
// Experimental.
type JunitOptions struct {
	// Java pom.
	// Experimental.
	Pom Pom `json:"pom" yaml:"pom"`
	// Java package for test sample.
	// Experimental.
	SampleJavaPackage *string `json:"sampleJavaPackage" yaml:"sampleJavaPackage"`
	// Junit version.
	// Experimental.
	Version *string `json:"version" yaml:"version"`
}

// Adds the maven-compiler plugin to a POM file and the `compile` task.
// Experimental.
type MavenCompile interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for MavenCompile
type jsiiProxy_MavenCompile struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenCompile) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenCompile(project projen.Project, pom Pom, options *MavenCompileOptions) MavenCompile {
	_init_.Initialize()

	j := jsiiProxy_MavenCompile{}

	_jsii_.Create(
		"projen.java.MavenCompile",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenCompile_Override(m MavenCompile, project projen.Project, pom Pom, options *MavenCompileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenCompile",
		[]interface{}{project, pom, options},
		m,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (m *jsiiProxy_MavenCompile) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (m *jsiiProxy_MavenCompile) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (m *jsiiProxy_MavenCompile) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `MavenCompile`.
// Experimental.
type MavenCompileOptions struct {
	// Source language version.
	// Experimental.
	Source *string `json:"source" yaml:"source"`
	// Target JVM version.
	// Experimental.
	Target *string `json:"target" yaml:"target"`
}

// Configures a maven project to produce a .jar archive with sources and javadocs.
// Experimental.
type MavenPackaging interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for MavenPackaging
type jsiiProxy_MavenPackaging struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenPackaging) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenPackaging(project projen.Project, pom Pom, options *MavenPackagingOptions) MavenPackaging {
	_init_.Initialize()

	j := jsiiProxy_MavenPackaging{}

	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenPackaging_Override(m MavenPackaging, project projen.Project, pom Pom, options *MavenPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		m,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (m *jsiiProxy_MavenPackaging) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (m *jsiiProxy_MavenPackaging) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (m *jsiiProxy_MavenPackaging) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `MavenPackage`.
// Experimental.
type MavenPackagingOptions struct {
	// Where to place the package output?
	// Experimental.
	Distdir *string `json:"distdir" yaml:"distdir"`
	// Include javadocs jar in package.
	// Experimental.
	Javadocs *bool `json:"javadocs" yaml:"javadocs"`
	// Exclude source files from docs.
	// Experimental.
	JavadocsExclude *[]*string `json:"javadocsExclude" yaml:"javadocsExclude"`
	// Include sources jar in package.
	// Experimental.
	Sources *bool `json:"sources" yaml:"sources"`
}

// Java code sample.
// Experimental.
type MavenSample interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for MavenSample
type jsiiProxy_MavenSample struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenSample) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenSample(project projen.Project, options *MavenSampleOptions) MavenSample {
	_init_.Initialize()

	j := jsiiProxy_MavenSample{}

	_jsii_.Create(
		"projen.java.MavenSample",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenSample_Override(m MavenSample, project projen.Project, options *MavenSampleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenSample",
		[]interface{}{project, options},
		m,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (m *jsiiProxy_MavenSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (m *jsiiProxy_MavenSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (m *jsiiProxy_MavenSample) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type MavenSampleOptions struct {
	// Project root java package.
	// Experimental.
	Package *string `json:"package" yaml:"package"`
}

// Plugin execution definition.
// Experimental.
type PluginExecution struct {
	// Which Maven goals this plugin should be associated with.
	// Experimental.
	Goals *[]*string `json:"goals" yaml:"goals"`
	// The ID.
	// Experimental.
	Id *string `json:"id" yaml:"id"`
}

// Options for Maven plugins.
// Experimental.
type PluginOptions struct {
	// Plugin key/value configuration.
	// Experimental.
	Configuration *map[string]interface{} `json:"configuration" yaml:"configuration"`
	// You could configure the dependencies for the plugin.
	//
	// Dependencies are in `<groupId>/<artifactId>@<semver>` format.
	// Experimental.
	Dependencies *[]*string `json:"dependencies" yaml:"dependencies"`
	// Plugin executions.
	// Experimental.
	Executions *[]*PluginExecution `json:"executions" yaml:"executions"`
}

// A Project Object Model or POM is the fundamental unit of work in Maven.
//
// It is
// an XML file that contains information about the project and configuration
// details used by Maven to build the project.
// Experimental.
type Pom interface {
	projen.Component
	ArtifactId() *string
	Description() *string
	FileName() *string
	GroupId() *string
	Name() *string
	Packaging() *string
	Project() projen.Project
	Url() *string
	Version() *string
	AddDependency(spec *string)
	AddPlugin(spec *string, options *PluginOptions) *projen.Dependency
	AddProperty(key *string, value *string)
	AddTestDependency(spec *string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Pom
type jsiiProxy_Pom struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Pom) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Packaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pom) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewPom(project projen.Project, options *PomOptions) Pom {
	_init_.Initialize()

	j := jsiiProxy_Pom{}

	_jsii_.Create(
		"projen.java.Pom",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPom_Override(p Pom, project projen.Project, options *PomOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Pom",
		[]interface{}{project, options},
		p,
	)
}

// Adds a runtime dependency.
// Experimental.
func (p *jsiiProxy_Pom) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

// Adds a build plugin to the pom.
//
// The plug in is also added as a BUILD dep to the project.
// Experimental.
func (p *jsiiProxy_Pom) AddPlugin(spec *string, options *PluginOptions) *projen.Dependency {
	var returns *projen.Dependency

	_jsii_.Invoke(
		p,
		"addPlugin",
		[]interface{}{spec, options},
		&returns,
	)

	return returns
}

// Adds a key/value property to the pom.
// Experimental.
func (p *jsiiProxy_Pom) AddProperty(key *string, value *string) {
	_jsii_.InvokeVoid(
		p,
		"addProperty",
		[]interface{}{key, value},
	)
}

// Adds a test dependency.
// Experimental.
func (p *jsiiProxy_Pom) AddTestDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addTestDependency",
		[]interface{}{spec},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Pom) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Pom) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (p *jsiiProxy_Pom) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Pom`.
// Experimental.
type PomOptions struct {
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
	ArtifactId *string `json:"artifactId" yaml:"artifactId"`
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
	GroupId *string `json:"groupId" yaml:"groupId"`
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
	Version *string `json:"version" yaml:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Project packaging format.
	// Experimental.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url *string `json:"url" yaml:"url"`
}

// Allows writing projenrc files in java.
//
// This will install `org.projen/projen` as a Maven dependency and will add a
// `synth` task which will compile & execute `main()` from
// `src/main/java/projenrc.java`.
// Experimental.
type Projenrc interface {
	projen.Component
	ClassName() *string
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Projenrc
type jsiiProxy_Projenrc struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Projenrc) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Projenrc) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrc(project projen.Project, pom Pom, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.java.Projenrc",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrc_Override(p Projenrc, project projen.Project, pom Pom, options *ProjenrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Projenrc",
		[]interface{}{project, pom, options},
		p,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (p *jsiiProxy_Projenrc) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Projenrc`.
// Experimental.
type ProjenrcOptions struct {
	// The name of the Java class which contains the `main()` method for projen.
	// Experimental.
	ClassName *string `json:"className" yaml:"className"`
	// The projen version to use.
	// Experimental.
	ProjenVersion *string `json:"projenVersion" yaml:"projenVersion"`
	// Defines projenrc under the test scope instead of the main scope, which is reserved to the app.
	//
	// This means that projenrc will be under
	// `src/test/java/projenrc.java` and projen will be defined as a test
	// dependency. This enforces that application code does not take a dependency
	// on projen code.
	//
	// If this is disabled, projenrc should be under
	// `src/main/java/projenrc.java`.
	// Experimental.
	TestScope *bool `json:"testScope" yaml:"testScope"`
}

