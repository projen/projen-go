package awscdk


// Language-specific AWS CDK package names.
// Experimental.
type AwsCdkPackageNames struct {
	// Fully qualified name of the assertions library package.
	// Experimental.
	Assertions *string `field:"required" json:"assertions" yaml:"assertions"`
	// Fully qualified name of the constructs library package.
	// Experimental.
	Constructs *string `field:"required" json:"constructs" yaml:"constructs"`
	// Fully qualified name of the core framework package for CDKv1.
	// Experimental.
	CoreV1 *string `field:"required" json:"coreV1" yaml:"coreV1"`
	// Fully qualified name of the core framework package for CDKv2.
	// Experimental.
	CoreV2 *string `field:"required" json:"coreV2" yaml:"coreV2"`
	// Fully qualified name of the assert library package Can be empty as it's only really available for javascript projects.
	// Experimental.
	Assert *string `field:"optional" json:"assert" yaml:"assert"`
}

