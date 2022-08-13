package java


// Options for `MavenCompile`.
// Experimental.
type MavenCompileOptions struct {
	// Source language version.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Target JVM version.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

