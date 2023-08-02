package awscdk


// Experimental.
type IntegrationTestCommonOptions struct {
	// Destroy the test app after a successful deployment.
	//
	// If disabled, leaves the
	// app deployed in the dev account.
	// Default: true.
	//
	// Experimental.
	DestroyAfterDeploy *bool `field:"optional" json:"destroyAfterDeploy" yaml:"destroyAfterDeploy"`
	// Enables path metadata, adding `aws:cdk:path`, with the defining construct's path, to the CloudFormation metadata for each synthesized resource.
	// Default: false.
	//
	// Experimental.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
}

