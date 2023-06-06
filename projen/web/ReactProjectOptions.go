package web

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
)

// Experimental.
type ReactProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to commit the managed files by default.
	// Experimental.
	CommitGenerated *bool `field:"optional" json:"commitGenerated" yaml:"commitGenerated"`
	// Configuration options for .gitignore file.
	// Experimental.
	GitIgnoreOptions *projen.IgnoreFileOptions `field:"optional" json:"gitIgnoreOptions" yaml:"gitIgnoreOptions"`
	// Configuration options for git.
	// Experimental.
	GitOptions *projen.GitOptions `field:"optional" json:"gitOptions" yaml:"gitOptions"`
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
	ProjenrcJsonOptions *projen.ProjenrcJsonOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Experimental.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Experimental.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
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
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
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
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `field:"optional" json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `field:"optional" json:"projectType" yaml:"projectType"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
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
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `field:"optional" json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `field:"optional" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `field:"optional" json:"authorName" yaml:"authorName"`
	// Is the author an organization.
	// Experimental.
	AuthorOrganization *bool `field:"optional" json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `field:"optional" json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `field:"optional" json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `field:"optional" json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `field:"optional" json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
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
	// Experimental.
	BundledDeps *[]*string `field:"optional" json:"bundledDeps" yaml:"bundledDeps"`
	// Options for npm packages using AWS CodeArtifact.
	//
	// This is required if publishing packages to, or installing scoped packages from AWS CodeArtifact.
	// Experimental.
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
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
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
	// Experimental.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `field:"optional" json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `field:"optional" json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `field:"optional" json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess javascript.NpmAccess `field:"optional" json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `field:"optional" json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `field:"optional" json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `field:"optional" json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager javascript.NodePackageManager `field:"optional" json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
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
	// Experimental.
	PeerDeps *[]*string `field:"optional" json:"peerDeps" yaml:"peerDeps"`
	// The version of PNPM to use if using PNPM as a package manager.
	// Experimental.
	PnpmVersion *string `field:"optional" json:"pnpmVersion" yaml:"pnpmVersion"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `field:"optional" json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// Options for privately hosted scoped packages.
	// Experimental.
	ScopedPackagesOptions *[]*javascript.ScopedPackagesOptions `field:"optional" json:"scopedPackagesOptions" yaml:"scopedPackagesOptions"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Also adds the script as a task.
	// Deprecated: use `project.addTask()` or `package.setScript()`
	Scripts *map[string]*string `field:"optional" json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `field:"optional" json:"stability" yaml:"stability"`
	// Version requirement of `publib` which is used to publish modules to npm.
	// Experimental.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Experimental.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// Minimal Major version to release.
	//
	// This can be useful to set to 1, as breaking changes before the 1.x major
	// release are not incrementing the major version number.
	//
	// Can not be set together with `majorVersion`.
	// Experimental.
	MinMajorVersion *float64 `field:"optional" json:"minMajorVersion" yaml:"minMajorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Experimental.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Experimental.
	PublishDryRun *bool `field:"optional" json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Experimental.
	ReleaseBranches *map[string]*release.BranchOptions `field:"optional" json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `field:"optional" json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Experimental.
	ReleaseFailureIssue *bool `field:"optional" json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Experimental.
	ReleaseFailureIssueLabel *string `field:"optional" json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `field:"optional" json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	ReleaseTagPrefix *string `field:"optional" json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Experimental.
	ReleaseTrigger release.ReleaseTrigger `field:"optional" json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Experimental.
	ReleaseWorkflowName *string `field:"optional" json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `field:"optional" json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Experimental.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch *string `field:"required" json:"defaultReleaseBranch" yaml:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Experimental.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveUpgrades *bool `field:"optional" json:"autoApproveUpgrades" yaml:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow *bool `field:"optional" json:"buildWorkflow" yaml:"buildWorkflow"`
	// Build workflow triggers.
	// Experimental.
	BuildWorkflowTriggers *workflows.Triggers `field:"optional" json:"buildWorkflowTriggers" yaml:"buildWorkflowTriggers"`
	// Options for `Bundler`.
	// Experimental.
	BundlerOptions *javascript.BundlerOptions `field:"optional" json:"bundlerOptions" yaml:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v3 A secret is required for private repos. Configured with `@codeCovTokenSecret`.
	// Experimental.
	CodeCov *bool `field:"optional" json:"codeCov" yaml:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret *string `field:"optional" json:"codeCovTokenSecret" yaml:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner *string `field:"optional" json:"copyrightOwner" yaml:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod *string `field:"optional" json:"copyrightPeriod" yaml:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Experimental.
	Dependabot *bool `field:"optional" json:"dependabot" yaml:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions *github.DependabotOptions `field:"optional" json:"dependabotOptions" yaml:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Experimental.
	DepsUpgrade *bool `field:"optional" json:"depsUpgrade" yaml:"depsUpgrade"`
	// Options for `UpgradeDependencies`.
	// Experimental.
	DepsUpgradeOptions *javascript.UpgradeDependenciesOptions `field:"optional" json:"depsUpgradeOptions" yaml:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore *[]*string `field:"optional" json:"gitignore" yaml:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest *bool `field:"optional" json:"jest" yaml:"jest"`
	// Jest options.
	// Experimental.
	JestOptions *javascript.JestOptions `field:"optional" json:"jestOptions" yaml:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `field:"optional" json:"npmignore" yaml:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled *bool `field:"optional" json:"npmignoreEnabled" yaml:"npmignoreEnabled"`
	// Configuration options for .npmignore file.
	// Experimental.
	NpmIgnoreOptions *projen.IgnoreFileOptions `field:"optional" json:"npmIgnoreOptions" yaml:"npmIgnoreOptions"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Experimental.
	Package *bool `field:"optional" json:"package" yaml:"package"`
	// Setup prettier.
	// Experimental.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Prettier options.
	// Experimental.
	PrettierOptions *javascript.PrettierOptions `field:"optional" json:"prettierOptions" yaml:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency *bool `field:"optional" json:"projenDevDependency" yaml:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Experimental.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options for .projenrc.js.
	// Experimental.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate *bool `field:"optional" json:"pullRequestTemplate" yaml:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents *[]*string `field:"optional" json:"pullRequestTemplateContents" yaml:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Experimental.
	Release *bool `field:"optional" json:"release" yaml:"release"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm *bool `field:"optional" json:"releaseToNpm" yaml:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `field:"optional" json:"releaseWorkflow" yaml:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps *[]*workflows.JobStep `field:"optional" json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Experimental.
	WorkflowGitIdentity *github.GitIdentity `field:"optional" json:"workflowGitIdentity" yaml:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Enable Node.js package cache in GitHub workflows.
	// Experimental.
	WorkflowPackageCache *bool `field:"optional" json:"workflowPackageCache" yaml:"workflowPackageCache"`
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
	//   config.module.unknownContextCritical = false;
	// }
	// ```.
	// See: https://github.com/timarney/react-app-rewired
	//
	// Experimental.
	Rewire *map[string]interface{} `field:"optional" json:"rewire" yaml:"rewire"`
	// Generate one-time sample in `src/` and `public/` if there are no files there.
	// Experimental.
	SampleCode *bool `field:"optional" json:"sampleCode" yaml:"sampleCode"`
	// Source directory.
	// Experimental.
	Srcdir *string `field:"optional" json:"srcdir" yaml:"srcdir"`
}

