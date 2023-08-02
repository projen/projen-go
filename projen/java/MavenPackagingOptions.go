package java


// Options for `MavenPackage`.
// Experimental.
type MavenPackagingOptions struct {
	// Where to place the package output?
	// Default: "dist/java".
	//
	// Experimental.
	Distdir *string `field:"optional" json:"distdir" yaml:"distdir"`
	// Include javadocs jar in package.
	// Default: true.
	//
	// Experimental.
	Javadocs *bool `field:"optional" json:"javadocs" yaml:"javadocs"`
	// Exclude source files from docs.
	// Default: [].
	//
	// Experimental.
	JavadocsExclude *[]*string `field:"optional" json:"javadocsExclude" yaml:"javadocsExclude"`
	// Include sources jar in package.
	// Default: true.
	//
	// Experimental.
	Sources *bool `field:"optional" json:"sources" yaml:"sources"`
}

