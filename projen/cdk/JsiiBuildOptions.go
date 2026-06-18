package cdk

import (
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/release"
)

// Options for `JsiiBuild`.
// Experimental.
type JsiiBuildOptions struct {
	// Options for publishing npm package to AWS CodeArtifact.
	// Default: - undefined.
	//
	// Experimental.
	CodeArtifactOptions *release.CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Automatically run API compatibility test against the latest version published to npm after compilation.
	//
	// - You can manually run compatibility tests using `yarn compat` if this feature is disabled.
	// - You can ignore compatibility failures by adding lines to a ".compatignore" file.
	// Default: false.
	//
	// Experimental.
	Compat *bool `field:"optional" json:"compat" yaml:"compat"`
	// Name of the ignore file for API compatibility tests.
	// Default: ".compatignore"
	//
	// Experimental.
	CompatIgnore *string `field:"optional" json:"compatIgnore" yaml:"compatIgnore"`
	// Emit a compressed version of the assembly.
	// Default: false.
	//
	// Experimental.
	CompressAssembly *bool `field:"optional" json:"compressAssembly" yaml:"compressAssembly"`
	// Generate a MarkDown file describing the jsii API.
	// Default: true.
	//
	// Experimental.
	Docgen *bool `field:"optional" json:"docgen" yaml:"docgen"`
	// File path for generated docs.
	// Default: "API.md"
	//
	// Experimental.
	DocgenFilePath *string `field:"optional" json:"docgenFilePath" yaml:"docgenFilePath"`
	// Accepts a list of glob patterns.
	//
	// Files matching any of those patterns will be excluded from the TypeScript compiler input.
	//
	// By default, jsii will include all *.ts files (except .d.ts files) in the TypeScript compiler input.
	// This can be problematic for example when the package's build or test procedure generates .ts files
	// that cannot be compiled with jsii's compiler settings.
	// Experimental.
	ExcludeTypescript *[]*string `field:"optional" json:"excludeTypescript" yaml:"excludeTypescript"`
	// Version of the jsii compiler to use.
	//
	// Set to "*" if you want to manually manage the version of jsii in your
	// project by managing updates to `package.json` on your own.
	//
	// NOTE: The jsii compiler releases since 5.0.0 are not semantically versioned
	// and should remain on the same minor, so we recommend using a `~` dependency
	// (e.g. `~5.0.0`).
	// Default: "~5.9.0"
	//
	// Experimental.
	JsiiVersion *string `field:"optional" json:"jsiiVersion" yaml:"jsiiVersion"`
	// Whether to use trusted publishing for npm.
	// Default: false.
	//
	// Experimental.
	NpmTrustedPublishing *bool `field:"optional" json:"npmTrustedPublishing" yaml:"npmTrustedPublishing"`
	// Publish Go bindings to a git repository.
	// Default: - no publishing.
	//
	// Experimental.
	PublishToGo *JsiiGoTarget `field:"optional" json:"publishToGo" yaml:"publishToGo"`
	// Publish to maven.
	// Default: - no publishing.
	//
	// Experimental.
	PublishToMaven *JsiiJavaTarget `field:"optional" json:"publishToMaven" yaml:"publishToMaven"`
	// Publish to NuGet.
	// Default: - no publishing.
	//
	// Experimental.
	PublishToNuget *JsiiDotNetTarget `field:"optional" json:"publishToNuget" yaml:"publishToNuget"`
	// Publish to pypi.
	// Default: - no publishing.
	//
	// Experimental.
	PublishToPypi *JsiiPythonTarget `field:"optional" json:"publishToPypi" yaml:"publishToPypi"`
	// Whether to publish to npm.
	// Default: true.
	//
	// Experimental.
	ReleaseToNpm *bool `field:"optional" json:"releaseToNpm" yaml:"releaseToNpm"`
	// The stability of the package.
	// Default: "stable".
	//
	// Experimental.
	Stability *string `field:"optional" json:"stability" yaml:"stability"`
	// The TypescriptConfig that jsii compiles against.
	//
	// Provide a dedicated config (e.g. one named `tsconfig.jsii.json`) to keep
	// jsii compilation separate from the project's editor-facing `tsconfig.json`.
	// jsii requires a comment-free, strict-valid config, so `JsiiBuild` disables
	// comments on whichever config it is given.
	// Default: - the project's tsconfig (i.e. `project.tsconfig`)
	//
	// Experimental.
	Tsconfig javascript.TypescriptConfig `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Level of tsconfig validation jsii should perform on the user-provided tsconfig.
	//
	// Only relevant when the project synthesizes its own tsconfig
	// (i.e. `disableTsconfig` is not set on the TypeScriptProject).
	// See: https://aws.github.io/jsii/user-guides/lib-author/configuration/#validatetsconfig
	//
	// Default: ValidateTsconfig.STRICT
	//
	// Experimental.
	ValidateTsconfig ValidateTsconfig `field:"optional" json:"validateTsconfig" yaml:"validateTsconfig"`
	// Additional steps to run before packaging in workflows.
	// Default: [].
	//
	// Experimental.
	WorkflowBootstrapSteps *[]*workflows.Step `field:"optional" json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The node version to use in packaging workflows.
	// Default: "lts/*".
	//
	// Experimental.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Relative path of the package within the repository.
	//
	// This is used in monorepo setups where the package is not at the root.
	// Packaging steps will extract build artifacts into this subdirectory.
	// Default: "." - root of the repository
	//
	// Experimental.
	WorkspaceDirectory *string `field:"optional" json:"workspaceDirectory" yaml:"workspaceDirectory"`
}

