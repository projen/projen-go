package awscdk


// Language-specific AWS CDK package names.
// Experimental.
type AwsCdkPackageNames struct {
	// Fully qualified name of the constructs library package.
	// Experimental.
	Constructs *string `field:"required" json:"constructs" yaml:"constructs"`
	// Fully qualified name of the core framework package for CDKv2.
	// Experimental.
	CoreV2 *string `field:"required" json:"coreV2" yaml:"coreV2"`
}

