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
	// Experimental.
	Label() *string
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

func (a *jsiiProxy_AutoApprove) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoApprove) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	AllowedUsernames *[]*string `field:"optional" json:"allowedUsernames" yaml:"allowedUsernames"`
	// Only pull requests with this label will be auto-approved.
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
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
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
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
	// Experimental.
	Project() projen.Project
	// Adds conditions to the auto merge rule.
	// Experimental.
	AddConditions(conditions ...*string)
	// Adds conditions that will be rendered only during synthesis.
	// Experimental.
	AddConditionsLater(later IAddConditionsLater)
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

func (a *jsiiProxy_AutoMerge) AddConditionsLater(later IAddConditionsLater) {
	_jsii_.InvokeVoid(
		a,
		"addConditionsLater",
		[]interface{}{later},
	)
}

func (a *jsiiProxy_AutoMerge) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	ApprovedReviews *float64 `field:"optional" json:"approvedReviews" yaml:"approvedReviews"`
	// List of labels that will prevent auto-merging.
	// Experimental.
	BlockingLabels *[]*string `field:"optional" json:"blockingLabels" yaml:"blockingLabels"`
}

// Defines dependabot configuration for node projects.
//
// Since module versions are managed in projen, the versioning strategy will be
// configured to "lockfile-only" which means that only updates that can be done
// on the lockfile itself will be proposed.
// Experimental.
type Dependabot interface {
	projen.Component
	// The raw dependabot configuration.
	// See: https://docs.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates
	//
	// Experimental.
	Config() interface{}
	// Whether or not projen is also upgraded in this config,.
	// Experimental.
	IgnoresProjen() *bool
	// Experimental.
	Project() projen.Project
	// Ignores a dependency from automatic updates.
	// Experimental.
	AddIgnore(dependencyName *string, versions ...*string)
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

func (d *jsiiProxy_Dependabot) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependabot) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	DependencyName *string `field:"required" json:"dependencyName" yaml:"dependencyName"`
	// Use to ignore specific versions or ranges of versions.
	//
	// If you want to
	// define a range, use the standard pattern for the package manager (for
	// example: `^1.0.0` for npm, or `~> 2.0` for Bundler).
	// Experimental.
	Versions *[]*string `field:"optional" json:"versions" yaml:"versions"`
}

// Experimental.
type DependabotOptions struct {
	// You can use the `ignore` option to customize which dependencies are updated.
	//
	// The ignore option supports the following options.
	// Experimental.
	Ignore *[]*DependabotIgnore `field:"optional" json:"ignore" yaml:"ignore"`
	// Ignores updates to `projen`.
	//
	// This is required since projen updates may cause changes in committed files
	// and anti-tamper checks will fail.
	//
	// Projen upgrades are covered through the `ProjenUpgrade` class.
	// Experimental.
	IgnoreProjen *bool `field:"optional" json:"ignoreProjen" yaml:"ignoreProjen"`
	// List of labels to apply to the created PR's.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Map of package registries to use.
	// Experimental.
	Registries *map[string]*DependabotRegistry `field:"optional" json:"registries" yaml:"registries"`
	// How often to check for new versions and raise pull requests.
	// Experimental.
	ScheduleInterval DependabotScheduleInterval `field:"optional" json:"scheduleInterval" yaml:"scheduleInterval"`
	// The strategy to use when edits manifest and lock files.
	// Experimental.
	VersioningStrategy VersioningStrategy `field:"optional" json:"versioningStrategy" yaml:"versioningStrategy"`
}

