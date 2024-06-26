package java


// Options for `Junit`.
// Experimental.
type JunitOptions struct {
	// Java pom.
	// Experimental.
	Pom Pom `field:"required" json:"pom" yaml:"pom"`
	// Java package for test sample.
	// Default: "org.acme"
	//
	// Experimental.
	SampleJavaPackage *string `field:"optional" json:"sampleJavaPackage" yaml:"sampleJavaPackage"`
	// Junit version.
	// Default: "5.7.0"
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

