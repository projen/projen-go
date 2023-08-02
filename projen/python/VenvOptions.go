package python


// Options for venv.
// Experimental.
type VenvOptions struct {
	// Name of directory to store the environment in.
	// Default: ".env"
	//
	// Experimental.
	Envdir *string `field:"optional" json:"envdir" yaml:"envdir"`
	// Path to the python executable to use.
	// Default: "python".
	//
	// Experimental.
	PythonExec *string `field:"optional" json:"pythonExec" yaml:"pythonExec"`
}

