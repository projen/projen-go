// A new generation of project generators
package projen

import (
	_jsii_ "github.com/aws-cdk/jsii/jsii-experimental"
	_init_ "github.com/projen/projen-go/projen/jsii"
	"reflect"
	"github.com/projen/projen-go/projen/deps"
	"github.com/projen/projen-go/projen/tasks"
	"github.com/projen/projen-go/projen/vscode"
	"github.com/projen/projen-go/projen/github"
)

// Automatic bump modes.
// Experimental.
type AutoRelease string

const (
	AutoReleaseEveryCommit AutoRelease = "EVERY_COMMIT"
	AutoReleaseDaily AutoRelease = "DAILY"
)

// Class interface
type AwsCdkConstructLibraryIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	GetTwineRegistryUrl() string
	GetVersion() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
	AddCdkDependencies(deps string)
	AddCdkTestDependencies(deps string)
}

// AWS CDK construct library project.
// 
// A multi-language (jsii) construct library which vends constructs designed to
// use within the AWS CDK with a friendly workflow and automatic publishing to
// the construct catalog.
// 
// ```ts
// const project = new ConstructLibraryAws({
//    name: 'cdk-watchful',
//    description: 'Watching your CDK apps since 2019',
//    jsiiVersion: Semver.caret('1.7.0'),
//    authorName: 'Elad Ben-Israel',
//    authorEmail: 'elad.benisrael@gmail.com',
//    repository: 'https://github.com/eladb/cdk-watchful.git',
//    keywords: [
//      "cloudwatch",
//      "monitoring"
//    ],
// 
//    catalog: {
//      twitter: 'emeshbi'
//    },
// 
//    // creates PRs for projen upgrades
//    projenUpgradeSecret: 'PROJEN_GITHUB_TOKEN',
// 
//    cdkVersion: '1.54.0',
//    cdkDependencies: [
//      "@aws-cdk/aws-apigateway",
//      "@aws-cdk/aws-cloudwatch",
//      "@aws-cdk/aws-cloudwatch-actions",
//      "@aws-cdk/aws-dynamodb",
//      "@aws-cdk/aws-ecs",
//      "@aws-cdk/aws-ecs-patterns",
//      "@aws-cdk/aws-elasticloadbalancingv2",
//      "@aws-cdk/aws-events",
//      "@aws-cdk/aws-events-targets",
//      "@aws-cdk/aws-lambda",
//      "@aws-cdk/aws-rds",
//      "@aws-cdk/aws-sns",
//      "@aws-cdk/aws-sns-subscriptions",
//      "@aws-cdk/aws-sqs",
//      "@aws-cdk/core"
//    ],
//    devDependencies: {
//      "aws-sdk": Semver.caret("2.708.0")
//    },
// 
//    // jsii publishing
// 
//    java: {
//      javaPackage: 'com.github.eladb.watchful',
//      mavenGroupId: 'com.github.eladb',
//      mavenArtifactId: 'cdk-watchful'
//    },
//    python: {
//      distName: 'cdk-watchful',
//      module: 'cdk_watchful'
//    }
// });
// 
// project.synth();
// ```
// Experimental.
// Struct proxy
type AwsCdkConstructLibrary struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
	// Experimental.
	TwineRegistryUrl string `json:"twineRegistryUrl"`
	// The target CDK version for this library.
	// Experimental.
	Version string `json:"version"`
}

