package javascript


// Specifies architectures for which you'd like to install optional dependencies, even if they don't match the architecture of the system running the install.
// Experimental.
type PnpmWorkspaceYamlSchemaSupportedArchitectures struct {
	// Experimental.
	Cpu *[]*string `field:"optional" json:"cpu" yaml:"cpu"`
	// Experimental.
	Libc *[]*string `field:"optional" json:"libc" yaml:"libc"`
	// Experimental.
	Os *[]*string `field:"optional" json:"os" yaml:"os"`
}

