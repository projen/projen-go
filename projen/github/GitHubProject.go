package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
	"github.com/projen/projen-go/projen/vscode"
)

// GitHub-based project.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
type GitHubProject interface {
	projen.Project
	// Auto approve set up for this project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	AutoApprove() AutoApprove
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	BuildTask() projen.Task
	// Whether to commit the managed files by default.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	CommitGenerated() *bool
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	CompileTask() projen.Task
	// Returns all the components within this project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Components() *[]projen.Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	DefaultTask() projen.Task
	// Project dependencies.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Deps() projen.Dependencies
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	DevContainer() vscode.DevContainer
	// Whether or not the project is being ejected.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Ejected() *bool
	// All files in this project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Files() *[]projen.FileBase
	// The .gitattributes file for this repository.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Gitattributes() projen.GitAttributesFile
	// Access all github components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Github() GitHub
	// .gitignore.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Gitignore() projen.IgnoreFile
	// Access for Gitpod.
	//
	// This will be `undefined` if gitpod boolean is false.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Gitpod() projen.Gitpod
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	InitProject() *projen.InitProject
	// Logging utilities.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Logger() projen.Logger
	// Project name.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Name() *string
	// Absolute output directory of this project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Outdir() *string
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	PackageTask() projen.Task
	// A parent project.
	//
	// If undefined, this is the root project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Parent() projen.Project
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	PostCompileTask() projen.Task
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	PreCompileTask() projen.Task
	// Manages the build process of the project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	ProjectBuild() projen.ProjectBuild
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	ProjenCommand() *string
	// The root project.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Root() projen.Project
	// Project tasks.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Tasks() projen.Tasks
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	TestTask() projen.Task
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Vscode() vscode.VsCode
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	AddExcludeFromCleanup(globs ...*string)
	// Adds a .gitignore pattern.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	AddGitIgnore(pattern *string)
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	AddPackageIgnore(_pattern *string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
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
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	AnnotateGenerated(glob *string)
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	PostSynthesize()
	// Called before all components are synthesized.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	PreSynthesize()
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	RemoveTask(name *string) projen.Task
	// Returns the shell command to execute in order to run a task.
	//
	// By default, this is `npx projen@<version> <task>`.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	RunTaskCommand(task projen.Task) *string
	// Synthesize all project files into `outdir`.
	//
	// 1. Call "this.preSynthesize()"
	// 2. Delete all generated files
	// 3. Synthesize all sub-projects
	// 4. Synthesize all components of this project
	// 5. Call "postSynthesize()" for all components of this project
	// 6. Call "this.postSynthesize()"
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	Synth()
	// Finds a file at the specified relative path within this project and all its subprojects.
	//
	// Returns: a `FileBase` or undefined if there is no file in that path.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	TryFindFile(filePath *string) projen.FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) projen.JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	TryFindObjectFile(filePath *string) projen.ObjectFile
	// Finds a file at the specified relative path within this project and removes it.
	//
	// Returns: a `FileBase` if the file was found and removed, or undefined if
	// the file was not found.
	// Deprecated: This is a *temporary* class. At the moment, our base project
	// types such as `NodeProject` and `JavaProject` are derived from this, but we
	// want to be able to use these project types outside of GitHub as well. One of
	// the next steps to address this is to abstract workflows so that different
	// "engines" can be used to implement our CI/CD solutions.
	TryRemoveFile(filePath *string) projen.FileBase
}

// The jsii proxy struct for GitHubProject
type jsiiProxy_GitHubProject struct {
	internal.Type__projenProject
}

func (j *jsiiProxy_GitHubProject) AutoApprove() AutoApprove {
	var returns AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) CommitGenerated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"commitGenerated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Github() GitHub {
	var returns GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func NewGitHubProject(options *GitHubProjectOptions) GitHubProject {
	_init_.Initialize()

	if err := validateNewGitHubProjectParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubProject{}

	_jsii_.Create(
		"projen.github.GitHubProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func NewGitHubProject_Override(g GitHubProject, options *GitHubProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.GitHubProject",
		[]interface{}{options},
		g,
	)
}

func GitHubProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.github.GitHubProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GitHubProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addExcludeFromCleanup",
		args,
	)
}

func (g *jsiiProxy_GitHubProject) AddGitIgnore(pattern *string) {
	if err := g.validateAddGitIgnoreParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (g *jsiiProxy_GitHubProject) AddPackageIgnore(_pattern *string) {
	if err := g.validateAddPackageIgnoreParameters(_pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

func (g *jsiiProxy_GitHubProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	if err := g.validateAddTaskParameters(name, props); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		g,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubProject) AddTip(message *string) {
	if err := g.validateAddTipParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addTip",
		[]interface{}{message},
	)
}

func (g *jsiiProxy_GitHubProject) AnnotateGenerated(glob *string) {
	if err := g.validateAnnotateGeneratedParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (g *jsiiProxy_GitHubProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHubProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHubProject) RemoveTask(name *string) projen.Task {
	if err := g.validateRemoveTaskParameters(name); err != nil {
		panic(err)
	}
	var returns projen.Task

	_jsii_.Invoke(
		g,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubProject) RunTaskCommand(task projen.Task) *string {
	if err := g.validateRunTaskCommandParameters(task); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubProject) Synth() {
	_jsii_.InvokeVoid(
		g,
		"synth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHubProject) TryFindFile(filePath *string) projen.FileBase {
	if err := g.validateTryFindFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		g,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	if err := g.validateTryFindJsonFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.JsonFile

	_jsii_.Invoke(
		g,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	if err := g.validateTryFindObjectFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.ObjectFile

	_jsii_.Invoke(
		g,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubProject) TryRemoveFile(filePath *string) projen.FileBase {
	if err := g.validateTryRemoveFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns projen.FileBase

	_jsii_.Invoke(
		g,
		"tryRemoveFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

