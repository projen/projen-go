package python


// Declares any Python level dependencies that must be installed in order to run the project’s build system successfully.
// Experimental.
type BuildSystem struct {
	// List of strings following the version specifier specification, representing dependencies required to execute the build system.
	// Experimental.
	Requires *[]*string `field:"required" json:"requires" yaml:"requires"`
	// list of directories to prepend to `sys.path` when loading the build backend, relative to project root.
	// Experimental.
	BackendPath *[]*string `field:"optional" json:"backendPath" yaml:"backendPath"`
	// String is formatted following the same `module:object` syntax as a `setuptools` entry point.
	//
	// It’s also legal to leave out the `:object` part.
	// Experimental.
	BuildBackend *string `field:"optional" json:"buildBackend" yaml:"buildBackend"`
}

