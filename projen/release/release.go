package release

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/release/internal"
)

// Options for a release branch.
// Experimental.
type BranchOptions struct {
	// The major versions released from this branch.
	// Experimental.
	MajorVersion *float64 `json:"majorVersion"`
	// The npm distribution tag to use for this branch.
	// Experimental.
	NpmDistTag *string `json:"npmDistTag"`
	// Bump the version as a pre-release tag.
	// Experimental.
	Prerelease *string `json:"prerelease"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	TagPrefix *string `json:"tagPrefix"`
	// The name of the release workflow.
	// Experimental.
	WorkflowName *string `json:"workflowName"`
}

// Experimental.
type CodeArtifactOptions struct {
	// GitHub secret which contains the AWS access key ID to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`registry` contains AWS CodeArtifact URL).
	// Experimental.
	AccessKeyIdSecret *string `json:"accessKeyIdSecret"`
	// ARN of AWS role to be assumed prior to get authorization token from AWS CodeArtifact This property must be specified only when publishing to AWS CodeArtifact (`registry` contains AWS CodeArtifact URL).
	// Experimental.
	RoleToAssume *string `json:"roleToAssume"`
	// GitHub secret which contains the AWS secret access key to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`registry` contains AWS CodeArtifact URL).
	// Experimental.
	SecretAccessKeySecret *string `json:"secretAccessKeySecret"`
}

// Publishing options for GitHub releases.
// Experimental.
type GitHubReleasesPublishOptions struct {
	// The location of an .md file (relative to `dist/`) that includes the changelog for the release.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ChangelogFile *string `json:"changelogFile"`
	// The location of a text file (relative to `dist/`) that contains the release tag.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ReleaseTagFile *string `json:"releaseTagFile"`
	// The location of a text file (relative to `dist/`) that contains the version number.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	VersionFile *string `json:"versionFile"`
}

// Publishing options for Git releases.
// Experimental.
type GitPublishOptions struct {
	// The location of an .md file (relative to `dist/`) that includes the changelog for the release.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ChangelogFile *string `json:"changelogFile"`
	// The location of a text file (relative to `dist/`) that contains the release tag.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ReleaseTagFile *string `json:"releaseTagFile"`
	// The location of a text file (relative to `dist/`) that contains the version number.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	VersionFile *string `json:"versionFile"`
	// Branch to push to.
	// Experimental.
	GitBranch *string `json:"gitBranch"`
	// Override git-push command.
	//
	// Set to an empty string to disable pushing.
	// Experimental.
	GitPushCommand *string `json:"gitPushCommand"`
	// The location of an .md file that includes the project-level changelog.
	// Experimental.
	ProjectChangelogFile *string `json:"projectChangelogFile"`
}

// Deprecated: Use `GoPublishOptions` instead.
// export interface JsiiReleaseGo extends GoPublishOptions { }
// /**
// Options for Go releases.
type GoPublishOptions struct {
	// Branch to push to.
	// Deprecated: Use `GoPublishOptions` instead.
	// export interface JsiiReleaseGo extends GoPublishOptions { }
	// /**
	// Options for Go releases.
	GitBranch *string `json:"gitBranch"`
	// The commit message.
	// Deprecated: Use `GoPublishOptions` instead.
	// export interface JsiiReleaseGo extends GoPublishOptions { }
	// /**
	// Options for Go releases.
	GitCommitMessage *string `json:"gitCommitMessage"`
	// GitHub repository to push to.
	// Deprecated: Use `GoPublishOptions` instead.
	// export interface JsiiReleaseGo extends GoPublishOptions { }
	// /**
	// Options for Go releases.
	GithubRepo *string `json:"githubRepo"`
	// The name of the secret that includes a personal GitHub access token used to push to the GitHub repository.
	// Deprecated: Use `GoPublishOptions` instead.
	// export interface JsiiReleaseGo extends GoPublishOptions { }
	// /**
	// Options for Go releases.
	GithubTokenSecret *string `json:"githubTokenSecret"`
	// The email to use in the release git commit.
	// Deprecated: Use `GoPublishOptions` instead.
	// export interface JsiiReleaseGo extends GoPublishOptions { }
	// /**
	// Options for Go releases.
	GitUserEmail *string `json:"gitUserEmail"`
	// The user name to use for the release git commit.
	// Deprecated: Use `GoPublishOptions` instead.
	// export interface JsiiReleaseGo extends GoPublishOptions { }
	// /**
	// Options for Go releases.
	GitUserName *string `json:"gitUserName"`
}

// Deprecated: Use `MavenPublishOptions` instead.
type JsiiReleaseMaven struct {
	// URL of Nexus repository.
	//
	// if not set, defaults to https://oss.sonatype.org
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenEndpoint *string `json:"mavenEndpoint"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven packages. See instructions.
	// See: https://github.com/aws/jsii-release#maven
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenGpgPrivateKeyPassphrase *string `json:"mavenGpgPrivateKeyPassphrase"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven
	// packages. See instructions.
	// See: https://github.com/aws/jsii-release#maven
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenGpgPrivateKeySecret *string `json:"mavenGpgPrivateKeySecret"`
	// GitHub secret name which contains the Password for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenPassword *string `json:"mavenPassword"`
	// Deployment repository when not deploying to Maven Central.
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenRepositoryUrl *string `json:"mavenRepositoryUrl"`
	// Used in maven settings for credential lookup (e.g. use github when publishing to GitHub).
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenServerId *string `json:"mavenServerId"`
	// GitHub secret name which contains the Maven Central (sonatype) staging profile ID (e.g. 68a05363083174). Staging profile ID can be found in the URL of the "Releases" staging profile under "Staging Profiles" in https://oss.sonatype.org (e.g. https://oss.sonatype.org/#stagingProfiles;11a33451234521).
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenStagingProfileId *string `json:"mavenStagingProfileId"`
	// GitHub secret name which contains the Username for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Deprecated: Use `MavenPublishOptions` instead.
	MavenUsername *string `json:"mavenUsername"`
}

// Deprecated: Use `NpmPublishOptions` instead.
type JsiiReleaseNpm struct {
	// Options for publishing npm package to AWS CodeArtifact.
	// Deprecated: Use `NpmPublishOptions` instead.
	CodeArtifactOptions *CodeArtifactOptions `json:"codeArtifactOptions"`
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
	// Deprecated: Use `npmDistTag` for each release branch instead.
	DistTag *string `json:"distTag"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Deprecated: Use `NpmPublishOptions` instead.
	NpmTokenSecret *string `json:"npmTokenSecret"`
	// The domain name of the npm package registry.
	//
	// To publish to GitHub Packages, set this value to `"npm.pkg.github.com"`. In
	// this if `npmTokenSecret` is not specified, it will default to
	// `GITHUB_TOKEN` which means that you will be able to publish to the
	// repository's package store. In this case, make sure `repositoryUrl` is
	// correctly defined.
	//
	// TODO: EXAMPLE
	//
	// Deprecated: Use `NpmPublishOptions` instead.
	Registry *string `json:"registry"`
}

