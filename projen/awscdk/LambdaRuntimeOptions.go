package awscdk


// Options for the AWS Lambda function runtime.
// Experimental.
type LambdaRuntimeOptions struct {
	// Packages that are considered externals by default when bundling.
	// Default: ['@aws-sdk/*'].
	//
	// Experimental.
	DefaultExternals *[]*string `field:"optional" json:"defaultExternals" yaml:"defaultExternals"`
}

