package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/vscode"
)

// Auto approve pull requests that meet a criteria.
// Experimental.
type AutoApprove interface {
	projen.Component
	Label() *string
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for AutoApprove
type jsiiProxy_AutoApprove struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoApprove) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoApprove) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoApprove(github GitHub, options *AutoApproveOptions) AutoApprove {
	_init_.Initialize()

	j := jsiiProxy_AutoApprove{}

	_jsii_.Create(
		"projen.github.AutoApprove",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoApprove_Override(a AutoApprove, github GitHub, options *AutoApproveOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.AutoApprove",
		[]interface{}{github, options},
		a,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AutoApprove) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (a *jsiiProxy_AutoApprove) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (a *jsiiProxy_AutoApprove) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Options for 'AutoApprove'.
// Experimental.
type AutoApproveOptions struct {
	// Only pull requests authored by these Github usernames will be auto-approved.
	// Experimental.
	AllowedUsernames *[]*string `json:"allowedUsernames" yaml:"allowedUsernames"`
	// Only pull requests with this label will be auto-approved.
	// Experimental.
	Label *string `json:"label" yaml:"label"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// A GitHub secret name which contains a GitHub Access Token with write permissions for the `pull_request` scope.
	//
	// This token is used to approve pull requests.
	//
	// Github forbids an identity to approve its own pull request.
	// If your project produces automated pull requests using the Github default token -
	// {@link https://docs.github.com/en/actions/reference/authentication-in-a-workflow `GITHUB_TOKEN` }
	// - that you would like auto approved, such as when using the `depsUpgrade` property in
	// `NodeProjectOptions`, then you must use a different token here.
	// Experimental.
	Secret *string `json:"secret" yaml:"secret"`
}

// Sets up mergify to merging approved pull requests.
//
// If `buildJob` is specified, the specified GitHub workflow job ID is required
// to succeed in order for the PR to be merged.
//
// `approvedReviews` specified the number of code review approvals required for
// the PR to be merged.
// Experimental.
type AutoMerge interface {
	projen.Component
	Project() projen.Project
	AddConditions(conditions ...*string)
	AddConditionsLater(later IAddConditionsLater)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for AutoMerge
type jsiiProxy_AutoMerge struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoMerge) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoMerge(github GitHub, options *AutoMergeOptions) AutoMerge {
	_init_.Initialize()

	j := jsiiProxy_AutoMerge{}

	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoMerge_Override(a AutoMerge, github GitHub, options *AutoMergeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{github, options},
		a,
	)
}

// Adds conditions to the auto merge rule.
// Experimental.
func (a *jsiiProxy_AutoMerge) AddConditions(conditions ...*string) {
	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addConditions",
		args,
	)
}

// Adds conditions that will be rendered only during synthesis.
// Experimental.
func (a *jsiiProxy_AutoMerge) AddConditionsLater(later IAddConditionsLater) {
	_jsii_.InvokeVoid(
		a,
		"addConditionsLater",
		[]interface{}{later},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (a *jsiiProxy_AutoMerge) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (a *jsiiProxy_AutoMerge) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (a *jsiiProxy_AutoMerge) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type AutoMergeOptions struct {
	// Number of approved code reviews.
	// Experimental.
	ApprovedReviews *float64 `json:"approvedReviews" yaml:"approvedReviews"`
	// List of labels that will prevent auto-merging.
	// Experimental.
	BlockingLabels *[]*string `json:"blockingLabels" yaml:"blockingLabels"`
}

// Defines dependabot configuration for node projects.
//
// Since module versions are managed in projen, the versioning strategy will be
// configured to "lockfile-only" which means that only updates that can be done
// on the lockfile itself will be proposed.
// Experimental.
type Dependabot interface {
	projen.Component
	Config() interface{}
	IgnoresProjen() *bool
	Project() projen.Project
	AddIgnore(dependencyName *string, versions ...*string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Dependabot
type jsiiProxy_Dependabot struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Dependabot) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) IgnoresProjen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoresProjen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDependabot(github GitHub, options *DependabotOptions) Dependabot {
	_init_.Initialize()

	j := jsiiProxy_Dependabot{}

	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewDependabot_Override(d Dependabot, github GitHub, options *DependabotOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		d,
	)
}

// Ignores a dependency from automatic updates.
// Experimental.
func (d *jsiiProxy_Dependabot) AddIgnore(dependencyName *string, versions ...*string) {
	args := []interface{}{dependencyName}
	for _, a := range versions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addIgnore",
		args,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (d *jsiiProxy_Dependabot) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (d *jsiiProxy_Dependabot) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (d *jsiiProxy_Dependabot) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

// You can use the `ignore` option to customize which dependencies are updated.
//
// The ignore option supports the following options.
// Experimental.
type DependabotIgnore struct {
	// Use to ignore updates for dependencies with matching names, optionally using `*` to match zero or more characters.
	//
	// For Java dependencies, the format of the dependency-name attribute is:
	// `groupId:artifactId`, for example: `org.kohsuke:github-api`.
	// Experimental.
	DependencyName *string `json:"dependencyName" yaml:"dependencyName"`
	// Use to ignore specific versions or ranges of versions.
	//
	// If you want to
	// define a range, use the standard pattern for the package manager (for
	// example: `^1.0.0` for npm, or `~> 2.0` for Bundler).
	// Experimental.
	Versions *[]*string `json:"versions" yaml:"versions"`
}

// Experimental.
type DependabotOptions struct {
	// You can use the `ignore` option to customize which dependencies are updated.
	//
	// The ignore option supports the following options.
	// Experimental.
	Ignore *[]*DependabotIgnore `json:"ignore" yaml:"ignore"`
	// Ignores updates to `projen`.
	//
	// This is required since projen updates may cause changes in committed files
	// and anti-tamper checks will fail.
	//
	// Projen upgrades are covered through the `ProjenUpgrade` class.
	// Experimental.
	IgnoreProjen *bool `json:"ignoreProjen" yaml:"ignoreProjen"`
	// List of labels to apply to the created PR's.
	// Experimental.
	Labels *[]*string `json:"labels" yaml:"labels"`
	// Map of package registries to use.
	// Experimental.
	Registries *map[string]*DependabotRegistry `json:"registries" yaml:"registries"`
	// How often to check for new versions and raise pull requests.
	// Experimental.
	ScheduleInterval DependabotScheduleInterval `json:"scheduleInterval" yaml:"scheduleInterval"`
	// The strategy to use when edits manifest and lock files.
	// Experimental.
	VersioningStrategy VersioningStrategy `json:"versioningStrategy" yaml:"versioningStrategy"`
}

// Use to add private registry support for dependabot.
// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#configuration-options-for-private-registries
//
// Experimental.
type DependabotRegistry struct {
	// Registry type e.g. 'npm-registry' or 'docker-registry'.
	// Experimental.
	Type DependabotRegistryType `json:"type" yaml:"type"`
	// Url for the registry e.g. 'https://npm.pkg.github.com' or 'registry.hub.docker.com'.
	// Experimental.
	Url *string `json:"url" yaml:"url"`
	// A reference to a Dependabot secret containing an access key for this registry.
	// Experimental.
	Key *string `json:"key" yaml:"key"`
	// Used with the hex-organization registry type.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#hex-organization
	//
	// Experimental.
	Organization *string `json:"organization" yaml:"organization"`
	// A reference to a Dependabot secret containing the password for the specified user.
	// Experimental.
	Password *string `json:"password" yaml:"password"`
	// For registries with type: python-index, if the boolean value is true, pip esolves dependencies by using the specified URL rather than the base URL of the Python Package Index (by default https://pypi.org/simple).
	// Experimental.
	ReplacesBase *bool `json:"replacesBase" yaml:"replacesBase"`
	// Secret token for dependabot access e.g. '${{ secrets.DEPENDABOT_PACKAGE_TOKEN }}'.
	// Experimental.
	Token *string `json:"token" yaml:"token"`
	// The username that Dependabot uses to access the registry.
	// Experimental.
	Username *string `json:"username" yaml:"username"`
}

// Each configuration type requires you to provide particular settings.
//
// Some types allow more than one way to connect
// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#configuration-options-for-private-registries
//
// Experimental.
type DependabotRegistryType string

const (
	DependabotRegistryType_COMPOSER_REGISTRY DependabotRegistryType = "COMPOSER_REGISTRY"
	DependabotRegistryType_DOCKER_REGISTRY DependabotRegistryType = "DOCKER_REGISTRY"
	DependabotRegistryType_GIT DependabotRegistryType = "GIT"
	DependabotRegistryType_HEX_ORGANIZATION DependabotRegistryType = "HEX_ORGANIZATION"
	DependabotRegistryType_MAVEN_REPOSITORY DependabotRegistryType = "MAVEN_REPOSITORY"
	DependabotRegistryType_NPM_REGISTRY DependabotRegistryType = "NPM_REGISTRY"
	DependabotRegistryType_NUGET_FEED DependabotRegistryType = "NUGET_FEED"
	DependabotRegistryType_PYTHON_INDEX DependabotRegistryType = "PYTHON_INDEX"
	DependabotRegistryType_RUBYGEMS_SERVER DependabotRegistryType = "RUBYGEMS_SERVER"
	DependabotRegistryType_TERRAFORM_REGISTRY DependabotRegistryType = "TERRAFORM_REGISTRY"
)

// How often to check for new versions and raise pull requests for version updates.
// Experimental.
type DependabotScheduleInterval string

const (
	DependabotScheduleInterval_DAILY DependabotScheduleInterval = "DAILY"
	DependabotScheduleInterval_WEEKLY DependabotScheduleInterval = "WEEKLY"
	DependabotScheduleInterval_MONTHLY DependabotScheduleInterval = "MONTHLY"
)

// Experimental.
type GitHub interface {
	projen.Component
	Mergify() Mergify
	Project() projen.Project
	ProjenTokenSecret() *string
	Workflows() *[]GithubWorkflow
	WorkflowsEnabled() *bool
	AddDependabot(options *DependabotOptions) Dependabot
	AddPullRequestTemplate(content ...*string) PullRequestTemplate
	AddWorkflow(name *string) GithubWorkflow
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	TryFindWorkflow(name *string) GithubWorkflow
}

// The jsii proxy struct for GitHub
type jsiiProxy_GitHub struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_GitHub) Mergify() Mergify {
	var returns Mergify
	_jsii_.Get(
		j,
		"mergify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) ProjenTokenSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenTokenSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) Workflows() *[]GithubWorkflow {
	var returns *[]GithubWorkflow
	_jsii_.Get(
		j,
		"workflows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHub) WorkflowsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"workflowsEnabled",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHub(project projen.Project, options *GitHubOptions) GitHub {
	_init_.Initialize()

	j := jsiiProxy_GitHub{}

	_jsii_.Create(
		"projen.github.GitHub",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHub_Override(g GitHub, project projen.Project, options *GitHubOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.GitHub",
		[]interface{}{project, options},
		g,
	)
}

// Returns the `GitHub` component of a project or `undefined` if the project does not have a GitHub component.
// Experimental.
func GitHub_Of(project projen.Project) GitHub {
	_init_.Initialize()

	var returns GitHub

	_jsii_.StaticInvoke(
		"projen.github.GitHub",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GitHub) AddDependabot(options *DependabotOptions) Dependabot {
	var returns Dependabot

	_jsii_.Invoke(
		g,
		"addDependabot",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GitHub) AddPullRequestTemplate(content ...*string) PullRequestTemplate {
	args := []interface{}{}
	for _, a := range content {
		args = append(args, a)
	}

	var returns PullRequestTemplate

	_jsii_.Invoke(
		g,
		"addPullRequestTemplate",
		args,
		&returns,
	)

	return returns
}

// Adds a workflow to the project.
//
// Returns: a GithubWorkflow instance
// Experimental.
func (g *jsiiProxy_GitHub) AddWorkflow(name *string) GithubWorkflow {
	var returns GithubWorkflow

	_jsii_.Invoke(
		g,
		"addWorkflow",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (g *jsiiProxy_GitHub) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (g *jsiiProxy_GitHub) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (g *jsiiProxy_GitHub) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

// Finds a GitHub workflow by name.
//
// Returns `undefined` if the workflow cannot be found.
// Experimental.
func (g *jsiiProxy_GitHub) TryFindWorkflow(name *string) GithubWorkflow {
	var returns GithubWorkflow

	_jsii_.Invoke(
		g,
		"tryFindWorkflow",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Experimental.
type GitHubOptions struct {
	// Whether mergify should be enabled on this repository or not.
	// Experimental.
	Mergify *bool `json:"mergify" yaml:"mergify"`
	// Options for Mergify.
	// Experimental.
	MergifyOptions *MergifyOptions `json:"mergifyOptions" yaml:"mergifyOptions"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Experimental.
	ProjenTokenSecret *string `json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// Add a workflow that performs basic checks for pull requests, like validating that PRs follow Conventional Commits.
	// Experimental.
	PullRequestLint *bool `json:"pullRequestLint" yaml:"pullRequestLint"`
	// Options for configuring a pull request linter.
	// Experimental.
	PullRequestLintOptions *PullRequestLintOptions `json:"pullRequestLintOptions" yaml:"pullRequestLintOptions"`
	// Enables GitHub workflows.
	//
	// If this is set to `false`, workflows will not be created.
	// Experimental.
	Workflows *bool `json:"workflows" yaml:"workflows"`
}

// GitHub-based project.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
type GitHubProject interface {
	projen.Project
	AutoApprove() AutoApprove
	BuildTask() projen.Task
	CompileTask() projen.Task
	Components() *[]projen.Component
	DefaultTask() projen.Task
	Deps() projen.Dependencies
	DevContainer() vscode.DevContainer
	Ejected() *bool
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	Logger() projen.Logger
	Name() *string
	Outdir() *string
	PackageTask() projen.Task
	Parent() projen.Project
	PostCompileTask() projen.Task
	PreCompileTask() projen.Task
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Root() projen.Project
	Tasks() projen.Tasks
	TestTask() projen.Task
	Vscode() vscode.VsCode
	AddExcludeFromCleanup(globs ...*string)
	AddGitIgnore(pattern *string)
	AddPackageIgnore(_pattern *string)
	AddTask(name *string, props *projen.TaskOptions) projen.Task
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

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
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

// Adds a .gitignore pattern.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		g,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		g,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		g,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (g *jsiiProxy_GitHubProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		g,
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
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		g,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		g,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Returns the shell command to execute in order to run a task.
//
// By default, this is `npx projen@<version> <task>`
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) Synth() {
	_jsii_.InvokeVoid(
		g,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		g,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
func (g *jsiiProxy_GitHubProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		g,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Deprecated: This is a *temporary* class. At the moment, our base project
// types such as `NodeProject` and `JavaProject` are derived from this, but we
// want to be able to use these project types outside of GitHub as well. One of
// the next steps to address this is to abstract workflows so that different
// "engines" can be used to implement our CI/CD solutions.
func (g *jsiiProxy_GitHubProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		g,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Options for `GitHubProject`.
// Experimental.
type GitHubProjectOptions struct {
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
	AutoApproveOptions *AutoApproveOptions `json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *AutoMergeOptions `json:"autoMergeOptions" yaml:"autoMergeOptions"`
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
	GithubOptions *GitHubOptions `json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *MergifyOptions `json:"mergifyOptions" yaml:"mergifyOptions"`
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
	StaleOptions *StaleOptions `json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode" yaml:"vscode"`
}

// Represents the git identity.
// Experimental.
type GitIdentity struct {
	// The email address of the git user.
	// Experimental.
	Email *string `json:"email" yaml:"email"`
	// The name of the user.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// Workflow for GitHub.
//
// A workflow is a configurable automated process made up of one or more jobs.
// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions
//
// Experimental.
type GithubWorkflow interface {
	projen.Component
	Concurrency() *string
	File() projen.YamlFile
	Name() *string
	Project() projen.Project
	ProjenTokenSecret() *string
	AddJob(id *string, job *workflows.Job)
	AddJobs(jobs *map[string]*workflows.Job)
	On(events *workflows.Triggers)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for GithubWorkflow
type jsiiProxy_GithubWorkflow struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_GithubWorkflow) Concurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) ProjenTokenSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenTokenSecret",
		&returns,
	)
	return returns
}