// Deprecated: Use `NugetPublishOptions` instead.
type JsiiReleaseNuget struct {
	// GitHub secret which contains the API key for NuGet.
	// Deprecated: Use `NugetPublishOptions` instead.
	NugetApiKeySecret *string `json:"nugetApiKeySecret"`
}

// Deprecated: Use `PyPiPublishOptions` instead.
type JsiiReleasePyPi struct {
	// The GitHub secret which contains PyPI password.
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwinePasswordSecret *string `json:"twinePasswordSecret"`
	// The registry url to use when releasing packages.
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwineRegistryUrl *string `json:"twineRegistryUrl"`
	// The GitHub secret which contains PyPI user name.
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwineUsernameSecret *string `json:"twineUsernameSecret"`
}

// Experimental.
type ManualReleaseOptions struct {
	// Maintain a project-level changelog.
	// Experimental.
	Changelog *bool `json:"changelog"`
	// Project-level changelog file path.
	//
	// Ignored if `changelog` is false.
	// Experimental.
	ChangelogPath *string `json:"changelogPath"`
	// Override git-push command.
	//
	// Set to an empty string to disable pushing.
	// Experimental.
	GitPushCommand *string `json:"gitPushCommand"`
}

// Options for Maven releases.
// Experimental.
type MavenPublishOptions struct {
	// URL of Nexus repository.
	//
	// if not set, defaults to https://oss.sonatype.org
	// Experimental.
	MavenEndpoint *string `json:"mavenEndpoint"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven packages. See instructions.
	// See: https://github.com/aws/jsii-release#maven
	//
	// Experimental.
	MavenGpgPrivateKeyPassphrase *string `json:"mavenGpgPrivateKeyPassphrase"`
	// GitHub secret name which contains the GPG private key or file that includes it.
	//
	// This is used to sign your Maven
	// packages. See instructions.
	// See: https://github.com/aws/jsii-release#maven
	//
	// Experimental.
	MavenGpgPrivateKeySecret *string `json:"mavenGpgPrivateKeySecret"`
	// GitHub secret name which contains the Password for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Experimental.
	MavenPassword *string `json:"mavenPassword"`
	// Deployment repository when not deploying to Maven Central.
	// Experimental.
	MavenRepositoryUrl *string `json:"mavenRepositoryUrl"`
	// Used in maven settings for credential lookup (e.g. use github when publishing to GitHub).
	// Experimental.
	MavenServerId *string `json:"mavenServerId"`
	// GitHub secret name which contains the Maven Central (sonatype) staging profile ID (e.g. 68a05363083174). Staging profile ID can be found in the URL of the "Releases" staging profile under "Staging Profiles" in https://oss.sonatype.org (e.g. https://oss.sonatype.org/#stagingProfiles;11a33451234521).
	// Experimental.
	MavenStagingProfileId *string `json:"mavenStagingProfileId"`
	// GitHub secret name which contains the Username for maven repository.
	//
	// For Maven Central, you will need to Create JIRA account and then request a
	// new project (see links).
	// See: https://issues.sonatype.org/secure/CreateIssue.jspa?issuetype=21&pid=10134
	//
	// Experimental.
	MavenUsername *string `json:"mavenUsername"`
}

// Options for npm release.
// Experimental.
type NpmPublishOptions struct {
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *CodeArtifactOptions `json:"codeArtifactOptions"`
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
	// Deprecated: Use `npmDistTag` for each release branch instead.
	DistTag *string `json:"distTag"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `json:"npmTokenSecret"`
	// The domain name of the npm package registry.
	//
	// To publish to GitHub Packages, set this value to `"npm.pkg.github.com"`. In
	// this if `npmTokenSecret` is not specified, it will default to
	// `GITHUB_TOKEN` which means that you will be able to publish to the
	// repository's package store. In this case, make sure `repositoryUrl` is
	// correctly defined.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Registry *string `json:"registry"`
}

