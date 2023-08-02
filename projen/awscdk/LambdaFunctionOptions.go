package awscdk

import (
	"github.com/projen/projen-go/projen/javascript"
)

// Options for `Function`.
// Experimental.
type LambdaFunctionOptions struct {
	// Whether to automatically reuse TCP connections when working with the AWS SDK for JavaScript.
	//
	// This sets the `AWS_NODEJS_CONNECTION_REUSE_ENABLED` environment variable
	// to `1`.
	//
	// Not applicable when `edgeLambda` is set to `true` because environment
	// variables are not supported in Lambda@Edge.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/node-reusing-connections.html
	//
	// Default: true.
	//
	// Experimental.
	AwsSdkConnectionReuse *bool `field:"optional" json:"awsSdkConnectionReuse" yaml:"awsSdkConnectionReuse"`
	// Bundling options for this AWS Lambda function.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Default: - defaults.
	//
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// Whether to create a `cloudfront.experimental.EdgeFunction` instead of a `lambda.Function`.
	// Default: false.
	//
	// Experimental.
	EdgeLambda *bool `field:"optional" json:"edgeLambda" yaml:"edgeLambda"`
	// The node.js version to target.
	// Default: Runtime.NODEJS_16_X
	//
	// Experimental.
	Runtime LambdaRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// A path from the project root directory to a TypeScript file which contains the AWS Lambda handler entrypoint (exports a `handler` function).
	//
	// This is relative to the root directory of the project.
	//
	// Example:
	//   "src/subdir/foo.lambda.ts"
	//
	// Experimental.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// The name of the generated TypeScript source file.
	//
	// This file should also be
	// under the source tree.
	// Default: - The name of the entrypoint file, with the `-function.ts` suffix
	// instead of `.lambda.ts`.
	//
	// Experimental.
	ConstructFile *string `field:"optional" json:"constructFile" yaml:"constructFile"`
	// The name of the generated `lambda.Function` subclass.
	// Default: - A pascal cased version of the name of the entrypoint file, with
	// the extension `Function` (e.g. `ResizeImageFunction`).
	//
	// Experimental.
	ConstructName *string `field:"optional" json:"constructName" yaml:"constructName"`
}

