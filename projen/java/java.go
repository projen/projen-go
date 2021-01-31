package java

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen/deps"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/tasks"
	"github.com/projen/projen-go/projen/vscode"
	"github.com/projen/projen-go/projen/github"
)

// Class interface
type JavaProjectIface interface {
	GetComponents() []projen.ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []projen.FileBaseIface
	GetGitignore() projen.IgnoreFileIface
	GetLogger() projen.LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() projen.ProjectType
	GetRoot() projen.ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() projen.GitpodIface
	GetJsiiFqn() string
	GetParent() projen.ProjectIface
	GetVscode() vscode.VsCodeIface
	GetCompile() MavenCompileIface
	GetDistdir() string
	GetPackaging() MavenPackagingIface
	GetPom() PomIface
	GetJunit() JunitIface
	GetProjenrc() ProjenrcIface
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) projen.FileBaseIface
	TryFindJsonFile(filePath string) projen.JsonFileIface
	TryFindObjectFile(filePath string) projen.ObjectFileIface
	AddDependency(spec string)
	AddPlugin(spec string, options PluginOptionsIface) deps.DependencyIface
	AddTestDependency(spec string)
}

// Java project.
// Experimental.
// Struct proxy
type JavaProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []projen.ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []projen.FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore projen.IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger projen.LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root projen.ProjectIface `json:"root"`
	// Experimental.
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Experimental.
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Experimental.
	Gitpod projen.GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Compile component.
	// Experimental.
	Compile MavenCompileIface `json:"compile"`
	// Maven artifact output directory.
	// Experimental.
	Distdir string `json:"distdir"`
	// Packaging component.
	// Experimental.
	Packaging MavenPackagingIface `json:"packaging"`
	// API for managing `pom.xml`.
	// Experimental.
	Pom PomIface `json:"pom"`
	// JUnit component.
	// Experimental.
	Junit JunitIface `json:"junit"`
	// Projenrc component.
	// Experimental.
	Projenrc ProjenrcIface `json:"projenrc"`
}

