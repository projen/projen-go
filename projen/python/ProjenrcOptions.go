package python


// Options for `Projenrc`.
// Experimental.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Default: ".projenrc.py"
	//
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// The projen version to use.
	// Default: - current version.
	//
	// Experimental.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Path to the python executable to use.
	// Default: "python".
	//
	// Experimental.
	PythonExec *string `field:"optional" json:"pythonExec" yaml:"pythonExec"`
}

