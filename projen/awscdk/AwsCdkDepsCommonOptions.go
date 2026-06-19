package awscdk


// Options for `AwsCdkDeps`.
// Experimental.
type AwsCdkDepsCommonOptions struct {
	// Minimum version of the AWS CDK to depend on.
	// Default: "2.189.1"
	//
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
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
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Default: - for CDK 1.x the default is "3.2.27", for CDK 2.x the default is
	// "10.5.1".
	//
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
}