// Options for NuGet releases.
// Experimental.
type NugetPublishOptions struct {
	// GitHub secret which contains the API key for NuGet.
	// Experimental.
	NugetApiKeySecret *string `json:"nugetApiKeySecret"`
}

// Implements GitHub jobs for publishing modules to package managers.
//
// Under the hood, it uses https://github.com/aws/jsii-release
// Experimental.
type Publisher interface {
	projen.Component
	ArtifactName() *string
	BuildJobId() *string
	Condition() *string
	JsiiReleaseVersion() *string
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
	PublishToGit(options *GitPublishOptions) projen.Task
	PublishToGitHubReleases(options *GitHubReleasesPublishOptions)
	PublishToGo(options *GoPublishOptions)
	PublishToMaven(options *MavenPublishOptions)
	PublishToNpm(options *NpmPublishOptions)
	PublishToNuget(options *NugetPublishOptions)
	PublishToPyPi(options *PyPiPublishOptions)
	Synthesize()
}

// The jsii proxy struct for Publisher
type jsiiProxy_Publisher struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Publisher) ArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) BuildJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) JsiiReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsiiReleaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Publisher) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPublisher(project projen.Project, options *PublisherOptions) Publisher {
	_init_.Initialize()

	j := jsiiProxy_Publisher{}

	_jsii_.Create(
		"projen.release.Publisher",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPublisher_Override(p Publisher, project projen.Project, options *PublisherOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.release.Publisher",
		[]interface{}{project, options},
		p,
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Publisher) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Publisher) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Publish to git.
//
// This includes generating a project-level changelog and release tags.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToGit(options *GitPublishOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		p,
		"publishToGit",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a GitHub Release.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToGitHubReleases(options *GitHubReleasesPublishOptions) {
	_jsii_.InvokeVoid(
		p,
		"publishToGitHubReleases",
		[]interface{}{options},
	)
}

// Adds a go publishing job.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToGo(options *GoPublishOptions) {
	_jsii_.InvokeVoid(
		p,
		"publishToGo",
		[]interface{}{options},
	)
}

