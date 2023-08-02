package java


// Options for `MavenCompile`.
// Experimental.
type MavenCompileOptions struct {
	// Source language version.
	// Default: "1.8"
	//
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Target JVM version.
	// Default: "1.8"
	//
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