// Experimental.
func NewGithubWorkflow(github GitHub, name *string, options *GithubWorkflowOptions) GithubWorkflow {
	_init_.Initialize()

	j := jsiiProxy_GithubWorkflow{}

	_jsii_.Create(
		"projen.github.GithubWorkflow",
		[]interface{}{github, name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGithubWorkflow_Override(g GithubWorkflow, github GitHub, name *string, options *GithubWorkflowOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.GithubWorkflow",
		[]interface{}{github, name, options},
		g,
	)
}

// Adds a single job to the workflow.
// Experimental.
func (g *jsiiProxy_GithubWorkflow) AddJob(id *string, job *workflows.Job) {
	_jsii_.InvokeVoid(
		g,
		"addJob",
		[]interface{}{id, job},
	)
}

// Add jobs to the workflow.
// Experimental.
func (g *jsiiProxy_GithubWorkflow) AddJobs(jobs *map[string]*workflows.Job) {
	_jsii_.InvokeVoid(
		g,
		"addJobs",
		[]interface{}{jobs},
	)
}

// Add events to triggers the workflow.
// Experimental.
func (g *jsiiProxy_GithubWorkflow) On(events *workflows.Triggers) {
	_jsii_.InvokeVoid(
		g,
		"on",
		[]interface{}{events},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (g *jsiiProxy_GithubWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (g *jsiiProxy_GithubWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (g *jsiiProxy_GithubWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `GithubWorkflow`.
// Experimental.
type GithubWorkflowOptions struct {
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// Currently in beta.
	// See: https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#concurrency
	//
	// Experimental.
	Concurrency *string `json:"concurrency" yaml:"concurrency"`
	// Force the creation of the workflow even if `workflows` is disabled in `GitHub`.
	// Experimental.
	Force *bool `json:"force" yaml:"force"`
}

// Experimental.
type IAddConditionsLater interface {
	// Experimental.
	Render() *[]*string
}

// The jsii proxy for IAddConditionsLater
type jsiiProxy_IAddConditionsLater struct {
	_ byte // padding
}

func (i *jsiiProxy_IAddConditionsLater) Render() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type IJobProvider interface {
	// Generates a collection of named GitHub workflow jobs.
	// Experimental.
	RenderJobs() *map[string]*workflows.Job
}

// The jsii proxy for IJobProvider
type jsiiProxy_IJobProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IJobProvider) RenderJobs() *map[string]*workflows.Job {
	var returns *map[string]*workflows.Job

	_jsii_.Invoke(
		i,
		"renderJobs",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type Mergify interface {
	projen.Component
	Project() projen.Project
	AddQueue(queue *MergifyQueue)
	AddRule(rule *MergifyRule)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Mergify
type jsiiProxy_Mergify struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Mergify) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMergify(github GitHub, options *MergifyOptions) Mergify {
	_init_.Initialize()

	j := jsiiProxy_Mergify{}

	_jsii_.Create(
		"projen.github.Mergify",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMergify_Override(m Mergify, github GitHub, options *MergifyOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Mergify",
		[]interface{}{github, options},
		m,
	)
}

// Experimental.
func (m *jsiiProxy_Mergify) AddQueue(queue *MergifyQueue) {
	_jsii_.InvokeVoid(
		m,
		"addQueue",
		[]interface{}{queue},
	)
}

// Experimental.
func (m *jsiiProxy_Mergify) AddRule(rule *MergifyRule) {
	_jsii_.InvokeVoid(
		m,
		"addRule",
		[]interface{}{rule},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (m *jsiiProxy_Mergify) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (m *jsiiProxy_Mergify) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (m *jsiiProxy_Mergify) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

// The Mergify conditional operators that can be used are: `or` and `and`.
//
// Note: The number of nested conditions is limited to 3.
// See: https://docs.mergify.io/conditions/#combining-conditions-with-operators
//
// Experimental.
type MergifyConditionalOperator struct {
	// Experimental.
	And *[]interface{} `json:"and" yaml:"and"`
	// Experimental.
	Or *[]interface{} `json:"or" yaml:"or"`
}

// Experimental.
type MergifyOptions struct {
	// Experimental.
	Queues *[]*MergifyQueue `json:"queues" yaml:"queues"`
	// Experimental.
	Rules *[]*MergifyRule `json:"rules" yaml:"rules"`
}

// Experimental.
type MergifyQueue struct {
	// A list of Conditions string that must match against the pull request for the pull request to be added to the queue.
	// See: https://docs.mergify.com/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `json:"conditions" yaml:"conditions"`
	// The name of the queue.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// Experimental.
type MergifyRule struct {
	// A dictionary made of Actions that will be executed on the matching pull requests.
	// See: https://docs.mergify.io/actions/#actions
	//
	// Experimental.
	Actions *map[string]interface{} `json:"actions" yaml:"actions"`
	// A list of Conditions string that must match against the pull request for the rule to be applied.
	// See: https://docs.mergify.io/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `json:"conditions" yaml:"conditions"`
	// The name of the rule.
	//
	// This is not used by the engine directly,
	// but is used when reporting information about a rule.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// Configure validations to run on GitHub pull requests.
//
// Only generates a file if at least one linter is configured.
// Experimental.
type PullRequestLint interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for PullRequestLint
type jsiiProxy_PullRequestLint struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PullRequestLint) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPullRequestLint(github GitHub, options *PullRequestLintOptions) PullRequestLint {
	_init_.Initialize()

	j := jsiiProxy_PullRequestLint{}

	_jsii_.Create(
		"projen.github.PullRequestLint",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPullRequestLint_Override(p PullRequestLint, github GitHub, options *PullRequestLintOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.PullRequestLint",
		[]interface{}{github, options},
		p,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_PullRequestLint) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_PullRequestLint) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (p *jsiiProxy_PullRequestLint) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for PullRequestLint.
// Experimental.
type PullRequestLintOptions struct {
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// Validate that pull request titles follow Conventional Commits.
	// See: https://www.conventionalcommits.org/
	//
	// Experimental.
	SemanticTitle *bool `json:"semanticTitle" yaml:"semanticTitle"`
	// Options for validating the conventional commit title linter.
	// Experimental.
	SemanticTitleOptions *SemanticTitleOptions `json:"semanticTitleOptions" yaml:"semanticTitleOptions"`
}

// Template for GitHub pull requests.
// Experimental.
type PullRequestTemplate interface {
	projen.TextFile
	AbsolutePath() *string
	Changed() *bool
	Executable() *bool
	SetExecutable(val *bool)
	Marker() *string
	Path() *string
	Project() projen.Project
	Readonly() *bool
	SetReadonly(val *bool)
	AddLine(line *string)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
	SynthesizeContent(_arg projen.IResolver) *string
}

// The jsii proxy struct for PullRequestTemplate
type jsiiProxy_PullRequestTemplate struct {
	internal.Type__projenTextFile
}

func (j *jsiiProxy_PullRequestTemplate) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewPullRequestTemplate(github GitHub, options *PullRequestTemplateOptions) PullRequestTemplate {
	_init_.Initialize()

	j := jsiiProxy_PullRequestTemplate{}

	_jsii_.Create(
		"projen.github.PullRequestTemplate",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPullRequestTemplate_Override(p PullRequestTemplate, github GitHub, options *PullRequestTemplateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.PullRequestTemplate",
		[]interface{}{github, options},
		p,
	)
}

func (j *jsiiProxy_PullRequestTemplate) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_PullRequestTemplate) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

// Adds a line to the text file.
// Experimental.
func (p *jsiiProxy_PullRequestTemplate) AddLine(line *string) {
	_jsii_.InvokeVoid(
		p,
		"addLine",
		[]interface{}{line},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_PullRequestTemplate) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_PullRequestTemplate) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Writes the file to the project's output directory.
// Experimental.
func (p *jsiiProxy_PullRequestTemplate) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Implemented by derived classes and returns the contents of the file to emit.
// Experimental.
func (p *jsiiProxy_PullRequestTemplate) SynthesizeContent(_arg projen.IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// Options for `PullRequestTemplate`.
// Experimental.
type PullRequestTemplateOptions struct {
	// The contents of the template.
	//
	// You can use `addLine()` to add additional lines.
	// Experimental.
	Lines *[]*string `json:"lines" yaml:"lines"`
}

// Options for linting that PR titles follow Conventional Commits.
// See: https://www.conventionalcommits.org/
//
// Experimental.
type SemanticTitleOptions struct {
	// Configure that a scope must always be provided.
	//
	// e.g. feat(ui), fix(core)
	// Experimental.
	RequireScope *bool `json:"requireScope" yaml:"requireScope"`
	// Configure a list of commit types that are allowed.
	// Experimental.
	Types *[]*string `json:"types" yaml:"types"`
}

// Warns and then closes issues and PRs that have had no activity for a specified amount of time.
//
// The default configuration will:
//
//   * Add a "Stale" label to pull requests after 14 days and closed after 2 days
//   * Add a "Stale" label to issues after 60 days and closed after 7 days
//   * If a comment is added, the label will be removed and timer is restarted.
// See: https://github.com/actions/stale
//
// Experimental.
type Stale interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Stale
type jsiiProxy_Stale struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Stale) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewStale(github GitHub, options *StaleOptions) Stale {
	_init_.Initialize()

	j := jsiiProxy_Stale{}

	_jsii_.Create(
		"projen.github.Stale",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewStale_Override(s Stale, github GitHub, options *StaleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Stale",
		[]interface{}{github, options},
		s,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (s *jsiiProxy_Stale) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (s *jsiiProxy_Stale) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (s *jsiiProxy_Stale) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

// Stale behavior.
// Experimental.
type StaleBehavior struct {
	// The comment to add to the issue/PR when it's closed.
	// Experimental.
	CloseMessage *string `json:"closeMessage" yaml:"closeMessage"`
	// Days until the issue/PR is closed after it is marked as "Stale".
	//
	// Set to -1 to disable.
	// Experimental.
	DaysBeforeClose *float64 `json:"daysBeforeClose" yaml:"daysBeforeClose"`
	// How many days until the issue or pull request is marked as "Stale".
	//
	// Set to -1 to disable.
	// Experimental.
	DaysBeforeStale *float64 `json:"daysBeforeStale" yaml:"daysBeforeStale"`
	// Determines if this behavior is enabled.
	//
	// Same as setting `daysBeforeStale` and `daysBeforeClose` to `-1`.
	// Experimental.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// Label which exempt an issue/PR from becoming stale.
	//
	// Set to `[]` to disable.
	// Experimental.
	ExemptLabels *[]*string `json:"exemptLabels" yaml:"exemptLabels"`
	// The label to apply to the issue/PR when it becomes stale.
	// Experimental.
	StaleLabel *string `json:"staleLabel" yaml:"staleLabel"`
	// The comment to add to the issue/PR when it becomes stale.
	// Experimental.
	StaleMessage *string `json:"staleMessage" yaml:"staleMessage"`
}

// Options for `Stale`.
// Experimental.
type StaleOptions struct {
	// How to handle stale issues.
	// Experimental.
	Issues *StaleBehavior `json:"issues" yaml:"issues"`
	// How to handle stale pull requests.
	// Experimental.
	PullRequest *StaleBehavior `json:"pullRequest" yaml:"pullRequest"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
}

// A GitHub workflow for common build tasks within a project.
// Experimental.
type TaskWorkflow interface {
	GithubWorkflow
	ArtifactsDirectory() *string
	Concurrency() *string
	File() projen.YamlFile
	JobId() *string
	Name() *string
	Project() projen.Project
	ProjenTokenSecret() *string
	AddJob(id *string, job *workflows.Job)
	AddJobs(jobs *map[string]*workflows.Job)
	On(events *workflows.Triggers)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for TaskWorkflow
type jsiiProxy_TaskWorkflow struct {
	jsiiProxy_GithubWorkflow
}

func (j *jsiiProxy_TaskWorkflow) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Concurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) ProjenTokenSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenTokenSecret",
		&returns,
	)
	return returns
}


// Experimental.
func NewTaskWorkflow(github GitHub, options *TaskWorkflowOptions) TaskWorkflow {
	_init_.Initialize()

	j := jsiiProxy_TaskWorkflow{}

	_jsii_.Create(
		"projen.github.TaskWorkflow",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTaskWorkflow_Override(t TaskWorkflow, github GitHub, options *TaskWorkflowOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.TaskWorkflow",
		[]interface{}{github, options},
		t,
	)
}

// Adds a single job to the workflow.
// Experimental.
func (t *jsiiProxy_TaskWorkflow) AddJob(id *string, job *workflows.Job) {
	_jsii_.InvokeVoid(
		t,
		"addJob",
		[]interface{}{id, job},
	)
}

// Add jobs to the workflow.
// Experimental.
func (t *jsiiProxy_TaskWorkflow) AddJobs(jobs *map[string]*workflows.Job) {
	_jsii_.InvokeVoid(
		t,
		"addJobs",
		[]interface{}{jobs},
	)
}

// Add events to triggers the workflow.
// Experimental.
func (t *jsiiProxy_TaskWorkflow) On(events *workflows.Triggers) {
	_jsii_.InvokeVoid(
		t,
		"on",
		[]interface{}{events},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (t *jsiiProxy_TaskWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (t *jsiiProxy_TaskWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (t *jsiiProxy_TaskWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type TaskWorkflowOptions struct {
	// The workflow name.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Permissions for the build job.
	// Experimental.
	Permissions *workflows.JobPermissions `json:"permissions" yaml:"permissions"`
	// The main task to be executed.
	// Experimental.
	Task projen.Task `json:"task" yaml:"task"`
	// A directory name which contains artifacts to be uploaded (e.g. `dist`). If this is set, the contents of this directory will be uploaded as an artifact at the end of the workflow run, even if other steps fail.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Override for the `with` property of the source code checkout step.
	// Experimental.
	CheckoutWith *map[string]interface{} `json:"checkoutWith" yaml:"checkoutWith"`
	// Adds an 'if' condition to the workflow.
	// Experimental.
	Condition *string `json:"condition" yaml:"condition"`
	// Experimental.
	Container *workflows.ContainerOptions `json:"container" yaml:"container"`
	// Workflow environment variables.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// The git identity to use in this workflow.
	// Experimental.
	GitIdentity *GitIdentity `json:"gitIdentity" yaml:"gitIdentity"`
	// The primary job id.
	// Experimental.
	JobId *string `json:"jobId" yaml:"jobId"`
	// Mapping of job output names to values/expressions.
	// Experimental.
	Outputs *map[string]*workflows.JobStepOutput `json:"outputs" yaml:"outputs"`
	// Actions to run after the main build step.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps" yaml:"postBuildSteps"`
	// Steps to run before the main build step.
	// Experimental.
	PreBuildSteps *[]*workflows.JobStep `json:"preBuildSteps" yaml:"preBuildSteps"`
	// Initial steps to run before the source code checkout.
	// Experimental.
	PreCheckoutSteps *[]*workflows.JobStep `json:"preCheckoutSteps" yaml:"preCheckoutSteps"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// The triggers for the workflow.
	// Experimental.
	Triggers *workflows.Triggers `json:"triggers" yaml:"triggers"`
}

// The strategy to use when edits manifest and lock files.
// Experimental.
type VersioningStrategy string

const (
	VersioningStrategy_LOCKFILE_ONLY VersioningStrategy = "LOCKFILE_ONLY"
	VersioningStrategy_AUTO VersioningStrategy = "AUTO"
	VersioningStrategy_WIDEN VersioningStrategy = "WIDEN"
	VersioningStrategy_INCREASE VersioningStrategy = "INCREASE"
	VersioningStrategy_INCREASE_IF_NECESSARY VersioningStrategy = "INCREASE_IF_NECESSARY"
)

