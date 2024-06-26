package awscdk

import (
	"github.com/projen/projen-go/projen/javascript"
)

// Common options for `LambdaFunction`.
//
// Applies to all functions in
// auto-discovery.
// Experimental.
type LambdaFunctionCommonOptions struct {
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
	// Default: Runtime.NODEJS_18_X
	//
	// Experimental.
	Runtime LambdaRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

