package java


// Options for `MavenPackage`.
// Experimental.
type MavenPackagingOptions struct {
	// Where to place the package output?
	// Experimental.
	Distdir *string `field:"optional" json:"distdir" yaml:"distdir"`
	// Include javadocs jar in package.
	// Experimental.
	Javadocs *bool `field:"optional" json:"javadocs" yaml:"javadocs"`
	// Exclude source files from docs.
	// Experimental.
	JavadocsExclude *[]*string `field:"optional" json:"javadocsExclude" yaml:"javadocsExclude"`
	// Include sources jar in package.
	// Experimental.
	Sources *bool `field:"optional" json:"sources" yaml:"sources"`
}

