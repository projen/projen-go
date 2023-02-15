package cdk8s


// Options for `Cdk8sDeps`.
// Experimental.
type Cdk8sDepsCommonOptions struct {
	// Minumum version of the cdk8s to depend on.
	// Experimental.
	Cdk8sVersion *string `field:"required" json:"cdk8sVersion" yaml:"cdk8sVersion"`
	// Minumum version of the cdk8s-cli to depend on.
	// Experimental.
	Cdk8sCliVersion *string `field:"optional" json:"cdk8sCliVersion" yaml:"cdk8sCliVersion"`
	// Use pinned version instead of caret version for cdk8s-cli.
	//
	// You can use this to prevent yarn to mix versions for your CDK8s package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	Cdk8sCliVersionPinning *bool `field:"optional" json:"cdk8sCliVersionPinning" yaml:"cdk8sCliVersionPinning"`
	// Include cdk8s-plus.
	// Experimental.
	Cdk8sPlus *bool `field:"optional" json:"cdk8sPlus" yaml:"cdk8sPlus"`
	// Minumum version of the cdk8s-plus-XX to depend on.
	// Experimental.
	Cdk8sPlusVersion *string `field:"optional" json:"cdk8sPlusVersion" yaml:"cdk8sPlusVersion"`
	// Use pinned version instead of caret version for cdk8s-plus-17.
	//
	// You can use this to prevent yarn to mix versions for your CDK8s package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	Cdk8sPlusVersionPinning *bool `field:"optional" json:"cdk8sPlusVersionPinning" yaml:"cdk8sPlusVersionPinning"`
	// Use pinned version instead of caret version for cdk8s.
	//
	// You can use this to prevent yarn to mix versions for your CDK8s package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	Cdk8sVersionPinning *bool `field:"optional" json:"cdk8sVersionPinning" yaml:"cdk8sVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// Use pinned version instead of caret version for constructs.
	//
	// You can use this to prevent yarn to mix versions for your consructs package and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	ConstructsVersionPinning *bool `field:"optional" json:"constructsVersionPinning" yaml:"constructsVersionPinning"`
	// The cdk8s-plus library depends of Kubernetes minor version For example, cdk8s-plus-22 targets kubernetes version 1.22.0 cdk8s-plus-21 targets kubernetes version 1.21.0.
	// Experimental.
	K8sMinorVersion *float64 `field:"optional" json:"k8sMinorVersion" yaml:"k8sMinorVersion"`
}