// Publishes artifacts from `java/**` to Maven.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToMaven(options *MavenPublishOptions) {
	_jsii_.InvokeVoid(
		p,
		"publishToMaven",
		[]interface{}{options},
	)
}

// Publishes artifacts from `js/**` to npm.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToNpm(options *NpmPublishOptions) {
	_jsii_.InvokeVoid(
		p,
		"publishToNpm",
		[]interface{}{options},
	)
}

// Publishes artifacts from `dotnet/**` to NuGet Gallery.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToNuget(options *NugetPublishOptions) {
	_jsii_.InvokeVoid(
		p,
		"publishToNuget",
		[]interface{}{options},
	)
}

// Publishes wheel artifacts from `python` to PyPI.
// Experimental.
func (p *jsiiProxy_Publisher) PublishToPyPi(options *PyPiPublishOptions) {
	_jsii_.InvokeVoid(
		p,
		"publishToPyPi",
		[]interface{}{options},
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (p *jsiiProxy_Publisher) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Publisher`.
// Experimental.
type PublisherOptions struct {
	// The name of the artifact to download (e.g. `dist`).
	//
	// The artifact is expected to include a subdirectory for each release target:
	// `go` (GitHub), `dotnet` (NuGet), `java` (Maven), `js` (npm), `python`
	// (PyPI).
	// See: https://github.com/aws/jsii-release
	//
	// Experimental.
	ArtifactName *string `json:"artifactName"`
	// The job ID that produces the build artifacts.
	//
	// All publish jobs will take a dependency on this job.
	// Experimental.
	BuildJobId *string `json:"buildJobId"`
	// A GitHub workflow expression used as a condition for publishers.
	// Experimental.
	Condition *string `json:"condition"`
	// Create an issue when a publish task fails.
	// Experimental.
	FailureIssue *bool `json:"failureIssue"`
	// The label to apply to the issue marking failed publish tasks.
	//
	// Only applies if `failureIssue` is true.
	// Experimental.
	FailureIssueLabel *string `json:"failureIssueLabel"`
	// Version requirement for `jsii-release`.
	// Experimental.
	JsiiReleaseVersion *string `json:"jsiiReleaseVersion"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `json:"publishTasks"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `json:"workflowRunsOn"`
}

// Options for PyPI release.
// Experimental.
type PyPiPublishOptions struct {
	// The GitHub secret which contains PyPI password.
	// Experimental.
	TwinePasswordSecret *string `json:"twinePasswordSecret"`
	// The registry url to use when releasing packages.
	// Experimental.
	TwineRegistryUrl *string `json:"twineRegistryUrl"`
	// The GitHub secret which contains PyPI user name.
	// Experimental.
	TwineUsernameSecret *string `json:"twineUsernameSecret"`
}

// Manages releases (currently through GitHub workflows).
//
// By default, no branches are released. To add branches, call `addBranch()`.
// Experimental.
type Release interface {
	projen.Component
	Branches() *[]*string
	Project() projen.Project
	Publisher() Publisher
	AddBranch(branch *string, options *BranchOptions)
	AddJobs(jobs *map[string]*workflows.Job)
	PostSynthesize()
	PreSynthesize()
	Synthesize()
}

// The jsii proxy struct for Release
type jsiiProxy_Release struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Release) Branches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Publisher() Publisher {
	var returns Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}


// Experimental.
func NewRelease(project github.GitHubProject, options *ReleaseOptions) Release {
	_init_.Initialize()

	j := jsiiProxy_Release{}

	_jsii_.Create(
		"projen.release.Release",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRelease_Override(r Release, project github.GitHubProject, options *ReleaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.release.Release",
		[]interface{}{project, options},
		r,
	)
}

// Adds a release branch.
//
// It is a git branch from which releases are published. If a project has more than one release
// branch, we require that `majorVersion` is also specified for the primary branch in order to
// ensure branches always release the correct version.
// Experimental.
func (r *jsiiProxy_Release) AddBranch(branch *string, options *BranchOptions) {
	_jsii_.InvokeVoid(
		r,
		"addBranch",
		[]interface{}{branch, options},
	)
}

// Adds jobs to all release workflows.
// Experimental.
func (r *jsiiProxy_Release) AddJobs(jobs *map[string]*workflows.Job) {
	_jsii_.InvokeVoid(
		r,
		"addJobs",
		[]interface{}{jobs},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (r *jsiiProxy_Release) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (r *jsiiProxy_Release) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
func (r *jsiiProxy_Release) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Release`.
// Experimental.
type ReleaseOptions struct {
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
	ReleaseBranches *map[string]*BranchOptions `json:"releaseBranches"`
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
	ReleaseTrigger ReleaseTrigger `json:"releaseTrigger"`
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
	// The default branch name to release from.
	//
	// Use `majorVersion` to restrict this branch to only publish releases with a
	// specific major version.
	//
	// You can add additional branches using `addBranch()`.
	// Experimental.
	Branch *string `json:"branch"`
	// The task to execute in order to create the release artifacts.
	//
	// Artifacts are
	// expected to reside under `artifactsDirectory` (defaults to `dist/`) once
	// build is complete.
	// Experimental.
	Task projen.Task `json:"task"`
	// A name of a .json file to set the `version` field in after a bump.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	VersionFile *string `json:"versionFile"`
	// Create a GitHub release for each release.
	// Experimental.
	GithubRelease *bool `json:"githubRelease"`
}

// Project options for release.
// Experimental.
type ReleaseProjectOptions struct {
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
	ReleaseBranches *map[string]*BranchOptions `json:"releaseBranches"`
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
	ReleaseTrigger ReleaseTrigger `json:"releaseTrigger"`
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
}

// Used to manage release strategies.
//
// This includes release
// and release artifact automation
// Experimental.
type ReleaseTrigger interface {
	ChangelogPath() *string
	GitPushCommand() *string
	IsContinuous() *bool
	IsManual() *bool
	Schedule() *string
}

// The jsii proxy struct for ReleaseTrigger
type jsiiProxy_ReleaseTrigger struct {
	_ byte // padding
}

func (j *jsiiProxy_ReleaseTrigger) ChangelogPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changelogPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) GitPushCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitPushCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) IsContinuous() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isContinuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) IsManual() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseTrigger) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}


