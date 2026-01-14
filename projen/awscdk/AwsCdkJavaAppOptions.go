package awscdk

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/java"
)

// Experimental.
type AwsCdkJavaAppOptions struct {
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
	// Default: "my-app".
	//
	// Experimental.
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
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
	// Default: "org.acme"
	//
	// Experimental.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
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
	// Default: "0.1.0"
	//
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Description of a project is always good.
	//
	// Although this should not replace
	// formal documentation, a quick comment to any readers of the POM is always
	// helpful.
	// Default: undefined.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Project packaging format.
	// Default: "jar".
	//
	// Experimental.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// A Parent Pom can be used to have a child project inherit properties/plugins/ect in order to reduce duplication and keep standards across a large amount of repos.
	// Default: undefined.
	//
	// Experimental.
	ParentPom *java.ParentPom `field:"optional" json:"parentPom" yaml:"parentPom"`
	// The URL, like the name, is not required.
	//
	// This is a nice gesture for
	// projects users, however, so that they know where the project lives.
	// Default: undefined.
	//
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Compile options.
	// Default: - defaults.
	//
	// Experimental.
	CompileOptions *java.MavenCompileOptions `field:"optional" json:"compileOptions" yaml:"compileOptions"`
	// List of runtime dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addDependency()`.
	// Default: [].
	//
	// Experimental.
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// Final artifact output directory.
	// Default: "dist/java".
	//
	// Experimental.
	Distdir *string `field:"optional" json:"distdir" yaml:"distdir"`
	// Include junit tests.
	// Default: true.
	//
	// Experimental.
	Junit *bool `field:"optional" json:"junit" yaml:"junit"`
	// junit options.
	// Default: - defaults.
	//
	// Experimental.
	JunitOptions *java.JunitOptions `field:"optional" json:"junitOptions" yaml:"junitOptions"`
	// Packaging options.
	// Default: - defaults.
	//
	// Experimental.
	PackagingOptions *java.MavenPackagingOptions `field:"optional" json:"packagingOptions" yaml:"packagingOptions"`
	// Use projenrc in java.
	//
	// This will install `projen` as a java dependency and will add a `synth` task which
	// will compile & execute `main()` from `src/main/java/projenrc.java`.
	// Default: true.
	//
	// Experimental.
	ProjenrcJava *bool `field:"optional" json:"projenrcJava" yaml:"projenrcJava"`
	// Options related to projenrc in java.
	// Default: - default options.
	//
	// Experimental.
	ProjenrcJavaOptions *java.ProjenrcOptions `field:"optional" json:"projenrcJavaOptions" yaml:"projenrcJavaOptions"`
	// List of test dependencies for this project.
	//
	// Dependencies use the format: `<groupId>/<artifactId>@<semver>`
	//
	// Additional dependencies can be added via `project.addTestDependency()`.
	// Default: [].
	//
	// Experimental.
	TestDeps *[]*string `field:"optional" json:"testDeps" yaml:"testDeps"`
	// Include sample code and test if the relevant directories don't exist.
	// Default: true.
	//
	// Experimental.
	Sample *bool `field:"optional" json:"sample" yaml:"sample"`
	// The java package to use for the code sample.
	// Default: "org.acme"
	//
	// Experimental.
	SampleJavaPackage *string `field:"optional" json:"sampleJavaPackage" yaml:"sampleJavaPackage"`
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
	// The name of the Java class with the static `main()` method.
	//
	// This method
	// should call `app.synth()` on the CDK app.
	// Default: "org.acme.MyApp"
	//
	// Experimental.
	MainClass *string `field:"required" json:"mainClass" yaml:"mainClass"`
}

