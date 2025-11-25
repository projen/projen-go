package awscdk

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/python"
	"github.com/projen/projen-go/projen/typescript"
)

// Options for `AwsCdkPythonApp`.
// Experimental.
type AwsCdkPythonAppOptions struct {
	// This is the name of your project.
	// Default: $BASEDIR.
	//
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to commit the managed files by default.
	// Default: true.
	//
	// Experimental.
	CommitGenerated *bool `field:"optional" json:"commitGenerated" yaml:"commitGenerated"`
	// Configuration options for .gitignore file.
	// Experimental.
	GitIgnoreOptions *projen.IgnoreFileOptions `field:"optional" json:"gitIgnoreOptions" yaml:"gitIgnoreOptions"`
	// Configuration options for git.
	// Experimental.
	GitOptions *projen.GitOptions `field:"optional" json:"gitOptions" yaml:"gitOptions"`
	// Configure logging options such as verbosity.
	// Default: {}.
	//
	// Experimental.
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
	// Experimental.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `field:"optional" json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Default: "npx projen".
	//
	// Experimental.
	ProjenCommand *string `field:"optional" json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Default: false.
	//
	// Experimental.
	ProjenrcJson *bool `field:"optional" json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Default: - default options.
	//
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcJsonOptions `field:"optional" json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Use renovatebot to handle dependency upgrades.
	// Default: false.
	//
	// Experimental.
	Renovatebot *bool `field:"optional" json:"renovatebot" yaml:"renovatebot"`
	// Options for renovatebot.
	// Default: - default options.
	//
	// Experimental.
	RenovatebotOptions *projen.RenovatebotOptions `field:"optional" json:"renovatebotOptions" yaml:"renovatebotOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Default: - auto approve is disabled.
	//
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `field:"optional" json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Enable automatic merging on GitHub.
	//
	// Has no effect if `github.mergify`
	// is set to false.
	// Default: true.
	//
	// Experimental.
	AutoMerge *bool `field:"optional" json:"autoMerge" yaml:"autoMerge"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` or `autoMerge` is set to false.
	// Default: - see defaults in `AutoMergeOptions`.
	//
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `field:"optional" json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Default: - true, but false for subprojects.
	//
	// Experimental.
	Clobber *bool `field:"optional" json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Default: false.
	//
	// Experimental.
	DevContainer *bool `field:"optional" json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Default: true.
	//
	// Experimental.
	Github *bool `field:"optional" json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Default: - see GitHubOptions.
	//
	// Experimental.
	GithubOptions *github.GitHubOptions `field:"optional" json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Default: false.
	//
	// Experimental.
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
	// Experimental.
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
	// Experimental.
	Readme *projen.SampleReadmeProps `field:"optional" json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Default: false.
	//
	// Experimental.
	Stale *bool `field:"optional" json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Default: - see defaults in `StaleOptions`.
	//
	// Experimental.
	StaleOptions *github.StaleOptions `field:"optional" json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Default: true.
	//
	// Experimental.
	Vscode *bool `field:"optional" json:"vscode" yaml:"vscode"`
	// Author's e-mail.
	// Default: $GIT_USER_EMAIL.
	//
	// Experimental.
	AuthorEmail *string `field:"required" json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Default: $GIT_USER_NAME.
	//
	// Experimental.
	AuthorName *string `field:"required" json:"authorName" yaml:"authorName"`
	// Version of the package.
	// Default: "0.1.0"
	//
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// A list of PyPI trove classifiers that describe the project.
	// See: https://pypi.org/classifiers/
	//
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// A short description of the package.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A URL to the website of the project.
	// Experimental.
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
	// License of this package as an SPDX identifier.
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Package name.
	// Experimental.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Additional options to set for poetry if using poetry.
	// Experimental.
	PoetryOptions *python.PoetryPyprojectOptionsWithoutDeps `field:"optional" json:"poetryOptions" yaml:"poetryOptions"`
	// Additional fields to pass in the setup() function if using setuptools.
	// Experimental.
	SetupConfig *map[string]interface{} `field:"optional" json:"setupConfig" yaml:"setupConfig"`
	// Additional options to set for uv if using uv.
	// Experimental.
	UvOptions *python.UvOptions `field:"optional" json:"uvOptions" yaml:"uvOptions"`
	// Path to the python executable to use.
	// Default: "python".
	//
	// Experimental.
	PythonExec *string `field:"optional" json:"pythonExec" yaml:"pythonExec"`
	// Name of the python package as used in imports and filenames.
	//
	// Must only consist of alphanumeric characters and underscores.
	// Default: $PYTHON_MODULE_NAME.
	//
	// Experimental.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Default: [].
	//
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// List of dev dependencies for this project.
	//
	// Dependencies use the format: `<module>@<semver>`
	//
	// Additional dependencies can be added via `project.addDevDependency()`.
	// Default: [].
	//
	// Experimental.
	DevDeps *[]*string `field:"optional" json:"devDeps" yaml:"devDeps"`
	// Use pip with a requirements.txt file to track project dependencies.
	// Default: - true, unless poetry is true, then false.
	//
	// Experimental.
	Pip *bool `field:"optional" json:"pip" yaml:"pip"`
	// Use poetry to manage your project dependencies, virtual environment, and (optional) packaging/publishing.
	//
	// This feature is incompatible with pip, setuptools, or venv.
	// If you set this option to `true`, then pip, setuptools, and venv must be set to `false`.
	// Default: false.
	//
	// Experimental.
	Poetry *bool `field:"optional" json:"poetry" yaml:"poetry"`
	// Use projenrc in javascript.
	//
	// This will install `projen` as a JavaScript dependency and add a `synth`
	// task which will run `.projenrc.js`.
	// Default: false.
	//
	// Experimental.
	ProjenrcJs *bool `field:"optional" json:"projenrcJs" yaml:"projenrcJs"`
	// Options related to projenrc in JavaScript.
	// Default: - default options.
	//
	// Experimental.
	ProjenrcJsOptions *javascript.ProjenrcOptions `field:"optional" json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Use projenrc in Python.
	//
	// This will install `projen` as a Python dependency and add a `synth`
	// task which will run `.projenrc.py`.
	// Default: true.
	//
	// Experimental.
	ProjenrcPython *bool `field:"optional" json:"projenrcPython" yaml:"projenrcPython"`
	// Options related to projenrc in python.
	// Default: - default options.
	//
	// Experimental.
	ProjenrcPythonOptions *python.ProjenrcOptions `field:"optional" json:"projenrcPythonOptions" yaml:"projenrcPythonOptions"`
	// Use projenrc in TypeScript.
	//
	// This will create a tsconfig file (default: `tsconfig.projen.json`)
	// and use `ts-node` in the default task to parse the project source files.
	// Default: false.
	//
	// Experimental.
	ProjenrcTs *bool `field:"optional" json:"projenrcTs" yaml:"projenrcTs"`
	// Options related to projenrc in TypeScript.
	// Default: - default options.
	//
	// Experimental.
	ProjenrcTsOptions *typescript.ProjenrcTsOptions `field:"optional" json:"projenrcTsOptions" yaml:"projenrcTsOptions"`
	// Include pytest tests.
	// Default: true.
	//
	// Experimental.
	Pytest *bool `field:"optional" json:"pytest" yaml:"pytest"`
	// pytest options.
	// Default: - defaults.
	//
	// Experimental.
	PytestOptions *python.PytestOptions `field:"optional" json:"pytestOptions" yaml:"pytestOptions"`
	// Include sample code and test if the relevant directories don't exist.
	// Default: true.
	//
	// Experimental.
	Sample *bool `field:"optional" json:"sample" yaml:"sample"`
	// Location of sample tests.
	//
	// Typically the same directory where project tests will be located.
	// Default: "tests".
	//
	// Experimental.
	SampleTestdir *string `field:"optional" json:"sampleTestdir" yaml:"sampleTestdir"`
	// Use setuptools with a setup.py script for packaging and publishing.
	// Default: - true, unless poetry is true, then false.
	//
	// Experimental.
	Setuptools *bool `field:"optional" json:"setuptools" yaml:"setuptools"`
	// Use uv to manage your project dependencies, virtual environment, and (optional) packaging/publishing.
	// Default: false.
	//
	// Experimental.
	Uv *bool `field:"optional" json:"uv" yaml:"uv"`
	// Use venv to manage a virtual environment for installing dependencies inside.
	// Default: - true, unless poetry is true, then false.
	//
	// Experimental.
	Venv *bool `field:"optional" json:"venv" yaml:"venv"`
	// Venv options.
	// Default: - defaults.
	//
	// Experimental.
	VenvOptions *python.VenvOptions `field:"optional" json:"venvOptions" yaml:"venvOptions"`
	// A command to execute before synthesis.
	//
	// This command will be called when
	// running `cdk synth` or when `cdk watch` identifies a change in your source
	// code before redeployment.
	// Default: - no build command.
	//
	// Experimental.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// cdk.out directory.
	// Default: "cdk.out"
	//
	// Experimental.
	Cdkout *string `field:"optional" json:"cdkout" yaml:"cdkout"`
	// Additional context to include in `cdk.json`.
	// Default: - no additional context.
	//
	// Experimental.
	Context *map[string]interface{} `field:"optional" json:"context" yaml:"context"`
	// Feature flags that should be enabled in `cdk.json`.
	//
	// Make sure to double-check any changes to feature flags in `cdk.json` before deploying.
	// Unexpected changes may cause breaking changes in your CDK app.
	// You can overwrite any feature flag by passing it into the context field.
	// Default: - no feature flags are enabled by default.
	//
	// Experimental.
	FeatureFlags ICdkFeatureFlags `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// To protect you against unintended changes that affect your security posture, the AWS CDK Toolkit prompts you to approve security-related changes before deploying them.
	// Default: ApprovalLevel.BROADENING
	//
	// Experimental.
	RequireApproval ApprovalLevel `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Glob patterns to exclude from `cdk watch`.
	// Default: [].
	//
	// Experimental.
	WatchExcludes *[]*string `field:"optional" json:"watchExcludes" yaml:"watchExcludes"`
	// Glob patterns to include in `cdk watch`.
	// Default: [].
	//
	// Experimental.
	WatchIncludes *[]*string `field:"optional" json:"watchIncludes" yaml:"watchIncludes"`
	// Minimum version of the AWS CDK to depend on.
	// Default: "2.1.0"
	//
	// Experimental.
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
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Version range of the AWS CDK CLI to depend on.
	//
	// Can be either a specific version, or an NPM version range.
	//
	// By default, the latest 2.x version will be installed; you can use this
	// option to restrict it to a specific version or version range.
	// Default: "^2".
	//
	// Experimental.
	CdkCliVersion *string `field:"optional" json:"cdkCliVersion" yaml:"cdkCliVersion"`
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
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Default: - for CDK 1.x the default is "3.2.27", for CDK 2.x the default is
	// "10.0.5".
	//
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// The CDK app's entrypoint (relative to the source directory, which is "src" by default).
	// Default: "app.py"
	//
	// Experimental.
	AppEntrypoint *string `field:"optional" json:"appEntrypoint" yaml:"appEntrypoint"`
	// Python sources directory.
	// Default: "tests".
	//
	// Deprecated: Use `sampleTestdir` instead.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
}