// Creates a continuous release trigger.
//
// Automated releases will occur on every commit.
// Experimental.
func ReleaseTrigger_Continuous() ReleaseTrigger {
	_init_.Initialize()

	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"continuous",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a manual release trigger.
//
// Use this option if you want totally manual releases.
//
// This will give you a release task that, in addition to the normal
// release activities will trigger a `publish:git` task. This task will
// handle project-level changelog management, release tagging, and pushing
// these artifacts to origin.
//
// The command used for pushing can be customised by specifying
// `gitPushCommand`. Set to an empty string to disable pushing entirely.
//
// Simply run `yarn release` to trigger a manual release.
// Experimental.
func ReleaseTrigger_Manual(options *ManualReleaseOptions) ReleaseTrigger {
	_init_.Initialize()

	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"manual",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a scheduled release trigger.
//
// Automated releases will occur based on the provided cron schedule.
// Experimental.
func ReleaseTrigger_Scheduled(options *ScheduledReleaseOptions) ReleaseTrigger {
	_init_.Initialize()

	var returns ReleaseTrigger

	_jsii_.StaticInvoke(
		"projen.release.ReleaseTrigger",
		"scheduled",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Experimental.
type ScheduledReleaseOptions struct {
	// Cron schedule for releases.
	//
	// Only defined if this is a scheduled release.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Schedule *string `json:"schedule"`
}

