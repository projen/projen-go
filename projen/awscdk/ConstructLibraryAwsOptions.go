package awscdk

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/cdk"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/typescript"
)

// Deprecated: use `AwsCdkConstructLibraryOptions`.
type ConstructLibraryAwsOptions struct {
	// This is the name of your project.
	// Default: $BASEDIR.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to commit the managed files by default.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CommitGenerated *bool `field:"optional" json:"commitGenerated" yaml:"commitGenerated"`
	// Configuration options for .gitignore file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	GitIgnoreOptions *projen.IgnoreFileOptions `field:"optional" json:"gitIgnoreOptions" yaml:"gitIgnoreOptions"`
	// Configuration options for git.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	GitOptions *projen.GitOptions `field:"optional" json:"gitOptions" yaml:"gitOptions"`
	// Configure logging options such as verbosity.
	// Default: {}.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Logging *projen.LoggerOptions `field:"optional" json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// subprojects.
	// Default: "."
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Default: "npx projen".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJsonOptions *projen.ProjenrcJsonOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Default: - auto approve is disabled.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Default: - see defaults in `AutoMergeOptions`.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Default: - true, but false for subprojects.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Default: - see GitHubOptions.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Default: true.
	//
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Default: - default options.
	//
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Default: ProjectType.UNKNOWN
	//
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Default: - use a personal access token named PROJEN_GITHUB_TOKEN.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Default: "PROJEN_GITHUB_TOKEN".
	//
	// Deprecated: use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Default: - { filename: 'README.md', contents: '# replace this' }
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Default: - see defaults in `StaleOptions`.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AllowLibraryDependencies *bool `field:"optional" json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Is the author an organization.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorOrganization *bool `field:"optional" json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorUrl *string `field:"optional" json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoDetectBin *bool `field:"optional" json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Bin *map[string]*string `field:"optional" json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BugsEmail *string `field:"optional" json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BugsUrl *string `field:"optional" json:"bugsUrl" yaml:"bugsUrl"`
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
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BundledDeps *[]*string `field:"optional" json:"bundledDeps" yaml:"bundledDeps"`
	// Options for npm packages using AWS CodeArtifact.
	//
	// This is required if publishing packages to, or installing scoped packages from AWS CodeArtifact.
	// Default: - undefined.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CodeArtifactOptions *javascript.CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'express', 'lodash', 'foo@^2' ]
	//
	// Default: [].
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Description *string `field:"optional" json:"description" yaml:"description"`
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
	// Example:
	//   [ 'typescript', '@types/express' ]
	//
	// Default: [].
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Default: "lib/index.js"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Default: "Apache-2.0"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Licensed *bool `field:"optional" json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Default: - no max.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MaxNodeVersion *string `field:"optional" json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Default: - no "engines" specified.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MinNodeVersion *string `field:"optional" json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Default: - for scoped packages (e.g. `foo@bar`), the default is
	// `NpmAccess.RESTRICTED`, for non-scoped packages, the default is
	// `NpmAccess.PUBLIC`.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmAccess javascript.NpmAccess `field:"optional" json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `field:"optional" json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Default: "https://registry.npmjs.org"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmRegistryUrl *string `field:"optional" json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Default: "NPM_TOKEN".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmTokenSecret *string `field:"optional" json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Default: NodePackageManager.YARN_CLASSIC
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PackageManager javascript.NodePackageManager `field:"optional" json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Default: - defaults to project name.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PeerDependencyOptions *javascript.PeerDependencyOptions `field:"optional" json:"peerDependencyOptions" yaml:"peerDependencyOptions"`
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
	// Default: [].
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PeerDeps *[]*string `field:"optional" json:"peerDeps" yaml:"peerDeps"`
	// The version of PNPM to use if using PNPM as a package manager.
	// Default: "7".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PnpmVersion *string `field:"optional" json:"pnpmVersion" yaml:"pnpmVersion"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	RepositoryDirectory *string `field:"optional" json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// Options for privately hosted scoped packages.
	// Default: - fetch all scoped packages from the public npm registry.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ScopedPackagesOptions *[]*javascript.ScopedPackagesOptions `field:"optional" json:"scopedPackagesOptions" yaml:"scopedPackagesOptions"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Also adds the script as a task.
	// Default: {}.
	//
	// Deprecated: use `project.addTask()` or `package.setScript()`
	Scripts *map[string]*string `field:"optional" json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Stability *string `field:"optional" json:"stability" yaml:"stability"`
	// Options for Yarn Berry.
	// Default: - Yarn Berry v4 with all default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	YarnBerryOptions *javascript.YarnBerryOptions `field:"optional" json:"yarnBerryOptions" yaml:"yarnBerryOptions"`
	// Version requirement of `publib` which is used to publish modules to npm.
	// Default: "latest".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Default: - Major version is not enforced.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// Minimal Major version to release.
	//
	// This can be useful to set to 1, as breaking changes before the 1.x major
	// release are not incrementing the major version number.
	//
	// Can not be set together with `majorVersion`.
	// Default: - No minimum version is being enforced.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MinMajorVersion *float64 `field:"optional" json:"minMajorVersion" yaml:"minMajorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Default: "latest".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Default: [].
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Default: - normal semantic versions.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishDryRun *bool `field:"optional" json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Find commits that should be considered releasable Used to decide if a release is required.
	// Default: ReleasableCommits.everyCommit()
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleasableCommits projen.ReleasableCommits `field:"optional" json:"releasableCommits" yaml:"releasableCommits"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Default: - no additional branches are used for release. you can use
	// `addBranch()` to add additional branches.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseBranches *map[string]*release.BranchOptions `field:"optional" json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Default: true.
	//
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `field:"optional" json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseFailureIssue *bool `field:"optional" json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Default: "failed-release".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseFailureIssueLabel *string `field:"optional" json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Default: - no scheduled releases.
	//
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `field:"optional" json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Default: "v".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseTagPrefix *string `field:"optional" json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Default: - Continuous releases (`ReleaseTrigger.continuous()`)
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseTrigger release.ReleaseTrigger `field:"optional" json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Default: "Release".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseWorkflowName *string `field:"optional" json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `field:"optional" json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Default: - standard configuration applicable for GitHub repositories.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Default: - default image.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// Github Runner Group selection options.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowRunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"workflowRunsOnGroup" yaml:"workflowRunsOnGroup"`
	// The name of the main release branch.
	// Default: "main".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DefaultReleaseBranch *string `field:"required" json:"defaultReleaseBranch" yaml:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Default: "dist".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Default: - true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AutoApproveUpgrades *bool `field:"optional" json:"autoApproveUpgrades" yaml:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Default: - true if not a subproject.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BuildWorkflow *bool `field:"optional" json:"buildWorkflow" yaml:"buildWorkflow"`
	// Build workflow triggers.
	// Default: "{ pullRequest: {}, workflowDispatch: {} }".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BuildWorkflowTriggers *workflows.Triggers `field:"optional" json:"buildWorkflowTriggers" yaml:"buildWorkflowTriggers"`
	// Options for `Bundler`.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	BundlerOptions *javascript.BundlerOptions `field:"optional" json:"bundlerOptions" yaml:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v3 A secret is required for private repos. Configured with `@codeCovTokenSecret`.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CodeCov *bool `field:"optional" json:"codeCov" yaml:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Default: - if this option is not specified, only public repositories are supported.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CodeCovTokenSecret *string `field:"optional" json:"codeCovTokenSecret" yaml:"codeCovTokenSecret"`
	// License copyright owner.
	// Default: - defaults to the value of authorName or "" if `authorName` is undefined.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CopyrightOwner *string `field:"optional" json:"copyrightOwner" yaml:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Default: - current year.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CopyrightPeriod *string `field:"optional" json:"copyrightPeriod" yaml:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Dependabot *bool `field:"optional" json:"dependabot" yaml:"dependabot"`
	// Options for dependabot.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DependabotOptions *github.DependabotOptions `field:"optional" json:"dependabotOptions" yaml:"dependabotOptions"`
	// Use tasks and github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DepsUpgrade *bool `field:"optional" json:"depsUpgrade" yaml:"depsUpgrade"`
	// Options for `UpgradeDependencies`.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `field:"optional" json:"depsUpgradeOptions" yaml:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Gitignore *[]*string `field:"optional" json:"gitignore" yaml:"gitignore"`
	// Setup jest unit tests.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Jest *bool `field:"optional" json:"jest" yaml:"jest"`
	// Jest options.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	JestOptions *javascript.JestOptions `field:"optional" json:"jestOptions" yaml:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `field:"optional" json:"npmignore" yaml:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmignoreEnabled *bool `field:"optional" json:"npmignoreEnabled" yaml:"npmignoreEnabled"`
	// Configuration options for .npmignore file.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	NpmIgnoreOptions *projen.IgnoreFileOptions `field:"optional" json:"npmIgnoreOptions" yaml:"npmIgnoreOptions"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Package *bool `field:"optional" json:"package" yaml:"package"`
	// Setup prettier.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Prettier options.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PrettierOptions *javascript.PrettierOptions `field:"optional" json:"prettierOptions" yaml:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenDevDependency *bool `field:"optional" json:"projenDevDependency" yaml:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Default: - true if projenrcJson is false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options for .projenrc.js.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Version of projen to install.
	// Default: - Defaults to the latest version.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Include a GitHub pull request template.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PullRequestTemplate *bool `field:"optional" json:"pullRequestTemplate" yaml:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Default: - default content.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PullRequestTemplateContents *[]*string `field:"optional" json:"pullRequestTemplateContents" yaml:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Default: - true (false for subprojects).
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Release *bool `field:"optional" json:"release" yaml:"release"`
	// Automatically release to npm when new versions are introduced.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ReleaseToNpm *bool `field:"optional" json:"releaseToNpm" yaml:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Default: - true if not a subproject.
	//
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `field:"optional" json:"releaseWorkflow" yaml:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Default: "yarn install --frozen-lockfile && yarn projen".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowBootstrapSteps *[]*workflows.JobStep `field:"optional" json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Default: - GitHub Actions.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowGitIdentity *github.GitIdentity `field:"optional" json:"workflowGitIdentity" yaml:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Default: - same as `minNodeVersion`.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Enable Node.js package cache in GitHub workflows.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	WorkflowPackageCache *bool `field:"optional" json:"workflowPackageCache" yaml:"workflowPackageCache"`
	// Do not generate a `tsconfig.json` file (used by jsii projects since tsconfig.json is generated by the jsii compiler).
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DisableTsconfig *bool `field:"optional" json:"disableTsconfig" yaml:"disableTsconfig"`
	// Do not generate a `tsconfig.dev.json` file.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DisableTsconfigDev *bool `field:"optional" json:"disableTsconfigDev" yaml:"disableTsconfigDev"`
	// Docgen by Typedoc.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Docgen *bool `field:"optional" json:"docgen" yaml:"docgen"`
	// Docs directory.
	// Default: "docs".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DocsDirectory *string `field:"optional" json:"docsDirectory" yaml:"docsDirectory"`
	// The .d.ts file that includes the type declarations for this module.
	// Default: - .d.ts file derived from the project's entrypoint (usually lib/index.d.ts)
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	EntrypointTypes *string `field:"optional" json:"entrypointTypes" yaml:"entrypointTypes"`
	// Setup eslint.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Eslint *bool `field:"optional" json:"eslint" yaml:"eslint"`
	// Eslint options.
	// Default: - opinionated default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	EslintOptions *javascript.EslintOptions `field:"optional" json:"eslintOptions" yaml:"eslintOptions"`
	// Typescript  artifacts output directory.
	// Default: "lib".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Libdir *string `field:"optional" json:"libdir" yaml:"libdir"`
	// Use TypeScript for your projenrc file (`.projenrc.ts`).
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcTs *bool `field:"optional" json:"projenrcTs" yaml:"projenrcTs"`
	// Options for .projenrc.ts.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ProjenrcTsOptions *typescript.ProjenrcOptions `field:"optional" json:"projenrcTsOptions" yaml:"projenrcTsOptions"`
	// Generate one-time sample in `src/` and `test/` if there are no files there.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	SampleCode *bool `field:"optional" json:"sampleCode" yaml:"sampleCode"`
	// Typescript sources directory.
	// Default: "src".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Srcdir *string `field:"optional" json:"srcdir" yaml:"srcdir"`
	// Jest tests directory. Tests files should be named `xxx.test.ts`.
	//
	// If this directory is under `srcdir` (e.g. `src/test`, `src/__tests__`),
	// then tests are going to be compiled into `lib/` and executed as javascript.
	// If the test directory is outside of `src`, then we configure jest to
	// compile the code in-memory.
	// Default: "test".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
	// Custom TSConfig.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Tsconfig *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Custom tsconfig options for the development tsconfig.json file (used for testing).
	// Default: - use the production tsconfig options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	TsconfigDev *javascript.TypescriptConfigOptions `field:"optional" json:"tsconfigDev" yaml:"tsconfigDev"`
	// The name of the development tsconfig.json file.
	// Default: "tsconfig.dev.json"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	TsconfigDevFile *string `field:"optional" json:"tsconfigDevFile" yaml:"tsconfigDevFile"`
	// TypeScript version to use.
	//
	// NOTE: Typescript is not semantically versioned and should remain on the
	// same minor, so we recommend using a `~` dependency (e.g. `~1.2.3`).
	// Default: "latest".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	TypescriptVersion *string `field:"optional" json:"typescriptVersion" yaml:"typescriptVersion"`
	// The name of the library author.
	// Default: $GIT_USER_NAME.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Author *string `field:"required" json:"author" yaml:"author"`
	// Email or URL of the library author.
	// Default: $GIT_USER_EMAIL.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	AuthorAddress *string `field:"required" json:"authorAddress" yaml:"authorAddress"`
	// Git repository URL.
	// Default: $GIT_REMOTE.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	//
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Compat *bool `field:"optional" json:"compat" yaml:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Default: ".compatignore"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CompatIgnore *string `field:"optional" json:"compatIgnore" yaml:"compatIgnore"`
	// Emit a compressed version of the assembly.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CompressAssembly *bool `field:"optional" json:"compressAssembly" yaml:"compressAssembly"`
	// File path for generated docs.
	// Default: "API.md"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	DocgenFilePath *string `field:"optional" json:"docgenFilePath" yaml:"docgenFilePath"`
	// Deprecated: use `publishToNuget`.
	Dotnet *cdk.JsiiDotNetTarget `field:"optional" json:"dotnet" yaml:"dotnet"`
	// Accepts a list of glob patterns.
	//
	// Files matching any of those patterns will be excluded from the TypeScript compiler input.
	//
	// By default, jsii will include all *.ts files (except .d.ts files) in the TypeScript compiler input.
	// This can be problematic for example when the package's build or test procedure generates .ts files
	// that cannot be compiled with jsii's compiler settings.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ExcludeTypescript *[]*string `field:"optional" json:"excludeTypescript" yaml:"excludeTypescript"`
	// Version of the jsii compiler to use.
	//
	// Set to "*" if you want to manually manage the version of jsii in your
	// project by managing updates to `package.json` on your own.
	//
	// NOTE: The jsii compiler releases since 5.0.0 are not semantically versioned
	// and should remain on the same minor, so we recommend using a `~` dependency
	// (e.g. `~5.0.0`).
	// Default: "1.x"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	JsiiVersion *string `field:"optional" json:"jsiiVersion" yaml:"jsiiVersion"`
	// Publish Go bindings to a git repository.
	// Default: - no publishing.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToGo *cdk.JsiiGoTarget `field:"optional" json:"publishToGo" yaml:"publishToGo"`
	// Publish to maven.
	// Default: - no publishing.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToMaven *cdk.JsiiJavaTarget `field:"optional" json:"publishToMaven" yaml:"publishToMaven"`
	// Publish to NuGet.
	// Default: - no publishing.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToNuget *cdk.JsiiDotNetTarget `field:"optional" json:"publishToNuget" yaml:"publishToNuget"`
	// Publish to pypi.
	// Default: - no publishing.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	PublishToPypi *cdk.JsiiPythonTarget `field:"optional" json:"publishToPypi" yaml:"publishToPypi"`
	// Deprecated: use `publishToPyPi`.
	Python *cdk.JsiiPythonTarget `field:"optional" json:"python" yaml:"python"`
	// Default: "."
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Rootdir *string `field:"optional" json:"rootdir" yaml:"rootdir"`
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
	// Default: - new version will be announced.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	Catalog *cdk.Catalog `field:"optional" json:"catalog" yaml:"catalog"`
	// Minimum version of the AWS CDK to depend on.
	// Default: "2.1.0"
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the.
	// Default: - will be included by default for AWS CDK >= 1.0.0 < 2.0.0
	//
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Default: - will be included by default for AWS CDK >= 1.111.0 < 2.0.0
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Default: true.
	//
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Default: - for CDK 1.x the default is "3.2.27", for CDK 2.x the default is
	// "10.0.5".
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// Automatically adds an `cloudfront.experimental.EdgeFunction` for each `.edge-lambda.ts` handler in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	EdgeLambdaAutoDiscover *bool `field:"optional" json:"edgeLambdaAutoDiscover" yaml:"edgeLambdaAutoDiscover"`
	// Enable experimental support for the AWS CDK integ-runner.
	// Default: false.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	ExperimentalIntegRunner *bool `field:"optional" json:"experimentalIntegRunner" yaml:"experimentalIntegRunner"`
	// Automatically discovers and creates integration tests for each `.integ.ts` file in under your test directory.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
	// Automatically adds an `aws_lambda.Function` for each `.lambda.ts` handler in your source tree. If this is disabled, you either need to explicitly call `aws_lambda.Function.autoDiscover()` or define a `new aws_lambda.Function()` for each handler.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	LambdaAutoDiscover *bool `field:"optional" json:"lambdaAutoDiscover" yaml:"lambdaAutoDiscover"`
	// Automatically adds an `awscdk.LambdaExtension` for each `.lambda-extension.ts` entrypoint in your source tree. If this is disabled, you can manually add an `awscdk.AutoDiscover` component to your project.
	// Default: true.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	LambdaExtensionAutoDiscover *bool `field:"optional" json:"lambdaExtensionAutoDiscover" yaml:"lambdaExtensionAutoDiscover"`
	// Common options for all AWS Lambda functions.
	// Default: - default options.
	//
	// Deprecated: use `AwsCdkConstructLibraryOptions`.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
}