func (a *AwsCdkConstructLibrary) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		a,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		a,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		a,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		a,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		a,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetName() string {
	var returns string
	_jsii_.Get(
		a,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetOutdir() string {
	var returns string
	_jsii_.Get(
		a,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		a,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		a,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		a,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		a,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		a,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		a,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		a,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		a,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		a,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		a,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		a,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		a,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		a,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		a,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		a,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		a,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		a,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		a,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		a,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		a,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		a,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		a,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		a,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		a,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		a,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		a,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		a,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetLibdir() string {
	var returns string
	_jsii_.Get(
		a,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		a,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetTestdir() string {
	var returns string
	_jsii_.Get(
		a,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		a,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		a,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		a,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetTwineRegistryUrl() string {
	var returns string
	_jsii_.Get(
		a,
		"twineRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) GetVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewAwsCdkConstructLibrary(options AwsCdkConstructLibraryOptionsIface) AwsCdkConstructLibraryIface {
	_init_.Initialize()
	self := AwsCdkConstructLibrary{}
	_jsii_.Create(
		"projen.AwsCdkConstructLibrary",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func AwsCdkConstructLibrary_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.AwsCdkConstructLibrary",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		a,
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

func (a *AwsCdkConstructLibrary) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		a,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		a,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		a,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		a,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		a,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibrary) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddCdkDependencies(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addCdkDependencies",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkConstructLibrary) AddCdkTestDependencies(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addCdkTestDependencies",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// AwsCdkConstructLibraryOptionsIface is the public interface for the custom type AwsCdkConstructLibraryOptions
// Experimental.
type AwsCdkConstructLibraryOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetAuthor() string
	GetAuthorAddress() string
	GetRepositoryUrl() string
	GetCompat() bool
	GetCompatIgnore() string
	GetDocgen() bool
	GetDotnet() JsiiDotNetTargetIface
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetPublishToGo() JsiiGoTargetIface
	GetPublishToMaven() JsiiJavaTargetIface
	GetPublishToNuget() JsiiDotNetTargetIface
	GetPublishToPypi() JsiiPythonTargetIface
	GetPython() JsiiPythonTargetIface
	GetRootdir() string
	GetCatalog() CatalogIface
	GetCdkVersion() string
	GetCdkAssert() bool
	GetCdkDependencies() []string
	GetCdkTestDependencies() []string
	GetCdkVersionPinning() bool
}

// Options for the construct-lib-aws project.
// Experimental.
// Struct proxy
type AwsCdkConstructLibraryOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// The name of the library author.
	// Experimental.
	Author string `json:"author"`
	// Email or URL of the library author.
	// Experimental.
	AuthorAddress string `json:"authorAddress"`
	// Git repository URL.
	// Experimental.
	RepositoryUrl string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	// 
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Experimental.
	Compat bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Experimental.
	CompatIgnore string `json:"compatIgnore"`
	// Automatically generate API.md from jsii.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Deprecated: use `publishToNuget`
	Dotnet JsiiDotNetTargetIface `json:"dotnet"`
	// Install eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Publish Go bindings to a git repository.
	// Experimental.
	PublishToGo JsiiGoTargetIface `json:"publishToGo"`
	// Publish to maven.
	// Experimental.
	PublishToMaven JsiiJavaTargetIface `json:"publishToMaven"`
	// Publish to NuGet.
	// Experimental.
	PublishToNuget JsiiDotNetTargetIface `json:"publishToNuget"`
	// Publish to pypi.
	// Experimental.
	PublishToPypi JsiiPythonTargetIface `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python JsiiPythonTargetIface `json:"python"`
	// Experimental.
	Rootdir string `json:"rootdir"`
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
	Catalog CatalogIface `json:"catalog"`
	// Minimum target version this library is tested against.
	// Experimental.
	CdkVersion string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Experimental.
	CdkAssert bool `json:"cdkAssert"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Experimental.
	CdkDependencies []string `json:"cdkDependencies"`
	// AWS CDK modules required for testing.
	// Experimental.
	CdkTestDependencies []string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	// 
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning bool `json:"cdkVersionPinning"`
}

func (a *AwsCdkConstructLibraryOptions) GetName() string {
	var returns string
	_jsii_.Get(
		a,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		a,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		a,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		a,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		a,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		a,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		a,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		a,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		a,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		a,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		a,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		a,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		a,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		a,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		a,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		a,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		a,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		a,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		a,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		a,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		a,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		a,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		a,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		a,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		a,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		a,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		a,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		a,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		a,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		a,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		a,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		a,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		a,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		a,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		a,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		a,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		a,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		a,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		a,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		a,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		a,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		a,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		a,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		a,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		a,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		a,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		a,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		a,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		a,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		a,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		a,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		a,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		a,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		a,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		a,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		a,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		a,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		a,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		a,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		a,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		a,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		a,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		a,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		a,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		a,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		a,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		a,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAuthor() string {
	var returns string
	_jsii_.Get(
		a,
		"author",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetAuthorAddress() string {
	var returns string
	_jsii_.Get(
		a,
		"authorAddress",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetRepositoryUrl() string {
	var returns string
	_jsii_.Get(
		a,
		"repositoryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCompat() bool {
	var returns bool
	_jsii_.Get(
		a,
		"compat",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCompatIgnore() string {
	var returns string
	_jsii_.Get(
		a,
		"compatIgnore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		a,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetDotnet() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		a,
		"dotnet",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		a,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		a,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPublishToGo() JsiiGoTargetIface {
	var returns JsiiGoTargetIface
	_jsii_.Get(
		a,
		"publishToGo",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiGoTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiGoTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPublishToMaven() JsiiJavaTargetIface {
	var returns JsiiJavaTargetIface
	_jsii_.Get(
		a,
		"publishToMaven",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiJavaTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiJavaTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPublishToNuget() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		a,
		"publishToNuget",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPublishToPypi() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		a,
		"publishToPypi",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetPython() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		a,
		"python",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetRootdir() string {
	var returns string
	_jsii_.Get(
		a,
		"rootdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCatalog() CatalogIface {
	var returns CatalogIface
	_jsii_.Get(
		a,
		"catalog",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*CatalogIface)(nil)).Elem(): reflect.TypeOf((*Catalog)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCdkVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"cdkVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCdkAssert() bool {
	var returns bool
	_jsii_.Get(
		a,
		"cdkAssert",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCdkDependencies() []string {
	var returns []string
	_jsii_.Get(
		a,
		"cdkDependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCdkTestDependencies() []string {
	var returns []string
	_jsii_.Get(
		a,
		"cdkTestDependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkConstructLibraryOptions) GetCdkVersionPinning() bool {
	var returns bool
	_jsii_.Get(
		a,
		"cdkVersionPinning",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type AwsCdkTypeScriptAppIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	GetAppEntrypoint() string
	GetCdkConfig() interface{}
	GetCdkVersion() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
	AddCdkDependency(modules string)
}

// AWS CDK app in TypeScript.
// Experimental.
// Struct proxy
type AwsCdkTypeScriptApp struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
	// The CDK app entrypoint.
	// Experimental.
	AppEntrypoint string `json:"appEntrypoint"`
	// Contents of `cdk.json`.
	// Experimental.
	CdkConfig interface{} `json:"cdkConfig"`
	// The CDK version this app is using.
	// Experimental.
	CdkVersion string `json:"cdkVersion"`
}

func (a *AwsCdkTypeScriptApp) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		a,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		a,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		a,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		a,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		a,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetName() string {
	var returns string
	_jsii_.Get(
		a,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetOutdir() string {
	var returns string
	_jsii_.Get(
		a,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		a,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		a,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		a,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		a,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		a,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		a,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		a,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		a,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		a,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		a,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		a,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		a,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		a,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		a,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		a,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		a,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		a,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		a,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		a,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		a,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		a,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		a,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		a,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		a,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		a,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		a,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		a,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetLibdir() string {
	var returns string
	_jsii_.Get(
		a,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		a,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetTestdir() string {
	var returns string
	_jsii_.Get(
		a,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		a,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		a,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		a,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		a,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetAppEntrypoint() string {
	var returns string
	_jsii_.Get(
		a,
		"appEntrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetCdkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		a,
		"cdkConfig",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) GetCdkVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"cdkVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewAwsCdkTypeScriptApp(options AwsCdkTypeScriptAppOptionsIface) AwsCdkTypeScriptAppIface {
	_init_.Initialize()
	self := AwsCdkTypeScriptApp{}
	_jsii_.Create(
		"projen.AwsCdkTypeScriptApp",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func AwsCdkTypeScriptApp_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.AwsCdkTypeScriptApp",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		a,
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

func (a *AwsCdkTypeScriptApp) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		a,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		a,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		a,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		a,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		a,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptApp) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (a *AwsCdkTypeScriptApp) AddCdkDependency(modules string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addCdkDependency",
		[]interface{}{modules},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// AwsCdkTypeScriptAppOptionsIface is the public interface for the custom type AwsCdkTypeScriptAppOptions
// Experimental.
type AwsCdkTypeScriptAppOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetCompileBeforeTest() bool
	GetDisableTsconfig() bool
	GetDocgen() bool
	GetDocsDirectory() string
	GetEntrypointTypes() string
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetLibdir() string
	GetPackage() bool
	GetSampleCode() bool
	GetSrcdir() string
	GetTestdir() string
	GetTsconfig() TypescriptConfigOptionsIface
	GetTypescriptVersion() string
	GetCdkVersion() string
	GetAppEntrypoint() string
	GetCdkDependencies() []string
	GetCdkVersionPinning() bool
	GetContext() map[string]string
	GetRequireApproval() CdkApprovalLevel
}

// Experimental.
// Struct proxy
type AwsCdkTypeScriptAppOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Compile the code before running tests.
	// Experimental.
	CompileBeforeTest bool `json:"compileBeforeTest"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes string `json:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir string `json:"libdir"`
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Experimental.
	Package bool `json:"package"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	// 
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir string `json:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig TypescriptConfigOptionsIface `json:"tsconfig"`
	// TypeScript version to use.
	// Experimental.
	TypescriptVersion string `json:"typescriptVersion"`
	// AWS CDK version to use.
	// Experimental.
	CdkVersion string `json:"cdkVersion"`
	// The CDK app's entrypoint (relative to the source directory, which is "src" by default).
	// Experimental.
	AppEntrypoint string `json:"appEntrypoint"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") this app uses.
	// Experimental.
	CdkDependencies []string `json:"cdkDependencies"`
	// Use pinned version instead of caret version for CDK.
	// 
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning bool `json:"cdkVersionPinning"`
	// Additional context to include in `cdk.json`.
	// Experimental.
	Context map[string]string `json:"context"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Experimental.
	RequireApproval CdkApprovalLevel `json:"requireApproval"`
}

func (a *AwsCdkTypeScriptAppOptions) GetName() string {
	var returns string
	_jsii_.Get(
		a,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		a,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		a,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		a,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		a,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		a,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		a,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		a,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		a,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		a,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		a,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		a,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		a,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		a,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		a,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		a,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		a,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		a,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		a,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		a,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		a,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		a,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		a,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		a,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		a,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		a,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		a,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		a,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		a,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		a,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		a,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		a,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		a,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		a,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		a,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		a,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		a,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		a,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		a,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		a,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		a,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		a,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		a,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		a,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		a,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		a,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		a,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		a,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		a,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		a,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		a,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		a,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		a,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		a,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		a,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		a,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		a,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		a,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		a,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		a,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		a,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		a,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		a,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		a,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		a,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		a,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		a,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		a,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		a,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCompileBeforeTest() bool {
	var returns bool
	_jsii_.Get(
		a,
		"compileBeforeTest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDisableTsconfig() bool {
	var returns bool
	_jsii_.Get(
		a,
		"disableTsconfig",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		a,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		a,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetEntrypointTypes() string {
	var returns string
	_jsii_.Get(
		a,
		"entrypointTypes",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		a,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		a,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetLibdir() string {
	var returns string
	_jsii_.Get(
		a,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetPackage() bool {
	var returns bool
	_jsii_.Get(
		a,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		a,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		a,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetTestdir() string {
	var returns string
	_jsii_.Get(
		a,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetTsconfig() TypescriptConfigOptionsIface {
	var returns TypescriptConfigOptionsIface
	_jsii_.Get(
		a,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetTypescriptVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"typescriptVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCdkVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"cdkVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetAppEntrypoint() string {
	var returns string
	_jsii_.Get(
		a,
		"appEntrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCdkDependencies() []string {
	var returns []string
	_jsii_.Get(
		a,
		"cdkDependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetCdkVersionPinning() bool {
	var returns bool
	_jsii_.Get(
		a,
		"cdkVersionPinning",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetContext() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		a,
		"context",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AwsCdkTypeScriptAppOptions) GetRequireApproval() CdkApprovalLevel {
	var returns CdkApprovalLevel
	_jsii_.Get(
		a,
		"requireApproval",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*CdkApprovalLevel)(nil)).Elem(): reflect.TypeOf((*CdkApprovalLevel)(nil)).Elem(),
		},
	)
	return returns
}


// CatalogIface is the public interface for the custom type Catalog
// Experimental.
type CatalogIface interface {
	GetAnnounce() bool
	GetTwitter() string
}

// Experimental.
// Struct proxy
type Catalog struct {
	// Should we announce new versions?
	// Experimental.
	Announce bool `json:"announce"`
	// Twitter account to @mention in announcement tweet.
	// Experimental.
	Twitter string `json:"twitter"`
}

func (c *Catalog) GetAnnounce() bool {
	var returns bool
	_jsii_.Get(
		c,
		"announce",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *Catalog) GetTwitter() string {
	var returns string
	_jsii_.Get(
		c,
		"twitter",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Experimental.
type CdkApprovalLevel string

const (
	CdkApprovalLevelNever CdkApprovalLevel = "NEVER"
	CdkApprovalLevelAnyChange CdkApprovalLevel = "ANY_CHANGE"
	CdkApprovalLevelBroadening CdkApprovalLevel = "BROADENING"
)

// Class interface
type ComponentIface interface {
	GetProject() ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Represents a project component.
// Experimental.
// Struct proxy
type Component struct {
	// Experimental.
	Project ProjectIface `json:"project"`
}

func (c *Component) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewComponent(project ProjectIface) ComponentIface {
	_init_.Initialize()
	self := Component{}
	_jsii_.Create(
		"projen.Component",
		[]interface{}{project},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (c *Component) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *Component) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *Component) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type ConstructLibraryIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	GetTwineRegistryUrl() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// A multi-language library for CDK constructs.
// Experimental.
// Struct proxy
type ConstructLibrary struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
	// Experimental.
	TwineRegistryUrl string `json:"twineRegistryUrl"`
}

func (c *ConstructLibrary) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		c,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		c,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		c,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		c,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		c,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetOutdir() string {
	var returns string
	_jsii_.Get(
		c,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		c,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		c,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		c,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		c,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		c,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		c,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		c,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		c,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		c,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		c,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		c,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		c,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		c,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		c,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		c,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		c,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		c,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		c,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		c,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		c,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		c,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		c,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		c,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		c,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetLibdir() string {
	var returns string
	_jsii_.Get(
		c,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		c,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetTestdir() string {
	var returns string
	_jsii_.Get(
		c,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		c,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		c,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		c,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) GetTwineRegistryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"twineRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewConstructLibrary(options ConstructLibraryOptionsIface) ConstructLibraryIface {
	_init_.Initialize()
	self := ConstructLibrary{}
	_jsii_.Create(
		"projen.ConstructLibrary",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ConstructLibrary_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.ConstructLibrary",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		c,
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

func (c *ConstructLibrary) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		c,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		c,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		c,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibrary) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		c,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibrary) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		c,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibrary) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type ConstructLibraryAwsIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	GetTwineRegistryUrl() string
	GetVersion() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
	AddCdkDependencies(deps string)
	AddCdkTestDependencies(deps string)
}

// Deprecated: use `AwsCdkConstructLibrary`
// Struct proxy
type ConstructLibraryAws struct {
	// Returns all the components within this project.
	// Deprecated: use `AwsCdkConstructLibrary`
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Deprecated: use `AwsCdkConstructLibrary`
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Deprecated: use `AwsCdkConstructLibrary`
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Deprecated: use `AwsCdkConstructLibrary`
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Deprecated: use `AwsCdkConstructLibrary`
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Deprecated: use `AwsCdkConstructLibrary`
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Deprecated: use `AwsCdkConstructLibrary`
	Outdir string `json:"outdir"`
	// Deprecated: use `AwsCdkConstructLibrary`
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Deprecated: use `AwsCdkConstructLibrary`
	Root ProjectIface `json:"root"`
	// Deprecated: use `AwsCdkConstructLibrary`
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Deprecated: use `AwsCdkConstructLibrary`
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Deprecated: use `AwsCdkConstructLibrary`
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Deprecated: use `AwsCdkConstructLibrary`
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Deprecated: use `AwsCdkConstructLibrary`
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Deprecated: use `AwsCdkConstructLibrary`
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Deprecated: use `AwsCdkConstructLibrary`
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Deprecated: use `AwsCdkConstructLibrary`
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Deprecated: use `AwsCdkConstructLibrary`
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Deprecated: use `AwsCdkConstructLibrary`
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Deprecated: use `AwsCdkConstructLibrary`
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Deprecated: use `AwsCdkConstructLibrary`
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Deprecated: use `AwsCdkConstructLibrary`
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Deprecated: use `AwsCdkConstructLibrary`
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Deprecated: use `AwsCdkConstructLibrary`
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Deprecated: use `AwsCdkConstructLibrary`
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Deprecated: use `AwsCdkConstructLibrary`
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Deprecated: use `AwsCdkConstructLibrary`
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Deprecated: use `AwsCdkConstructLibrary`
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Deprecated: use `AwsCdkConstructLibrary`
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Deprecated: use `AwsCdkConstructLibrary`
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Deprecated: use `AwsCdkConstructLibrary`
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Deprecated: use `AwsCdkConstructLibrary`
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Deprecated: use `AwsCdkConstructLibrary`
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Deprecated: use `AwsCdkConstructLibrary`
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Deprecated: use `AwsCdkConstructLibrary`
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Deprecated: use `AwsCdkConstructLibrary`
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Deprecated: use `AwsCdkConstructLibrary`
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Deprecated: use `AwsCdkConstructLibrary`
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Deprecated: use `AwsCdkConstructLibrary`
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Deprecated: use `AwsCdkConstructLibrary`
	Docgen bool `json:"docgen"`
	// Deprecated: use `AwsCdkConstructLibrary`
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Deprecated: use `AwsCdkConstructLibrary`
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Deprecated: use `AwsCdkConstructLibrary`
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
	// Deprecated: use `AwsCdkConstructLibrary`
	TwineRegistryUrl string `json:"twineRegistryUrl"`
	// The target CDK version for this library.
	// Deprecated: use `AwsCdkConstructLibrary`
	Version string `json:"version"`
}

func (c *ConstructLibraryAws) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		c,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		c,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		c,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		c,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		c,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetOutdir() string {
	var returns string
	_jsii_.Get(
		c,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		c,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		c,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		c,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		c,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		c,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		c,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		c,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		c,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		c,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		c,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		c,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		c,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		c,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		c,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		c,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		c,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		c,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		c,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		c,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		c,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		c,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		c,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		c,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		c,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetLibdir() string {
	var returns string
	_jsii_.Get(
		c,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		c,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetTestdir() string {
	var returns string
	_jsii_.Get(
		c,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		c,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		c,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		c,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) GetTwineRegistryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"twineRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) GetVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewConstructLibraryAws(options AwsCdkConstructLibraryOptionsIface) ConstructLibraryAwsIface {
	_init_.Initialize()
	self := ConstructLibraryAws{}
	_jsii_.Create(
		"projen.ConstructLibraryAws",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ConstructLibraryAws_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.ConstructLibraryAws",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		c,
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

func (c *ConstructLibraryAws) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		c,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		c,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		c,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAws) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		c,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		c,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAws) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddCdkDependencies(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addCdkDependencies",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryAws) AddCdkTestDependencies(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addCdkTestDependencies",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ConstructLibraryAwsOptionsIface is the public interface for the custom type ConstructLibraryAwsOptions
// Deprecated: use `AwsCdkConstructLibraryOptions`
type ConstructLibraryAwsOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetAuthor() string
	GetAuthorAddress() string
	GetRepositoryUrl() string
	GetCompat() bool
	GetCompatIgnore() string
	GetDocgen() bool
	GetDotnet() JsiiDotNetTargetIface
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetPublishToGo() JsiiGoTargetIface
	GetPublishToMaven() JsiiJavaTargetIface
	GetPublishToNuget() JsiiDotNetTargetIface
	GetPublishToPypi() JsiiPythonTargetIface
	GetPython() JsiiPythonTargetIface
	GetRootdir() string
	GetCatalog() CatalogIface
	GetCdkVersion() string
	GetCdkAssert() bool
	GetCdkDependencies() []string
	GetCdkTestDependencies() []string
	GetCdkVersionPinning() bool
}

// Deprecated: use `AwsCdkConstructLibraryOptions`
// Struct proxy
type ConstructLibraryAwsOptions struct {
	// This is the name of your project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Logging LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	License string `json:"license"`
	// Indicates if a license should be added.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Jest bool `json:"jest"`
	// Jest options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// The name of the library author.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Author string `json:"author"`
	// Email or URL of the library author.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	AuthorAddress string `json:"authorAddress"`
	// Git repository URL.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	RepositoryUrl string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	// 
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Compat bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CompatIgnore string `json:"compatIgnore"`
	// Automatically generate API.md from jsii.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Docgen bool `json:"docgen"`
	// Deprecated: use `publishToNuget`
	Dotnet JsiiDotNetTargetIface `json:"dotnet"`
	// Install eslint.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Publish Go bindings to a git repository.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToGo JsiiGoTargetIface `json:"publishToGo"`
	// Publish to maven.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToMaven JsiiJavaTargetIface `json:"publishToMaven"`
	// Publish to NuGet.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToNuget JsiiDotNetTargetIface `json:"publishToNuget"`
	// Publish to pypi.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	PublishToPypi JsiiPythonTargetIface `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python JsiiPythonTargetIface `json:"python"`
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	Rootdir string `json:"rootdir"`
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
	Catalog CatalogIface `json:"catalog"`
	// Minimum target version this library is tested against.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkVersion string `json:"cdkVersion"`
	// Install the @aws-cdk/assert library?
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkAssert bool `json:"cdkAssert"`
	// Which AWS CDK modules (those that start with "@aws-cdk/") does this library require when consumed?
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkDependencies []string `json:"cdkDependencies"`
	// AWS CDK modules required for testing.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkTestDependencies []string `json:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	// 
	// You can use this to prevent yarn to mix versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Deprecated: use `AwsCdkConstructLibraryOptions`
	CdkVersionPinning bool `json:"cdkVersionPinning"`
}

func (c *ConstructLibraryAwsOptions) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		c,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		c,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		c,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		c,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		c,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		c,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		c,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		c,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		c,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		c,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		c,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		c,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		c,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		c,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		c,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		c,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		c,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		c,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		c,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		c,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		c,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		c,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		c,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		c,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		c,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		c,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		c,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		c,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		c,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		c,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		c,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		c,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		c,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		c,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		c,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		c,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		c,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		c,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		c,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		c,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		c,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		c,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		c,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		c,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		c,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		c,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		c,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		c,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		c,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		c,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		c,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		c,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		c,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		c,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		c,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		c,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		c,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAuthor() string {
	var returns string
	_jsii_.Get(
		c,
		"author",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetAuthorAddress() string {
	var returns string
	_jsii_.Get(
		c,
		"authorAddress",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetRepositoryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"repositoryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCompat() bool {
	var returns bool
	_jsii_.Get(
		c,
		"compat",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCompatIgnore() string {
	var returns string
	_jsii_.Get(
		c,
		"compatIgnore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		c,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetDotnet() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		c,
		"dotnet",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		c,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		c,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPublishToGo() JsiiGoTargetIface {
	var returns JsiiGoTargetIface
	_jsii_.Get(
		c,
		"publishToGo",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiGoTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiGoTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPublishToMaven() JsiiJavaTargetIface {
	var returns JsiiJavaTargetIface
	_jsii_.Get(
		c,
		"publishToMaven",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiJavaTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiJavaTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPublishToNuget() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		c,
		"publishToNuget",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPublishToPypi() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		c,
		"publishToPypi",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetPython() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		c,
		"python",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetRootdir() string {
	var returns string
	_jsii_.Get(
		c,
		"rootdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCatalog() CatalogIface {
	var returns CatalogIface
	_jsii_.Get(
		c,
		"catalog",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*CatalogIface)(nil)).Elem(): reflect.TypeOf((*Catalog)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCdkVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"cdkVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCdkAssert() bool {
	var returns bool
	_jsii_.Get(
		c,
		"cdkAssert",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCdkDependencies() []string {
	var returns []string
	_jsii_.Get(
		c,
		"cdkDependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCdkTestDependencies() []string {
	var returns []string
	_jsii_.Get(
		c,
		"cdkTestDependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryAwsOptions) GetCdkVersionPinning() bool {
	var returns bool
	_jsii_.Get(
		c,
		"cdkVersionPinning",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ConstructLibraryCdk8SIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	GetTwineRegistryUrl() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// CDK8s construct library project.
// 
// A multi-language (jsii) construct library which vends constructs designed to
// use within the CDK for Kubernetes (CDK8s), with a friendly workflow and
// automatic publishing to the construct catalog.
// Experimental.
// Struct proxy
type ConstructLibraryCdk8S struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
	// Experimental.
	TwineRegistryUrl string `json:"twineRegistryUrl"`
}

func (c *ConstructLibraryCdk8S) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		c,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		c,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		c,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		c,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		c,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetOutdir() string {
	var returns string
	_jsii_.Get(
		c,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		c,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		c,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		c,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		c,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		c,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		c,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		c,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		c,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		c,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		c,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		c,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		c,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		c,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		c,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		c,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		c,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		c,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		c,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		c,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		c,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		c,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		c,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		c,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		c,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetLibdir() string {
	var returns string
	_jsii_.Get(
		c,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		c,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetTestdir() string {
	var returns string
	_jsii_.Get(
		c,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		c,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		c,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		c,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		c,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) GetTwineRegistryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"twineRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewConstructLibraryCdk8S(options ConstructLibraryCdk8SOptionsIface) ConstructLibraryCdk8SIface {
	_init_.Initialize()
	self := ConstructLibraryCdk8S{}
	_jsii_.Create(
		"projen.ConstructLibraryCdk8s",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ConstructLibraryCdk8S_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.ConstructLibraryCdk8s",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		c,
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

func (c *ConstructLibraryCdk8S) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		c,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		c,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		c,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		c,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConstructLibraryCdk8S) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		c,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8S) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ConstructLibraryCdk8SOptionsIface is the public interface for the custom type ConstructLibraryCdk8SOptions
// Experimental.
type ConstructLibraryCdk8SOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetAuthor() string
	GetAuthorAddress() string
	GetRepositoryUrl() string
	GetCompat() bool
	GetCompatIgnore() string
	GetDocgen() bool
	GetDotnet() JsiiDotNetTargetIface
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetPublishToGo() JsiiGoTargetIface
	GetPublishToMaven() JsiiJavaTargetIface
	GetPublishToNuget() JsiiDotNetTargetIface
	GetPublishToPypi() JsiiPythonTargetIface
	GetPython() JsiiPythonTargetIface
	GetRootdir() string
	GetCatalog() CatalogIface
	GetCdk8SVersion() string
}

// Experimental.
// Struct proxy
type ConstructLibraryCdk8SOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// The name of the library author.
	// Experimental.
	Author string `json:"author"`
	// Email or URL of the library author.
	// Experimental.
	AuthorAddress string `json:"authorAddress"`
	// Git repository URL.
	// Experimental.
	RepositoryUrl string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	// 
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Experimental.
	Compat bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Experimental.
	CompatIgnore string `json:"compatIgnore"`
	// Automatically generate API.md from jsii.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Deprecated: use `publishToNuget`
	Dotnet JsiiDotNetTargetIface `json:"dotnet"`
	// Install eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Publish Go bindings to a git repository.
	// Experimental.
	PublishToGo JsiiGoTargetIface `json:"publishToGo"`
	// Publish to maven.
	// Experimental.
	PublishToMaven JsiiJavaTargetIface `json:"publishToMaven"`
	// Publish to NuGet.
	// Experimental.
	PublishToNuget JsiiDotNetTargetIface `json:"publishToNuget"`
	// Publish to pypi.
	// Experimental.
	PublishToPypi JsiiPythonTargetIface `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python JsiiPythonTargetIface `json:"python"`
	// Experimental.
	Rootdir string `json:"rootdir"`
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
	Catalog CatalogIface `json:"catalog"`
	// Minimum target version this library is tested against.
	// Experimental.
	Cdk8SVersion string `json:"cdk8sVersion"`
}

func (c *ConstructLibraryCdk8SOptions) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		c,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		c,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		c,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		c,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		c,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		c,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		c,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		c,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		c,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		c,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		c,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		c,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		c,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		c,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		c,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		c,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		c,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		c,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		c,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		c,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		c,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		c,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		c,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		c,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		c,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		c,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		c,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		c,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		c,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		c,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		c,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		c,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		c,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		c,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		c,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		c,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		c,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		c,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		c,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		c,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		c,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		c,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		c,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		c,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		c,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		c,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		c,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		c,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		c,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		c,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		c,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		c,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		c,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		c,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		c,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		c,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		c,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAuthor() string {
	var returns string
	_jsii_.Get(
		c,
		"author",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetAuthorAddress() string {
	var returns string
	_jsii_.Get(
		c,
		"authorAddress",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetRepositoryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"repositoryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCompat() bool {
	var returns bool
	_jsii_.Get(
		c,
		"compat",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCompatIgnore() string {
	var returns string
	_jsii_.Get(
		c,
		"compatIgnore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		c,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetDotnet() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		c,
		"dotnet",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		c,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		c,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPublishToGo() JsiiGoTargetIface {
	var returns JsiiGoTargetIface
	_jsii_.Get(
		c,
		"publishToGo",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiGoTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiGoTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPublishToMaven() JsiiJavaTargetIface {
	var returns JsiiJavaTargetIface
	_jsii_.Get(
		c,
		"publishToMaven",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiJavaTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiJavaTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPublishToNuget() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		c,
		"publishToNuget",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPublishToPypi() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		c,
		"publishToPypi",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetPython() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		c,
		"python",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetRootdir() string {
	var returns string
	_jsii_.Get(
		c,
		"rootdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCatalog() CatalogIface {
	var returns CatalogIface
	_jsii_.Get(
		c,
		"catalog",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*CatalogIface)(nil)).Elem(): reflect.TypeOf((*Catalog)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryCdk8SOptions) GetCdk8SVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"cdk8sVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ConstructLibraryOptionsIface is the public interface for the custom type ConstructLibraryOptions
// Experimental.
type ConstructLibraryOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetAuthor() string
	GetAuthorAddress() string
	GetRepositoryUrl() string
	GetCompat() bool
	GetCompatIgnore() string
	GetDocgen() bool
	GetDotnet() JsiiDotNetTargetIface
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetPublishToGo() JsiiGoTargetIface
	GetPublishToMaven() JsiiJavaTargetIface
	GetPublishToNuget() JsiiDotNetTargetIface
	GetPublishToPypi() JsiiPythonTargetIface
	GetPython() JsiiPythonTargetIface
	GetRootdir() string
	GetCatalog() CatalogIface
}

// Experimental.
// Struct proxy
type ConstructLibraryOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// The name of the library author.
	// Experimental.
	Author string `json:"author"`
	// Email or URL of the library author.
	// Experimental.
	AuthorAddress string `json:"authorAddress"`
	// Git repository URL.
	// Experimental.
	RepositoryUrl string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	// 
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Experimental.
	Compat bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Experimental.
	CompatIgnore string `json:"compatIgnore"`
	// Automatically generate API.md from jsii.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Deprecated: use `publishToNuget`
	Dotnet JsiiDotNetTargetIface `json:"dotnet"`
	// Install eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Publish Go bindings to a git repository.
	// Experimental.
	PublishToGo JsiiGoTargetIface `json:"publishToGo"`
	// Publish to maven.
	// Experimental.
	PublishToMaven JsiiJavaTargetIface `json:"publishToMaven"`
	// Publish to NuGet.
	// Experimental.
	PublishToNuget JsiiDotNetTargetIface `json:"publishToNuget"`
	// Publish to pypi.
	// Experimental.
	PublishToPypi JsiiPythonTargetIface `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python JsiiPythonTargetIface `json:"python"`
	// Experimental.
	Rootdir string `json:"rootdir"`
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
	Catalog CatalogIface `json:"catalog"`
}

func (c *ConstructLibraryOptions) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		c,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		c,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		c,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		c,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		c,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		c,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		c,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		c,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		c,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		c,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		c,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		c,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		c,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		c,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		c,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		c,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		c,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		c,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		c,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		c,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		c,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		c,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		c,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		c,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		c,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		c,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		c,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		c,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		c,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		c,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		c,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		c,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		c,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		c,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		c,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		c,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		c,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		c,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		c,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		c,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		c,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		c,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		c,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		c,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		c,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		c,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		c,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		c,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		c,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		c,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		c,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		c,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		c,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		c,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		c,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		c,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		c,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		c,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		c,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		c,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		c,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		c,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAuthor() string {
	var returns string
	_jsii_.Get(
		c,
		"author",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetAuthorAddress() string {
	var returns string
	_jsii_.Get(
		c,
		"authorAddress",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetRepositoryUrl() string {
	var returns string
	_jsii_.Get(
		c,
		"repositoryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCompat() bool {
	var returns bool
	_jsii_.Get(
		c,
		"compat",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCompatIgnore() string {
	var returns string
	_jsii_.Get(
		c,
		"compatIgnore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		c,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetDotnet() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		c,
		"dotnet",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		c,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		c,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPublishToGo() JsiiGoTargetIface {
	var returns JsiiGoTargetIface
	_jsii_.Get(
		c,
		"publishToGo",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiGoTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiGoTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPublishToMaven() JsiiJavaTargetIface {
	var returns JsiiJavaTargetIface
	_jsii_.Get(
		c,
		"publishToMaven",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiJavaTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiJavaTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPublishToNuget() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		c,
		"publishToNuget",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPublishToPypi() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		c,
		"publishToPypi",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetPython() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		c,
		"python",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetRootdir() string {
	var returns string
	_jsii_.Get(
		c,
		"rootdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConstructLibraryOptions) GetCatalog() CatalogIface {
	var returns CatalogIface
	_jsii_.Get(
		c,
		"catalog",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*CatalogIface)(nil)).Elem(): reflect.TypeOf((*Catalog)(nil)).Elem(),
		},
	)
	return returns
}


// CoverageThresholdIface is the public interface for the custom type CoverageThreshold
// Experimental.
type CoverageThresholdIface interface {
	GetBranches() float64
	GetFunctions() float64
	GetLines() float64
	GetStatements() float64
}

// Experimental.
// Struct proxy
type CoverageThreshold struct {
	// Experimental.
	Branches float64 `json:"branches"`
	// Experimental.
	Functions float64 `json:"functions"`
	// Experimental.
	Lines float64 `json:"lines"`
	// Experimental.
	Statements float64 `json:"statements"`
}

func (c *CoverageThreshold) GetBranches() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"branches",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *CoverageThreshold) GetFunctions() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"functions",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *CoverageThreshold) GetLines() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"lines",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *CoverageThreshold) GetStatements() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"statements",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type DevEnvironmentDockerImageIface interface {
	GetDockerFile() string
	GetImage() string
}

// Options for specifying the Docker image of the container.
// Experimental.
// Struct proxy
type DevEnvironmentDockerImage struct {
	// The relative path of a Dockerfile that defines the container contents.
	// Experimental.
	DockerFile string `json:"dockerFile"`
	// A publicly available Docker image.
	// Experimental.
	Image string `json:"image"`
}

func (d *DevEnvironmentDockerImage) GetDockerFile() string {
	var returns string
	_jsii_.Get(
		d,
		"dockerFile",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DevEnvironmentDockerImage) GetImage() string {
	var returns string
	_jsii_.Get(
		d,
		"image",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func DevEnvironmentDockerImage_FromFile(dockerFile string) DevEnvironmentDockerImageIface {
	_init_.Initialize()
	var returns DevEnvironmentDockerImageIface
	_jsii_.InvokeStatic(
		"projen.DevEnvironmentDockerImage",
		"fromFile",
		[]interface{}{dockerFile},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DevEnvironmentDockerImageIface)(nil)).Elem(): reflect.TypeOf((*DevEnvironmentDockerImage)(nil)).Elem(),
		},
	)
	return returns
}

func DevEnvironmentDockerImage_FromImage(image string) DevEnvironmentDockerImageIface {
	_init_.Initialize()
	var returns DevEnvironmentDockerImageIface
	_jsii_.InvokeStatic(
		"projen.DevEnvironmentDockerImage",
		"fromImage",
		[]interface{}{image},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DevEnvironmentDockerImageIface)(nil)).Elem(): reflect.TypeOf((*DevEnvironmentDockerImage)(nil)).Elem(),
		},
	)
	return returns
}

// DevEnvironmentOptionsIface is the public interface for the custom type DevEnvironmentOptions
// Experimental.
type DevEnvironmentOptionsIface interface {
	GetDockerImage() DevEnvironmentDockerImageIface
	GetPorts() []string
	GetTasks() []tasks.TaskIface
	GetVscodeExtensions() []string
}

// Base options for configuring a container-based development environemnt.
// Experimental.
// Struct proxy
type DevEnvironmentOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage DevEnvironmentDockerImageIface `json:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports []string `json:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks []tasks.TaskIface `json:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions []string `json:"vscodeExtensions"`
}

func (d *DevEnvironmentOptions) GetDockerImage() DevEnvironmentDockerImageIface {
	var returns DevEnvironmentDockerImageIface
	_jsii_.Get(
		d,
		"dockerImage",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DevEnvironmentDockerImageIface)(nil)).Elem(): reflect.TypeOf((*DevEnvironmentDockerImage)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevEnvironmentOptions) GetPorts() []string {
	var returns []string
	_jsii_.Get(
		d,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevEnvironmentOptions) GetTasks() []tasks.TaskIface {
	var returns []tasks.TaskIface
	_jsii_.Get(
		d,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DevEnvironmentOptions) GetVscodeExtensions() []string {
	var returns []string
	_jsii_.Get(
		d,
		"vscodeExtensions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type DockerComposeIface interface {
	GetProject() ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddService(serviceName string, description DockerComposeServiceDescriptionIface) DockerComposeServiceIface
}

// Create a docker-compose YAML file.
// Experimental.
// Struct proxy
type DockerCompose struct {
	// Experimental.
	Project ProjectIface `json:"project"`
}

func (d *DockerCompose) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		d,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewDockerCompose(project ProjectIface, props DockerComposePropsIface) DockerComposeIface {
	_init_.Initialize()
	self := DockerCompose{}
	_jsii_.Create(
		"projen.DockerCompose",
		[]interface{}{project, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func DockerCompose_BindVolume(sourcePath string, targetPath string) IDockerComposeVolumeBindingIface {
	_init_.Initialize()
	var returns IDockerComposeVolumeBindingIface
	_jsii_.InvokeStatic(
		"projen.DockerCompose",
		"bindVolume",
		[]interface{}{sourcePath, targetPath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeVolumeBindingIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeVolumeBinding)(nil)).Elem(),
		},
	)
	return returns
}

func DockerCompose_NamedVolume(volumeName string, targetPath string, options DockerComposeVolumeConfigIface) IDockerComposeVolumeBindingIface {
	_init_.Initialize()
	var returns IDockerComposeVolumeBindingIface
	_jsii_.InvokeStatic(
		"projen.DockerCompose",
		"namedVolume",
		[]interface{}{volumeName, targetPath, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeVolumeBindingIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeVolumeBinding)(nil)).Elem(),
		},
	)
	return returns
}

func DockerCompose_PortMapping(publishedPort float64, targetPort float64, options DockerComposePortMappingOptionsIface) DockerComposeServicePortIface {
	_init_.Initialize()
	var returns DockerComposeServicePortIface
	_jsii_.InvokeStatic(
		"projen.DockerCompose",
		"portMapping",
		[]interface{}{publishedPort, targetPort, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeServicePortIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeServicePort)(nil)).Elem(),
		},
	)
	return returns
}

func DockerCompose_ServiceName(serviceName string) IDockerComposeServiceNameIface {
	_init_.Initialize()
	var returns IDockerComposeServiceNameIface
	_jsii_.InvokeStatic(
		"projen.DockerCompose",
		"serviceName",
		[]interface{}{serviceName},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeServiceNameIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeServiceName)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerCompose) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DockerCompose) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DockerCompose) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DockerCompose) AddService(serviceName string, description DockerComposeServiceDescriptionIface) DockerComposeServiceIface {
	var returns DockerComposeServiceIface
	_jsii_.Invoke(
		d,
		"addService",
		[]interface{}{serviceName, description},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeServiceIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeService)(nil)).Elem(),
		},
	)
	return returns
}

// DockerComposeBuildIface is the public interface for the custom type DockerComposeBuild
// Experimental.
type DockerComposeBuildIface interface {
	GetContext() string
	GetArgs() map[string]string
	GetDockerfile() string
}

// Build arguments for creating a docker image.
// Experimental.
// Struct proxy
type DockerComposeBuild struct {
	// Docker build context directory.
	// Experimental.
	Context string `json:"context"`
	// Build args.
	// Experimental.
	Args map[string]string `json:"args"`
	// A dockerfile to build from.
	// Experimental.
	Dockerfile string `json:"dockerfile"`
}

func (d *DockerComposeBuild) GetContext() string {
	var returns string
	_jsii_.Get(
		d,
		"context",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeBuild) GetArgs() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		d,
		"args",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeBuild) GetDockerfile() string {
	var returns string
	_jsii_.Get(
		d,
		"dockerfile",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// DockerComposePortMappingOptionsIface is the public interface for the custom type DockerComposePortMappingOptions
// Experimental.
type DockerComposePortMappingOptionsIface interface {
	GetProtocol() DockerComposeProtocol
}

// Options for port mappings.
// Experimental.
// Struct proxy
type DockerComposePortMappingOptions struct {
	// Port mapping protocol.
	// Experimental.
	Protocol DockerComposeProtocol `json:"protocol"`
}

func (d *DockerComposePortMappingOptions) GetProtocol() DockerComposeProtocol {
	var returns DockerComposeProtocol
	_jsii_.Get(
		d,
		"protocol",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeProtocol)(nil)).Elem(): reflect.TypeOf((*DockerComposeProtocol)(nil)).Elem(),
		},
	)
	return returns
}


// DockerComposePropsIface is the public interface for the custom type DockerComposeProps
// Experimental.
type DockerComposePropsIface interface {
	GetNameSuffix() string
	GetSchemaVersion() string
	GetServices() map[string]DockerComposeServiceDescriptionIface
}

// Props for DockerCompose.
// Experimental.
// Struct proxy
type DockerComposeProps struct {
	// A name to add to the docker-compose.yml filename.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	NameSuffix string `json:"nameSuffix"`
	// Docker Compose schema version do be used.
	// Experimental.
	SchemaVersion string `json:"schemaVersion"`
	// Service descriptions.
	// Experimental.
	Services map[string]DockerComposeServiceDescriptionIface `json:"services"`
}

func (d *DockerComposeProps) GetNameSuffix() string {
	var returns string
	_jsii_.Get(
		d,
		"nameSuffix",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeProps) GetSchemaVersion() string {
	var returns string
	_jsii_.Get(
		d,
		"schemaVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeProps) GetServices() map[string]DockerComposeServiceDescriptionIface {
	var returns map[string]DockerComposeServiceDescriptionIface
	_jsii_.Get(
		d,
		"services",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeServiceDescriptionIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeServiceDescription)(nil)).Elem(),
		},
	)
	return returns
}


// Network protocol for port mapping.
// Experimental.
type DockerComposeProtocol string

const (
	DockerComposeProtocolTcp DockerComposeProtocol = "TCP"
	DockerComposeProtocolUdp DockerComposeProtocol = "UDP"
)

// Class interface
type DockerComposeServiceIface interface {
	IDockerComposeServiceNameIface
	GetDependsOn() []IDockerComposeServiceNameIface
	GetEnvironment() map[string]string
	GetPorts() []DockerComposeServicePortIface
	GetServiceName() string
	GetVolumes() []IDockerComposeVolumeBindingIface
	GetCommand() []string
	GetImage() string
	GetImageBuild() DockerComposeBuildIface
	AddDependsOn(serviceName IDockerComposeServiceNameIface)
	AddEnvironment(name string, value string)
	AddPort(publishedPort float64, targetPort float64, options DockerComposePortMappingOptionsIface)
	AddVolume(volume IDockerComposeVolumeBindingIface)
}

// A docker-compose service.
// Experimental.
// Struct proxy
type DockerComposeService struct {
	// Other services that this service depends on.
	// Experimental.
	DependsOn []IDockerComposeServiceNameIface `json:"dependsOn"`
	// Environment variables.
	// Experimental.
	Environment map[string]string `json:"environment"`
	// Published ports.
	// Experimental.
	Ports []DockerComposeServicePortIface `json:"ports"`
	// Name of the service.
	// Experimental.
	ServiceName string `json:"serviceName"`
	// Volumes mounted in the container.
	// Experimental.
	Volumes []IDockerComposeVolumeBindingIface `json:"volumes"`
	// Command to run in the container.
	// Experimental.
	Command []string `json:"command"`
	// Docker image.
	// Experimental.
	Image string `json:"image"`
	// Docker image build instructions.
	// Experimental.
	ImageBuild DockerComposeBuildIface `json:"imageBuild"`
}

func (d *DockerComposeService) GetDependsOn() []IDockerComposeServiceNameIface {
	var returns []IDockerComposeServiceNameIface
	_jsii_.Get(
		d,
		"dependsOn",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeServiceNameIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeServiceName)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeService) GetEnvironment() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		d,
		"environment",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeService) GetPorts() []DockerComposeServicePortIface {
	var returns []DockerComposeServicePortIface
	_jsii_.Get(
		d,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeServicePortIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeServicePort)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeService) GetServiceName() string {
	var returns string
	_jsii_.Get(
		d,
		"serviceName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeService) GetVolumes() []IDockerComposeVolumeBindingIface {
	var returns []IDockerComposeVolumeBindingIface
	_jsii_.Get(
		d,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeVolumeBindingIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeVolumeBinding)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeService) GetCommand() []string {
	var returns []string
	_jsii_.Get(
		d,
		"command",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeService) GetImage() string {
	var returns string
	_jsii_.Get(
		d,
		"image",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeService) GetImageBuild() DockerComposeBuildIface {
	var returns DockerComposeBuildIface
	_jsii_.Get(
		d,
		"imageBuild",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeBuildIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeBuild)(nil)).Elem(),
		},
	)
	return returns
}


func NewDockerComposeService(serviceName string, serviceDescription DockerComposeServiceDescriptionIface) DockerComposeServiceIface {
	_init_.Initialize()
	self := DockerComposeService{}
	_jsii_.Create(
		"projen.DockerComposeService",
		[]interface{}{serviceName, serviceDescription},
		[]_jsii_.FQN{"projen.IDockerComposeServiceName"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (d *DockerComposeService) AddDependsOn(serviceName IDockerComposeServiceNameIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addDependsOn",
		[]interface{}{serviceName},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DockerComposeService) AddEnvironment(name string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addEnvironment",
		[]interface{}{name, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DockerComposeService) AddPort(publishedPort float64, targetPort float64, options DockerComposePortMappingOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addPort",
		[]interface{}{publishedPort, targetPort, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *DockerComposeService) AddVolume(volume IDockerComposeVolumeBindingIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// DockerComposeServiceDescriptionIface is the public interface for the custom type DockerComposeServiceDescription
// Experimental.
type DockerComposeServiceDescriptionIface interface {
	GetCommand() []string
	GetDependsOn() []IDockerComposeServiceNameIface
	GetEnvironment() map[string]string
	GetImage() string
	GetImageBuild() DockerComposeBuildIface
	GetPorts() []DockerComposeServicePortIface
	GetVolumes() []IDockerComposeVolumeBindingIface
}

// Description of a docker-compose.yml service.
// Experimental.
// Struct proxy
type DockerComposeServiceDescription struct {
	// Provide a command to the docker container.
	// Experimental.
	Command []string `json:"command"`
	// Names of other services this service depends on.
	// Experimental.
	DependsOn []IDockerComposeServiceNameIface `json:"dependsOn"`
	// Add environment variables.
	// Experimental.
	Environment map[string]string `json:"environment"`
	// Use a docker image.
	// 
	// Note: You must specify either `build` or `image` key.
	// See: imageBuild
	//
	// Experimental.
	Image string `json:"image"`
	// Build a docker image.
	// 
	// Note: You must specify either `imageBuild` or `image` key.
	// See: image
	//
	// Experimental.
	ImageBuild DockerComposeBuildIface `json:"imageBuild"`
	// Map some ports.
	// Experimental.
	Ports []DockerComposeServicePortIface `json:"ports"`
	// Mount some volumes into the service.
	// 
	// Use one of the following to create volumes:
	// See: DockerCompose.namedVolume() to create & mount a named volume
	//
	// Experimental.
	Volumes []IDockerComposeVolumeBindingIface `json:"volumes"`
}

func (d *DockerComposeServiceDescription) GetCommand() []string {
	var returns []string
	_jsii_.Get(
		d,
		"command",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeServiceDescription) GetDependsOn() []IDockerComposeServiceNameIface {
	var returns []IDockerComposeServiceNameIface
	_jsii_.Get(
		d,
		"dependsOn",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeServiceNameIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeServiceName)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeServiceDescription) GetEnvironment() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		d,
		"environment",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeServiceDescription) GetImage() string {
	var returns string
	_jsii_.Get(
		d,
		"image",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeServiceDescription) GetImageBuild() DockerComposeBuildIface {
	var returns DockerComposeBuildIface
	_jsii_.Get(
		d,
		"imageBuild",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeBuildIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeBuild)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeServiceDescription) GetPorts() []DockerComposeServicePortIface {
	var returns []DockerComposeServicePortIface
	_jsii_.Get(
		d,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeServicePortIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeServicePort)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeServiceDescription) GetVolumes() []IDockerComposeVolumeBindingIface {
	var returns []IDockerComposeVolumeBindingIface
	_jsii_.Get(
		d,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IDockerComposeVolumeBindingIface)(nil)).Elem(): reflect.TypeOf((*IDockerComposeVolumeBinding)(nil)).Elem(),
		},
	)
	return returns
}


// DockerComposeServicePortIface is the public interface for the custom type DockerComposeServicePort
// Experimental.
type DockerComposeServicePortIface interface {
	GetMode() string
	GetProtocol() DockerComposeProtocol
	GetPublished() float64
	GetTarget() float64
}

// A service port mapping.
// Experimental.
// Struct proxy
type DockerComposeServicePort struct {
	// Port mapping mode.
	// Experimental.
	Mode string `json:"mode"`
	// Network protocol.
	// Experimental.
	Protocol DockerComposeProtocol `json:"protocol"`
	// Published port number.
	// Experimental.
	Published float64 `json:"published"`
	// Target port number.
	// Experimental.
	Target float64 `json:"target"`
}

func (d *DockerComposeServicePort) GetMode() string {
	var returns string
	_jsii_.Get(
		d,
		"mode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeServicePort) GetProtocol() DockerComposeProtocol {
	var returns DockerComposeProtocol
	_jsii_.Get(
		d,
		"protocol",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeProtocol)(nil)).Elem(): reflect.TypeOf((*DockerComposeProtocol)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeServicePort) GetPublished() float64 {
	var returns float64
	_jsii_.Get(
		d,
		"published",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeServicePort) GetTarget() float64 {
	var returns float64
	_jsii_.Get(
		d,
		"target",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// DockerComposeVolumeConfigIface is the public interface for the custom type DockerComposeVolumeConfig
// Experimental.
type DockerComposeVolumeConfigIface interface {
	GetDriver() string
	GetDriverOpts() map[string]string
	GetExternal() bool
	GetName() string
}

// Volume configuration.
// Experimental.
// Struct proxy
type DockerComposeVolumeConfig struct {
	// Driver to use for the volume.
	// Experimental.
	Driver string `json:"driver"`
	// Options to provide to the driver.
	// Experimental.
	DriverOpts map[string]string `json:"driverOpts"`
	// Set to true to indicate that the volume is externally created.
	// Experimental.
	External bool `json:"external"`
	// Name of the volume for when the volume name isn't going to work in YAML.
	// Experimental.
	Name string `json:"name"`
}

func (d *DockerComposeVolumeConfig) GetDriver() string {
	var returns string
	_jsii_.Get(
		d,
		"driver",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeVolumeConfig) GetDriverOpts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		d,
		"driverOpts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DockerComposeVolumeConfig) GetExternal() bool {
	var returns bool
	_jsii_.Get(
		d,
		"external",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeVolumeConfig) GetName() string {
	var returns string
	_jsii_.Get(
		d,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// DockerComposeVolumeMountIface is the public interface for the custom type DockerComposeVolumeMount
// Experimental.
type DockerComposeVolumeMountIface interface {
	GetSource() string
	GetTarget() string
	GetType() string
}

// Service volume mounting information.
// Experimental.
// Struct proxy
type DockerComposeVolumeMount struct {
	// Volume source.
	// Experimental.
	Source string `json:"source"`
	// Volume target.
	// Experimental.
	Target string `json:"target"`
	// Type of volume.
	// Experimental.
	Type string `json:"type"`
}

func (d *DockerComposeVolumeMount) GetSource() string {
	var returns string
	_jsii_.Get(
		d,
		"source",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeVolumeMount) GetTarget() string {
	var returns string
	_jsii_.Get(
		d,
		"target",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DockerComposeVolumeMount) GetType() string {
	var returns string
	_jsii_.Get(
		d,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type EslintIface interface {
	GetProject() ProjectIface
	GetConfig() interface{}
	GetIgnorePatterns() []string
	GetOverrides() []EslintOverrideIface
	GetRules() map[string][]interface{}
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddIgnorePattern(pattern string)
	AddOverride(override EslintOverrideIface)
	AddRules(rules map[string]interface{})
}

// Experimental.
// Struct proxy
type Eslint struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// Direct access to the eslint configuration (escape hatch).
	// Experimental.
	Config interface{} `json:"config"`
	// File patterns that should not be linted.
	// Experimental.
	IgnorePatterns []string `json:"ignorePatterns"`
	// eslint overrides.
	// Experimental.
	Overrides []EslintOverrideIface `json:"overrides"`
	// eslint rules.
	// Experimental.
	Rules map[string][]interface{} `json:"rules"`
}

func (e *Eslint) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		e,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (e *Eslint) GetConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		e,
		"config",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (e *Eslint) GetIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		e,
		"ignorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (e *Eslint) GetOverrides() []EslintOverrideIface {
	var returns []EslintOverrideIface
	_jsii_.Get(
		e,
		"overrides",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOverrideIface)(nil)).Elem(): reflect.TypeOf((*EslintOverride)(nil)).Elem(),
		},
	)
	return returns
}

func (e *Eslint) GetRules() map[string][]interface{} {
	var returns map[string][]interface{}
	_jsii_.Get(
		e,
		"rules",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*[]interface{})(nil)).Elem(): reflect.TypeOf((*[]interface{})(nil)).Elem(),
		},
	)
	return returns
}


func NewEslint(project NodeProjectIface, options EslintOptionsIface) EslintIface {
	_init_.Initialize()
	self := Eslint{}
	_jsii_.Create(
		"projen.Eslint",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (e *Eslint) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		e,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (e *Eslint) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		e,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (e *Eslint) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		e,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (e *Eslint) AddIgnorePattern(pattern string) {
	var returns interface{}
	_jsii_.Invoke(
		e,
		"addIgnorePattern",
		[]interface{}{pattern},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (e *Eslint) AddOverride(override EslintOverrideIface) {
	var returns interface{}
	_jsii_.Invoke(
		e,
		"addOverride",
		[]interface{}{override},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (e *Eslint) AddRules(rules map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		e,
		"addRules",
		[]interface{}{rules},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// EslintOptionsIface is the public interface for the custom type EslintOptions
// Experimental.
type EslintOptionsIface interface {
	GetDirs() []string
	GetDevdirs() []string
	GetFileExtensions() []string
	GetIgnorePatterns() []string
	GetLintProjenRc() bool
	GetTsconfigPath() string
}

// Experimental.
// Struct proxy
type EslintOptions struct {
	// Directories with source files to lint (e.g. [ "src" ]).
	// Experimental.
	Dirs []string `json:"dirs"`
	// Directories with source files that include tests and build tools.
	// 
	// These
	// sources are linted but may also import packages from `devDependencies`.
	// Experimental.
	Devdirs []string `json:"devdirs"`
	// File types that should be linted (e.g. [ ".js", ".ts" ]).
	// Experimental.
	FileExtensions []string `json:"fileExtensions"`
	// List of file patterns that should not be linted, using the same syntax as .gitignore patterns.
	// Experimental.
	IgnorePatterns []string `json:"ignorePatterns"`
	// Should we lint .projenrc.js.
	// Experimental.
	LintProjenRc bool `json:"lintProjenRc"`
	// Path to `tsconfig.json` which should be used by eslint.
	// Experimental.
	TsconfigPath string `json:"tsconfigPath"`
}

func (e *EslintOptions) GetDirs() []string {
	var returns []string
	_jsii_.Get(
		e,
		"dirs",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (e *EslintOptions) GetDevdirs() []string {
	var returns []string
	_jsii_.Get(
		e,
		"devdirs",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (e *EslintOptions) GetFileExtensions() []string {
	var returns []string
	_jsii_.Get(
		e,
		"fileExtensions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (e *EslintOptions) GetIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		e,
		"ignorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (e *EslintOptions) GetLintProjenRc() bool {
	var returns bool
	_jsii_.Get(
		e,
		"lintProjenRc",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (e *EslintOptions) GetTsconfigPath() string {
	var returns string
	_jsii_.Get(
		e,
		"tsconfigPath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// EslintOverrideIface is the public interface for the custom type EslintOverride
// Experimental.
type EslintOverrideIface interface {
	GetFiles() []string
	GetRules() map[string]interface{}
}

// eslint rules override.
// Experimental.
// Struct proxy
type EslintOverride struct {
	// Files or file patterns on which to apply the override.
	// Experimental.
	Files []string `json:"files"`
	// The overriden rules.
	// Experimental.
	Rules map[string]interface{} `json:"rules"`
}

func (e *EslintOverride) GetFiles() []string {
	var returns []string
	_jsii_.Get(
		e,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (e *EslintOverride) GetRules() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		e,
		"rules",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type FileBaseIface interface {
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
}

// Experimental.
// Struct proxy
type FileBase struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (f *FileBase) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		f,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (f *FileBase) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		f,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBase) GetPath() string {
	var returns string
	_jsii_.Get(
		f,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBase) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		f,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBase) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		f,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewFileBase(project ProjectIface, filePath string, options FileBaseOptionsIface) FileBaseIface {
	_init_.Initialize()
	self := FileBase{}
	_jsii_.Create(
		"projen.FileBase",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (f *FileBase) SetExecutable(val bool) {
	_jsii_.Set(
		f,
		"executable",
		val,
	)
}

func (f *FileBase) SetReadonly(val bool) {
	_jsii_.Set(
		f,
		"readonly",
		val,
	)
}

func FileBase_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.FileBase",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBase) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		f,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (f *FileBase) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		f,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (f *FileBase) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		f,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (f *FileBase) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		f,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// FileBaseOptionsIface is the public interface for the custom type FileBaseOptions
// Experimental.
type FileBaseOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
}

// Experimental.
// Struct proxy
type FileBaseOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (f *FileBaseOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		f,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBaseOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		f,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBaseOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		f,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (f *FileBaseOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		f,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type GitpodIface interface {
	IDevEnvironmentIface
	GetProject() ProjectIface
	GetConfig() interface{}
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddCustomTask(options GitpodTaskIface)
	AddDockerImage(image DevEnvironmentDockerImageIface)
	AddPorts(ports string)
	AddPrebuilds(config GitpodPrebuildsIface)
	AddTasks(tasks tasks.TaskIface)
	AddVscodeExtensions(extensions string)
}

// The Gitpod component which emits .gitpod.yml.
// Experimental.
// Struct proxy
type Gitpod struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// Direct access to the gitpod configuration (escape hatch).
	// Experimental.
	Config interface{} `json:"config"`
}

func (g *Gitpod) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		g,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (g *Gitpod) GetConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		g,
		"config",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewGitpod(project ProjectIface, options GitpodOptionsIface) GitpodIface {
	_init_.Initialize()
	self := Gitpod{}
	_jsii_.Create(
		"projen.Gitpod",
		[]interface{}{project, options},
		[]_jsii_.FQN{"projen.IDevEnvironment"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (g *Gitpod) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) AddCustomTask(options GitpodTaskIface) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addCustomTask",
		[]interface{}{options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) AddDockerImage(image DevEnvironmentDockerImageIface) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addDockerImage",
		[]interface{}{image},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) AddPorts(ports string) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addPorts",
		[]interface{}{ports},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) AddPrebuilds(config GitpodPrebuildsIface) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addPrebuilds",
		[]interface{}{config},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) AddTasks(tasks tasks.TaskIface) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addTasks",
		[]interface{}{tasks},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (g *Gitpod) AddVscodeExtensions(extensions string) {
	var returns interface{}
	_jsii_.Invoke(
		g,
		"addVscodeExtensions",
		[]interface{}{extensions},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// What to do when a service on a port is detected.
// Experimental.
type GitpodOnOpen string

const (
	GitpodOnOpenOpenBrowser GitpodOnOpen = "OPEN_BROWSER"
	GitpodOnOpenOpenPreview GitpodOnOpen = "OPEN_PREVIEW"
	GitpodOnOpenNotify GitpodOnOpen = "NOTIFY"
	GitpodOnOpenIgnore GitpodOnOpen = "IGNORE"
)

// Configure where in the IDE the terminal should be opened.
// Experimental.
type GitpodOpenIn string

const (
	GitpodOpenInBottom GitpodOpenIn = "BOTTOM"
	GitpodOpenInLeft GitpodOpenIn = "LEFT"
	GitpodOpenInRight GitpodOpenIn = "RIGHT"
	GitpodOpenInMain GitpodOpenIn = "MAIN"
)

// Configure how the terminal should be opened relative to the previous task.
// Experimental.
type GitpodOpenMode string

const (
	GitpodOpenModeTabAfter GitpodOpenMode = "TAB_AFTER"
	GitpodOpenModeTabBefore GitpodOpenMode = "TAB_BEFORE"
	GitpodOpenModeSplitRight GitpodOpenMode = "SPLIT_RIGHT"
	GitpodOpenModeSplitLeft GitpodOpenMode = "SPLIT_LEFT"
	GitpodOpenModeSplitTop GitpodOpenMode = "SPLIT_TOP"
	GitpodOpenModeSplitBottom GitpodOpenMode = "SPLIT_BOTTOM"
)

// GitpodOptionsIface is the public interface for the custom type GitpodOptions
// Experimental.
type GitpodOptionsIface interface {
	GetDockerImage() DevEnvironmentDockerImageIface
	GetPorts() []string
	GetTasks() []tasks.TaskIface
	GetVscodeExtensions() []string
	GetPrebuilds() GitpodPrebuildsIface
}

// Constructor options for the Gitpod component.
// 
// By default, Gitpod uses the 'gitpod/workspace-full' docker image.
// See: https://github.com/gitpod-io/workspace-images/blob/master/full/Dockerfile

By default, all tasks will be run in parallel. To run the tasks in sequence,
create a new task and specify the other tasks as subtasks.
//
// Experimental.
// Struct proxy
type GitpodOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage DevEnvironmentDockerImageIface `json:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports []string `json:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks []tasks.TaskIface `json:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions []string `json:"vscodeExtensions"`
	// Optional Gitpod's Github App integration for prebuilds If this is not set and Gitpod's Github App is installed, then Gitpod will apply these defaults: https://www.gitpod.io/docs/prebuilds/#configure-the-github-app.
	// Experimental.
	Prebuilds GitpodPrebuildsIface `json:"prebuilds"`
}

func (g *GitpodOptions) GetDockerImage() DevEnvironmentDockerImageIface {
	var returns DevEnvironmentDockerImageIface
	_jsii_.Get(
		g,
		"dockerImage",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DevEnvironmentDockerImageIface)(nil)).Elem(): reflect.TypeOf((*DevEnvironmentDockerImage)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodOptions) GetPorts() []string {
	var returns []string
	_jsii_.Get(
		g,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodOptions) GetTasks() []tasks.TaskIface {
	var returns []tasks.TaskIface
	_jsii_.Get(
		g,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodOptions) GetVscodeExtensions() []string {
	var returns []string
	_jsii_.Get(
		g,
		"vscodeExtensions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodOptions) GetPrebuilds() GitpodPrebuildsIface {
	var returns GitpodPrebuildsIface
	_jsii_.Get(
		g,
		"prebuilds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodPrebuildsIface)(nil)).Elem(): reflect.TypeOf((*GitpodPrebuilds)(nil)).Elem(),
		},
	)
	return returns
}


// GitpodPortIface is the public interface for the custom type GitpodPort
// Experimental.
type GitpodPortIface interface {
	GetOnOpen() GitpodOnOpen
	GetPort() string
	GetVisibility() GitpodPortVisibility
}

// Options for an exposed port on Gitpod.
// Experimental.
// Struct proxy
type GitpodPort struct {
	// What to do when a service on a port is detected.
	// Experimental.
	OnOpen GitpodOnOpen `json:"onOpen"`
	// A port that should be exposed (forwarded) from the container.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Port string `json:"port"`
	// Whether the port visibility should be private or public.
	// Experimental.
	Visibility GitpodPortVisibility `json:"visibility"`
}

func (g *GitpodPort) GetOnOpen() GitpodOnOpen {
	var returns GitpodOnOpen
	_jsii_.Get(
		g,
		"onOpen",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodOnOpen)(nil)).Elem(): reflect.TypeOf((*GitpodOnOpen)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodPort) GetPort() string {
	var returns string
	_jsii_.Get(
		g,
		"port",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPort) GetVisibility() GitpodPortVisibility {
	var returns GitpodPortVisibility
	_jsii_.Get(
		g,
		"visibility",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodPortVisibility)(nil)).Elem(): reflect.TypeOf((*GitpodPortVisibility)(nil)).Elem(),
		},
	)
	return returns
}


// Whether the port visibility should be private or public.
// Experimental.
type GitpodPortVisibility string

const (
	GitpodPortVisibilityPublic GitpodPortVisibility = "PUBLIC"
	GitpodPortVisibilityPrivate GitpodPortVisibility = "PRIVATE"
)

// GitpodPrebuildsIface is the public interface for the custom type GitpodPrebuilds
// Experimental.
type GitpodPrebuildsIface interface {
	GetAddBadge() bool
	GetAddCheck() bool
	GetAddComment() bool
	GetAddLabel() bool
	GetBranches() bool
	GetMaster() bool
	GetPullRequests() bool
	GetPullRequestsFromForks() bool
}

// Configure the Gitpod App for prebuilds.
// 
// Currently only GitHub is supported.
// See: https://www.gitpod.io/docs/prebuilds/
//
// Experimental.
// Struct proxy
type GitpodPrebuilds struct {
	// Add a "Review in Gitpod" button to the pull request's description.
	// Experimental.
	AddBadge bool `json:"addBadge"`
	// Add a check to pull requests.
	// Experimental.
	AddCheck bool `json:"addCheck"`
	// Add a "Review in Gitpod" button as a comment to pull requests.
	// Experimental.
	AddComment bool `json:"addComment"`
	// Add a label once the prebuild is ready to pull requests.
	// Experimental.
	AddLabel bool `json:"addLabel"`
	// Enable for all branches in this repo.
	// Experimental.
	Branches bool `json:"branches"`
	// Enable for the master/default branch.
	// Experimental.
	Master bool `json:"master"`
	// Enable for pull requests coming from this repo.
	// Experimental.
	PullRequests bool `json:"pullRequests"`
	// Enable for pull requests coming from forks.
	// Experimental.
	PullRequestsFromForks bool `json:"pullRequestsFromForks"`
}

func (g *GitpodPrebuilds) GetAddBadge() bool {
	var returns bool
	_jsii_.Get(
		g,
		"addBadge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetAddCheck() bool {
	var returns bool
	_jsii_.Get(
		g,
		"addCheck",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetAddComment() bool {
	var returns bool
	_jsii_.Get(
		g,
		"addComment",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetAddLabel() bool {
	var returns bool
	_jsii_.Get(
		g,
		"addLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetBranches() bool {
	var returns bool
	_jsii_.Get(
		g,
		"branches",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetMaster() bool {
	var returns bool
	_jsii_.Get(
		g,
		"master",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetPullRequests() bool {
	var returns bool
	_jsii_.Get(
		g,
		"pullRequests",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodPrebuilds) GetPullRequestsFromForks() bool {
	var returns bool
	_jsii_.Get(
		g,
		"pullRequestsFromForks",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// GitpodTaskIface is the public interface for the custom type GitpodTask
// Experimental.
type GitpodTaskIface interface {
	GetCommand() string
	GetBefore() string
	GetInit() string
	GetName() string
	GetOpenIn() GitpodOpenIn
	GetOpenMode() GitpodOpenMode
	GetPrebuild() string
}

// Configure options for a task to be run when opening a Gitpod workspace (e.g. running tests, or starting a dev server).
// 
// Start Mode         | Execution
// Fresh Workspace    | before && init && command
// Restart Workspace  | before && command
// Snapshot           | before && command
// Prebuild           | before && init && prebuild
// Experimental.
// Struct proxy
type GitpodTask struct {
	// Required.
	// 
	// The shell command to run
	// Experimental.
	Command string `json:"command"`
	// In case you need to run something even before init, that is a requirement for both init and command, you can use the before property.
	// Experimental.
	Before string `json:"before"`
	// The init property can be used to specify shell commands that should only be executed after a workspace was freshly cloned and needs to be initialized somehow.
	// 
	// Such tasks are usually builds or downloading
	// dependencies. Anything you only want to do once but not when you restart a workspace or start a snapshot.
	// Experimental.
	Init string `json:"init"`
	// A name for this task.
	// Experimental.
	Name string `json:"name"`
	// You can configure where in the IDE the terminal should be opened.
	// Experimental.
	OpenIn GitpodOpenIn `json:"openIn"`
	// You can configure how the terminal should be opened relative to the previous task.
	// Experimental.
	OpenMode GitpodOpenMode `json:"openMode"`
	// The optional prebuild command will be executed during prebuilds.
	// 
	// It is meant to run additional long running
	// processes that could be useful, e.g. running test suites.
	// Experimental.
	Prebuild string `json:"prebuild"`
}

func (g *GitpodTask) GetCommand() string {
	var returns string
	_jsii_.Get(
		g,
		"command",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodTask) GetBefore() string {
	var returns string
	_jsii_.Get(
		g,
		"before",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodTask) GetInit() string {
	var returns string
	_jsii_.Get(
		g,
		"init",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodTask) GetName() string {
	var returns string
	_jsii_.Get(
		g,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (g *GitpodTask) GetOpenIn() GitpodOpenIn {
	var returns GitpodOpenIn
	_jsii_.Get(
		g,
		"openIn",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodOpenIn)(nil)).Elem(): reflect.TypeOf((*GitpodOpenIn)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodTask) GetOpenMode() GitpodOpenMode {
	var returns GitpodOpenMode
	_jsii_.Get(
		g,
		"openMode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodOpenMode)(nil)).Elem(): reflect.TypeOf((*GitpodOpenMode)(nil)).Elem(),
		},
	)
	return returns
}

func (g *GitpodTask) GetPrebuild() string {
	var returns string
	_jsii_.Get(
		g,
		"prebuild",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// HasteConfigIface is the public interface for the custom type HasteConfig
// Experimental.
type HasteConfigIface interface {
	GetComputeSha1() bool
	GetDefaultPlatform() string
	GetHasteImplModulePath() string
	GetPlatforms() []string
	GetThrowOnModuleCollision() bool
}

// Experimental.
// Struct proxy
type HasteConfig struct {
	// Experimental.
	ComputeSha1 bool `json:"computeSha1"`
	// Experimental.
	DefaultPlatform string `json:"defaultPlatform"`
	// Experimental.
	HasteImplModulePath string `json:"hasteImplModulePath"`
	// Experimental.
	Platforms []string `json:"platforms"`
	// Experimental.
	ThrowOnModuleCollision bool `json:"throwOnModuleCollision"`
}

func (h *HasteConfig) GetComputeSha1() bool {
	var returns bool
	_jsii_.Get(
		h,
		"computeSha1",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (h *HasteConfig) GetDefaultPlatform() string {
	var returns string
	_jsii_.Get(
		h,
		"defaultPlatform",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (h *HasteConfig) GetHasteImplModulePath() string {
	var returns string
	_jsii_.Get(
		h,
		"hasteImplModulePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (h *HasteConfig) GetPlatforms() []string {
	var returns []string
	_jsii_.Get(
		h,
		"platforms",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (h *HasteConfig) GetThrowOnModuleCollision() bool {
	var returns bool
	_jsii_.Get(
		h,
		"throwOnModuleCollision",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Abstract interface for container-based development environments, such as Gitpod and GitHub Codespaces.
// Experimental.
type IDevEnvironmentIface interface {
	// Add a custom Docker image or Dockerfile for the container.
	// Experimental.
	AddDockerImage(image DevEnvironmentDockerImageIface)
	// Adds ports that should be exposed (forwarded) from the container.
	// Experimental.
	AddPorts(ports string)
	// Adds tasks to run when the container starts.
	// Experimental.
	AddTasks(tasks tasks.TaskIface)
	// Adds a list of VSCode extensions that should be automatically installed in the container.
	// Experimental.
	AddVscodeExtensions(extensions string)
}

type IDevEnvironment struct {}

func (i *IDevEnvironment) AddDockerImage(image DevEnvironmentDockerImageIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addDockerImage",
		[]interface{}{image},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IDevEnvironment) AddPorts(ports string) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addPorts",
		[]interface{}{ports},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IDevEnvironment) AddTasks(tasks tasks.TaskIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addTasks",
		[]interface{}{tasks},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IDevEnvironment) AddVscodeExtensions(extensions string) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addVscodeExtensions",
		[]interface{}{extensions},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// An interface providing the name of a docker compose service.
// Experimental.
type IDockerComposeServiceNameIface interface {
	// The name of the docker compose service.
	// Experimental.
	GetServiceName() string
}

type IDockerComposeServiceName struct {}

func (i *IDockerComposeServiceName) GetServiceName() string {
	var returns string
	_jsii_.Get(
		i,
		"serviceName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Volume binding information.
// Experimental.
type IDockerComposeVolumeBindingIface interface {
	// Binds the requested volume to the docker-compose volume configuration and provide mounting instructions for synthesis.
	//
	// Returns: mounting instructions for the service.
	// Experimental.
	Bind(volumeConfig IDockerComposeVolumeConfigIface) DockerComposeVolumeMountIface
}

type IDockerComposeVolumeBinding struct {}

func (i *IDockerComposeVolumeBinding) Bind(volumeConfig IDockerComposeVolumeConfigIface) DockerComposeVolumeMountIface {
	var returns DockerComposeVolumeMountIface
	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{volumeConfig},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DockerComposeVolumeMountIface)(nil)).Elem(): reflect.TypeOf((*DockerComposeVolumeMount)(nil)).Elem(),
		},
	)
	return returns
}

// Storage for volume configuration.
// Experimental.
type IDockerComposeVolumeConfigIface interface {
	// Add volume configuration to the repository.
	// Experimental.
	AddVolumeConfiguration(volumeName string, configuration DockerComposeVolumeConfigIface)
}

type IDockerComposeVolumeConfig struct {}

func (i *IDockerComposeVolumeConfig) AddVolumeConfiguration(volumeName string, configuration DockerComposeVolumeConfigIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addVolumeConfiguration",
		[]interface{}{volumeName, configuration},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Files that may include the Projen marker.
// Experimental.
type IMarkableFileIface interface {
	// Adds the projen marker to the file.
	// Experimental.
	GetMarker() bool
}

type IMarkableFile struct {}

func (i *IMarkableFile) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		i,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// API for resolving tokens when synthesizing file content.
// Experimental.
type IResolverIface interface {
	// Given a value (object/string/array/whatever, looks up any functions inside the object and returns an object where all functions are called.
	// Experimental.
	Resolve(value interface{}, options ResolveOptionsIface) interface{}
}

type IResolver struct {}

func (i *IResolver) Resolve(value interface{}, options ResolveOptionsIface) interface{} {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{value, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Class interface
type IgnoreFileIface interface {
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddPatterns(patterns string)
	Exclude(patterns string)
	Include(patterns string)
	RemovePatterns(patterns string)
}

// Experimental.
// Struct proxy
type IgnoreFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (i *IgnoreFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		i,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IgnoreFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		i,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IgnoreFile) GetPath() string {
	var returns string
	_jsii_.Get(
		i,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IgnoreFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		i,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IgnoreFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		i,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewIgnoreFile(project ProjectIface, filePath string) IgnoreFileIface {
	_init_.Initialize()
	self := IgnoreFile{}
	_jsii_.Create(
		"projen.IgnoreFile",
		[]interface{}{project, filePath},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (i *IgnoreFile) SetExecutable(val bool) {
	_jsii_.Set(
		i,
		"executable",
		val,
	)
}

func (i *IgnoreFile) SetReadonly(val bool) {
	_jsii_.Set(
		i,
		"readonly",
		val,
	)
}

func IgnoreFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.IgnoreFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IgnoreFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IgnoreFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IgnoreFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IgnoreFile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		i,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IgnoreFile) AddPatterns(patterns string) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addPatterns",
		[]interface{}{patterns},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IgnoreFile) Exclude(patterns string) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"exclude",
		[]interface{}{patterns},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IgnoreFile) Include(patterns string) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"include",
		[]interface{}{patterns},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IgnoreFile) RemovePatterns(patterns string) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"removePatterns",
		[]interface{}{patterns},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type JestIface interface {
	GetConfig() interface{}
	AddIgnorePattern(pattern string)
	AddReporter(reporter interface{})
	AddSnapshotResolver(file string)
	AddTestMatch(pattern string)
	AddWatchIgnorePattern(pattern string)
	GenerateTypescriptConfig(options TypescriptConfigOptionsIface) TypescriptConfigIface
}

// Installs the following npm scripts:.
// 
// - `test` will run `jest --passWithNoTests`
// - `test:watch` will run `jest --watch`
// - `test:update` will run `jest -u`
// Experimental.
// Struct proxy
type Jest struct {
	// Escape hatch.
	// Experimental.
	Config interface{} `json:"config"`
}

func (j *Jest) GetConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewJest(project NodeProjectIface, options JestOptionsIface) JestIface {
	_init_.Initialize()
	self := Jest{}
	_jsii_.Create(
		"projen.Jest",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (j *Jest) AddIgnorePattern(pattern string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addIgnorePattern",
		[]interface{}{pattern},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Jest) AddReporter(reporter interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addReporter",
		[]interface{}{reporter},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Jest) AddSnapshotResolver(file string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addSnapshotResolver",
		[]interface{}{file},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Jest) AddTestMatch(pattern string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addTestMatch",
		[]interface{}{pattern},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Jest) AddWatchIgnorePattern(pattern string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addWatchIgnorePattern",
		[]interface{}{pattern},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Jest) GenerateTypescriptConfig(options TypescriptConfigOptionsIface) TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Invoke(
		j,
		"generateTypescriptConfig",
		[]interface{}{options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

// JestConfigOptionsIface is the public interface for the custom type JestConfigOptions
// Experimental.
type JestConfigOptionsIface interface {
	GetAutomock() bool
	GetBail() interface{}
	GetCacheDirectory() string
	GetClearMocks() bool
	GetCollectCoverage() bool
	GetCollectCoverageFrom() []string
	GetCoverageDirectory() string
	GetCoveragePathIgnorePatterns() []string
	GetCoverageProvider() string
	GetCoverageReporters() []string
	GetCoverageThreshold() CoverageThresholdIface
	GetDependencyExtractor() string
	GetDisplayName() interface{}
	GetErrorOnDeprecated() bool
	GetExtraGlobals() []string
	GetForceCoverageMatch() []string
	GetGlobals() interface{}
	GetGlobalSetup() string
	GetGlobalTeardown() string
	GetHaste() HasteConfigIface
	GetInjectGlobals() bool
	GetMaxConcurrency() float64
	GetModuleDirectories() []string
	GetModuleFileExtensions() []string
	GetModuleNameMapper() map[string]interface{}
	GetModulePathIgnorePatterns() []string
	GetModulePaths() []string
	GetNotify() bool
	GetNotifyMode() string
	GetPreset() string
	GetPrettierPath() string
	GetProjects() []interface{}
	GetReporters() []interface{}
	GetResetMocks() bool
	GetResetModules() bool
	GetResolver() string
	GetRestoreMocks() bool
	GetRootDir() string
	GetRoots() []string
	GetRunner() string
	GetSetupFiles() []string
	GetSetupFilesAfterEnv() []string
	GetSlowTestThreshold() float64
	GetSnapshotResolver() string
	GetSnapshotSerializers() []string
	GetTestEnvironment() string
	GetTestEnvironmentOptions() interface{}
	GetTestFailureExitCode() float64
	GetTestMatch() []string
	GetTestPathIgnorePatterns() []string
	GetTestRegex() interface{}
	GetTestResultsProcessor() string
	GetTestRunner() string
	GetTestSequencer() string
	GetTestTimeout() float64
	GetTestUrl() string
	GetTimers() string
	GetTransform() map[string]interface{}
	GetTransformIgnorePatterns() []string
	GetUnmockedModulePathPatterns() []string
	GetVerbose() bool
	GetWatchman() bool
	GetWatchPathIgnorePatterns() []string
	GetWatchPlugins() map[string]interface{}
}

// Experimental.
// Struct proxy
type JestConfigOptions struct {
	// This option tells Jest that all imported modules in your tests should be mocked automatically.
	// 
	// All modules used in your tests will have a replacement implementation, keeping the API surface
	// Experimental.
	Automock bool `json:"automock"`
	// By default, Jest runs all tests and produces all errors into the console upon completion.
	// 
	// The bail config option can be used here to have Jest stop running tests after n failures.
	// Setting bail to true is the same as setting bail to 1.
	// Experimental.
	Bail interface{} `json:"bail"`
	// The directory where Jest should store its cached dependency information.
	// Experimental.
	CacheDirectory string `json:"cacheDirectory"`
	// Automatically clear mock calls and instances before every test.
	// 
	// Equivalent to calling jest.clearAllMocks() before each test.
	// This does not remove any mock implementation that may have been provided
	// Experimental.
	ClearMocks bool `json:"clearMocks"`
	// Indicates whether the coverage information should be collected while executing the test.
	// 
	// Because this retrofits all executed files with coverage collection statements,
	// it may significantly slow down your tests
	// Experimental.
	CollectCoverage bool `json:"collectCoverage"`
	// An array of glob patterns indicating a set of files for which coverage information should be collected.
	// Experimental.
	CollectCoverageFrom []string `json:"collectCoverageFrom"`
	// The directory where Jest should output its coverage files.
	// Experimental.
	CoverageDirectory string `json:"coverageDirectory"`
	// An array of regexp pattern strings that are matched against all file paths before executing the test.
	// 
	// If the file path matches any of the patterns, coverage information will be skipped
	// Experimental.
	CoveragePathIgnorePatterns []string `json:"coveragePathIgnorePatterns"`
	// Indicates which provider should be used to instrument code for coverage.
	// 
	// Allowed values are babel (default) or v8
	// Experimental.
	CoverageProvider string `json:"coverageProvider"`
	// A list of reporter names that Jest uses when writing coverage reports.
	// 
	// Any istanbul reporter can be used
	// Experimental.
	CoverageReporters []string `json:"coverageReporters"`
	// Specify the global coverage thresholds.
	// 
	// This will be used to configure minimum threshold enforcement
	// for coverage results. Thresholds can be specified as global, as a glob, and as a directory or file path.
	// If thresholds aren't met, jest will fail.
	// Experimental.
	CoverageThreshold CoverageThresholdIface `json:"coverageThreshold"`
	// This option allows the use of a custom dependency extractor.
	// 
	// It must be a node module that exports an object with an extract function
	// Experimental.
	DependencyExtractor string `json:"dependencyExtractor"`
	// Allows for a label to be printed alongside a test while it is running.
	// Experimental.
	DisplayName interface{} `json:"displayName"`
	// Make calling deprecated APIs throw helpful error messages.
	// 
	// Useful for easing the upgrade process.
	// Experimental.
	ErrorOnDeprecated bool `json:"errorOnDeprecated"`
	// Test files run inside a vm, which slows calls to global context properties (e.g. Math). With this option you can specify extra properties to be defined inside the vm for faster lookups.
	// Experimental.
	ExtraGlobals []string `json:"extraGlobals"`
	// Test files are normally ignored from collecting code coverage.
	// 
	// With this option, you can overwrite this behavior and include otherwise ignored files in code coverage.
	// Experimental.
	ForceCoverageMatch []string `json:"forceCoverageMatch"`
	// A set of global variables that need to be available in all test environments.
	// Experimental.
	Globals interface{} `json:"globals"`
	// This option allows the use of a custom global setup module which exports an async function that is triggered once before all test suites.
	// 
	// This function gets Jest's globalConfig object as a parameter.
	// Experimental.
	GlobalSetup string `json:"globalSetup"`
	// This option allows the use of a custom global teardown module which exports an async function that is triggered once after all test suites.
	// 
	// This function gets Jest's globalConfig object as a parameter.
	// Experimental.
	GlobalTeardown string `json:"globalTeardown"`
	// This will be used to configure the behavior of jest-haste-map, Jest's internal file crawler/cache system.
	// Experimental.
	Haste HasteConfigIface `json:"haste"`
	// Insert Jest's globals (expect, test, describe, beforeEach etc.) into the global environment. If you set this to false, you should import from @jest/globals.
	// Experimental.
	InjectGlobals bool `json:"injectGlobals"`
	// A number limiting the number of tests that are allowed to run at the same time when using test.concurrent. Any test above this limit will be queued and executed once a slot is released.
	// Experimental.
	MaxConcurrency float64 `json:"maxConcurrency"`
	// An array of directory names to be searched recursively up from the requiring module's location.
	// 
	// Setting this option will override the default, if you wish to still search node_modules for packages
	// include it along with any other options: ["node_modules", "bower_components"]
	// Experimental.
	ModuleDirectories []string `json:"moduleDirectories"`
	// An array of file extensions your modules use.
	// 
	// If you require modules without specifying a file extension,
	// these are the extensions Jest will look for, in left-to-right order.
	// Experimental.
	ModuleFileExtensions []string `json:"moduleFileExtensions"`
	// A map from regular expressions to module names or to arrays of module names that allow to stub out resources, like images or styles with a single module.
	// Experimental.
	ModuleNameMapper map[string]interface{} `json:"moduleNameMapper"`
	// An array of regexp pattern strings that are matched against all module paths before those paths are to be considered 'visible' to the module loader.
	// 
	// If a given module's path matches any of the patterns,
	// it will not be require()-able in the test environment.
	// Experimental.
	ModulePathIgnorePatterns []string `json:"modulePathIgnorePatterns"`
	// An alternative API to setting the NODE_PATH env variable, modulePaths is an array of absolute paths to additional locations to search when resolving modules.
	// 
	// Use the <rootDir> string token to include
	// the path to your project's root directory. Example: ["<rootDir>/app/"].
	// Experimental.
	ModulePaths []string `json:"modulePaths"`
	// Activates notifications for test results.
	// Experimental.
	Notify bool `json:"notify"`
	// Specifies notification mode.
	// 
	// Requires notify: true
	// Experimental.
	NotifyMode string `json:"notifyMode"`
	// A preset that is used as a base for Jest's configuration.
	// 
	// A preset should point to an npm module
	// that has a jest-preset.json or jest-preset.js file at the root.
	// Experimental.
	Preset string `json:"preset"`
	// Sets the path to the prettier node module used to update inline snapshots.
	// Experimental.
	PrettierPath string `json:"prettierPath"`
	// When the projects configuration is provided with an array of paths or glob patterns, Jest will run tests in all of the specified projects at the same time.
	// 
	// This is great for monorepos or
	// when working on multiple projects at the same time.
	// Experimental.
	Projects []interface{} `json:"projects"`
	// Use this configuration option to add custom reporters to Jest.
	// 
	// A custom reporter is a class
	// that implements onRunStart, onTestStart, onTestResult, onRunComplete methods that will be
	// called when any of those events occurs.
	// Experimental.
	Reporters []interface{} `json:"reporters"`
	// Automatically reset mock state before every test.
	// 
	// Equivalent to calling jest.resetAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed but
	// does not restore their initial implementation.
	// Experimental.
	ResetMocks bool `json:"resetMocks"`
	// By default, each test file gets its own independent module registry.
	// 
	// Enabling resetModules
	// goes a step further and resets the module registry before running each individual test.
	// Experimental.
	ResetModules bool `json:"resetModules"`
	// This option allows the use of a custom resolver.
	// 
	// https://jestjs.io/docs/en/configuration#resolver-string
	// Experimental.
	Resolver string `json:"resolver"`
	// Automatically restore mock state before every test.
	// 
	// Equivalent to calling jest.restoreAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed and
	// restores their initial implementation.
	// Experimental.
	RestoreMocks bool `json:"restoreMocks"`
	// The root directory that Jest should scan for tests and modules within.
	// 
	// If you put your Jest
	// config inside your package.json and want the root directory to be the root of your repo, the
	// value for this config param will default to the directory of the package.json.
	// Experimental.
	RootDir string `json:"rootDir"`
	// A list of paths to directories that Jest should use to search for files in.
	// Experimental.
	Roots []string `json:"roots"`
	// This option allows you to use a custom runner instead of Jest's default test runner.
	// Experimental.
	Runner string `json:"runner"`
	// A list of paths to modules that run some code to configure or set up the testing environment.
	// 
	// Each setupFile will be run once per test file. Since every test runs in its own environment,
	// these scripts will be executed in the testing environment immediately before executing the
	// test code itself.
	// Experimental.
	SetupFiles []string `json:"setupFiles"`
	// A list of paths to modules that run some code to configure or set up the testing framework before each test file in the suite is executed.
	// 
	// Since setupFiles executes before the test
	// framework is installed in the environment, this script file presents you the opportunity of
	// running some code immediately after the test framework has been installed in the environment.
	// Experimental.
	SetupFilesAfterEnv []string `json:"setupFilesAfterEnv"`
	// The number of seconds after which a test is considered as slow and reported as such in the results.
	// Experimental.
	SlowTestThreshold float64 `json:"slowTestThreshold"`
	// The path to a module that can resolve test<->snapshot path.
	// 
	// This config option lets you customize
	// where Jest stores snapshot files on disk.
	// Experimental.
	SnapshotResolver string `json:"snapshotResolver"`
	// A list of paths to snapshot serializer modules Jest should use for snapshot testing.
	// Experimental.
	SnapshotSerializers []string `json:"snapshotSerializers"`
	// The test environment that will be used for testing.
	// 
	// The default environment in Jest is a
	// browser-like environment through jsdom. If you are building a node service, you can use the node
	// option to use a node-like environment instead.
	// Experimental.
	TestEnvironment string `json:"testEnvironment"`
	// Test environment options that will be passed to the testEnvironment.
	// 
	// The relevant options depend on the environment.
	// Experimental.
	TestEnvironmentOptions interface{} `json:"testEnvironmentOptions"`
	// The exit code Jest returns on test failure.
	// Experimental.
	TestFailureExitCode float64 `json:"testFailureExitCode"`
	// The glob patterns Jest uses to detect test files.
	// 
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Experimental.
	TestMatch []string `json:"testMatch"`
	// An array of regexp pattern strings that are matched against all test paths before executing the test.
	// 
	// If the test path matches any of the patterns, it will be skipped.
	// Experimental.
	TestPathIgnorePatterns []string `json:"testPathIgnorePatterns"`
	// The pattern or patterns Jest uses to detect test files.
	// 
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Experimental.
	TestRegex interface{} `json:"testRegex"`
	// This option allows the use of a custom results processor.
	// Experimental.
	TestResultsProcessor string `json:"testResultsProcessor"`
	// This option allows the use of a custom test runner.
	// 
	// The default is jasmine2. A custom test runner
	// can be provided by specifying a path to a test runner implementation.
	// Experimental.
	TestRunner string `json:"testRunner"`
	// This option allows you to use a custom sequencer instead of Jest's default.
	// 
	// Sort may optionally return a Promise.
	// Experimental.
	TestSequencer string `json:"testSequencer"`
	// Default timeout of a test in milliseconds.
	// Experimental.
	TestTimeout float64 `json:"testTimeout"`
	// This option sets the URL for the jsdom environment.
	// 
	// It is reflected in properties such as location.href.
	// Experimental.
	TestUrl string `json:"testURL"`
	// Setting this value to legacy or fake allows the use of fake timers for functions such as setTimeout.
	// 
	// Fake timers are useful when a piece of code sets a long timeout that we don't want to wait for in a test.
	// Experimental.
	Timers string `json:"timers"`
	// A map from regular expressions to paths to transformers.
	// 
	// A transformer is a module that provides a
	// synchronous function for transforming source files.
	// Experimental.
	Transform map[string]interface{} `json:"transform"`
	// An array of regexp pattern strings that are matched against all source file paths before transformation.
	// 
	// If the test path matches any of the patterns, it will not be transformed.
	// Experimental.
	TransformIgnorePatterns []string `json:"transformIgnorePatterns"`
	// An array of regexp pattern strings that are matched against all modules before the module loader will automatically return a mock for them.
	// 
	// If a module's path matches any of the patterns in this list, it
	// will not be automatically mocked by the module loader.
	// Experimental.
	UnmockedModulePathPatterns []string `json:"unmockedModulePathPatterns"`
	// Indicates whether each individual test should be reported during the run.
	// 
	// All errors will also
	// still be shown on the bottom after execution. Note that if there is only one test file being run
	// it will default to true.
	// Experimental.
	Verbose bool `json:"verbose"`
	// Whether to use watchman for file crawling.
	// Experimental.
	Watchman bool `json:"watchman"`
	// An array of RegExp patterns that are matched against all source file paths before re-running tests in watch mode.
	// 
	// If the file path matches any of the patterns, when it is updated, it will not trigger
	// a re-run of tests.
	// Experimental.
	WatchPathIgnorePatterns []string `json:"watchPathIgnorePatterns"`
	// Experimental.
	WatchPlugins map[string]interface{} `json:"watchPlugins"`
}

func (j *JestConfigOptions) GetAutomock() bool {
	var returns bool
	_jsii_.Get(
		j,
		"automock",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetBail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetCacheDirectory() string {
	var returns string
	_jsii_.Get(
		j,
		"cacheDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetClearMocks() bool {
	var returns bool
	_jsii_.Get(
		j,
		"clearMocks",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetCollectCoverage() bool {
	var returns bool
	_jsii_.Get(
		j,
		"collectCoverage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetCollectCoverageFrom() []string {
	var returns []string
	_jsii_.Get(
		j,
		"collectCoverageFrom",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetCoverageDirectory() string {
	var returns string
	_jsii_.Get(
		j,
		"coverageDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetCoveragePathIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"coveragePathIgnorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetCoverageProvider() string {
	var returns string
	_jsii_.Get(
		j,
		"coverageProvider",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetCoverageReporters() []string {
	var returns []string
	_jsii_.Get(
		j,
		"coverageReporters",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetCoverageThreshold() CoverageThresholdIface {
	var returns CoverageThresholdIface
	_jsii_.Get(
		j,
		"coverageThreshold",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*CoverageThresholdIface)(nil)).Elem(): reflect.TypeOf((*CoverageThreshold)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetDependencyExtractor() string {
	var returns string
	_jsii_.Get(
		j,
		"dependencyExtractor",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetDisplayName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"displayName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetErrorOnDeprecated() bool {
	var returns bool
	_jsii_.Get(
		j,
		"errorOnDeprecated",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetExtraGlobals() []string {
	var returns []string
	_jsii_.Get(
		j,
		"extraGlobals",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetForceCoverageMatch() []string {
	var returns []string
	_jsii_.Get(
		j,
		"forceCoverageMatch",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetGlobals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globals",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetGlobalSetup() string {
	var returns string
	_jsii_.Get(
		j,
		"globalSetup",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetGlobalTeardown() string {
	var returns string
	_jsii_.Get(
		j,
		"globalTeardown",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetHaste() HasteConfigIface {
	var returns HasteConfigIface
	_jsii_.Get(
		j,
		"haste",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*HasteConfigIface)(nil)).Elem(): reflect.TypeOf((*HasteConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetInjectGlobals() bool {
	var returns bool
	_jsii_.Get(
		j,
		"injectGlobals",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetMaxConcurrency() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetModuleDirectories() []string {
	var returns []string
	_jsii_.Get(
		j,
		"moduleDirectories",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetModuleFileExtensions() []string {
	var returns []string
	_jsii_.Get(
		j,
		"moduleFileExtensions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetModuleNameMapper() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		j,
		"moduleNameMapper",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetModulePathIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"modulePathIgnorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetModulePaths() []string {
	var returns []string
	_jsii_.Get(
		j,
		"modulePaths",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetNotify() bool {
	var returns bool
	_jsii_.Get(
		j,
		"notify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetNotifyMode() string {
	var returns string
	_jsii_.Get(
		j,
		"notifyMode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetPreset() string {
	var returns string
	_jsii_.Get(
		j,
		"preset",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetPrettierPath() string {
	var returns string
	_jsii_.Get(
		j,
		"prettierPath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetProjects() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		j,
		"projects",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetReporters() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		j,
		"reporters",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetResetMocks() bool {
	var returns bool
	_jsii_.Get(
		j,
		"resetMocks",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetResetModules() bool {
	var returns bool
	_jsii_.Get(
		j,
		"resetModules",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetResolver() string {
	var returns string
	_jsii_.Get(
		j,
		"resolver",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetRestoreMocks() bool {
	var returns bool
	_jsii_.Get(
		j,
		"restoreMocks",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetRootDir() string {
	var returns string
	_jsii_.Get(
		j,
		"rootDir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetRoots() []string {
	var returns []string
	_jsii_.Get(
		j,
		"roots",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetRunner() string {
	var returns string
	_jsii_.Get(
		j,
		"runner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetSetupFiles() []string {
	var returns []string
	_jsii_.Get(
		j,
		"setupFiles",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetSetupFilesAfterEnv() []string {
	var returns []string
	_jsii_.Get(
		j,
		"setupFilesAfterEnv",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetSlowTestThreshold() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"slowTestThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetSnapshotResolver() string {
	var returns string
	_jsii_.Get(
		j,
		"snapshotResolver",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetSnapshotSerializers() []string {
	var returns []string
	_jsii_.Get(
		j,
		"snapshotSerializers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetTestEnvironment() string {
	var returns string
	_jsii_.Get(
		j,
		"testEnvironment",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestEnvironmentOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testEnvironmentOptions",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestFailureExitCode() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"testFailureExitCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestMatch() []string {
	var returns []string
	_jsii_.Get(
		j,
		"testMatch",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetTestPathIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"testPathIgnorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetTestRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testRegex",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestResultsProcessor() string {
	var returns string
	_jsii_.Get(
		j,
		"testResultsProcessor",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestRunner() string {
	var returns string
	_jsii_.Get(
		j,
		"testRunner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestSequencer() string {
	var returns string
	_jsii_.Get(
		j,
		"testSequencer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestTimeout() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"testTimeout",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTestUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"testURL",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTimers() string {
	var returns string
	_jsii_.Get(
		j,
		"timers",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetTransform() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		j,
		"transform",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetTransformIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"transformIgnorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetUnmockedModulePathPatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"unmockedModulePathPatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetVerbose() bool {
	var returns bool
	_jsii_.Get(
		j,
		"verbose",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetWatchman() bool {
	var returns bool
	_jsii_.Get(
		j,
		"watchman",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestConfigOptions) GetWatchPathIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"watchPathIgnorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestConfigOptions) GetWatchPlugins() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		j,
		"watchPlugins",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// JestOptionsIface is the public interface for the custom type JestOptions
// Experimental.
type JestOptionsIface interface {
	GetCoverage() bool
	GetCoverageText() bool
	GetIgnorePatterns() []string
	GetJestConfig() JestConfigOptionsIface
	GetJestVersion() string
	GetJunitReporting() bool
	GetPreserveDefaultReporters() bool
	GetTypescriptConfig() TypescriptConfigOptionsIface
}

// Experimental.
// Struct proxy
type JestOptions struct {
	// Collect coverage.
	// 
	// Deprecated
	// Deprecated: use jestConfig.collectCoverage
	Coverage bool `json:"coverage"`
	// Include the `text` coverage reporter, which means that coverage summary is printed at the end of the jest execution.
	// Experimental.
	CoverageText bool `json:"coverageText"`
	// Defines `testPathIgnorePatterns` and `coveragePathIgnorePatterns`.
	// Deprecated: use jestConfig.coveragePathIgnorePatterns or jestConfig.testPathIgnorePatterns respectively
	IgnorePatterns []string `json:"ignorePatterns"`
	// Experimental.
	JestConfig JestConfigOptionsIface `json:"jestConfig"`
	// The version of jest to use.
	// Experimental.
	JestVersion string `json:"jestVersion"`
	// Result processing with jest-junit.
	// 
	// Output directory is `test-reports/`.
	// Experimental.
	JunitReporting bool `json:"junitReporting"`
	// Preserve the default Jest reporter when additional reporters are added.
	// Experimental.
	PreserveDefaultReporters bool `json:"preserveDefaultReporters"`
	// Experimental.
	TypescriptConfig TypescriptConfigOptionsIface `json:"typescriptConfig"`
}

func (j *JestOptions) GetCoverage() bool {
	var returns bool
	_jsii_.Get(
		j,
		"coverage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestOptions) GetCoverageText() bool {
	var returns bool
	_jsii_.Get(
		j,
		"coverageText",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestOptions) GetIgnorePatterns() []string {
	var returns []string
	_jsii_.Get(
		j,
		"ignorePatterns",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestOptions) GetJestConfig() JestConfigOptionsIface {
	var returns JestConfigOptionsIface
	_jsii_.Get(
		j,
		"jestConfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JestOptions) GetJestVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"jestVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestOptions) GetJunitReporting() bool {
	var returns bool
	_jsii_.Get(
		j,
		"junitReporting",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestOptions) GetPreserveDefaultReporters() bool {
	var returns bool
	_jsii_.Get(
		j,
		"preserveDefaultReporters",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JestOptions) GetTypescriptConfig() TypescriptConfigOptionsIface {
	var returns TypescriptConfigOptionsIface
	_jsii_.Get(
		j,
		"typescriptConfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}


// JsiiDotNetTargetIface is the public interface for the custom type JsiiDotNetTarget
// Experimental.
type JsiiDotNetTargetIface interface {
	GetDotNetNamespace() string
	GetPackageId() string
}

// Experimental.
// Struct proxy
type JsiiDotNetTarget struct {
	// Experimental.
	DotNetNamespace string `json:"dotNetNamespace"`
	// Experimental.
	PackageId string `json:"packageId"`
}

func (j *JsiiDotNetTarget) GetDotNetNamespace() string {
	var returns string
	_jsii_.Get(
		j,
		"dotNetNamespace",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiDotNetTarget) GetPackageId() string {
	var returns string
	_jsii_.Get(
		j,
		"packageId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// JsiiGoTargetIface is the public interface for the custom type JsiiGoTarget
// Experimental.
type JsiiGoTargetIface interface {
	GetModuleName() string
	GetGitBranch() string
	GetGitCommitMessage() string
	GetGithubRepo() string
	GetGithubTokenSecret() string
	GetGitUserEmail() string
	GetGitUserName() string
}

// Go target configuration.
// Experimental.
// Struct proxy
type JsiiGoTarget struct {
	// The name of the target go module.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ModuleName string `json:"moduleName"`
	// Branch to push to.
	// Experimental.
	GitBranch string `json:"gitBranch"`
	// The commit message.
	// Experimental.
	GitCommitMessage string `json:"gitCommitMessage"`
	// GitHub repository to push to.
	// Experimental.
	GithubRepo string `json:"githubRepo"`
	// The name of the secret that includes a personal GitHub access token used to push to the GitHub repository.
	// Experimental.
	GithubTokenSecret string `json:"githubTokenSecret"`
	// The email to use in the release git commit.
	// Experimental.
	GitUserEmail string `json:"gitUserEmail"`
	// The user name to use for the release git commit.
	// Experimental.
	GitUserName string `json:"gitUserName"`
}

func (j *JsiiGoTarget) GetModuleName() string {
	var returns string
	_jsii_.Get(
		j,
		"moduleName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiGoTarget) GetGitBranch() string {
	var returns string
	_jsii_.Get(
		j,
		"gitBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiGoTarget) GetGitCommitMessage() string {
	var returns string
	_jsii_.Get(
		j,
		"gitCommitMessage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiGoTarget) GetGithubRepo() string {
	var returns string
	_jsii_.Get(
		j,
		"githubRepo",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiGoTarget) GetGithubTokenSecret() string {
	var returns string
	_jsii_.Get(
		j,
		"githubTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiGoTarget) GetGitUserEmail() string {
	var returns string
	_jsii_.Get(
		j,
		"gitUserEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiGoTarget) GetGitUserName() string {
	var returns string
	_jsii_.Get(
		j,
		"gitUserName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// JsiiJavaTargetIface is the public interface for the custom type JsiiJavaTarget
// Experimental.
type JsiiJavaTargetIface interface {
	GetJavaPackage() string
	GetMavenArtifactId() string
	GetMavenGroupId() string
	GetMavenRepositoryUrl() string
	GetMavenServerId() string
}

// Experimental.
// Struct proxy
type JsiiJavaTarget struct {
	// Experimental.
	JavaPackage string `json:"javaPackage"`
	// Experimental.
	MavenArtifactId string `json:"mavenArtifactId"`
	// Experimental.
	MavenGroupId string `json:"mavenGroupId"`
	// Deployment repository when not deploying to Maven Central.
	// Experimental.
	MavenRepositoryUrl string `json:"mavenRepositoryUrl"`
	// Used in maven settings for credential lookup (e.g. use github when publishing to GitHub).
	// Experimental.
	MavenServerId string `json:"mavenServerId"`
}

func (j *JsiiJavaTarget) GetJavaPackage() string {
	var returns string
	_jsii_.Get(
		j,
		"javaPackage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiJavaTarget) GetMavenArtifactId() string {
	var returns string
	_jsii_.Get(
		j,
		"mavenArtifactId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiJavaTarget) GetMavenGroupId() string {
	var returns string
	_jsii_.Get(
		j,
		"mavenGroupId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiJavaTarget) GetMavenRepositoryUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"mavenRepositoryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiJavaTarget) GetMavenServerId() string {
	var returns string
	_jsii_.Get(
		j,
		"mavenServerId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type JsiiProjectIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	GetTwineRegistryUrl() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// Multi-language jsii library project.
// Experimental.
// Struct proxy
type JsiiProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
	// Experimental.
	TwineRegistryUrl string `json:"twineRegistryUrl"`
}

func (j *JsiiProject) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		j,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetDeps() deps.DependenciesIface {
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

func (j *JsiiProject) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		j,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		j,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetName() string {
	var returns string
	_jsii_.Get(
		j,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		j,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetTasks() tasks.TasksIface {
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

func (j *JsiiProject) GetDevContainer() vscode.DevContainerIface {
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

func (j *JsiiProject) GetGithub() github.GitHubIface {
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

func (j *JsiiProject) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		j,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		j,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetVscode() vscode.VsCodeIface {
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

func (j *JsiiProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		j,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		j,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		j,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		j,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		j,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		j,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		j,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		j,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		j,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		j,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		j,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetLibdir() string {
	var returns string
	_jsii_.Get(
		j,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		j,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetTestdir() string {
	var returns string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		j,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		j,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) GetTwineRegistryUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"twineRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewJsiiProject(options JsiiProjectOptionsIface) JsiiProjectIface {
	_init_.Initialize()
	self := JsiiProject{}
	_jsii_.Create(
		"projen.JsiiProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func JsiiProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.JsiiProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) AddExcludeFromCleanup(globs string) {
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

func (j *JsiiProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
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

func (j *JsiiProject) AddTip(message string) {
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

func (j *JsiiProject) PostSynthesize() {
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

func (j *JsiiProject) PreSynthesize() {
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

func (j *JsiiProject) Synth() {
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

func (j *JsiiProject) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		j,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		j,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		j,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		j,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsiiProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		j,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// JsiiProjectOptionsIface is the public interface for the custom type JsiiProjectOptions
// Experimental.
type JsiiProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetAuthor() string
	GetAuthorAddress() string
	GetRepositoryUrl() string
	GetCompat() bool
	GetCompatIgnore() string
	GetDocgen() bool
	GetDotnet() JsiiDotNetTargetIface
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetPublishToGo() JsiiGoTargetIface
	GetPublishToMaven() JsiiJavaTargetIface
	GetPublishToNuget() JsiiDotNetTargetIface
	GetPublishToPypi() JsiiPythonTargetIface
	GetPython() JsiiPythonTargetIface
	GetRootdir() string
}

// Experimental.
// Struct proxy
type JsiiProjectOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// The name of the library author.
	// Experimental.
	Author string `json:"author"`
	// Email or URL of the library author.
	// Experimental.
	AuthorAddress string `json:"authorAddress"`
	// Git repository URL.
	// Experimental.
	RepositoryUrl string `json:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	// 
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Experimental.
	Compat bool `json:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Experimental.
	CompatIgnore string `json:"compatIgnore"`
	// Automatically generate API.md from jsii.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Deprecated: use `publishToNuget`
	Dotnet JsiiDotNetTargetIface `json:"dotnet"`
	// Install eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Publish Go bindings to a git repository.
	// Experimental.
	PublishToGo JsiiGoTargetIface `json:"publishToGo"`
	// Publish to maven.
	// Experimental.
	PublishToMaven JsiiJavaTargetIface `json:"publishToMaven"`
	// Publish to NuGet.
	// Experimental.
	PublishToNuget JsiiDotNetTargetIface `json:"publishToNuget"`
	// Publish to pypi.
	// Experimental.
	PublishToPypi JsiiPythonTargetIface `json:"publishToPypi"`
	// Deprecated: use `publishToPyPi`
	Python JsiiPythonTargetIface `json:"python"`
	// Experimental.
	Rootdir string `json:"rootdir"`
}

func (j *JsiiProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		j,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		j,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		j,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		j,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		j,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		j,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		j,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		j,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		j,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		j,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		j,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		j,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDeps() []string {
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

func (j *JsiiProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		j,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		j,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		j,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		j,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		j,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		j,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		j,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		j,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		j,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		j,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		j,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		j,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		j,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		j,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		j,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		j,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		j,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		j,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		j,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		j,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		j,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		j,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		j,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		j,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		j,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		j,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		j,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		j,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		j,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		j,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		j,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		j,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		j,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		j,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		j,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		j,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		j,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		j,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		j,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		j,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		j,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		j,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		j,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		j,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		j,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		j,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		j,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAuthor() string {
	var returns string
	_jsii_.Get(
		j,
		"author",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetAuthorAddress() string {
	var returns string
	_jsii_.Get(
		j,
		"authorAddress",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetRepositoryUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetCompat() bool {
	var returns bool
	_jsii_.Get(
		j,
		"compat",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetCompatIgnore() string {
	var returns string
	_jsii_.Get(
		j,
		"compatIgnore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		j,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetDotnet() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		j,
		"dotnet",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		j,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiProjectOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		j,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPublishToGo() JsiiGoTargetIface {
	var returns JsiiGoTargetIface
	_jsii_.Get(
		j,
		"publishToGo",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiGoTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiGoTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPublishToMaven() JsiiJavaTargetIface {
	var returns JsiiJavaTargetIface
	_jsii_.Get(
		j,
		"publishToMaven",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiJavaTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiJavaTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPublishToNuget() JsiiDotNetTargetIface {
	var returns JsiiDotNetTargetIface
	_jsii_.Get(
		j,
		"publishToNuget",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiDotNetTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiDotNetTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPublishToPypi() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		j,
		"publishToPypi",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetPython() JsiiPythonTargetIface {
	var returns JsiiPythonTargetIface
	_jsii_.Get(
		j,
		"python",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsiiPythonTargetIface)(nil)).Elem(): reflect.TypeOf((*JsiiPythonTarget)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsiiProjectOptions) GetRootdir() string {
	var returns string
	_jsii_.Get(
		j,
		"rootdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// JsiiPythonTargetIface is the public interface for the custom type JsiiPythonTarget
// Experimental.
type JsiiPythonTargetIface interface {
	GetDistName() string
	GetModule() string
	GetTwineRegistryUrl() string
}

// Experimental.
// Struct proxy
type JsiiPythonTarget struct {
	// Experimental.
	DistName string `json:"distName"`
	// Experimental.
	Module string `json:"module"`
	// The registry url to use when releasing packages.
	// Experimental.
	TwineRegistryUrl string `json:"twineRegistryUrl"`
}

func (j *JsiiPythonTarget) GetDistName() string {
	var returns string
	_jsii_.Get(
		j,
		"distName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiPythonTarget) GetModule() string {
	var returns string
	_jsii_.Get(
		j,
		"module",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsiiPythonTarget) GetTwineRegistryUrl() string {
	var returns string
	_jsii_.Get(
		j,
		"twineRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type JsonFileIface interface {
	IMarkableFileIface
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	GetMarker() bool
	GetOmitEmpty() bool
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddDeletionOverride(path string)
	AddOverride(path string, value interface{})
}

// Represents a JSON file.
// Experimental.
// Struct proxy
type JsonFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Indicates if the projen marker JSON-comment will be added to the output object.
	// Experimental.
	Marker bool `json:"marker"`
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (j *JsonFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		j,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JsonFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) GetPath() string {
	var returns string
	_jsii_.Get(
		j,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewJsonFile(project ProjectIface, filePath string, options JsonFileOptionsIface) JsonFileIface {
	_init_.Initialize()
	self := JsonFile{}
	_jsii_.Create(
		"projen.JsonFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (j *JsonFile) SetExecutable(val bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *JsonFile) SetReadonly(val bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func JsonFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.JsonFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) PostSynthesize() {
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

func (j *JsonFile) PreSynthesize() {
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

func (j *JsonFile) Synthesize() {
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

func (j *JsonFile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		j,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFile) AddDeletionOverride(path string) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addDeletionOverride",
		[]interface{}{path},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *JsonFile) AddOverride(path string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addOverride",
		[]interface{}{path, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// JsonFileOptionsIface is the public interface for the custom type JsonFileOptions
// Experimental.
type JsonFileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetMarker() bool
	GetObj() interface{}
	GetOmitEmpty() bool
}

// Options for `JsonFile`.
// Experimental.
// Struct proxy
type JsonFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker bool `json:"marker"`
	// The object that will be serialized.
	// 
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (j *JsonFileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		j,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		j,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFileOptions) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		j,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFileOptions) GetObj() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"obj",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JsonFileOptions) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		j,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type LicenseIface interface {
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_ IResolverIface) string
}

// Experimental.
// Struct proxy
type License struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (l *License) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		l,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (l *License) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		l,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (l *License) GetPath() string {
	var returns string
	_jsii_.Get(
		l,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (l *License) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		l,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (l *License) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		l,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewLicense(project ProjectIface, spdx string, options LicenseOptionsIface) LicenseIface {
	_init_.Initialize()
	self := License{}
	_jsii_.Create(
		"projen.License",
		[]interface{}{project, spdx, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (l *License) SetExecutable(val bool) {
	_jsii_.Set(
		l,
		"executable",
		val,
	)
}

func (l *License) SetReadonly(val bool) {
	_jsii_.Set(
		l,
		"readonly",
		val,
	)
}

func License_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.License",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (l *License) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *License) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *License) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *License) SynthesizeContent(_ IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		l,
		"synthesizeContent",
		[]interface{}{_},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// LicenseOptionsIface is the public interface for the custom type LicenseOptions
// Experimental.
type LicenseOptionsIface interface {
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
}

// Experimental.
// Struct proxy
type LicenseOptions struct {
	// Copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// Period of license (e.g. "1998-2023").
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
}

func (l *LicenseOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		l,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (l *LicenseOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		l,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Logging verbosity.
// Experimental.
type LogLevel string

const (
	LogLevelOff LogLevel = "OFF"
	LogLevelError LogLevel = "ERROR"
	LogLevelWarn LogLevel = "WARN"
	LogLevelInfo LogLevel = "INFO"
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelVerbose LogLevel = "VERBOSE"
)

// Class interface
type LoggerIface interface {
	GetProject() ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	Debug(text interface{})
	Error(text interface{})
	Info(text interface{})
	Log(level LogLevel, text interface{})
	Verbose(text interface{})
	Warn(text interface{})
}

// Project-level logging utilities.
// Experimental.
// Struct proxy
type Logger struct {
	// Experimental.
	Project ProjectIface `json:"project"`
}

func (l *Logger) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		l,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewLogger(project ProjectIface, options LoggerOptionsIface) LoggerIface {
	_init_.Initialize()
	self := Logger{}
	_jsii_.Create(
		"projen.Logger",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (l *Logger) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Debug(text interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"debug",
		[]interface{}{text},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Error(text interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"error",
		[]interface{}{text},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Info(text interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"info",
		[]interface{}{text},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Log(level LogLevel, text interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"log",
		[]interface{}{level, text},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Verbose(text interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"verbose",
		[]interface{}{text},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (l *Logger) Warn(text interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"warn",
		[]interface{}{text},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// LoggerOptionsIface is the public interface for the custom type LoggerOptions
// Experimental.
type LoggerOptionsIface interface {
	GetLevel() LogLevel
	GetUsePrefix() bool
}

// Options for logging utilities.
// Experimental.
// Struct proxy
type LoggerOptions struct {
	// The logging verbosity.
	// 
	// The levels available (in increasing verbosity) are
	// OFF, ERROR, WARN, INFO, DEBUG, and VERBOSE.
	// Experimental.
	Level LogLevel `json:"level"`
	// Include a prefix for all logging messages with the project name.
	// Experimental.
	UsePrefix bool `json:"usePrefix"`
}

func (l *LoggerOptions) GetLevel() LogLevel {
	var returns LogLevel
	_jsii_.Get(
		l,
		"level",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LogLevel)(nil)).Elem(): reflect.TypeOf((*LogLevel)(nil)).Elem(),
		},
	)
	return returns
}

func (l *LoggerOptions) GetUsePrefix() bool {
	var returns bool
	_jsii_.Get(
		l,
		"usePrefix",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type MakefileIface interface {
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	GetRules() []RuleIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddAll(target string) MakefileIface
	AddAlls(targets string) MakefileIface
	AddRule(rule RuleIface) MakefileIface
	AddRules(rules RuleIface) MakefileIface
}

// Minimal Makefile.
// Experimental.
// Struct proxy
type Makefile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
	// List of rule definitions.
	// Experimental.
	Rules []RuleIface `json:"rules"`
}

func (m *Makefile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		m,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (m *Makefile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		m,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *Makefile) GetPath() string {
	var returns string
	_jsii_.Get(
		m,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *Makefile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		m,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *Makefile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		m,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *Makefile) GetRules() []RuleIface {
	var returns []RuleIface
	_jsii_.Get(
		m,
		"rules",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RuleIface)(nil)).Elem(): reflect.TypeOf((*Rule)(nil)).Elem(),
		},
	)
	return returns
}


func NewMakefile(project ProjectIface, filePath string, options MakefileOptionsIface) MakefileIface {
	_init_.Initialize()
	self := Makefile{}
	_jsii_.Create(
		"projen.Makefile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (m *Makefile) SetExecutable(val bool) {
	_jsii_.Set(
		m,
		"executable",
		val,
	)
}

func (m *Makefile) SetReadonly(val bool) {
	_jsii_.Set(
		m,
		"readonly",
		val,
	)
}

func Makefile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.Makefile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *Makefile) PostSynthesize() {
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

func (m *Makefile) PreSynthesize() {
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

func (m *Makefile) Synthesize() {
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

func (m *Makefile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		m,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *Makefile) AddAll(target string) MakefileIface {
	var returns MakefileIface
	_jsii_.Invoke(
		m,
		"addAll",
		[]interface{}{target},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MakefileIface)(nil)).Elem(): reflect.TypeOf((*Makefile)(nil)).Elem(),
		},
	)
	return returns
}

func (m *Makefile) AddAlls(targets string) MakefileIface {
	var returns MakefileIface
	_jsii_.Invoke(
		m,
		"addAlls",
		[]interface{}{targets},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MakefileIface)(nil)).Elem(): reflect.TypeOf((*Makefile)(nil)).Elem(),
		},
	)
	return returns
}

func (m *Makefile) AddRule(rule RuleIface) MakefileIface {
	var returns MakefileIface
	_jsii_.Invoke(
		m,
		"addRule",
		[]interface{}{rule},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MakefileIface)(nil)).Elem(): reflect.TypeOf((*Makefile)(nil)).Elem(),
		},
	)
	return returns
}

func (m *Makefile) AddRules(rules RuleIface) MakefileIface {
	var returns MakefileIface
	_jsii_.Invoke(
		m,
		"addRules",
		[]interface{}{rules},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MakefileIface)(nil)).Elem(): reflect.TypeOf((*Makefile)(nil)).Elem(),
		},
	)
	return returns
}

// MakefileOptionsIface is the public interface for the custom type MakefileOptions
// Experimental.
type MakefileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetAll() []string
	GetRules() []RuleIface
}

// Options for Makefiles.
// Experimental.
// Struct proxy
type MakefileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// List of targets to build when Make is invoked without specifying any targets.
	// Experimental.
	All []string `json:"all"`
	// Rules to include in the Makefile.
	// Experimental.
	Rules []RuleIface `json:"rules"`
}

func (m *MakefileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		m,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MakefileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		m,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MakefileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		m,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MakefileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		m,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MakefileOptions) GetAll() []string {
	var returns []string
	_jsii_.Get(
		m,
		"all",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (m *MakefileOptions) GetRules() []RuleIface {
	var returns []RuleIface
	_jsii_.Get(
		m,
		"rules",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RuleIface)(nil)).Elem(): reflect.TypeOf((*Rule)(nil)).Elem(),
		},
	)
	return returns
}


// MarkableFileOptionsIface is the public interface for the custom type MarkableFileOptions
// Experimental.
type MarkableFileOptionsIface interface {
	GetMarker() bool
}

// Options for files that may include the Projen marker.
// Experimental.
// Struct proxy
type MarkableFileOptions struct {
	// Adds the projen marker to the file.
	// Experimental.
	Marker bool `json:"marker"`
}

func (m *MarkableFileOptions) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		m,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type NodePackageIface interface {
	GetProject() ProjectIface
	GetAllowLibraryDependencies() bool
	GetEntrypoint() string
	GetInstallCommand() string
	GetManifest() interface{}
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetProjenCommand() string
	GetLicense() string
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	AddBin(bins map[string]string)
	AddBundledDeps(deps string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddEngine(engine string, version string)
	AddField(name string, value interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddVersion(version string)
	HasScript(name string) bool
	RemoveScript(name string)
	SetScript(name string, command string)
}

// Represents the npm `package.json` file.
// Experimental.
// Struct proxy
type NodePackage struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// Allow project to take library dependencies.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// The module's entrypoint (e.g. `lib/index.js`).
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Returns the command to execute in order to install all dependencies (always frozen).
	// Experimental.
	InstallCommand string `json:"installCommand"`
	// Deprecated: use `addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// npm package access level.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// npm distribution tag.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The npm registry host (e.g. `registry.npmjs.org`).
	// Experimental.
	NpmRegistry string `json:"npmRegistry"`
	// npm registry (e.g. `https://registry.npmjs.org`). Use `npmRegistryHost` to get just the host name.
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm/pnpm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The package manager to use.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The name of the npm package.
	// Experimental.
	PackageName string `json:"packageName"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The SPDX license of this module.
	// 
	// `undefined` if this package is not licensed.
	// Experimental.
	License string `json:"license"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
}

func (n *NodePackage) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		n,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackage) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetInstallCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"installCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		n,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		n,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackage) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackage) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackage) GetPackageName() string {
	var returns string
	_jsii_.Get(
		n,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetLicense() string {
	var returns string
	_jsii_.Get(
		n,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewNodePackage(project ProjectIface, options NodePackageOptionsIface) NodePackageIface {
	_init_.Initialize()
	self := NodePackage{}
	_jsii_.Create(
		"projen.NodePackage",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (n *NodePackage) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddBin(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBin",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddEngine(engine string, version string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addEngine",
		[]interface{}{engine, version},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddField(name string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addField",
		[]interface{}{name, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) AddVersion(version string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addVersion",
		[]interface{}{version},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackage) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodePackage) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// The node package manager to use.
// Experimental.
type NodePackageManager string

const (
	NodePackageManagerYarn NodePackageManager = "YARN"
	NodePackageManagerNpm NodePackageManager = "NPM"
	NodePackageManagerPnpm NodePackageManager = "PNPM"
)

// NodePackageOptionsIface is the public interface for the custom type NodePackageOptions
// Experimental.
type NodePackageOptionsIface interface {
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
}

// Experimental.
// Struct proxy
type NodePackageOptions struct {
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
}

func (n *NodePackageOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		n,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		n,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		n,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		n,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		n,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		n,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		n,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		n,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		n,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		n,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		n,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		n,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		n,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		n,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodePackageOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodePackageOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		n,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type NodeProjectIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// Node.js project.
// Experimental.
// Struct proxy
type NodeProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
}

func (n *NodeProject) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		n,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		n,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		n,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		n,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetName() string {
	var returns string
	_jsii_.Get(
		n,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		n,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		n,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		n,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		n,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		n,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		n,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		n,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		n,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		n,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		n,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		n,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		n,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		n,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		n,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		n,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		n,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		n,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		n,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		n,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		n,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewNodeProject(options NodeProjectOptionsIface) NodeProjectIface {
	_init_.Initialize()
	self := NodeProject{}
	_jsii_.Create(
		"projen.NodeProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func NodeProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.NodeProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		n,
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

func (n *NodeProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		n,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		n,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		n,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *NodeProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		n,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// NodeProjectOptionsIface is the public interface for the custom type NodeProjectOptions
// Experimental.
type NodeProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
}

// Experimental.
// Struct proxy
type NodeProjectOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
}

func (n *NodeProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		n,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		n,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		n,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		n,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		n,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		n,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		n,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		n,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		n,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		n,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		n,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		n,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		n,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		n,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		n,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		n,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		n,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		n,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		n,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		n,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		n,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		n,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		n,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		n,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		n,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		n,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		n,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		n,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		n,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		n,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		n,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		n,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		n,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		n,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		n,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		n,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		n,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		n,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		n,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		n,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		n,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		n,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		n,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		n,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		n,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		n,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		n,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		n,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		n,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		n,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		n,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		n,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		n,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		n,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		n,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		n,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		n,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		n,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		n,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		n,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		n,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		n,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *NodeProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		n,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// NodeWorkflowStepsIface is the public interface for the custom type NodeWorkflowSteps
// Experimental.
type NodeWorkflowStepsIface interface {
	GetAntitamper() []interface{}
	GetInstall() []interface{}
}

// Experimental.
// Struct proxy
type NodeWorkflowSteps struct {
	// Experimental.
	Antitamper []interface{} `json:"antitamper"`
	// Experimental.
	Install []interface{} `json:"install"`
}

func (n *NodeWorkflowSteps) GetAntitamper() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (n *NodeWorkflowSteps) GetInstall() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		n,
		"install",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}


// Npm package access level.
// Experimental.
type NpmAccess string

const (
	NpmAccessPublic NpmAccess = "PUBLIC"
	NpmAccessRestricted NpmAccess = "RESTRICTED"
)

// Experimental.
type NpmTaskExecution string

const (
	NpmTaskExecutionProjen NpmTaskExecution = "PROJEN"
	NpmTaskExecutionShell NpmTaskExecution = "SHELL"
)

// Class interface
type ObjectFileIface interface {
	IMarkableFileIface
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	GetMarker() bool
	GetOmitEmpty() bool
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddDeletionOverride(path string)
	AddOverride(path string, value interface{})
}

// Represents an Object file.
// Experimental.
// Struct proxy
type ObjectFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Indicates if the projen marker JSON-comment will be added to the output object.
	// Experimental.
	Marker bool `json:"marker"`
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (o *ObjectFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		o,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (o *ObjectFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		o,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) GetPath() string {
	var returns string
	_jsii_.Get(
		o,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		o,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		o,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		o,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		o,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewObjectFile(project ProjectIface, filePath string, options ObjectFileOptionsIface) ObjectFileIface {
	_init_.Initialize()
	self := ObjectFile{}
	_jsii_.Create(
		"projen.ObjectFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{"projen.IMarkableFile"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (o *ObjectFile) SetExecutable(val bool) {
	_jsii_.Set(
		o,
		"executable",
		val,
	)
}

func (o *ObjectFile) SetReadonly(val bool) {
	_jsii_.Set(
		o,
		"readonly",
		val,
	)
}

func ObjectFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.ObjectFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		o,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (o *ObjectFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		o,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (o *ObjectFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		o,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (o *ObjectFile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		o,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFile) AddDeletionOverride(path string) {
	var returns interface{}
	_jsii_.Invoke(
		o,
		"addDeletionOverride",
		[]interface{}{path},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (o *ObjectFile) AddOverride(path string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		o,
		"addOverride",
		[]interface{}{path, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ObjectFileOptionsIface is the public interface for the custom type ObjectFileOptions
// Experimental.
type ObjectFileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetMarker() bool
	GetObj() interface{}
	GetOmitEmpty() bool
}

// Options for `ObjectFile`.
// Experimental.
// Struct proxy
type ObjectFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker bool `json:"marker"`
	// The object that will be serialized.
	// 
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (o *ObjectFileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		o,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		o,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		o,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		o,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFileOptions) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		o,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFileOptions) GetObj() interface{} {
	var returns interface{}
	_jsii_.Get(
		o,
		"obj",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (o *ObjectFileOptions) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		o,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// PeerDependencyOptionsIface is the public interface for the custom type PeerDependencyOptions
// Experimental.
type PeerDependencyOptionsIface interface {
	GetPinnedDevDependency() bool
}

// Experimental.
// Struct proxy
type PeerDependencyOptions struct {
	// Automatically add a pinned dev dependency.
	// Experimental.
	PinnedDevDependency bool `json:"pinnedDevDependency"`
}

func (p *PeerDependencyOptions) GetPinnedDevDependency() bool {
	var returns bool
	_jsii_.Get(
		p,
		"pinnedDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ProjectIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
}

// Base project.
// Experimental.
// Struct proxy
type Project struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root *ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent *ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
}

func (p *Project) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		p,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		p,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		p,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		p,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		p,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetName() string {
	var returns string
	_jsii_.Get(
		p,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Project) GetOutdir() string {
	var returns string
	_jsii_.Get(
		p,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Project) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		p,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		p,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		p,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		p,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		p,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		p,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		p,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Project) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		p,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		p,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}


func NewProject(options ProjectOptionsIface) ProjectIface {
	_init_.Initialize()
	self := Project{}
	_jsii_.Create(
		"projen.Project",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func Project_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.Project",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Project) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Project) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		p,
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

func (p *Project) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Project) PostSynthesize() {
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

func (p *Project) PreSynthesize() {
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

func (p *Project) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Project) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		p,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		p,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Project) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		p,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

// ProjectOptionsIface is the public interface for the custom type ProjectOptions
// Experimental.
type ProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
}

// Experimental.
// Struct proxy
type ProjectOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
}

func (p *ProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		p,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		p,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		p,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		p,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		p,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjectOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		p,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (p *ProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		p,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProjectOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		p,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (p *ProjectOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		p,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (p *ProjectOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		p,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}


// Which type of project this is.
// Experimental.
type ProjectType string

const (
	ProjectTypeUnknown ProjectType = "UNKNOWN"
	ProjectTypeLib ProjectType = "LIB"
	ProjectTypeApp ProjectType = "APP"
)

// ResolveOptionsIface is the public interface for the custom type ResolveOptions
// Experimental.
type ResolveOptionsIface interface {
	GetArgs() []interface{}
	GetOmitEmpty() bool
}

// Resolve options.
// Experimental.
// Struct proxy
type ResolveOptions struct {
	// Context arguments.
	// Experimental.
	Args []interface{} `json:"args"`
	// Omits empty arrays and objects.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (r *ResolveOptions) GetArgs() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		r,
		"args",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (r *ResolveOptions) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		r,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// RuleIface is the public interface for the custom type Rule
// Experimental.
type RuleIface interface {
	GetTargets() []string
	GetPhony() bool
	GetPrerequisites() []string
	GetRecipe() []string
}

// A Make rule.
// Experimental.
// Struct proxy
type Rule struct {
	// Files to be created or updated by this rule.
	// 
	// If the rule is phony then instead this represents the command's name(s).
	// Experimental.
	Targets []string `json:"targets"`
	// Marks whether the target is phony.
	// Experimental.
	Phony bool `json:"phony"`
	// Files that are used as inputs to create a target.
	// Experimental.
	Prerequisites []string `json:"prerequisites"`
	// Commands that are run (using prerequisites as inputs) to create a target.
	// Experimental.
	Recipe []string `json:"recipe"`
}

func (r *Rule) GetTargets() []string {
	var returns []string
	_jsii_.Get(
		r,
		"targets",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *Rule) GetPhony() bool {
	var returns bool
	_jsii_.Get(
		r,
		"phony",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (r *Rule) GetPrerequisites() []string {
	var returns []string
	_jsii_.Get(
		r,
		"prerequisites",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *Rule) GetRecipe() []string {
	var returns []string
	_jsii_.Get(
		r,
		"recipe",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type SampleDirIface interface {
	GetProject() ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Renders the given files into the directory if the directory does not exist.
// 
// Use this to create sample code files
// Experimental.
// Struct proxy
type SampleDir struct {
	// Experimental.
	Project ProjectIface `json:"project"`
}

func (s *SampleDir) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		s,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}


// Create sample files in the given directory if the given directory does not exist.
func NewSampleDir(project ProjectIface, dir string, options SampleDirOptionsIface) SampleDirIface {
	_init_.Initialize()
	self := SampleDir{}
	_jsii_.Create(
		"projen.SampleDir",
		[]interface{}{project, dir, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (s *SampleDir) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *SampleDir) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *SampleDir) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// SampleDirOptionsIface is the public interface for the custom type SampleDirOptions
// Experimental.
type SampleDirOptionsIface interface {
	GetFiles() map[string]string
}

// SampleDir options.
// Experimental.
// Struct proxy
type SampleDirOptions struct {
	// The files to render into the directory.
	// Experimental.
	Files map[string]string `json:"files"`
}

func (s *SampleDirOptions) GetFiles() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		s,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type SampleFileIface interface {
	GetProject() ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Produces a file with the given contents but only once, if the file doesn't already exist.
// 
// Use this for creating example code files or other resources.
// Experimental.
// Struct proxy
type SampleFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
}

func (s *SampleFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		s,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}


// Creates a new SampleFile object.
func NewSampleFile(project ProjectIface, filePath string, options SampleFileOptionsIface) SampleFileIface {
	_init_.Initialize()
	self := SampleFile{}
	_jsii_.Create(
		"projen.SampleFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (s *SampleFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *SampleFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *SampleFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// SampleFileOptionsIface is the public interface for the custom type SampleFileOptions
// Experimental.
type SampleFileOptionsIface interface {
	GetContents() string
}

// Options for the SampleFile object.
// Experimental.
// Struct proxy
type SampleFileOptions struct {
	// The contents of the file to write.
	// Experimental.
	Contents string `json:"contents"`
}

func (s *SampleFileOptions) GetContents() string {
	var returns string
	_jsii_.Get(
		s,
		"contents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type SampleReadmeIface interface {
	GetProject() ProjectIface
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Represents a README.md sample file. You are expected to manage this file after creation.
// Experimental.
// Struct proxy
type SampleReadme struct {
	// Experimental.
	Project ProjectIface `json:"project"`
}

func (s *SampleReadme) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		s,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}


func NewSampleReadme(project ProjectIface, props SampleReadmePropsIface) SampleReadmeIface {
	_init_.Initialize()
	self := SampleReadme{}
	_jsii_.Create(
		"projen.SampleReadme",
		[]interface{}{project, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (s *SampleReadme) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *SampleReadme) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *SampleReadme) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// SampleReadmePropsIface is the public interface for the custom type SampleReadmeProps
// Experimental.
type SampleReadmePropsIface interface {
	GetContents() string
	GetFilename() string
}

// SampleReadme Properties.
// Experimental.
// Struct proxy
type SampleReadmeProps struct {
	// The contents.
	// Experimental.
	Contents string `json:"contents"`
	// The name of the README.md file.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Filename string `json:"filename"`
}

func (s *SampleReadmeProps) GetContents() string {
	var returns string
	_jsii_.Get(
		s,
		"contents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *SampleReadmeProps) GetFilename() string {
	var returns string
	_jsii_.Get(
		s,
		"filename",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type SemverIface interface {
	GetSpec() string
	GetMode() string
	GetVersion() string
}

// Deprecated: This class will be removed in upcoming releases. if you wish to
specify semver requirements in `deps`, `devDeps`, etc, specify them like so
`express@^2.1`.
// Struct proxy
type Semver struct {
	// Deprecated: This class will be removed in upcoming releases. if you wish to
specify semver requirements in `deps`, `devDeps`, etc, specify them like so
`express@^2.1`.
	Spec string `json:"spec"`
	// Deprecated: This class will be removed in upcoming releases. if you wish to
specify semver requirements in `deps`, `devDeps`, etc, specify them like so
`express@^2.1`.
	Mode string `json:"mode"`
	// Deprecated: This class will be removed in upcoming releases. if you wish to
specify semver requirements in `deps`, `devDeps`, etc, specify them like so
`express@^2.1`.
	Version string `json:"version"`
}

func (s *Semver) GetSpec() string {
	var returns string
	_jsii_.Get(
		s,
		"spec",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *Semver) GetMode() string {
	var returns string
	_jsii_.Get(
		s,
		"mode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *Semver) GetVersion() string {
	var returns string
	_jsii_.Get(
		s,
		"version",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func Semver_Caret(version string) SemverIface {
	_init_.Initialize()
	var returns SemverIface
	_jsii_.InvokeStatic(
		"projen.Semver",
		"caret",
		[]interface{}{version},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func Semver_Latest() SemverIface {
	_init_.Initialize()
	var returns SemverIface
	_jsii_.InvokeStatic(
		"projen.Semver",
		"latest",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func Semver_Of(spec string) SemverIface {
	_init_.Initialize()
	var returns SemverIface
	_jsii_.InvokeStatic(
		"projen.Semver",
		"of",
		[]interface{}{spec},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func Semver_Pinned(version string) SemverIface {
	_init_.Initialize()
	var returns SemverIface
	_jsii_.InvokeStatic(
		"projen.Semver",
		"pinned",
		[]interface{}{version},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func Semver_Tilde(version string) SemverIface {
	_init_.Initialize()
	var returns SemverIface
	_jsii_.InvokeStatic(
		"projen.Semver",
		"tilde",
		[]interface{}{version},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

// Experimental.
type Stability string

const (
	StabilityExperimental Stability = "EXPERIMENTAL"
	StabilityStable Stability = "STABLE"
	StabilityDeprecated Stability = "DEPRECATED"
)

// Class interface
type TextFileIface interface {
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_ IResolverIface) string
	AddLine(line string)
}

// A text file.
// Experimental.
// Struct proxy
type TextFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
}

func (t *TextFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TextFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		t,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFile) GetPath() string {
	var returns string
	_jsii_.Get(
		t,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		t,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		t,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Defines a text file.
func NewTextFile(project ProjectIface, filePath string, options TextFileOptionsIface) TextFileIface {
	_init_.Initialize()
	self := TextFile{}
	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (t *TextFile) SetExecutable(val bool) {
	_jsii_.Set(
		t,
		"executable",
		val,
	)
}

func (t *TextFile) SetReadonly(val bool) {
	_jsii_.Set(
		t,
		"readonly",
		val,
	)
}

func TextFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.TextFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TextFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TextFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TextFile) SynthesizeContent(_ IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		t,
		"synthesizeContent",
		[]interface{}{_},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFile) AddLine(line string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addLine",
		[]interface{}{line},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// TextFileOptionsIface is the public interface for the custom type TextFileOptions
// Experimental.
type TextFileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetLines() []string
}

// Options for `TextFile`.
// Experimental.
// Struct proxy
type TextFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// The contents of the text file.
	// 
	// You can use `addLine()` to append lines.
	// Experimental.
	Lines []string `json:"lines"`
}

func (t *TextFileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		t,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		t,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		t,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		t,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TextFileOptions) GetLines() []string {
	var returns []string
	_jsii_.Get(
		t,
		"lines",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type TomlFileIface interface {
	IMarkableFileIface
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	GetMarker() bool
	GetOmitEmpty() bool
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddDeletionOverride(path string)
	AddOverride(path string, value interface{})
}

// Represents a TOML file.
// Experimental.
// Struct proxy
type TomlFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Indicates if the projen marker JSON-comment will be added to the output object.
	// Experimental.
	Marker bool `json:"marker"`
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (t *TomlFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TomlFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		t,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) GetPath() string {
	var returns string
	_jsii_.Get(
		t,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		t,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		t,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		t,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		t,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewTomlFile(project ProjectIface, filePath string, options TomlFileOptionsIface) TomlFileIface {
	_init_.Initialize()
	self := TomlFile{}
	_jsii_.Create(
		"projen.TomlFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (t *TomlFile) SetExecutable(val bool) {
	_jsii_.Set(
		t,
		"executable",
		val,
	)
}

func (t *TomlFile) SetReadonly(val bool) {
	_jsii_.Set(
		t,
		"readonly",
		val,
	)
}

func TomlFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.TomlFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TomlFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TomlFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TomlFile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		t,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFile) AddDeletionOverride(path string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDeletionOverride",
		[]interface{}{path},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TomlFile) AddOverride(path string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addOverride",
		[]interface{}{path, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// TomlFileOptionsIface is the public interface for the custom type TomlFileOptions
// Experimental.
type TomlFileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetMarker() bool
	GetObj() interface{}
	GetOmitEmpty() bool
}

// Options for `TomlFile`.
// Experimental.
// Struct proxy
type TomlFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker bool `json:"marker"`
	// The object that will be serialized.
	// 
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (t *TomlFileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		t,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		t,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		t,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		t,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFileOptions) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		t,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFileOptions) GetObj() interface{} {
	var returns interface{}
	_jsii_.Get(
		t,
		"obj",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TomlFileOptions) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		t,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type TypeScriptAppProjectIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// TypeScript app.
// Experimental.
// Struct proxy
type TypeScriptAppProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
}

func (t *TypeScriptAppProject) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		t,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		t,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		t,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		t,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		t,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		t,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		t,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		t,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		t,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		t,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		t,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		t,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		t,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		t,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		t,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		t,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		t,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		t,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		t,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		t,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		t,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		t,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		t,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		t,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		t,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		t,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		t,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetLibdir() string {
	var returns string
	_jsii_.Get(
		t,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		t,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetTestdir() string {
	var returns string
	_jsii_.Get(
		t,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		t,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		t,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		t,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}


func NewTypeScriptAppProject(options TypeScriptProjectOptionsIface) TypeScriptAppProjectIface {
	_init_.Initialize()
	self := TypeScriptAppProject{}
	_jsii_.Create(
		"projen.TypeScriptAppProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func TypeScriptAppProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.TypeScriptAppProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		t,
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

func (t *TypeScriptAppProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		t,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		t,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		t,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptAppProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		t,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptAppProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		t,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptAppProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// TypeScriptCompilerOptionsIface is the public interface for the custom type TypeScriptCompilerOptions
// Experimental.
type TypeScriptCompilerOptionsIface interface {
	GetAllowJs() bool
	GetAllowSyntheticDefaultImports() bool
	GetAlwaysStrict() bool
	GetDeclaration() bool
	GetDeclarationDir() string
	GetEsModuleInterop() bool
	GetExperimentalDecorators() bool
	GetForceConsistentCasingInFileNames() bool
	GetInlineSourceMap() bool
	GetInlineSources() bool
	GetIsolatedModules() bool
	GetJsx() TypeScriptJsxMode
	GetLib() []string
	GetModule() string
	GetModuleResolution() TypeScriptModuleResolution
	GetNoEmit() bool
	GetNoEmitOnError() bool
	GetNoFallthroughCasesInSwitch() bool
	GetNoImplicitAny() bool
	GetNoImplicitReturns() bool
	GetNoImplicitThis() bool
	GetNoUnusedLocals() bool
	GetNoUnusedParameters() bool
	GetOutDir() string
	GetResolveJsonModule() bool
	GetRootDir() string
	GetSkipLibCheck() bool
	GetStrict() bool
	GetStrictNullChecks() bool
	GetStrictPropertyInitialization() bool
	GetStripInternal() bool
	GetTarget() string
}

// Experimental.
// Struct proxy
type TypeScriptCompilerOptions struct {
	// Allow JavaScript files to be compiled.
	// Experimental.
	AllowJs bool `json:"allowJs"`
	// Allow default imports from modules with no default export.
	// 
	// This does not affect code emit, just typechecking.
	// Experimental.
	AllowSyntheticDefaultImports bool `json:"allowSyntheticDefaultImports"`
	// Ensures that your files are parsed in the ECMAScript strict mode, and emit “use strict” for each source file.
	// Experimental.
	AlwaysStrict bool `json:"alwaysStrict"`
	// To be specified along with the above.
	// Experimental.
	Declaration bool `json:"declaration"`
	// Offers a way to configure the root directory for where declaration files are emitted.
	// Experimental.
	DeclarationDir string `json:"declarationDir"`
	// Emit __importStar and __importDefault helpers for runtime babel ecosystem compatibility and enable --allowSyntheticDefaultImports for typesystem compatibility.
	// Experimental.
	EsModuleInterop bool `json:"esModuleInterop"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	// Experimental.
	ExperimentalDecorators bool `json:"experimentalDecorators"`
	// Disallow inconsistently-cased references to the same file.
	// Experimental.
	ForceConsistentCasingInFileNames bool `json:"forceConsistentCasingInFileNames"`
	// When set, instead of writing out a .js.map file to provide source maps, TypeScript will embed the source map content in the .js files.
	// Experimental.
	InlineSourceMap bool `json:"inlineSourceMap"`
	// When set, TypeScript will include the original content of the .ts file as an embedded string in the source map. This is often useful in the same cases as inlineSourceMap.
	// Experimental.
	InlineSources bool `json:"inlineSources"`
	// Perform additional checks to ensure that separate compilation (such as with transpileModule or @babel/plugin-transform-typescript) would be safe.
	// Experimental.
	IsolatedModules bool `json:"isolatedModules"`
	// Support JSX in .tsx files: "react", "preserve", "react-native".
	// Experimental.
	Jsx TypeScriptJsxMode `json:"jsx"`
	// Reference for type definitions / libraries to use (eg.
	// 
	// ES2016, ES5, ES2018).
	// Experimental.
	Lib []string `json:"lib"`
	// Sets the module system for the program.
	// 
	// See https://www.typescriptlang.org/docs/handbook/modules.html#ambient-modules.
	// Experimental.
	Module string `json:"module"`
	// Determine how modules get resolved.
	// 
	// Either "Node" for Node.js/io.js style resolution, or "Classic".
	// Experimental.
	ModuleResolution TypeScriptModuleResolution `json:"moduleResolution"`
	// Do not emit outputs.
	// Experimental.
	NoEmit bool `json:"noEmit"`
	// Do not emit compiler output files like JavaScript source code, source-maps or declarations if any errors were reported.
	// Experimental.
	NoEmitOnError bool `json:"noEmitOnError"`
	// Report errors for fallthrough cases in switch statements.
	// 
	// Ensures that any non-empty
	// case inside a switch statement includes either break or return. This means you won’t
	// accidentally ship a case fallthrough bug.
	// Experimental.
	NoFallthroughCasesInSwitch bool `json:"noFallthroughCasesInSwitch"`
	// In some cases where no type annotations are present, TypeScript will fall back to a type of any for a variable when it cannot infer the type.
	// Experimental.
	NoImplicitAny bool `json:"noImplicitAny"`
	// When enabled, TypeScript will check all code paths in a function to ensure they return a value.
	// Experimental.
	NoImplicitReturns bool `json:"noImplicitReturns"`
	// Raise error on ‘this’ expressions with an implied ‘any’ type.
	// Experimental.
	NoImplicitThis bool `json:"noImplicitThis"`
	// Report errors on unused local variables.
	// Experimental.
	NoUnusedLocals bool `json:"noUnusedLocals"`
	// Report errors on unused parameters in functions.
	// Experimental.
	NoUnusedParameters bool `json:"noUnusedParameters"`
	// Output directory for the compiled files.
	// Experimental.
	OutDir string `json:"outDir"`
	// Allows importing modules with a ‘.json’ extension, which is a common practice in node projects. This includes generating a type for the import based on the static JSON shape.
	// Experimental.
	ResolveJsonModule bool `json:"resolveJsonModule"`
	// Specifies the root directory of input files.
	// 
	// Only use to control the output directory structure with `outDir`.
	// Experimental.
	RootDir string `json:"rootDir"`
	// Skip type checking of all declaration files (*.d.ts).
	// Experimental.
	SkipLibCheck bool `json:"skipLibCheck"`
	// The strict flag enables a wide range of type checking behavior that results in stronger guarantees of program correctness.
	// 
	// Turning this on is equivalent to enabling all of the strict mode family
	// options, which are outlined below. You can then turn off individual strict mode family checks as
	// needed.
	// Experimental.
	Strict bool `json:"strict"`
	// When strictNullChecks is false, null and undefined are effectively ignored by the language.
	// 
	// This can lead to unexpected errors at runtime.
	// When strictNullChecks is true, null and undefined have their own distinct types and you’ll
	// get a type error if you try to use them where a concrete value is expected.
	// Experimental.
	StrictNullChecks bool `json:"strictNullChecks"`
	// When set to true, TypeScript will raise an error when a class property was declared but not set in the constructor.
	// Experimental.
	StrictPropertyInitialization bool `json:"strictPropertyInitialization"`
	// Do not emit declarations for code that has an @internal annotation in it’s JSDoc comment.
	// Experimental.
	StripInternal bool `json:"stripInternal"`
	// Modern browsers support all ES6 features, so ES6 is a good choice.
	// 
	// You might choose to set
	// a lower target if your code is deployed to older environments, or a higher target if your
	// code is guaranteed to run in newer environments.
	// Experimental.
	Target string `json:"target"`
}

func (t *TypeScriptCompilerOptions) GetAllowJs() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowJs",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetAllowSyntheticDefaultImports() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowSyntheticDefaultImports",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetAlwaysStrict() bool {
	var returns bool
	_jsii_.Get(
		t,
		"alwaysStrict",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetDeclaration() bool {
	var returns bool
	_jsii_.Get(
		t,
		"declaration",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetDeclarationDir() string {
	var returns string
	_jsii_.Get(
		t,
		"declarationDir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetEsModuleInterop() bool {
	var returns bool
	_jsii_.Get(
		t,
		"esModuleInterop",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetExperimentalDecorators() bool {
	var returns bool
	_jsii_.Get(
		t,
		"experimentalDecorators",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetForceConsistentCasingInFileNames() bool {
	var returns bool
	_jsii_.Get(
		t,
		"forceConsistentCasingInFileNames",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetInlineSourceMap() bool {
	var returns bool
	_jsii_.Get(
		t,
		"inlineSourceMap",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetInlineSources() bool {
	var returns bool
	_jsii_.Get(
		t,
		"inlineSources",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetIsolatedModules() bool {
	var returns bool
	_jsii_.Get(
		t,
		"isolatedModules",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetJsx() TypeScriptJsxMode {
	var returns TypeScriptJsxMode
	_jsii_.Get(
		t,
		"jsx",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypeScriptJsxMode)(nil)).Elem(): reflect.TypeOf((*TypeScriptJsxMode)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetLib() []string {
	var returns []string
	_jsii_.Get(
		t,
		"lib",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetModule() string {
	var returns string
	_jsii_.Get(
		t,
		"module",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetModuleResolution() TypeScriptModuleResolution {
	var returns TypeScriptModuleResolution
	_jsii_.Get(
		t,
		"moduleResolution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypeScriptModuleResolution)(nil)).Elem(): reflect.TypeOf((*TypeScriptModuleResolution)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoEmit() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noEmit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoEmitOnError() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noEmitOnError",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoFallthroughCasesInSwitch() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noFallthroughCasesInSwitch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoImplicitAny() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noImplicitAny",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoImplicitReturns() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noImplicitReturns",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoImplicitThis() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noImplicitThis",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoUnusedLocals() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noUnusedLocals",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetNoUnusedParameters() bool {
	var returns bool
	_jsii_.Get(
		t,
		"noUnusedParameters",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetOutDir() string {
	var returns string
	_jsii_.Get(
		t,
		"outDir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetResolveJsonModule() bool {
	var returns bool
	_jsii_.Get(
		t,
		"resolveJsonModule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetRootDir() string {
	var returns string
	_jsii_.Get(
		t,
		"rootDir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetSkipLibCheck() bool {
	var returns bool
	_jsii_.Get(
		t,
		"skipLibCheck",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetStrict() bool {
	var returns bool
	_jsii_.Get(
		t,
		"strict",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetStrictNullChecks() bool {
	var returns bool
	_jsii_.Get(
		t,
		"strictNullChecks",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetStrictPropertyInitialization() bool {
	var returns bool
	_jsii_.Get(
		t,
		"strictPropertyInitialization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetStripInternal() bool {
	var returns bool
	_jsii_.Get(
		t,
		"stripInternal",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptCompilerOptions) GetTarget() string {
	var returns string
	_jsii_.Get(
		t,
		"target",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Determines how JSX should get transformed into valid JavaScript.
// See: https://www.typescriptlang.org/docs/handbook/jsx.html
//
// Experimental.
type TypeScriptJsxMode string

const (
	TypeScriptJsxModePreserve TypeScriptJsxMode = "PRESERVE"
	TypeScriptJsxModeReact TypeScriptJsxMode = "REACT"
	TypeScriptJsxModeReactNative TypeScriptJsxMode = "REACT_NATIVE"
)

// Class interface
type TypeScriptLibraryProjectIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// Deprecated: use `TypeScriptProject`
// Struct proxy
type TypeScriptLibraryProject struct {
	// Returns all the components within this project.
	// Deprecated: use `TypeScriptProject`
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Deprecated: use `TypeScriptProject`
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Deprecated: use `TypeScriptProject`
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Deprecated: use `TypeScriptProject`
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Deprecated: use `TypeScriptProject`
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Deprecated: use `TypeScriptProject`
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Deprecated: use `TypeScriptProject`
	Outdir string `json:"outdir"`
	// Deprecated: use `TypeScriptProject`
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Deprecated: use `TypeScriptProject`
	Root ProjectIface `json:"root"`
	// Deprecated: use `TypeScriptProject`
	Tasks tasks.TasksIface `json:"tasks"`
	// Access for .devcontainer.json (used for GitHub Codespaces).
	// 
	// This will be `undefined` if devContainer boolean is false
	// Deprecated: use `TypeScriptProject`
	DevContainer vscode.DevContainerIface `json:"devContainer"`
	// Access all github components.
	// 
	// This will be `undefined` for subprojects.
	// Deprecated: use `TypeScriptProject`
	Github github.GitHubIface `json:"github"`
	// Access for Gitpod.
	// 
	// This will be `undefined` if gitpod boolean is false
	// Deprecated: use `TypeScriptProject`
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Deprecated: use `TypeScriptProject`
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Deprecated: use `TypeScriptProject`
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Deprecated: use `TypeScriptProject`
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Deprecated: use `TypeScriptProject`
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Deprecated: use `TypeScriptProject`
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Deprecated: use `TypeScriptProject`
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Deprecated: use `TypeScriptProject`
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Deprecated: use `TypeScriptProject`
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Deprecated: use `TypeScriptProject`
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Deprecated: use `TypeScriptProject`
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Deprecated: use `TypeScriptProject`
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Deprecated: use `TypeScriptProject`
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Deprecated: use `TypeScriptProject`
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Deprecated: use `TypeScriptProject`
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Deprecated: use `TypeScriptProject`
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Deprecated: use `TypeScriptProject`
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Deprecated: use `TypeScriptProject`
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Deprecated: use `TypeScriptProject`
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Deprecated: use `TypeScriptProject`
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Deprecated: use `TypeScriptProject`
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Deprecated: use `TypeScriptProject`
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Deprecated: use `TypeScriptProject`
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Deprecated: use `TypeScriptProject`
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Deprecated: use `TypeScriptProject`
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Deprecated: use `TypeScriptProject`
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Deprecated: use `TypeScriptProject`
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Deprecated: use `TypeScriptProject`
	Docgen bool `json:"docgen"`
	// Deprecated: use `TypeScriptProject`
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Deprecated: use `TypeScriptProject`
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Deprecated: use `TypeScriptProject`
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
}

func (t *TypeScriptLibraryProject) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		t,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		t,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		t,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		t,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		t,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		t,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		t,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		t,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		t,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		t,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		t,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		t,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		t,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		t,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		t,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		t,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		t,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		t,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		t,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		t,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		t,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		t,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		t,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		t,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		t,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		t,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		t,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetLibdir() string {
	var returns string
	_jsii_.Get(
		t,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		t,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetTestdir() string {
	var returns string
	_jsii_.Get(
		t,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		t,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		t,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		t,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}


func NewTypeScriptLibraryProject(options TypeScriptProjectOptionsIface) TypeScriptLibraryProjectIface {
	_init_.Initialize()
	self := TypeScriptLibraryProject{}
	_jsii_.Create(
		"projen.TypeScriptLibraryProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func TypeScriptLibraryProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.TypeScriptLibraryProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		t,
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

func (t *TypeScriptLibraryProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		t,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		t,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		t,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		t,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptLibraryProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		t,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// TypeScriptLibraryProjectOptionsIface is the public interface for the custom type TypeScriptLibraryProjectOptions
// Deprecated: use TypeScriptProjectOptions
type TypeScriptLibraryProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetCompileBeforeTest() bool
	GetDisableTsconfig() bool
	GetDocgen() bool
	GetDocsDirectory() string
	GetEntrypointTypes() string
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetLibdir() string
	GetPackage() bool
	GetSampleCode() bool
	GetSrcdir() string
	GetTestdir() string
	GetTsconfig() TypescriptConfigOptionsIface
	GetTypescriptVersion() string
}

// Deprecated: use TypeScriptProjectOptions
// Struct proxy
type TypeScriptLibraryProjectOptions struct {
	// This is the name of your project.
	// Deprecated: use TypeScriptProjectOptions
	Name string `json:"name"`
	// Add a `clobber` task which resets the repo to origin.
	// Deprecated: use TypeScriptProjectOptions
	Clobber bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Deprecated: use TypeScriptProjectOptions
	DevContainer bool `json:"devContainer"`
	// Add a Gitpod development environment.
	// Deprecated: use TypeScriptProjectOptions
	Gitpod bool `json:"gitpod"`
	// The JSII FQN (fully qualified name) of the project class.
	// Deprecated: use TypeScriptProjectOptions
	JsiiFqn string `json:"jsiiFqn"`
	// Configure logging options such as verbosity.
	// Deprecated: use TypeScriptProjectOptions
	Logging LoggerOptionsIface `json:"logging"`
	// The root directory of the project.
	// 
	// Relative to this directory, all files are synthesized.
	// 
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Deprecated: use TypeScriptProjectOptions
	Outdir string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Deprecated: use TypeScriptProjectOptions
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Deprecated: use TypeScriptProjectOptions
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Deprecated: use TypeScriptProjectOptions
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Deprecated: use TypeScriptProjectOptions
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Deprecated: use TypeScriptProjectOptions
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Deprecated: use TypeScriptProjectOptions
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Deprecated: use TypeScriptProjectOptions
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Deprecated: use TypeScriptProjectOptions
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Deprecated: use TypeScriptProjectOptions
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Deprecated: use TypeScriptProjectOptions
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Deprecated: use TypeScriptProjectOptions
	BundledDeps []string `json:"bundledDeps"`
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
	// Deprecated: use TypeScriptProjectOptions
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Deprecated: use TypeScriptProjectOptions
	Description string `json:"description"`
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
	// Deprecated: use TypeScriptProjectOptions
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Deprecated: use TypeScriptProjectOptions
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Deprecated: use TypeScriptProjectOptions
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Deprecated: use TypeScriptProjectOptions
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Deprecated: use TypeScriptProjectOptions
	License string `json:"license"`
	// Indicates if a license should be added.
	// Deprecated: use TypeScriptProjectOptions
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Deprecated: use TypeScriptProjectOptions
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Deprecated: use TypeScriptProjectOptions
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Deprecated: use TypeScriptProjectOptions
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Deprecated: use TypeScriptProjectOptions
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Deprecated: use TypeScriptProjectOptions
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use TypeScriptProjectOptions
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Deprecated: use TypeScriptProjectOptions
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Deprecated: use TypeScriptProjectOptions
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Deprecated: use TypeScriptProjectOptions
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	// Deprecated: use TypeScriptProjectOptions
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Deprecated: use TypeScriptProjectOptions
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Deprecated: use TypeScriptProjectOptions
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Deprecated: use TypeScriptProjectOptions
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Deprecated: use TypeScriptProjectOptions
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Deprecated: use TypeScriptProjectOptions
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Deprecated: use TypeScriptProjectOptions
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Deprecated: use TypeScriptProjectOptions
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Deprecated: use TypeScriptProjectOptions
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Deprecated: use TypeScriptProjectOptions
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Deprecated: use TypeScriptProjectOptions
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Deprecated: use TypeScriptProjectOptions
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Deprecated: use TypeScriptProjectOptions
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Deprecated: use TypeScriptProjectOptions
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Deprecated: use TypeScriptProjectOptions
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Deprecated: use TypeScriptProjectOptions
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Deprecated: use TypeScriptProjectOptions
	Jest bool `json:"jest"`
	// Jest options.
	// Deprecated: use TypeScriptProjectOptions
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Deprecated: use TypeScriptProjectOptions
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Deprecated: use TypeScriptProjectOptions
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Deprecated: use TypeScriptProjectOptions
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Deprecated: use TypeScriptProjectOptions
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Deprecated: use TypeScriptProjectOptions
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Deprecated: use TypeScriptProjectOptions
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Deprecated: use TypeScriptProjectOptions
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Deprecated: use TypeScriptProjectOptions
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	// 
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	// 
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	// 
	// To create a personal access token see https://github.com/settings/tokens
	// Deprecated: use TypeScriptProjectOptions
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Deprecated: use TypeScriptProjectOptions
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Deprecated: use TypeScriptProjectOptions
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Deprecated: use TypeScriptProjectOptions
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Deprecated: use TypeScriptProjectOptions
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Deprecated: use TypeScriptProjectOptions
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Deprecated: use TypeScriptProjectOptions
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: use TypeScriptProjectOptions
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Deprecated: use TypeScriptProjectOptions
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Deprecated: use TypeScriptProjectOptions
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Deprecated: use TypeScriptProjectOptions
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Deprecated: use TypeScriptProjectOptions
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Deprecated: use TypeScriptProjectOptions
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Deprecated: use TypeScriptProjectOptions
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Compile the code before running tests.
	// Deprecated: use TypeScriptProjectOptions
	CompileBeforeTest bool `json:"compileBeforeTest"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Deprecated: use TypeScriptProjectOptions
	DisableTsconfig bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Deprecated: use TypeScriptProjectOptions
	Docgen bool `json:"docgen"`
	// Docs directory.
	// Deprecated: use TypeScriptProjectOptions
	DocsDirectory string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Deprecated: use TypeScriptProjectOptions
	EntrypointTypes string `json:"entrypointTypes"`
	// Setup eslint.
	// Deprecated: use TypeScriptProjectOptions
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Deprecated: use TypeScriptProjectOptions
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Deprecated: use TypeScriptProjectOptions
	Libdir string `json:"libdir"`
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Deprecated: use TypeScriptProjectOptions
	Package bool `json:"package"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Deprecated: use TypeScriptProjectOptions
	SampleCode bool `json:"sampleCode"`
	// Typescript sources directory.
	// Deprecated: use TypeScriptProjectOptions
	Srcdir string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	// 
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Deprecated: use TypeScriptProjectOptions
	Testdir string `json:"testdir"`
	// Custom TSConfig.
	// Deprecated: use TypeScriptProjectOptions
	Tsconfig TypescriptConfigOptionsIface `json:"tsconfig"`
	// TypeScript version to use.
	// Deprecated: use TypeScriptProjectOptions
	TypescriptVersion string `json:"typescriptVersion"`
}

func (t *TypeScriptLibraryProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		t,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		t,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		t,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		t,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		t,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		t,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		t,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		t,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		t,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		t,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		t,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		t,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		t,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		t,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		t,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		t,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		t,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		t,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		t,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		t,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		t,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		t,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		t,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		t,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		t,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		t,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		t,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		t,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		t,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		t,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		t,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		t,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		t,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		t,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		t,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		t,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		t,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		t,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		t,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		t,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		t,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		t,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		t,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		t,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		t,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		t,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		t,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		t,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		t,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		t,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		t,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		t,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		t,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		t,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		t,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		t,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		t,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		t,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetCompileBeforeTest() bool {
	var returns bool
	_jsii_.Get(
		t,
		"compileBeforeTest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDisableTsconfig() bool {
	var returns bool
	_jsii_.Get(
		t,
		"disableTsconfig",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		t,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetEntrypointTypes() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypointTypes",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		t,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		t,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetLibdir() string {
	var returns string
	_jsii_.Get(
		t,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetPackage() bool {
	var returns bool
	_jsii_.Get(
		t,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		t,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		t,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetTestdir() string {
	var returns string
	_jsii_.Get(
		t,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetTsconfig() TypescriptConfigOptionsIface {
	var returns TypescriptConfigOptionsIface
	_jsii_.Get(
		t,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptLibraryProjectOptions) GetTypescriptVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"typescriptVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Determines how modules get resolved.
// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html
//
// Experimental.
type TypeScriptModuleResolution string

const (
	TypeScriptModuleResolutionClassic TypeScriptModuleResolution = "CLASSIC"
	TypeScriptModuleResolutionNode TypeScriptModuleResolution = "NODE"
)

// Class interface
type TypeScriptProjectIface interface {
	GetComponents() []ComponentIface
	GetDeps() deps.DependenciesIface
	GetFiles() []FileBaseIface
	GetGitignore() IgnoreFileIface
	GetLogger() LoggerIface
	GetName() string
	GetOutdir() string
	GetProjectType() ProjectType
	GetRoot() ProjectIface
	GetTasks() tasks.TasksIface
	GetDevContainer() vscode.DevContainerIface
	GetGithub() github.GitHubIface
	GetGitpod() GitpodIface
	GetJsiiFqn() string
	GetParent() ProjectIface
	GetVscode() vscode.VsCodeIface
	GetAllowLibraryDependencies() bool
	GetAntitamper() bool
	GetBuildTask() tasks.TaskIface
	GetCompileTask() tasks.TaskIface
	GetEntrypoint() string
	GetInstallWorkflowSteps() []interface{}
	GetManifest() interface{}
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackage() NodePackageIface
	GetPackageManager() NodePackageManager
	GetProjenCommand() string
	GetRunScriptCommand() string
	GetTestCompileTask() tasks.TaskIface
	GetTestTask() tasks.TaskIface
	GetAutoMerge() github.AutoMergeIface
	GetBuildWorkflow() github.GithubWorkflowIface
	GetBuildWorkflowJobId() string
	GetJest() JestIface
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmignore() IgnoreFileIface
	GetReleaseWorkflow() github.GithubWorkflowIface
	GetReleaseWorkflowJobId() string
	GetDocsDirectory() string
	GetLibdir() string
	GetSrcdir() string
	GetTestdir() string
	GetWatchTask() tasks.TaskIface
	GetDocgen() bool
	GetEslint() EslintIface
	GetPackageTask() tasks.TaskIface
	GetTsconfig() TypescriptConfigIface
	AddExcludeFromCleanup(globs string)
	AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface
	AddTip(message string)
	PostSynthesize()
	PreSynthesize()
	Synth()
	TryFindFile(filePath string) FileBaseIface
	TryFindJsonFile(filePath string) JsonFileIface
	TryFindObjectFile(filePath string) ObjectFileIface
	AddBins(bins map[string]string)
	AddBuildCommand(commands string)
	AddBundledDeps(deps string)
	AddCompileCommand(commands string)
	AddDeps(deps string)
	AddDevDeps(deps string)
	AddFields(fields map[string]interface{})
	AddKeywords(keywords string)
	AddPeerDeps(deps string)
	AddTestCommand(commands string)
	HasScript(name string) bool
	RemoveScript(name string)
	RunTaskCommand(task tasks.TaskIface) string
	SetScript(name string, command string)
}

// TypeScript project.
// Experimental.
// Struct proxy
type TypeScriptProject struct {
	// Returns all the components within this project.
	// Experimental.
	Components []ComponentIface `json:"components"`
	// Project dependencies.
	// Experimental.
	Deps deps.DependenciesIface `json:"deps"`
	// All files in this project.
	// Experimental.
	Files []FileBaseIface `json:"files"`
	// .gitignore.
	// Experimental.
	Gitignore IgnoreFileIface `json:"gitignore"`
	// Logging utilities.
	// Experimental.
	Logger LoggerIface `json:"logger"`
	// Project name.
	// Experimental.
	Name string `json:"name"`
	// Absolute output directory of this project.
	// Experimental.
	Outdir string `json:"outdir"`
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The root project.
	// Experimental.
	Root ProjectIface `json:"root"`
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
	Gitpod GitpodIface `json:"gitpod"`
	// The JSII FQN of the project type (if known).
	// Experimental.
	JsiiFqn string `json:"jsiiFqn"`
	// A parent project.
	// 
	// If undefined, this is the root project.
	// Experimental.
	Parent ProjectIface `json:"parent"`
	// Access all VSCode components.
	// 
	// This will be `undefined` for subprojects.
	// Experimental.
	Vscode vscode.VsCodeIface `json:"vscode"`
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Indicates if workflows have anti-tamper checks.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// The task responsible for a full release build.
	// 
	// It spawns: compile + test + release + package
	// Experimental.
	BuildTask tasks.TaskIface `json:"buildTask"`
	// Compiles the code.
	// 
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask tasks.TaskIface `json:"compileTask"`
	// Deprecated: use `package.entrypoint`
	Entrypoint string `json:"entrypoint"`
	// Experimental.
	InstallWorkflowSteps []interface{} `json:"installWorkflowSteps"`
	// Deprecated: use `package.addField(x, y)`
	Manifest interface{} `json:"manifest"`
	// Deprecated: use `package.npmDistTag`
	NpmDistTag string `json:"npmDistTag"`
	// Deprecated: use `package.npmRegistry`
	NpmRegistry string `json:"npmRegistry"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Deprecated: use `package.npmTaskExecution`
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// API for managing the node package.
	// Experimental.
	Package NodePackageIface `json:"package"`
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager NodePackageManager `json:"packageManager"`
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand string `json:"runScriptCommand"`
	// Compiles the test code.
	// Experimental.
	TestCompileTask tasks.TaskIface `json:"testCompileTask"`
	// Tests the code.
	// Experimental.
	TestTask tasks.TaskIface `json:"testTask"`
	// Automatic PR merges.
	// Experimental.
	AutoMerge github.AutoMergeIface `json:"autoMerge"`
	// The PR build GitHub workflow.
	// 
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow github.GithubWorkflowIface `json:"buildWorkflow"`
	// Experimental.
	BuildWorkflowJobId string `json:"buildWorkflowJobId"`
	// The Jest configuration (if enabled).
	// Experimental.
	Jest JestIface `json:"jest"`
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// The .npmignore file.
	// Experimental.
	Npmignore IgnoreFileIface `json:"npmignore"`
	// The release GitHub workflow.
	// 
	// `undefined` if `releaseWorkflow` is disabled.
	// Experimental.
	ReleaseWorkflow github.GithubWorkflowIface `json:"releaseWorkflow"`
	// Experimental.
	ReleaseWorkflowJobId string `json:"releaseWorkflowJobId"`
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The directory in which compiled .js files reside.
	// Experimental.
	Libdir string `json:"libdir"`
	// The directory in which the .ts sources reside.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// The directory in which tests reside.
	// Experimental.
	Testdir string `json:"testdir"`
	// The "watch" task.
	// Experimental.
	WatchTask tasks.TaskIface `json:"watchTask"`
	// Experimental.
	Docgen bool `json:"docgen"`
	// Experimental.
	Eslint EslintIface `json:"eslint"`
	// The "package" task (or undefined if `package` is set to `false`).
	// Experimental.
	PackageTask tasks.TaskIface `json:"packageTask"`
	// Experimental.
	Tsconfig TypescriptConfigIface `json:"tsconfig"`
}

func (t *TypeScriptProject) GetComponents() []ComponentIface {
	var returns []ComponentIface
	_jsii_.Get(
		t,
		"components",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ComponentIface)(nil)).Elem(): reflect.TypeOf((*Component)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetDeps() deps.DependenciesIface {
	var returns deps.DependenciesIface
	_jsii_.Get(
		t,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*deps.DependenciesIface)(nil)).Elem(): reflect.TypeOf((*deps.Dependencies)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetFiles() []FileBaseIface {
	var returns []FileBaseIface
	_jsii_.Get(
		t,
		"files",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetGitignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		t,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetLogger() LoggerIface {
	var returns LoggerIface
	_jsii_.Get(
		t,
		"logger",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerIface)(nil)).Elem(): reflect.TypeOf((*Logger)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetOutdir() string {
	var returns string
	_jsii_.Get(
		t,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		t,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetRoot() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetTasks() tasks.TasksIface {
	var returns tasks.TasksIface
	_jsii_.Get(
		t,
		"tasks",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TasksIface)(nil)).Elem(): reflect.TypeOf((*tasks.Tasks)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetDevContainer() vscode.DevContainerIface {
	var returns vscode.DevContainerIface
	_jsii_.Get(
		t,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.DevContainerIface)(nil)).Elem(): reflect.TypeOf((*vscode.DevContainer)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetGithub() github.GitHubIface {
	var returns github.GitHubIface
	_jsii_.Get(
		t,
		"github",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GitHubIface)(nil)).Elem(): reflect.TypeOf((*github.GitHub)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetGitpod() GitpodIface {
	var returns GitpodIface
	_jsii_.Get(
		t,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*GitpodIface)(nil)).Elem(): reflect.TypeOf((*Gitpod)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		t,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetVscode() vscode.VsCodeIface {
	var returns vscode.VsCodeIface
	_jsii_.Get(
		t,
		"vscode",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*vscode.VsCodeIface)(nil)).Elem(): reflect.TypeOf((*vscode.VsCode)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		t,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetBuildTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"buildTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"compileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetInstallWorkflowSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		t,
		"installWorkflowSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetManifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		t,
		"manifest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		t,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		t,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetPackage() NodePackageIface {
	var returns NodePackageIface
	_jsii_.Get(
		t,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageIface)(nil)).Elem(): reflect.TypeOf((*NodePackage)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		t,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetRunScriptCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"runScriptCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetTestCompileTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"testCompileTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetTestTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"testTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetAutoMerge() github.AutoMergeIface {
	var returns github.AutoMergeIface
	_jsii_.Get(
		t,
		"autoMerge",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.AutoMergeIface)(nil)).Elem(): reflect.TypeOf((*github.AutoMerge)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetBuildWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		t,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetBuildWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		t,
		"buildWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetJest() JestIface {
	var returns JestIface
	_jsii_.Get(
		t,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestIface)(nil)).Elem(): reflect.TypeOf((*Jest)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetNpmignore() IgnoreFileIface {
	var returns IgnoreFileIface
	_jsii_.Get(
		t,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IgnoreFileIface)(nil)).Elem(): reflect.TypeOf((*IgnoreFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetReleaseWorkflow() github.GithubWorkflowIface {
	var returns github.GithubWorkflowIface
	_jsii_.Get(
		t,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.GithubWorkflowIface)(nil)).Elem(): reflect.TypeOf((*github.GithubWorkflow)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetReleaseWorkflowJobId() string {
	var returns string
	_jsii_.Get(
		t,
		"releaseWorkflowJobId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetLibdir() string {
	var returns string
	_jsii_.Get(
		t,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		t,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetTestdir() string {
	var returns string
	_jsii_.Get(
		t,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetWatchTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"watchTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		t,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) GetEslint() EslintIface {
	var returns EslintIface
	_jsii_.Get(
		t,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintIface)(nil)).Elem(): reflect.TypeOf((*Eslint)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetPackageTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		t,
		"packageTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) GetTsconfig() TypescriptConfigIface {
	var returns TypescriptConfigIface
	_jsii_.Get(
		t,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfig)(nil)).Elem(),
		},
	)
	return returns
}


func NewTypeScriptProject(options TypeScriptProjectOptionsIface) TypeScriptProjectIface {
	_init_.Initialize()
	self := TypeScriptProject{}
	_jsii_.Create(
		"projen.TypeScriptProject",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func TypeScriptProject_DefaultTask() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.TypeScriptProject",
		"DEFAULT_TASK",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) AddExcludeFromCleanup(globs string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addExcludeFromCleanup",
		[]interface{}{globs},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddTask(name string, props tasks.TaskOptionsIface) tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Invoke(
		t,
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

func (t *TypeScriptProject) AddTip(message string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addTip",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"synth",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) TryFindFile(filePath string) FileBaseIface {
	var returns FileBaseIface
	_jsii_.Invoke(
		t,
		"tryFindFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*FileBaseIface)(nil)).Elem(): reflect.TypeOf((*FileBase)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) TryFindJsonFile(filePath string) JsonFileIface {
	var returns JsonFileIface
	_jsii_.Invoke(
		t,
		"tryFindJsonFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) TryFindObjectFile(filePath string) ObjectFileIface {
	var returns ObjectFileIface
	_jsii_.Invoke(
		t,
		"tryFindObjectFile",
		[]interface{}{filePath},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ObjectFileIface)(nil)).Elem(): reflect.TypeOf((*ObjectFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProject) AddBins(bins map[string]string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBins",
		[]interface{}{bins},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddBuildCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBuildCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddBundledDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addBundledDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddCompileCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addCompileCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddDevDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addDevDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddFields(fields map[string]interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addFields",
		[]interface{}{fields},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddKeywords(keywords string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addKeywords",
		[]interface{}{keywords},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddPeerDeps(deps string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addPeerDeps",
		[]interface{}{deps},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) AddTestCommand(commands string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"addTestCommand",
		[]interface{}{commands},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) HasScript(name string) bool {
	var returns bool
	_jsii_.Invoke(
		t,
		"hasScript",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) RemoveScript(name string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"removeScript",
		[]interface{}{name},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (t *TypeScriptProject) RunTaskCommand(task tasks.TaskIface) string {
	var returns string
	_jsii_.Invoke(
		t,
		"runTaskCommand",
		[]interface{}{task},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProject) SetScript(name string, command string) {
	var returns interface{}
	_jsii_.Invoke(
		t,
		"setScript",
		[]interface{}{name, command},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// TypeScriptProjectOptionsIface is the public interface for the custom type TypeScriptProjectOptions
// Experimental.
type TypeScriptProjectOptionsIface interface {
	GetName() string
	GetClobber() bool
	GetDevContainer() bool
	GetGitpod() bool
	GetJsiiFqn() string
	GetLogging() LoggerOptionsIface
	GetOutdir() string
	GetParent() ProjectIface
	GetProjectType() ProjectType
	GetReadme() SampleReadmePropsIface
	GetAllowLibraryDependencies() bool
	GetAuthorEmail() string
	GetAuthorName() string
	GetAuthorOrganization() bool
	GetAuthorUrl() string
	GetAutoDetectBin() bool
	GetBin() map[string]string
	GetBundledDeps() []string
	GetDeps() []string
	GetDescription() string
	GetDevDeps() []string
	GetEntrypoint() string
	GetHomepage() string
	GetKeywords() []string
	GetLicense() string
	GetLicensed() bool
	GetMaxNodeVersion() string
	GetMinNodeVersion() string
	GetNpmAccess() NpmAccess
	GetNpmDistTag() string
	GetNpmRegistry() string
	GetNpmRegistryUrl() string
	GetNpmTaskExecution() NpmTaskExecution
	GetPackageManager() NodePackageManager
	GetPackageName() string
	GetPeerDependencyOptions() PeerDependencyOptionsIface
	GetPeerDeps() []string
	GetProjenCommand() string
	GetRepository() string
	GetRepositoryDirectory() string
	GetScripts() map[string]string
	GetStability() string
	GetAntitamper() bool
	GetBuildWorkflow() bool
	GetCodeCov() bool
	GetCodeCovTokenSecret() string
	GetCopyrightOwner() string
	GetCopyrightPeriod() string
	GetDefaultReleaseBranch() string
	GetDependabot() bool
	GetDependabotOptions() github.DependabotOptionsIface
	GetGitignore() []string
	GetJest() bool
	GetJestOptions() JestOptionsIface
	GetMergify() bool
	GetMergifyAutoMergeLabel() string
	GetMergifyOptions() github.MergifyOptionsIface
	GetNpmignore() []string
	GetNpmignoreEnabled() bool
	GetProjenDevDependency() bool
	GetProjenUpgradeAutoMerge() bool
	GetProjenUpgradeSchedule() []string
	GetProjenUpgradeSecret() string
	GetProjenVersion() SemverIface
	GetPullRequestTemplate() bool
	GetPullRequestTemplateContents() string
	GetRebuildBot() bool
	GetRebuildBotCommand() string
	GetReleaseBranches() []string
	GetReleaseEveryCommit() bool
	GetReleaseSchedule() string
	GetReleaseToNpm() bool
	GetReleaseWorkflow() bool
	GetWorkflowBootstrapSteps() []interface{}
	GetWorkflowContainerImage() string
	GetWorkflowNodeVersion() string
	GetCompileBeforeTest() bool
	GetDisableTsconfig() bool
	GetDocgen() bool
	GetDocsDirectory() string
	GetEntrypointTypes() string
	GetEslint() bool
	GetEslintOptions() EslintOptionsIface
	GetLibdir() string
	GetPackage() bool
	GetSampleCode() bool
	GetSrcdir() string
	GetTestdir() string
	GetTsconfig() TypescriptConfigOptionsIface
	GetTypescriptVersion() string
}

// Experimental.
// Struct proxy
type TypeScriptProjectOptions struct {
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
	Logging LoggerOptionsIface `json:"logging"`
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
	Parent ProjectIface `json:"parent"`
	// Which type of project this is (library/app).
	// Experimental.
	ProjectType ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme SampleReadmePropsIface `json:"readme"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	// 
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	// 
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin map[string]string `json:"bin"`
	// List of dependencies to bundle into this module.
	// 
	// These modules will be
	// added both to the `dependencies` section and `peerDependencies` section of
	// your `package.json`.
	// 
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps []string `json:"bundledDeps"`
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
	Deps []string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	// 
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description string `json:"description"`
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
	DevDeps []string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	// 
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords []string `json:"keywords"`
	// License's SPDX identifier.
	// 
	// See https://github.com/projen/projen/tree/master/license-text for a list of supported licenses.
	// Experimental.
	License string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// Tags can be used to provide an alias instead of version numbers.
	// 
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	// 
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	// 
	// The `next` tag is used by some projects to identify the upcoming version.
	// Experimental.
	NpmDistTag string `json:"npmDistTag"`
	// The host name of the npm registry to publish to.
	// 
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	// 
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl string `json:"npmRegistryUrl"`
	// Determines how tasks are executed when invoked as npm scripts (yarn/npm run xyz).
	// Experimental.
	NpmTaskExecution NpmTaskExecution `json:"npmTaskExecution"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions PeerDependencyOptionsIface `json:"peerDependencyOptions"`
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
	PeerDeps []string `json:"peerDeps"`
	// The shell command to use in order to run the projen CLI.
	// 
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand string `json:"projenCommand"`
	// The repository is the location where the actual code for your package lives.
	// 
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory string `json:"repositoryDirectory"`
	// npm scripts to include.
	// 
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts map[string]string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper bool `json:"antitamper"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow bool `json:"buildWorkflow"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod string `json:"copyrightPeriod"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch string `json:"defaultReleaseBranch"`
	// Include dependabot configuration.
	// Experimental.
	Dependabot bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions github.DependabotOptionsIface `json:"dependabotOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore []string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions JestOptionsIface `json:"jestOptions"`
	// Adds mergify configuration.
	// Experimental.
	Mergify bool `json:"mergify"`
	// Automatically merge PRs that build successfully and have this label.
	// 
	// To disable, set this value to an empty string.
	// Experimental.
	MergifyAutoMergeLabel string `json:"mergifyAutoMergeLabel"`
	// Options for mergify.
	// Experimental.
	MergifyOptions github.MergifyOptionsIface `json:"mergifyOptions"`
	// Additional entries to .npmignore.
	// Experimental.
	Npmignore []string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled bool `json:"npmignoreEnabled"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency bool `json:"projenDevDependency"`
	// Automatically merge projen upgrade PRs when build passes.
	// 
	// Applies the `mergifyAutoMergeLabel` to the PR if enabled.
	// Experimental.
	ProjenUpgradeAutoMerge bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule []string `json:"projenUpgradeSchedule"`
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
	ProjenUpgradeSecret string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion SemverIface `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents string `json:"pullRequestTemplateContents"`
	// Installs a GitHub workflow which is triggered when the comment "@projen rebuild" is added to a pull request.
	// 
	// The workflow will run a full build and
	// commit the changes to the pull request branch. This is useful for updating
	// test snapshots and other generated files like API.md.
	// Experimental.
	RebuildBot bool `json:"rebuildBot"`
	// The pull request bot command to use in order to trigger a rebuild and commit of the contents of the branch.
	// 
	// The command must be prefixed by "@projen", e.g. "@projen rebuild"
	// `gh pr review $pr --comment -b "@projen rebuild"`
	// Experimental.
	RebuildBotCommand string `json:"rebuildBotCommand"`
	// Branches which trigger a release.
	// 
	// Default value is based on defaultReleaseBranch.
	// Experimental.
	ReleaseBranches []string `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Experimental.
	ReleaseEveryCommit bool `json:"releaseEveryCommit"`
	// CRON schedule to trigger new releases.
	// Experimental.
	ReleaseSchedule string `json:"releaseSchedule"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm bool `json:"releaseToNpm"`
	// Define a GitHub workflow for releasing from "master" when new versions are bumped.
	// 
	// Requires that `version` will be undefined.
	// Experimental.
	ReleaseWorkflow bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps []interface{} `json:"workflowBootstrapSteps"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage string `json:"workflowContainerImage"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion string `json:"workflowNodeVersion"`
	// Compile the code before running tests.
	// Experimental.
	CompileBeforeTest bool `json:"compileBeforeTest"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Experimental.
	DisableTsconfig bool `json:"disableTsconfig"`
	// Docgen by Typedoc.
	// Experimental.
	Docgen bool `json:"docgen"`
	// Docs directory.
	// Experimental.
	DocsDirectory string `json:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Experimental.
	EntrypointTypes string `json:"entrypointTypes"`
	// Setup eslint.
	// Experimental.
	Eslint bool `json:"eslint"`
	// Eslint options.
	// Experimental.
	EslintOptions EslintOptionsIface `json:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Experimental.
	Libdir string `json:"libdir"`
	// Defines a `yarn package` command that will produce a tarball and place it under `dist/js`.
	// Experimental.
	Package bool `json:"package"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Experimental.
	SampleCode bool `json:"sampleCode"`
	// Typescript sources directory.
	// Experimental.
	Srcdir string `json:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	// 
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Experimental.
	Testdir string `json:"testdir"`
	// Custom TSConfig.
	// Experimental.
	Tsconfig TypescriptConfigOptionsIface `json:"tsconfig"`
	// TypeScript version to use.
	// Experimental.
	TypescriptVersion string `json:"typescriptVersion"`
}

func (t *TypeScriptProjectOptions) GetName() string {
	var returns string
	_jsii_.Get(
		t,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetClobber() bool {
	var returns bool
	_jsii_.Get(
		t,
		"clobber",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDevContainer() bool {
	var returns bool
	_jsii_.Get(
		t,
		"devContainer",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetGitpod() bool {
	var returns bool
	_jsii_.Get(
		t,
		"gitpod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetJsiiFqn() string {
	var returns string
	_jsii_.Get(
		t,
		"jsiiFqn",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetLogging() LoggerOptionsIface {
	var returns LoggerOptionsIface
	_jsii_.Get(
		t,
		"logging",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*LoggerOptionsIface)(nil)).Elem(): reflect.TypeOf((*LoggerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		t,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetParent() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		t,
		"parent",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjectType() ProjectType {
	var returns ProjectType
	_jsii_.Get(
		t,
		"projectType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectType)(nil)).Elem(): reflect.TypeOf((*ProjectType)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetReadme() SampleReadmePropsIface {
	var returns SampleReadmePropsIface
	_jsii_.Get(
		t,
		"readme",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SampleReadmePropsIface)(nil)).Elem(): reflect.TypeOf((*SampleReadmeProps)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAllowLibraryDependencies() bool {
	var returns bool
	_jsii_.Get(
		t,
		"allowLibraryDependencies",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAuthorEmail() string {
	var returns string
	_jsii_.Get(
		t,
		"authorEmail",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAuthorName() string {
	var returns string
	_jsii_.Get(
		t,
		"authorName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAuthorOrganization() bool {
	var returns bool
	_jsii_.Get(
		t,
		"authorOrganization",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAuthorUrl() string {
	var returns string
	_jsii_.Get(
		t,
		"authorUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAutoDetectBin() bool {
	var returns bool
	_jsii_.Get(
		t,
		"autoDetectBin",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetBin() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"bin",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetBundledDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"bundledDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"deps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDescription() string {
	var returns string
	_jsii_.Get(
		t,
		"description",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDevDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"devDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetEntrypoint() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypoint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetHomepage() string {
	var returns string
	_jsii_.Get(
		t,
		"homepage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetKeywords() []string {
	var returns []string
	_jsii_.Get(
		t,
		"keywords",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetLicense() string {
	var returns string
	_jsii_.Get(
		t,
		"license",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetLicensed() bool {
	var returns bool
	_jsii_.Get(
		t,
		"licensed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetMaxNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"maxNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetMinNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"minNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		t,
		"npmAccess",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmAccess)(nil)).Elem(): reflect.TypeOf((*NpmAccess)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmDistTag() string {
	var returns string
	_jsii_.Get(
		t,
		"npmDistTag",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmRegistry() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistry",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmRegistryUrl() string {
	var returns string
	_jsii_.Get(
		t,
		"npmRegistryUrl",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmTaskExecution() NpmTaskExecution {
	var returns NpmTaskExecution
	_jsii_.Get(
		t,
		"npmTaskExecution",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(): reflect.TypeOf((*NpmTaskExecution)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		t,
		"packageManager",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodePackageManager)(nil)).Elem(): reflect.TypeOf((*NodePackageManager)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPackageName() string {
	var returns string
	_jsii_.Get(
		t,
		"packageName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPeerDependencyOptions() PeerDependencyOptionsIface {
	var returns PeerDependencyOptionsIface
	_jsii_.Get(
		t,
		"peerDependencyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PeerDependencyOptionsIface)(nil)).Elem(): reflect.TypeOf((*PeerDependencyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPeerDeps() []string {
	var returns []string
	_jsii_.Get(
		t,
		"peerDeps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjenCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"projenCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetRepository() string {
	var returns string
	_jsii_.Get(
		t,
		"repository",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetRepositoryDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"repositoryDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetScripts() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		t,
		"scripts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetStability() string {
	var returns string
	_jsii_.Get(
		t,
		"stability",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetAntitamper() bool {
	var returns bool
	_jsii_.Get(
		t,
		"antitamper",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetBuildWorkflow() bool {
	var returns bool
	_jsii_.Get(
		t,
		"buildWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetCodeCov() bool {
	var returns bool
	_jsii_.Get(
		t,
		"codeCov",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetCodeCovTokenSecret() string {
	var returns string
	_jsii_.Get(
		t,
		"codeCovTokenSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetCopyrightOwner() string {
	var returns string
	_jsii_.Get(
		t,
		"copyrightOwner",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetCopyrightPeriod() string {
	var returns string
	_jsii_.Get(
		t,
		"copyrightPeriod",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDefaultReleaseBranch() string {
	var returns string
	_jsii_.Get(
		t,
		"defaultReleaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDependabot() bool {
	var returns bool
	_jsii_.Get(
		t,
		"dependabot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDependabotOptions() github.DependabotOptionsIface {
	var returns github.DependabotOptionsIface
	_jsii_.Get(
		t,
		"dependabotOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.DependabotOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.DependabotOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetGitignore() []string {
	var returns []string
	_jsii_.Get(
		t,
		"gitignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetJest() bool {
	var returns bool
	_jsii_.Get(
		t,
		"jest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetJestOptions() JestOptionsIface {
	var returns JestOptionsIface
	_jsii_.Get(
		t,
		"jestOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JestOptionsIface)(nil)).Elem(): reflect.TypeOf((*JestOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetMergify() bool {
	var returns bool
	_jsii_.Get(
		t,
		"mergify",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetMergifyAutoMergeLabel() string {
	var returns string
	_jsii_.Get(
		t,
		"mergifyAutoMergeLabel",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetMergifyOptions() github.MergifyOptionsIface {
	var returns github.MergifyOptionsIface
	_jsii_.Get(
		t,
		"mergifyOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*github.MergifyOptionsIface)(nil)).Elem(): reflect.TypeOf((*github.MergifyOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmignore() []string {
	var returns []string
	_jsii_.Get(
		t,
		"npmignore",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetNpmignoreEnabled() bool {
	var returns bool
	_jsii_.Get(
		t,
		"npmignoreEnabled",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjenDevDependency() bool {
	var returns bool
	_jsii_.Get(
		t,
		"projenDevDependency",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjenUpgradeAutoMerge() bool {
	var returns bool
	_jsii_.Get(
		t,
		"projenUpgradeAutoMerge",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjenUpgradeSchedule() []string {
	var returns []string
	_jsii_.Get(
		t,
		"projenUpgradeSchedule",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjenUpgradeSecret() string {
	var returns string
	_jsii_.Get(
		t,
		"projenUpgradeSecret",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetProjenVersion() SemverIface {
	var returns SemverIface
	_jsii_.Get(
		t,
		"projenVersion",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*SemverIface)(nil)).Elem(): reflect.TypeOf((*Semver)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPullRequestTemplate() bool {
	var returns bool
	_jsii_.Get(
		t,
		"pullRequestTemplate",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPullRequestTemplateContents() string {
	var returns string
	_jsii_.Get(
		t,
		"pullRequestTemplateContents",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetRebuildBot() bool {
	var returns bool
	_jsii_.Get(
		t,
		"rebuildBot",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetRebuildBotCommand() string {
	var returns string
	_jsii_.Get(
		t,
		"rebuildBotCommand",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetReleaseBranches() []string {
	var returns []string
	_jsii_.Get(
		t,
		"releaseBranches",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetReleaseEveryCommit() bool {
	var returns bool
	_jsii_.Get(
		t,
		"releaseEveryCommit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetReleaseSchedule() string {
	var returns string
	_jsii_.Get(
		t,
		"releaseSchedule",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetReleaseToNpm() bool {
	var returns bool
	_jsii_.Get(
		t,
		"releaseToNpm",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetReleaseWorkflow() bool {
	var returns bool
	_jsii_.Get(
		t,
		"releaseWorkflow",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetWorkflowBootstrapSteps() []interface{} {
	var returns []interface{}
	_jsii_.Get(
		t,
		"workflowBootstrapSteps",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetWorkflowContainerImage() string {
	var returns string
	_jsii_.Get(
		t,
		"workflowContainerImage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetWorkflowNodeVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"workflowNodeVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetCompileBeforeTest() bool {
	var returns bool
	_jsii_.Get(
		t,
		"compileBeforeTest",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDisableTsconfig() bool {
	var returns bool
	_jsii_.Get(
		t,
		"disableTsconfig",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDocgen() bool {
	var returns bool
	_jsii_.Get(
		t,
		"docgen",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetDocsDirectory() string {
	var returns string
	_jsii_.Get(
		t,
		"docsDirectory",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetEntrypointTypes() string {
	var returns string
	_jsii_.Get(
		t,
		"entrypointTypes",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetEslint() bool {
	var returns bool
	_jsii_.Get(
		t,
		"eslint",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetEslintOptions() EslintOptionsIface {
	var returns EslintOptionsIface
	_jsii_.Get(
		t,
		"eslintOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EslintOptionsIface)(nil)).Elem(): reflect.TypeOf((*EslintOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetLibdir() string {
	var returns string
	_jsii_.Get(
		t,
		"libdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetPackage() bool {
	var returns bool
	_jsii_.Get(
		t,
		"package",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetSampleCode() bool {
	var returns bool
	_jsii_.Get(
		t,
		"sampleCode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetSrcdir() string {
	var returns string
	_jsii_.Get(
		t,
		"srcdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetTestdir() string {
	var returns string
	_jsii_.Get(
		t,
		"testdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetTsconfig() TypescriptConfigOptionsIface {
	var returns TypescriptConfigOptionsIface
	_jsii_.Get(
		t,
		"tsconfig",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypescriptConfigOptionsIface)(nil)).Elem(): reflect.TypeOf((*TypescriptConfigOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypeScriptProjectOptions) GetTypescriptVersion() string {
	var returns string
	_jsii_.Get(
		t,
		"typescriptVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type TypescriptConfigIface interface {
	GetCompilerOptions() TypeScriptCompilerOptionsIface
	GetExclude() []string
	GetFile() JsonFileIface
	GetFileName() string
	GetInclude() []string
}

// Experimental.
// Struct proxy
type TypescriptConfig struct {
	// Experimental.
	CompilerOptions TypeScriptCompilerOptionsIface `json:"compilerOptions"`
	// Experimental.
	Exclude []string `json:"exclude"`
	// Experimental.
	File JsonFileIface `json:"file"`
	// Experimental.
	FileName string `json:"fileName"`
	// Experimental.
	Include []string `json:"include"`
}

func (t *TypescriptConfig) GetCompilerOptions() TypeScriptCompilerOptionsIface {
	var returns TypeScriptCompilerOptionsIface
	_jsii_.Get(
		t,
		"compilerOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypeScriptCompilerOptionsIface)(nil)).Elem(): reflect.TypeOf((*TypeScriptCompilerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypescriptConfig) GetExclude() []string {
	var returns []string
	_jsii_.Get(
		t,
		"exclude",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypescriptConfig) GetFile() JsonFileIface {
	var returns JsonFileIface
	_jsii_.Get(
		t,
		"file",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*JsonFileIface)(nil)).Elem(): reflect.TypeOf((*JsonFile)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypescriptConfig) GetFileName() string {
	var returns string
	_jsii_.Get(
		t,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypescriptConfig) GetInclude() []string {
	var returns []string
	_jsii_.Get(
		t,
		"include",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


func NewTypescriptConfig(project NodeProjectIface, options TypescriptConfigOptionsIface) TypescriptConfigIface {
	_init_.Initialize()
	self := TypescriptConfig{}
	_jsii_.Create(
		"projen.TypescriptConfig",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

// TypescriptConfigOptionsIface is the public interface for the custom type TypescriptConfigOptions
// Experimental.
type TypescriptConfigOptionsIface interface {
	GetCompilerOptions() TypeScriptCompilerOptionsIface
	GetExclude() []string
	GetFileName() string
	GetInclude() []string
}

// Experimental.
// Struct proxy
type TypescriptConfigOptions struct {
	// Compiler options to use.
	// Experimental.
	CompilerOptions TypeScriptCompilerOptionsIface `json:"compilerOptions"`
	// Filters results from the "include" option.
	// Experimental.
	Exclude []string `json:"exclude"`
	// Experimental.
	FileName string `json:"fileName"`
	// Specifies a list of glob patterns that match TypeScript files to be included in compilation.
	// Experimental.
	Include []string `json:"include"`
}

func (t *TypescriptConfigOptions) GetCompilerOptions() TypeScriptCompilerOptionsIface {
	var returns TypeScriptCompilerOptionsIface
	_jsii_.Get(
		t,
		"compilerOptions",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*TypeScriptCompilerOptionsIface)(nil)).Elem(): reflect.TypeOf((*TypeScriptCompilerOptions)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypescriptConfigOptions) GetExclude() []string {
	var returns []string
	_jsii_.Get(
		t,
		"exclude",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (t *TypescriptConfigOptions) GetFileName() string {
	var returns string
	_jsii_.Get(
		t,
		"fileName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (t *TypescriptConfigOptions) GetInclude() []string {
	var returns []string
	_jsii_.Get(
		t,
		"include",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type VersionIface interface {
	GetProject() ProjectIface
	GetBumpTask() tasks.TaskIface
	GetCurrentVersion() interface{}
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// Experimental.
// Struct proxy
type Version struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// Experimental.
	BumpTask tasks.TaskIface `json:"bumpTask"`
	// Returns the current version of the project.
	// Experimental.
	CurrentVersion interface{} `json:"currentVersion"`
}

func (v *Version) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		v,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (v *Version) GetBumpTask() tasks.TaskIface {
	var returns tasks.TaskIface
	_jsii_.Get(
		v,
		"bumpTask",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*tasks.TaskIface)(nil)).Elem(): reflect.TypeOf((*tasks.Task)(nil)).Elem(),
		},
	)
	return returns
}

func (v *Version) GetCurrentVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		v,
		"currentVersion",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewVersion(project NodeProjectIface, options VersionOptionsIface) VersionIface {
	_init_.Initialize()
	self := Version{}
	_jsii_.Create(
		"projen.Version",
		[]interface{}{project, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (v *Version) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *Version) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (v *Version) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		v,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// VersionOptionsIface is the public interface for the custom type VersionOptions
// Experimental.
type VersionOptionsIface interface {
	GetReleaseBranch() string
}

// Experimental.
// Struct proxy
type VersionOptions struct {
	// The name of the release branch where the code and tags are pushed to.
	// Experimental.
	ReleaseBranch string `json:"releaseBranch"`
}

func (v *VersionOptions) GetReleaseBranch() string {
	var returns string
	_jsii_.Get(
		v,
		"releaseBranch",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type XmlFileIface interface {
	IMarkableFileIface
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	GetMarker() bool
	GetOmitEmpty() bool
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddDeletionOverride(path string)
	AddOverride(path string, value interface{})
}

// Represents an XML file.
// 
// Objects passed in will be synthesized using the npm "xml" library.
// See: https://www.npmjs.com/package/xml
//
// Experimental.
// Struct proxy
type XmlFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Indicates if the projen marker JSON-comment will be added to the output object.
	// Experimental.
	Marker bool `json:"marker"`
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (x *XmlFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		x,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (x *XmlFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		x,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) GetPath() string {
	var returns string
	_jsii_.Get(
		x,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		x,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		x,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		x,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		x,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewXmlFile(project ProjectIface, filePath string, options XmlFileOptionsIface) XmlFileIface {
	_init_.Initialize()
	self := XmlFile{}
	_jsii_.Create(
		"projen.XmlFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (x *XmlFile) SetExecutable(val bool) {
	_jsii_.Set(
		x,
		"executable",
		val,
	)
}

func (x *XmlFile) SetReadonly(val bool) {
	_jsii_.Set(
		x,
		"readonly",
		val,
	)
}

func XmlFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.XmlFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		x,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (x *XmlFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		x,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (x *XmlFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		x,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (x *XmlFile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		x,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFile) AddDeletionOverride(path string) {
	var returns interface{}
	_jsii_.Invoke(
		x,
		"addDeletionOverride",
		[]interface{}{path},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (x *XmlFile) AddOverride(path string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		x,
		"addOverride",
		[]interface{}{path, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// XmlFileOptionsIface is the public interface for the custom type XmlFileOptions
// Experimental.
type XmlFileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetMarker() bool
	GetObj() interface{}
	GetOmitEmpty() bool
}

// Options for `XmlFile`.
// Experimental.
// Struct proxy
type XmlFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker bool `json:"marker"`
	// The object that will be serialized.
	// 
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (x *XmlFileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		x,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		x,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		x,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		x,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFileOptions) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		x,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFileOptions) GetObj() interface{} {
	var returns interface{}
	_jsii_.Get(
		x,
		"obj",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (x *XmlFileOptions) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		x,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type YamlFileIface interface {
	IMarkableFileIface
	GetProject() ProjectIface
	GetAbsolutePath() string
	GetPath() string
	GetExecutable() bool
	SetExecutable(val bool)
	GetReadonly() bool
	SetReadonly(val bool)
	GetMarker() bool
	GetOmitEmpty() bool
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(resolver IResolverIface) string
	AddDeletionOverride(path string)
	AddOverride(path string, value interface{})
}

// Represents a YAML file.
// Experimental.
// Struct proxy
type YamlFile struct {
	// Experimental.
	Project ProjectIface `json:"project"`
	// The absolute path of this file.
	// Experimental.
	AbsolutePath string `json:"absolutePath"`
	// The file path, relative to the project root.
	// Experimental.
	Path string `json:"path"`
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Indicates if the projen marker JSON-comment will be added to the output object.
	// Experimental.
	Marker bool `json:"marker"`
	// Indicates if empty objects and arrays are omitted from the output object.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (y *YamlFile) GetProject() ProjectIface {
	var returns ProjectIface
	_jsii_.Get(
		y,
		"project",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProjectIface)(nil)).Elem(): reflect.TypeOf((*Project)(nil)).Elem(),
		},
	)
	return returns
}

func (y *YamlFile) GetAbsolutePath() string {
	var returns string
	_jsii_.Get(
		y,
		"absolutePath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) GetPath() string {
	var returns string
	_jsii_.Get(
		y,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		y,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		y,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		y,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		y,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewYamlFile(project ProjectIface, filePath string, options YamlFileOptionsIface) YamlFileIface {
	_init_.Initialize()
	self := YamlFile{}
	_jsii_.Create(
		"projen.YamlFile",
		[]interface{}{project, filePath, options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (y *YamlFile) SetExecutable(val bool) {
	_jsii_.Set(
		y,
		"executable",
		val,
	)
}

func (y *YamlFile) SetReadonly(val bool) {
	_jsii_.Set(
		y,
		"readonly",
		val,
	)
}

func YamlFile_ProjenMarker() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"projen.YamlFile",
		"PROJEN_MARKER",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) PostSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		y,
		"postSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (y *YamlFile) PreSynthesize() {
	var returns interface{}
	_jsii_.Invoke(
		y,
		"preSynthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (y *YamlFile) Synthesize() {
	var returns interface{}
	_jsii_.Invoke(
		y,
		"synthesize",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (y *YamlFile) SynthesizeContent(resolver IResolverIface) string {
	var returns string
	_jsii_.Invoke(
		y,
		"synthesizeContent",
		[]interface{}{resolver},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFile) AddDeletionOverride(path string) {
	var returns interface{}
	_jsii_.Invoke(
		y,
		"addDeletionOverride",
		[]interface{}{path},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (y *YamlFile) AddOverride(path string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		y,
		"addOverride",
		[]interface{}{path, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// YamlFileOptionsIface is the public interface for the custom type YamlFileOptions
// Experimental.
type YamlFileOptionsIface interface {
	GetCommitted() bool
	GetEditGitignore() bool
	GetExecutable() bool
	GetReadonly() bool
	GetMarker() bool
	GetObj() interface{}
	GetOmitEmpty() bool
}

// Options for `JsonFile`.
// Experimental.
// Struct proxy
type YamlFileOptions struct {
	// Indicates whether this file should be committed to git or ignored.
	// 
	// By
	// default, all generated files are committed and anti-tamper is used to
	// protect against manual modifications.
	// Experimental.
	Committed bool `json:"committed"`
	// Update the project's .gitignore file.
	// Experimental.
	EditGitignore bool `json:"editGitignore"`
	// Whether the generated file should be marked as executable.
	// Experimental.
	Executable bool `json:"executable"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly bool `json:"readonly"`
	// Adds the projen marker to the file.
	// Experimental.
	Marker bool `json:"marker"`
	// The object that will be serialized.
	// 
	// You can modify the object's contents
	// before synthesis.
	// Experimental.
	Obj interface{} `json:"obj"`
	// Omits empty objects and arrays.
	// Experimental.
	OmitEmpty bool `json:"omitEmpty"`
}

func (y *YamlFileOptions) GetCommitted() bool {
	var returns bool
	_jsii_.Get(
		y,
		"committed",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFileOptions) GetEditGitignore() bool {
	var returns bool
	_jsii_.Get(
		y,
		"editGitignore",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFileOptions) GetExecutable() bool {
	var returns bool
	_jsii_.Get(
		y,
		"executable",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFileOptions) GetReadonly() bool {
	var returns bool
	_jsii_.Get(
		y,
		"readonly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFileOptions) GetMarker() bool {
	var returns bool
	_jsii_.Get(
		y,
		"marker",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFileOptions) GetObj() interface{} {
	var returns interface{}
	_jsii_.Get(
		y,
		"obj",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (y *YamlFileOptions) GetOmitEmpty() bool {
	var returns bool
	_jsii_.Get(
		y,
		"omitEmpty",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