func (j *JavaProject) GetComponents() []projen.ComponentIface {
	var returns []projen.ComponentIface
	_jsii_.Get(
		j,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ComponentIface)(nil)).Elem(): reflect.TypeOf((*projen.Component)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		j,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetFiles() []projen.FileBaseIface {
	var returns []projen.FileBaseIface
	_jsii_.Get(
		j,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetGitignore() projen.IgnoreFileIface {
	var returns projen.IgnoreFileIface
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*projen.IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetLogger() projen.LoggerIface {
	var returns projen.LoggerIface
	_jsii_.Get(
		j,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerIface)(nil)).Elem(): reflect.TypeOf((*projen.Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetName() string {
	var returns string
	_jsii_.Get(
		j,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProject) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetRoot() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		j,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		j,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		j,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetGitpod() projen.GitpodIface {
	var returns projen.GitpodIface
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.GitpodIface)(nil)).Elem(): reflect.TypeOf((*projen.Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		j,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProject) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		j,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		j,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetCompile() MavenCompileIface {
	var returns MavenCompileIface
	_jsii_.Get(
		j,
		"compile",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MavenCompileIface)(nil)).Elem(): reflect.TypeOf((*MavenCompile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetDistdir() string {
	var returns string
	_jsii_.Get(
		j,
		"distdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProject) GetPackaging() MavenPackagingIface {
	var returns MavenPackagingIface
	_jsii_.Get(
		j,
		"packaging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MavenPackagingIface)(nil)).Elem(): reflect.TypeOf((*MavenPackaging)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetPom() PomIface {
	var returns PomIface
	_jsii_.Get(
		j,
		"pom",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PomIface)(nil)).Elem(): reflect.TypeOf((*Pom)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetJunit() JunitIface {
	var returns JunitIface
	_jsii_.Get(
		j,
		"junit",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JunitIface)(nil)).Elem(): reflect.TypeOf((*Junit)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) GetProjenrc() ProjenrcIface {
	var returns ProjenrcIface
	_jsii_.Get(
		j,
		"projenrc",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjenrcIface)(nil)).Elem(): reflect.TypeOf((*Projenrc)(nil)).Elem(),
		},
	)
	return returns
}


func NewJavaProject(options JavaProjectOptionsIface) JavaProjectIface {
	_init_.Initialize()
	self := JavaProject{}
	_jsii_.Create(
		"projen.java.JavaProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func JavaProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.java.JavaProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JavaProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		j,
		"addTask",
		[]interface{}{name, props},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JavaProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JavaProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JavaProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JavaProject) TryFindFile(filePath string) projen.FileBaseIface {
	var returns projen.FileBaseIface
	_jsii_.Invoke(
		j,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.FileBaseIface)(nil)).Elem(): reflect.TypeOf((*projen.FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) TryFindJsonFile(filePath string) projen.JsonFileIface {
	var returns projen.JsonFileIface
	_jsii_.Invoke(
		j,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.JsonFileIface)(nil)).Elem(): reflect.TypeOf((*projen.JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) TryFindObjectFile(filePath string) projen.ObjectFileIface {
	var returns projen.ObjectFileIface
	_jsii_.Invoke(
		j,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*projen.ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) AddDependency(spec string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addDependency",
		[]interface{}{spec},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JavaProject) AddPlugin(spec string, options PluginOptionsIface) deps.DependencyIface {
	var returns deps.DependencyIface
	_jsii_.Invoke(
		j,
		"addPlugin",
		[]interface{}{spec, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependencyIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependency)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProject) AddTestDependency(spec string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addTestDependency",
		[]interface{}{spec},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// JavaProjectOptionsIface is the public interface for the custom type JavaProjectOptions
// Experimental.
type JavaProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() projen.LoggerOptionsIface
	GetOutdir() string
	GetParent() projen.ProjectIface
	GetProjectType() projen.ProjectType
	GetReadme() projen.SampleReadmePropsIface
	GetArtifactId() string
	GetGroupId() string
	GetVersion() string
	GetDescription() string
	GetPackaging() string
	GetUrl() string
	GetCompileOptions() MavenCompileOptionsIface
	GetDeps() []string
	GetDistdir() string
	GetJunit() bool
	GetJunitOptions() JunitOptionsIface
	GetPackagingOptions() MavenPackagingOptionsIface
	GetProjenrcJava() bool
	GetProjenrcJavaOptions() ProjenrcCommonOptionsIface
	GetSample() bool
	GetSampleJavaPackage() string
	GetTestDeps() []string
}

// Options for `JavaProject`.
// Experimental.
// Struct proxy
type JavaProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging projen.LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType projen.ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme projen.SampleReadmePropsIface `json:"readme"`
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
	ArtifactId string `json:"artifactId"`
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
	GroupId string `json:"groupId"`
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
	Version string `json:"version"`
	// Description of a project is always good.
	// 
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description string `json:"description"`
	// Project packaging format.
	// Experimental.
	Packaging string `json:"packaging"`
	// The URL, like the name, is not required.
	// 
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url string `json:"url"`
	// Compile options.
	// Experimental.
	CompileOptions MavenCompileOptionsIface `json:"compileOptions"`
	// List of runtime dependencies for this project.
	// 
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	// 
	// Additional dependencies can be added via `project.addDependency()`.
	// Experimental.
	Deps []string `json:"deps"`
	// Final artifact output directory.
	// Experimental.
	Distdir string `json:"distdir"`
	// Include junit tests.
	// Experimental.
	Junit bool `json:"junit"`
	// junit options.
	// Experimental.
	JunitOptions JunitOptionsIface `json:"junitOptions"`
	// Packaging options.
	// Experimental.
	PackagingOptions MavenPackagingOptionsIface `json:"packagingOptions"`
	// Use projenrc in java.
	// 
	// This will install `projen` as a java depedency and will add a `synth` task which
	// will compile & execute `main()` from `src/main/java/projenrc.java`.
	// Experimental.
	ProjenrcJava bool `json:"projenrcJava"`
	// Options related to projenrc in java.
	// Experimental.
	ProjenrcJavaOptions ProjenrcCommonOptionsIface `json:"projenrcJavaOptions"`
	// Include sample code and test if the relevant directories don't exist.
	// Experimental.
	Sample bool `json:"sample"`
	// The java package to use for the code sample.
	// Experimental.
	SampleJavaPackage string `json:"sampleJavaPackage"`
	// List of test dependencies for this project.
	// 
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	// 
	// Additional dependencies can be added via `project.addTestDependency()`.
	// Experimental.
	TestDeps []string `json:"testDeps"`
}

func (j *JavaProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		j,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		j,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		j,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetLogging() projen.LoggerOptionsIface {
	var returns projen.LoggerOptionsIface
	_jsii_.Get(
		j,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*projen.LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetParent() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		j,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectType)(nil)).Elem(): reflect.TypeOf((*projen.ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetReadme() projen.SampleReadmePropsIface {
	var returns projen.SampleReadmePropsIface
	_jsii_.Get(
		j,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*projen.SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetArtifactId() string {
	var returns string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetGroupId() string {
	var returns string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		j,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetPackaging() string {
	var returns string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"url",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetCompileOptions() MavenCompileOptionsIface {
	var returns MavenCompileOptionsIface
	_jsii_.Get(
		j,
		"compileOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MavenCompileOptionsIface)(nil)).Elem(): reflect.TypeOf((*MavenCompileOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		j,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetDistdir() string {
	var returns string
	_jsii_.Get(
		j,
		"distdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetJunit() bool {
	var returns bool
	_jsii_.Get(
		j,
		"junit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetJunitOptions() JunitOptionsIface {
	var returns JunitOptionsIface
	_jsii_.Get(
		j,
		"junitOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JunitOptionsIface)(nil)).Elem(): reflect.TypeOf((*JunitOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetPackagingOptions() MavenPackagingOptionsIface {
	var returns MavenPackagingOptionsIface
	_jsii_.Get(
		j,
		"packagingOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MavenPackagingOptionsIface)(nil)).Elem(): reflect.TypeOf((*MavenPackagingOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetProjenrcJava() bool {
	var returns bool
	_jsii_.Get(
		j,
		"projenrcJava",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetProjenrcJavaOptions() ProjenrcCommonOptionsIface {
	var returns ProjenrcCommonOptionsIface
	_jsii_.Get(
		j,
		"projenrcJavaOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjenrcCommonOptionsIface)(nil)).Elem(): reflect.TypeOf((*ProjenrcCommonOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JavaProjectOptions) GetSample() bool {
	var returns bool
	_jsii_.Get(
		j,
		"sample",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetSampleJavaPackage() string {
	var returns string
	_jsii_.Get(
		j,
		"sampleJavaPackage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JavaProjectOptions) GetTestDeps() []string {
	var returns []string
	_jsii_.Get(
		j,
		"testDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type JunitIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Implements JUnit-based testing.
// Experimental.
// Struct proxy
type Junit struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (j *Junit) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		j,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewJunit(project projen.ProjectIface, options JunitOptionsIface) JunitIface {
	_init_.Initialize()
	self := Junit{}
	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (j *Junit) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Junit) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Junit) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// JunitOptionsIface is the public interface for the custom type JunitOptions
// Experimental.
type JunitOptionsIface interface {
	GetPom() PomIface
	GetSampleJavaPackage() string
	GetVersion() string
}

// Options for `Junit`.
// Experimental.
// Struct proxy
type JunitOptions struct {
	// Java pom.
	// Experimental.
	Pom PomIface `json:"pom"`
	// Java package for test sample.
	// Experimental.
	SampleJavaPackage string `json:"sampleJavaPackage"`
	// Junit version.
	// Experimental.
	Version string `json:"version"`
}

func (j *JunitOptions) GetPom() PomIface {
	var returns PomIface
	_jsii_.Get(
		j,
		"pom",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PomIface)(nil)).Elem(): reflect.TypeOf((*Pom)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JunitOptions) GetSampleJavaPackage() string {
	var returns string
	_jsii_.Get(
		j,
		"sampleJavaPackage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JunitOptions) GetVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type MavenCompileIface interface {
	GetProject() projen.ProjectIface
	GetCompileTask() tasks.TaskIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Adds the maven-compiler plugin to a POM file and the `compile` task.
// Experimental.
// Struct proxy
type MavenCompile struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
}

func (m *MavenCompile) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		m,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (m *MavenCompile) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		m,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}


func NewMavenCompile(project projen.ProjectIface, pom PomIface, options MavenCompileOptionsIface) MavenCompileIface {
	_init_.Initialize()
	self := MavenCompile{}
	_jsii_.Create(
		"projen.java.MavenCompile",
		[]interface{}{project, pom, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (m *MavenCompile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *MavenCompile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *MavenCompile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// MavenCompileOptionsIface is the public interface for the custom type MavenCompileOptions
// Experimental.
type MavenCompileOptionsIface interface {
	GetSource() string
	GetTarget() string
}

// Options for `MavenCompile`.
// Experimental.
// Struct proxy
type MavenCompileOptions struct {
	// Source language version.
	// Experimental.
	Source string `json:"source"`
	// Target JVM version.
	// Experimental.
	Target string `json:"target"`
}

func (m *MavenCompileOptions) GetSource() string {
	var returns string
	_jsii_.Get(
		m,
		"source",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MavenCompileOptions) GetTarget() string {
	var returns string
	_jsii_.Get(
		m,
		"target",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type MavenPackagingIface interface {
	GetProject() projen.ProjectIface
	GetTask() tasks.TaskIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Configures a maven project to produce a .jar archive with sources and javadocs.
// Experimental.
// Struct proxy
type MavenPackaging struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// The "package" task.
	// Experimental.
	Task tasks.TaskIface `json:"task"`
}

func (m *MavenPackaging) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		m,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (m *MavenPackaging) GetTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		m,
		"task",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}


func NewMavenPackaging(project projen.ProjectIface, pom PomIface, options MavenPackagingOptionsIface) MavenPackagingIface {
	_init_.Initialize()
	self := MavenPackaging{}
	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (m *MavenPackaging) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *MavenPackaging) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *MavenPackaging) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// MavenPackagingOptionsIface is the public interface for the custom type MavenPackagingOptions
// Experimental.
type MavenPackagingOptionsIface interface {
	GetDistdir() string
	GetJavadocs() bool
	GetJavadocsExclude() []string
	GetSources() bool
}

// Options for `MavenPackage`.
// Experimental.
// Struct proxy
type MavenPackagingOptions struct {
	// Where to place the package output?
	// Experimental.
	Distdir string `json:"distdir"`
	// Include javadocs jar in package.
	// Experimental.
	Javadocs bool `json:"javadocs"`
	// Exclude source files from docs.
	// Experimental.
	JavadocsExclude []string `json:"javadocsExclude"`
	// Include sources jar in package.
	// Experimental.
	Sources bool `json:"sources"`
}

func (m *MavenPackagingOptions) GetDistdir() string {
	var returns string
	_jsii_.Get(
		m,
		"distdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MavenPackagingOptions) GetJavadocs() bool {
	var returns bool
	_jsii_.Get(
		m,
		"javadocs",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MavenPackagingOptions) GetJavadocsExclude() []string {
	var returns []string
	_jsii_.Get(
		m,
		"javadocsExclude",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (m *MavenPackagingOptions) GetSources() bool {
	var returns bool
	_jsii_.Get(
		m,
		"sources",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type MavenSampleIface interface {
	GetProject() projen.ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Java code sample.
// Experimental.
// Struct proxy
type MavenSample struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
}

func (m *MavenSample) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		m,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewMavenSample(project projen.ProjectIface, options MavenSampleOptionsIface) MavenSampleIface {
	_init_.Initialize()
	self := MavenSample{}
	_jsii_.Create(
		"projen.java.MavenSample",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (m *MavenSample) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *MavenSample) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (m *MavenSample) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		m,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// MavenSampleOptionsIface is the public interface for the custom type MavenSampleOptions
// Experimental.
type MavenSampleOptionsIface interface {
	GetPackage() string
}

// Experimental.
// Struct proxy
type MavenSampleOptions struct {
	// Project root java package.
	// Experimental.
	Package string `json:"package"`
}

func (m *MavenSampleOptions) GetPackage() string {
	var returns string
	_jsii_.Get(
		m,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// PluginExecutionIface is the public interface for the custom type PluginExecution
// Experimental.
type PluginExecutionIface interface {
	GetGoals() []string
	GetId() string
}

// Plugin execution definition.
// Experimental.
// Struct proxy
type PluginExecution struct {
	// Which Maven goals this plugin should be associated with.
	// Experimental.
	Goals []string `json:"goals"`
	// The ID.
	// Experimental.
	Id string `json:"id"`
}

func (p *PluginExecution) GetGoals() []string {
	var returns []string
	_jsii_.Get(
		p,
		"goals",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PluginExecution) GetId() string {
	var returns string
	_jsii_.Get(
		p,
		"id",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// PluginOptionsIface is the public interface for the custom type PluginOptions
// Experimental.
type PluginOptionsIface interface {
	GetConfiguration() map[string]interface{}
	GetDependencies() []string
	GetExecutions() []PluginExecutionIface
}

// Options for Maven plugins.
// Experimental.
// Struct proxy
type PluginOptions struct {
	// Plugin key/value configuration.
	// Experimental.
	Configuration map[string]interface{} `json:"configuration"`
	// You could configure the dependencies for the plugin.
	// 
	// Dependencies are in `<groupId>/<artifactId>@<semver>` format.
	// Experimental.
	Dependencies []string `json:"dependencies"`
	// Plugin executions.
	// Experimental.
	Executions []PluginExecutionIface `json:"executions"`
}

func (p *PluginOptions) GetConfiguration() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		p,
		"configuration",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (p *PluginOptions) GetDependencies() []string {
	var returns []string
	_jsii_.Get(
		p,
		"dependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PluginOptions) GetExecutions() []PluginExecutionIface {
	var returns []PluginExecutionIface
	_jsii_.Get(
		p,
		"executions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PluginExecutionIface)(nil)).Elem(): reflect.TypeOf((*PluginExecution)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type PomIface interface {
	GetProject() projen.ProjectIface
	GetArtifactId() string
	GetFileName() string
	GetGroupId() string
	GetPackaging() string
	GetVersion() string
	GetDescription() string
	GetName() string
	GetUrl() string
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddDependency(spec string)
	AddPlugin(spec string, options PluginOptionsIface) deps.DependencyIface
	AddProperty(key string, value string)
	AddTestDependency(spec string)
}

// A Project Object Model or POM is the fundamental unit of work in Maven.
// 
// It is
// an XML file that contains information about the project and configuration
// details used by Maven to build the project.
// Experimental.
// Struct proxy
type Pom struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// Maven artifact ID.
	// Experimental.
	ArtifactId string `json:"artifactId"`
	// The name of the pom file.
	// Experimental.
	FileName string `json:"fileName"`
	// Maven group ID.
	// Experimental.
	GroupId string `json:"groupId"`
	// Maven packaging format.
	// Experimental.
	Packaging string `json:"packaging"`
	// Project version.
	// Experimental.
	Version string `json:"version"`
	// Project description.
	// Experimental.
	Description string `json:"description"`
	// Project display name.
	// Experimental.
	Name string `json:"name"`
	// Project URL.
	// Experimental.
	Url string `json:"url"`
}

func (p *Pom) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		p,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pom) GetArtifactId() string {
	var returns string
	_jsii_.Get(
		p,
		"artifactId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetFileName() string {
	var returns string
	_jsii_.Get(
		p,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetGroupId() string {
	var returns string
	_jsii_.Get(
		p,
		"groupId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetPackaging() string {
	var returns string
	_jsii_.Get(
		p,
		"packaging",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetVersion() string {
	var returns string
	_jsii_.Get(
		p,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetDescription() string {
	var returns string
	_jsii_.Get(
		p,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetName() string {
	var returns string
	_jsii_.Get(
		p,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pom) GetUrl() string {
	var returns string
	_jsii_.Get(
		p,
		"url",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewPom(project projen.ProjectIface, options PomOptionsIface) PomIface {
	_init_.Initialize()
	self := Pom{}
	_jsii_.Create(
		"projen.java.Pom",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (p *Pom) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pom) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pom) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pom) AddDependency(spec string) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addDependency",
		[]interface{}{spec},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pom) AddPlugin(spec string, options PluginOptionsIface) deps.DependencyIface {
	var returns deps.DependencyIface
	_jsii_.Invoke(
		p,
		"addPlugin",
		[]interface{}{spec, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependencyIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependency)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pom) AddProperty(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addProperty",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pom) AddTestDependency(spec string) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addTestDependency",
		[]interface{}{spec},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// PomOptionsIface is the public interface for the custom type PomOptions
// Experimental.
type PomOptionsIface interface {
	GetArtifactId() string
	GetGroupId() string
	GetVersion() string
	GetDescription() string
	GetPackaging() string
	GetUrl() string
}

// Options for `Pom`.
// Experimental.
// Struct proxy
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
	ArtifactId string `json:"artifactId"`
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
	GroupId string `json:"groupId"`
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
	Version string `json:"version"`
	// Description of a project is always good.
	// 
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Experimental.
	Description string `json:"description"`
	// Project packaging format.
	// Experimental.
	Packaging string `json:"packaging"`
	// The URL, like the name, is not required.
	// 
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Experimental.
	Url string `json:"url"`
}

func (p *PomOptions) GetArtifactId() string {
	var returns string
	_jsii_.Get(
		p,
		"artifactId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PomOptions) GetGroupId() string {
	var returns string
	_jsii_.Get(
		p,
		"groupId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PomOptions) GetVersion() string {
	var returns string
	_jsii_.Get(
		p,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PomOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		p,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PomOptions) GetPackaging() string {
	var returns string
	_jsii_.Get(
		p,
		"packaging",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PomOptions) GetUrl() string {
	var returns string
	_jsii_.Get(
		p,
		"url",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ProjenrcIface interface {
	GetProject() projen.ProjectIface
	GetClassName() string
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Allows writing projenrc files in java.
// 
// This will install `org.projen/projen` as a Maven dependency and will add a
// `synth` task which will compile & execute `main()` from
// `src/main/java/projenrc.java`.
// Experimental.
// Struct proxy
type Projenrc struct {
	// Experimental.
	Project projen.ProjectIface `json:"project"`
	// The name of the java class that includes the projen entrypoint.
	// Experimental.
	ClassName string `json:"className"`
}

func (p *Projenrc) GetProject() projen.ProjectIface {
	var returns projen.ProjectIface
	_jsii_.Get(
		p,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*projen.ProjectIface)(nil)).Elem(): reflect.TypeOf((*projen.Project)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Projenrc) GetClassName() string {
	var returns string
	_jsii_.Get(
		p,
		"className",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewProjenrc(project projen.ProjectIface, pom PomIface, options ProjenrcOptionsIface) ProjenrcIface {
	_init_.Initialize()
	self := Projenrc{}
	_jsii_.Create(
		"projen.java.Projenrc",
		[]interface{}{project, pom, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (p *Projenrc) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Projenrc) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Projenrc) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ProjenrcCommonOptionsIface is the public interface for the custom type ProjenrcCommonOptions
// Experimental.
type ProjenrcCommonOptionsIface interface {
	GetClassName() string
	GetProjenVersion() string
	GetTestScope() bool
}

// Options for `Projenrc`.
// Experimental.
// Struct proxy
type ProjenrcCommonOptions struct {
	// The name of the Java class which contains the `main()` method for projen.
	// Experimental.
	ClassName string `json:"className"`
	// The projen version to use.
	// Experimental.
	ProjenVersion string `json:"projenVersion"`
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
	TestScope bool `json:"testScope"`
}

func (p *ProjenrcCommonOptions) GetClassName() string {
	var returns string
	_jsii_.Get(
		p,
		"className",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjenrcCommonOptions) GetProjenVersion() string {
	var returns string
	_jsii_.Get(
		p,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjenrcCommonOptions) GetTestScope() bool {
	var returns bool
	_jsii_.Get(
		p,
		"testScope",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ProjenrcOptionsIface is the public interface for the custom type ProjenrcOptions
// Experimental.
type ProjenrcOptionsIface interface {
	GetClassName() string
	GetProjenVersion() string
	GetTestScope() bool
	GetInitializationOptions() map[string]interface{}
}

// Experimental.
// Struct proxy
type ProjenrcOptions struct {
	// The name of the Java class which contains the `main()` method for projen.
	// Experimental.
	ClassName string `json:"className"`
	// The projen version to use.
	// Experimental.
	ProjenVersion string `json:"projenVersion"`
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
	TestScope bool `json:"testScope"`
	// Project initialization options.
	// Experimental.
	InitializationOptions map[string]interface{} `json:"initializationOptions"`
}

func (p *ProjenrcOptions) GetClassName() string {
	var returns string
	_jsii_.Get(
		p,
		"className",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjenrcOptions) GetProjenVersion() string {
	var returns string
	_jsii_.Get(
		p,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjenrcOptions) GetTestScope() bool {
	var returns bool
	_jsii_.Get(
		p,
		"testScope",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjenrcOptions) GetInitializationOptions() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		p,
		"initializationOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}


