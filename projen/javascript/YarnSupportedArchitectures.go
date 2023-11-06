package javascript


// https://yarnpkg.com/configuration/yarnrc#supportedArchitectures.
// Experimental.
type YarnSupportedArchitectures struct {
	// Experimental.
	Cpu *[]*string `field:"optional" json:"cpu" yaml:"cpu"`
	// Experimental.
	Libc *[]*string `field:"optional" json:"libc" yaml:"libc"`
	// Experimental.
	Os *[]*string `field:"optional" json:"os" yaml:"os"`
}

