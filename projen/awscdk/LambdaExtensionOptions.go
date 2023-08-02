package awscdk

import (
	"github.com/projen/projen-go/projen/javascript"
)

// Options for creating lambda extensions.
// Experimental.
type LambdaExtensionOptions struct {
	// Bundling options for this AWS Lambda extension.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Default: - defaults.
	//
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// The extension's compatible runtimes.
	// Experimental.
	CompatibleRuntimes *[]LambdaRuntime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// A path from the project root directory to a TypeScript file which contains the AWS Lambda extension entrypoint (stand-alone script).
	//
	// This is relative to the root directory of the project.
	//
	// Example:
	//   "src/subdir/foo.lambda-extension.ts"
	//
	// Experimental.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// The name of the generated TypeScript source file.
	//
	// This file should also be
	// under the source tree.
	// Default: - The name of the entrypoint file, with the `-layer-version.ts`
	// suffix instead of `.lambda-extension.ts`.
	//
	// Experimental.
	ConstructFile *string `field:"optional" json:"constructFile" yaml:"constructFile"`
	// The name of the generated `lambda.LayerVersion` subclass.
	// Default: - A pascal cased version of the name of the entrypoint file, with
	// the extension `LayerVersion` (e.g. `AppConfigLayerVersion`).
	//
	// Experimental.
	ConstructName *string `field:"optional" json:"constructName" yaml:"constructName"`
	// Name of the extension.
	// Default: - Derived from the entrypoint filename.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