// Use to add private registry support for dependabot.
// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#configuration-options-for-private-registries
//
// Experimental.
type DependabotRegistry struct {
	// Registry type e.g. 'npm-registry' or 'docker-registry'.
	// Experimental.
	Type DependabotRegistryType `field:"required" json:"type" yaml:"type"`
	// Url for the registry e.g. 'https://npm.pkg.github.com' or 'registry.hub.docker.com'.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// A reference to a Dependabot secret containing an access key for this registry.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Used with the hex-organization registry type.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#hex-organization
	//
	// Experimental.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A reference to a Dependabot secret containing the password for the specified user.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// For registries with type: python-index, if the boolean value is true, pip esolves dependencies by using the specified URL rather than the base URL of the Python Package Index (by default https://pypi.org/simple).
	// Experimental.
	ReplacesBase *bool `field:"optional" json:"replacesBase" yaml:"replacesBase"`
	// Secret token for dependabot access e.g. '${{ secrets.DEPENDABOT_PACKAGE_TOKEN }}'.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username that Dependabot uses to access the registry.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

// Each configuration type requires you to provide particular settings.
//
// Some types allow more than one way to connect.
// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#configuration-options-for-private-registries
//
// Experimental.
type DependabotRegistryType string

const (
	// The composer-repository type supports username and password.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#composer-repository
	//
	// Experimental.
	DependabotRegistryType_COMPOSER_REGISTRY DependabotRegistryType = "COMPOSER_REGISTRY"
	// The docker-registry type supports username and password.
	//
	// The docker-registry type can also be used to pull from Amazon ECR using static AWS credentials.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#docker-registry
	//
	// Experimental.
	DependabotRegistryType_DOCKER_REGISTRY DependabotRegistryType = "DOCKER_REGISTRY"
	// The git type supports username and password.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#git
	//
	// Experimental.
	DependabotRegistryType_GIT DependabotRegistryType = "GIT"
	// The hex-organization type supports organization and key.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#hex-organization
	//
	// Experimental.
	DependabotRegistryType_HEX_ORGANIZATION DependabotRegistryType = "HEX_ORGANIZATION"
	// The maven-repository type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#maven-repository
	//
	// Experimental.
	DependabotRegistryType_MAVEN_REPOSITORY DependabotRegistryType = "MAVEN_REPOSITORY"
	// The npm-registry type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#npm-registry
	//
	// Experimental.
	DependabotRegistryType_NPM_REGISTRY DependabotRegistryType = "NPM_REGISTRY"
	// The nuget-feed type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#nuget-feed
	//
	// Experimental.
	DependabotRegistryType_NUGET_FEED DependabotRegistryType = "NUGET_FEED"
	// The python-index type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#python-index
	//
	// Experimental.
	DependabotRegistryType_PYTHON_INDEX DependabotRegistryType = "PYTHON_INDEX"
	// The rubygems-server type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#rubygems-server
	//
	// Experimental.
	DependabotRegistryType_RUBYGEMS_SERVER DependabotRegistryType = "RUBYGEMS_SERVER"
	// The terraform-registry type supports a token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#terraform-registry
	//
	// Experimental.
	DependabotRegistryType_TERRAFORM_REGISTRY DependabotRegistryType = "TERRAFORM_REGISTRY"
)

// How often to check for new versions and raise pull requests for version updates.
// Experimental.
type DependabotScheduleInterval string

const (
	// Runs on every weekday, Monday to Friday.
	// Experimental.
	DependabotScheduleInterval_DAILY DependabotScheduleInterval = "DAILY"
	// Runs once each week.
	//
	// By default, this is on Monday.
	// Experimental.
	DependabotScheduleInterval_WEEKLY DependabotScheduleInterval = "WEEKLY"
	// Runs once each month.
	//
	// This is on the first day of the month.
	// Experimental.
	DependabotScheduleInterval_MONTHLY DependabotScheduleInterval = "MONTHLY"
)

// Experimental.
type GitHub interface {
	projen.Component
	// The `Mergify` configured on this repository.
	//
	// This is `undefined` if Mergify
	// was not enabled when creating the repository.
	// Experimental.
	Mergify() Mergify
	// Experimental.
	Project() projen.Project
	// GitHub API authentication method used by projen workflows.
	// Experimental.
	ProjenCredentials() GithubCredentials
	// All workflows.
	// Experimental.
	Workflows() *[]GithubWorkflow
	// Are workflows enabled?
	// Experimental.
	WorkflowsEnabled() *bool
	// Experimental.
	AddDependabot(options *DependabotOptions) Dependabot
	// Experimental.
	AddPullRequestTemplate(content ...*string) PullRequestTemplate
	// Adds a workflow to the project.
	//
	// Returns: a GithubWorkflow instance.
	// Experimental.
	AddWorkflow(name *string) GithubWorkflow
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
	// Finds a GitHub workflow by name.
	//
	// Returns `undefined` if the workflow cannot be found.
	// Experimental.
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

func (j *jsiiProxy_GitHub) ProjenCredentials() GithubCredentials {
	var returns GithubCredentials
	_jsii_.Get(
		j,
		"projenCredentials",
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

func (g *jsiiProxy_GitHub) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHub) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitHub) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

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
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for Mergify.
	// Experimental.
	MergifyOptions *MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: - use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// Add a workflow that performs basic checks for pull requests, like validating that PRs follow Conventional Commits.
	// Experimental.
	PullRequestLint *bool `field:"optional" json:"pullRequestLint" yaml:"pullRequestLint"`
	// Options for configuring a pull request linter.
	// Experimental.
	PullRequestLintOptions *PullRequestLintOptions `field:"optional" json:"pullRequestLintOptions" yaml:"pullRequestLintOptions"`
	// Enables GitHub workflows.
	//
	// If this is set to `false`, workflows will not be created.
	// Experimental.
	Workflows *bool `field:"optional" json:"workflows" yaml:"workflows"`
}

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
	_jsii_.InvokeVoid(
		g,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (g *jsiiProxy_GitHubProject) AddPackageIgnore(_pattern *string) {
	_jsii_.InvokeVoid(
		g,
		"addPackageIgnore",
		[]interface{}{_pattern},
	)
}

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

func (g *jsiiProxy_GitHubProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		g,
		"addTip",
		[]interface{}{message},
	)
}

func (g *jsiiProxy_GitHubProject) AnnotateGenerated(glob *string) {
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
	var returns projen.FileBase

	_jsii_.Invoke(
		g,
		"tryRemoveFile",
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
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to commit the managed files by default.
	// Experimental.
	CommitGenerated *bool `field:"optional" json:"commitGenerated" yaml:"commitGenerated"`
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
	AutoApproveOptions *AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
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
	AutoMergeOptions *AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
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
	GithubOptions *GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
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
	StaleOptions *StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
}

// Represents the git identity.
// Experimental.
type GitIdentity struct {
	// The email address of the git user.
	// Experimental.
	Email *string `field:"required" json:"email" yaml:"email"`
	// The name of the user.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Represents a method of providing GitHub API access for projen workflows.
// Experimental.
type GithubCredentials interface {
	// Setup steps to obtain GitHub credentials.
	// Experimental.
	SetupSteps() *[]*workflows.JobStep
	// The value to use in a workflow when a GitHub token is expected.
	//
	// This
	// typically looks like "${{ some.path.to.a.value }}".
	// Experimental.
	TokenRef() *string
}

// The jsii proxy struct for GithubCredentials
type jsiiProxy_GithubCredentials struct {
	_ byte // padding
}

func (j *jsiiProxy_GithubCredentials) SetupSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"setupSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubCredentials) TokenRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenRef",
		&returns,
	)
	return returns
}


// Provide API access through a GitHub App.
//
// The GitHub App must be installed on the GitHub repo, its App ID and a
// private key must be added as secrets to the repo. The name of the secrets
// can be specified here.
// See: https://docs.github.com/en/developers/apps/building-github-apps/creating-a-github-app
//
// Experimental.
func GithubCredentials_FromApp(options *GithubCredentialsAppOptions) GithubCredentials {
	_init_.Initialize()

	var returns GithubCredentials

	_jsii_.StaticInvoke(
		"projen.github.GithubCredentials",
		"fromApp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Provide API access through a GitHub personal access token.
//
// The token must be added as a secret to the GitHub repo, and the name of the
// secret can be specified here.
// See: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
//
// Experimental.
func GithubCredentials_FromPersonalAccessToken(options *GithubCredentialsPersonalAccessTokenOptions) GithubCredentials {
	_init_.Initialize()

	var returns GithubCredentials

	_jsii_.StaticInvoke(
		"projen.github.GithubCredentials",
		"fromPersonalAccessToken",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Options for `GithubCredentials.fromApp`.
// Experimental.
type GithubCredentialsAppOptions struct {
	// Experimental.
	AppIdSecret *string `field:"optional" json:"appIdSecret" yaml:"appIdSecret"`
	// Experimental.
	PrivateKeySecret *string `field:"optional" json:"privateKeySecret" yaml:"privateKeySecret"`
}

// Options for `GithubCredentials.fromPersonalAccessToken`.
// Experimental.
type GithubCredentialsPersonalAccessTokenOptions struct {
	// Experimental.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

// Workflow for GitHub.
//
// A workflow is a configurable automated process made up of one or more jobs.
// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions
//
// Experimental.
type GithubWorkflow interface {
	projen.Component
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	// Experimental.
	Concurrency() *string
	// The workflow YAML file.
	//
	// May not exist if `workflowsEnabled` is false on `GitHub`.
	// Experimental.
	File() projen.YamlFile
	// The name of the workflow.
	// Experimental.
	Name() *string
	// Experimental.
	Project() projen.Project
	// GitHub API authentication method used by projen workflows.
	// Experimental.
	ProjenCredentials() GithubCredentials
	// Adds a single job to the workflow.
	// Experimental.
	AddJob(id *string, job interface{})
	// Add jobs to the workflow.
	// Experimental.
	AddJobs(jobs *map[string]interface{})
	// Add events to triggers the workflow.
	// Experimental.
	On(events *workflows.Triggers)
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

func (j *jsiiProxy_GithubWorkflow) ProjenCredentials() GithubCredentials {
	var returns GithubCredentials
	_jsii_.Get(
		j,
		"projenCredentials",
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

func (g *jsiiProxy_GithubWorkflow) AddJob(id *string, job interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addJob",
		[]interface{}{id, job},
	)
}

func (g *jsiiProxy_GithubWorkflow) AddJobs(jobs *map[string]interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (g *jsiiProxy_GithubWorkflow) On(events *workflows.Triggers) {
	_jsii_.InvokeVoid(
		g,
		"on",
		[]interface{}{events},
	)
}

func (g *jsiiProxy_GithubWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	Concurrency *string `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Force the creation of the workflow even if `workflows` is disabled in `GitHub`.
	// Experimental.
	Force *bool `field:"optional" json:"force" yaml:"force"`
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
	// Experimental.
	Project() projen.Project
	// Experimental.
	AddQueue(queue *MergifyQueue)
	// Experimental.
	AddRule(rule *MergifyRule)
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

func (m *jsiiProxy_Mergify) AddQueue(queue *MergifyQueue) {
	_jsii_.InvokeVoid(
		m,
		"addQueue",
		[]interface{}{queue},
	)
}

func (m *jsiiProxy_Mergify) AddRule(rule *MergifyRule) {
	_jsii_.InvokeVoid(
		m,
		"addRule",
		[]interface{}{rule},
	)
}

func (m *jsiiProxy_Mergify) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mergify) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	And *[]interface{} `field:"optional" json:"and" yaml:"and"`
	// Experimental.
	Or *[]interface{} `field:"optional" json:"or" yaml:"or"`
}

// Experimental.
type MergifyOptions struct {
	// Experimental.
	Queues *[]*MergifyQueue `field:"optional" json:"queues" yaml:"queues"`
	// Experimental.
	Rules *[]*MergifyRule `field:"optional" json:"rules" yaml:"rules"`
}

// Experimental.
type MergifyQueue struct {
	// A list of Conditions string that must match against the pull request for the pull request to be added to the queue.
	// See: https://docs.mergify.com/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The name of the queue.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Experimental.
type MergifyRule struct {
	// A dictionary made of Actions that will be executed on the matching pull requests.
	// See: https://docs.mergify.io/actions/#actions
	//
	// Experimental.
	Actions *map[string]interface{} `field:"required" json:"actions" yaml:"actions"`
	// A list of Conditions string that must match against the pull request for the rule to be applied.
	// See: https://docs.mergify.io/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The name of the rule.
	//
	// This is not used by the engine directly,
	// but is used when reporting information about a rule.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Configure validations to run on GitHub pull requests.
//
// Only generates a file if at least one linter is configured.
// Experimental.
type PullRequestLint interface {
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

func (p *jsiiProxy_PullRequestLint) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestLint) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Validate that pull request titles follow Conventional Commits.
	// See: https://www.conventionalcommits.org/
	//
	// Experimental.
	SemanticTitle *bool `field:"optional" json:"semanticTitle" yaml:"semanticTitle"`
	// Options for validating the conventional commit title linter.
	// Experimental.
	SemanticTitleOptions *SemanticTitleOptions `field:"optional" json:"semanticTitleOptions" yaml:"semanticTitleOptions"`
}

// Template for GitHub pull requests.
// Experimental.
type PullRequestTemplate interface {
	projen.TextFile
	// The absolute path of this file.
	// Experimental.
	AbsolutePath() *string
	// Indicates if the file has been changed during synthesis.
	//
	// This property is
	// only available in `postSynthesize()` hooks. If this is `undefined`, the
	// file has not been synthesized yet.
	// Experimental.
	Changed() *bool
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable() *bool
	// Experimental.
	SetExecutable(val *bool)
	// The projen marker, used to identify files as projen-generated.
	//
	// Value is undefined if the project is being ejected.
	// Experimental.
	Marker() *string
	// The file path, relative to the project root.
	// Experimental.
	Path() *string
	// Experimental.
	Project() projen.Project
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly() *bool
	// Experimental.
	SetReadonly(val *bool)
	// Adds a line to the text file.
	// Experimental.
	AddLine(line *string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Writes the file to the project's output directory.
	// Experimental.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Experimental.
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

func (p *jsiiProxy_PullRequestTemplate) AddLine(line *string) {
	_jsii_.InvokeVoid(
		p,
		"addLine",
		[]interface{}{line},
	)
}

func (p *jsiiProxy_PullRequestTemplate) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestTemplate) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestTemplate) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

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
	Lines *[]*string `field:"optional" json:"lines" yaml:"lines"`
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
	RequireScope *bool `field:"optional" json:"requireScope" yaml:"requireScope"`
	// Configure a list of commit types that are allowed.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
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

func (s *jsiiProxy_Stale) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stale) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	CloseMessage *string `field:"optional" json:"closeMessage" yaml:"closeMessage"`
	// Days until the issue/PR is closed after it is marked as "Stale".
	//
	// Set to -1 to disable.
	// Experimental.
	DaysBeforeClose *float64 `field:"optional" json:"daysBeforeClose" yaml:"daysBeforeClose"`
	// How many days until the issue or pull request is marked as "Stale".
	//
	// Set to -1 to disable.
	// Experimental.
	DaysBeforeStale *float64 `field:"optional" json:"daysBeforeStale" yaml:"daysBeforeStale"`
	// Determines if this behavior is enabled.
	//
	// Same as setting `daysBeforeStale` and `daysBeforeClose` to `-1`.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Label which exempt an issue/PR from becoming stale.
	//
	// Set to `[]` to disable.
	// Experimental.
	ExemptLabels *[]*string `field:"optional" json:"exemptLabels" yaml:"exemptLabels"`
	// The label to apply to the issue/PR when it becomes stale.
	// Experimental.
	StaleLabel *string `field:"optional" json:"staleLabel" yaml:"staleLabel"`
	// The comment to add to the issue/PR when it becomes stale.
	// Experimental.
	StaleMessage *string `field:"optional" json:"staleMessage" yaml:"staleMessage"`
}

// Options for `Stale`.
// Experimental.
type StaleOptions struct {
	// How to handle stale issues.
	// Experimental.
	Issues *StaleBehavior `field:"optional" json:"issues" yaml:"issues"`
	// How to handle stale pull requests.
	// Experimental.
	PullRequest *StaleBehavior `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
}

// A GitHub workflow for common build tasks within a project.
// Experimental.
type TaskWorkflow interface {
	GithubWorkflow
	// Experimental.
	ArtifactsDirectory() *string
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	// Experimental.
	Concurrency() *string
	// The workflow YAML file.
	//
	// May not exist if `workflowsEnabled` is false on `GitHub`.
	// Experimental.
	File() projen.YamlFile
	// Experimental.
	JobId() *string
	// The name of the workflow.
	// Experimental.
	Name() *string
	// Experimental.
	Project() projen.Project
	// GitHub API authentication method used by projen workflows.
	// Experimental.
	ProjenCredentials() GithubCredentials
	// Adds a single job to the workflow.
	// Experimental.
	AddJob(id *string, job interface{})
	// Add jobs to the workflow.
	// Experimental.
	AddJobs(jobs *map[string]interface{})
	// Add events to triggers the workflow.
	// Experimental.
	On(events *workflows.Triggers)
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

func (j *jsiiProxy_TaskWorkflow) ProjenCredentials() GithubCredentials {
	var returns GithubCredentials
	_jsii_.Get(
		j,
		"projenCredentials",
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

func (t *jsiiProxy_TaskWorkflow) AddJob(id *string, job interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addJob",
		[]interface{}{id, job},
	)
}

func (t *jsiiProxy_TaskWorkflow) AddJobs(jobs *map[string]interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (t *jsiiProxy_TaskWorkflow) On(events *workflows.Triggers) {
	_jsii_.InvokeVoid(
		t,
		"on",
		[]interface{}{events},
	)
}

func (t *jsiiProxy_TaskWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

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
	Name *string `field:"required" json:"name" yaml:"name"`
	// Permissions for the build job.
	// Experimental.
	Permissions *workflows.JobPermissions `field:"required" json:"permissions" yaml:"permissions"`
	// The main task to be executed.
	// Experimental.
	Task projen.Task `field:"required" json:"task" yaml:"task"`
	// A directory name which contains artifacts to be uploaded (e.g. `dist`). If this is set, the contents of this directory will be uploaded as an artifact at the end of the workflow run, even if other steps fail.
	// Experimental.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Override for the `with` property of the source code checkout step.
	// Experimental.
	CheckoutWith *map[string]interface{} `field:"optional" json:"checkoutWith" yaml:"checkoutWith"`
	// Adds an 'if' condition to the workflow.
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Experimental.
	Container *workflows.ContainerOptions `field:"optional" json:"container" yaml:"container"`
	// Workflow environment variables.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// The git identity to use in this workflow.
	// Experimental.
	GitIdentity *GitIdentity `field:"optional" json:"gitIdentity" yaml:"gitIdentity"`
	// The primary job id.
	// Experimental.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Mapping of job output names to values/expressions.
	// Experimental.
	Outputs *map[string]*workflows.JobStepOutput `field:"optional" json:"outputs" yaml:"outputs"`
	// Actions to run after the main build step.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Steps to run before the main build step.
	// Experimental.
	PreBuildSteps *[]*workflows.JobStep `field:"optional" json:"preBuildSteps" yaml:"preBuildSteps"`
	// Initial steps to run before the source code checkout.
	// Experimental.
	PreCheckoutSteps *[]*workflows.JobStep `field:"optional" json:"preCheckoutSteps" yaml:"preCheckoutSteps"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// The triggers for the workflow.
	// Experimental.
	Triggers *workflows.Triggers `field:"optional" json:"triggers" yaml:"triggers"`
}

// The strategy to use when edits manifest and lock files.
// Experimental.
type VersioningStrategy string

const (
	// Only create pull requests to update lockfiles updates.
	//
	// Ignore any new
	// versions that would require package manifest changes.
	// Experimental.
	VersioningStrategy_LOCKFILE_ONLY VersioningStrategy = "LOCKFILE_ONLY"
	// - For apps, the version requirements are increased.
	//
	// - For libraries, the range of versions is widened.
	// Experimental.
	VersioningStrategy_AUTO VersioningStrategy = "AUTO"
	// Relax the version requirement to include both the new and old version, when possible.
	// Experimental.
	VersioningStrategy_WIDEN VersioningStrategy = "WIDEN"
	// Always increase the version requirement to match the new version.
	// Experimental.
	VersioningStrategy_INCREASE VersioningStrategy = "INCREASE"
	// Increase the version requirement only when required by the new version.
	// Experimental.
	VersioningStrategy_INCREASE_IF_NECESSARY VersioningStrategy = "INCREASE_IF_NECESSARY"
)

