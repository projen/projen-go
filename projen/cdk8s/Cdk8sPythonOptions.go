package cdk8s

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/python"
	"github.com/projen/projen-go/projen/typescript"
)

// Options for `Cdk8sPythonApp`.
// Experimental.
type Cdk8sPythonOptions struct {
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
	// Generate a project tree file (`.projen/tree.json`) that shows all components and their relationships. Useful for understanding your project structure and debugging.
	// Default: false.
	//
	// Experimental.
	ProjectTree *bool `field:"optional" json:"projectTree" yaml:"projectTree"`
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
	// Minimum version of the cdk8s to depend on.
	// Default: "2.3.33"
	//
	// Experimental.
	Cdk8sVersion *string `field:"required" json:"cdk8sVersion" yaml:"cdk8sVersion"`
	// Minimum version of the cdk8s-cli to depend on.
	// Default: "2.0.28"
	//
	// Experimental.
	Cdk8sCliVersion *string `field:"optional" json:"cdk8sCliVersion" yaml:"cdk8sCliVersion"`
	// Use pinned version instead of caret version for cdk8s-cli.
	//
	// You can use this to prevent yarn to mix versions for your CDK8s package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Default: false.
	//
	// Experimental.
	Cdk8sCliVersionPinning *bool `field:"optional" json:"cdk8sCliVersionPinning" yaml:"cdk8sCliVersionPinning"`
	// Include cdk8s-plus.
	// Default: true.
	//
	// Experimental.
	Cdk8sPlus *bool `field:"optional" json:"cdk8sPlus" yaml:"cdk8sPlus"`
	// Minimum version of the cdk8s-plus-XX to depend on.
	// Default: "2.0.0-rc.26"
	//
	// Experimental.
	Cdk8sPlusVersion *string `field:"optional" json:"cdk8sPlusVersion" yaml:"cdk8sPlusVersion"`
	// Use pinned version instead of caret version for cdk8s-plus-17.
	//
	// You can use this to prevent yarn to mix versions for your CDK8s package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Default: false.
	//
	// Experimental.
	Cdk8sPlusVersionPinning *bool `field:"optional" json:"cdk8sPlusVersionPinning" yaml:"cdk8sPlusVersionPinning"`
	// Use pinned version instead of caret version for cdk8s.
	//
	// You can use this to prevent yarn to mix versions for your CDK8s package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Default: false.
	//
	// Experimental.
	Cdk8sVersionPinning *bool `field:"optional" json:"cdk8sVersionPinning" yaml:"cdk8sVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Default: "10.1.42"
	//
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// Use pinned version instead of caret version for constructs.
	//
	// You can use this to prevent yarn to mix versions for your consructs package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Default: false.
	//
	// Experimental.
	ConstructsVersionPinning *bool `field:"optional" json:"constructsVersionPinning" yaml:"constructsVersionPinning"`
	// The cdk8s-plus library depends of Kubernetes minor version For example, cdk8s-plus-22 targets kubernetes version 1.22.0 cdk8s-plus-21 targets kubernetes version 1.21.0.
	// Default: 22.
	//
	// Experimental.
	K8sMinorVersion *float64 `field:"optional" json:"k8sMinorVersion" yaml:"k8sMinorVersion"`
	// The CDK8s app's entrypoint.
	// Default: "app.py"
	//
	// Experimental.
	AppEntrypoint *string `field:"optional" json:"appEntrypoint" yaml:"appEntrypoint"`
	// Import additional specs.
	// Default: - no additional specs imported.
	//
	// Experimental.
	Cdk8sImports *[]*string `field:"optional" json:"cdk8sImports" yaml:"cdk8sImports"`
	// Import a specific Kubernetes spec version.
	// Default: - Use the cdk8s default.
	//
	// Experimental.
	K8sSpecVersion *string `field:"optional" json:"k8sSpecVersion" yaml:"k8sSpecVersion"`
}

