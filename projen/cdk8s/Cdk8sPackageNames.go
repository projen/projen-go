package cdk8s


// Experimental.
type Cdk8sPackageNames struct {
	// Fully qualified name of the core framework package.
	// Experimental.
	Cdk8s *string `field:"required" json:"cdk8s" yaml:"cdk8s"`
	// Fully qualified name of the cdk9s-plus-XX library package.
	// Experimental.
	Cdk8sPlus *string `field:"required" json:"cdk8sPlus" yaml:"cdk8sPlus"`
	// Fully qualified name of the constructs library package.
	// Experimental.
	Constructs *string `field:"required" json:"constructs" yaml:"constructs"`
	// Fully qualified name of the client package.
	//
	// Used only on Node projects.
	// Experimental.
	Cdk8sClient *string `field:"optional" json:"cdk8sClient" yaml:"cdk8sClient"`
}

