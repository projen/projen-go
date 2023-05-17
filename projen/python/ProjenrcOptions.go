package python


// Options for `Projenrc`.
// Experimental.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// The projen version to use.
	// Experimental.
	ProjenVersion *string `field:"optional" json:"projenVersion" yaml:"projenVersion"`
	// Path to the python executable to use.
	// Experimental.
	PythonExec *string `field:"optional" json:"pythonExec" yaml:"pythonExec"`
}

