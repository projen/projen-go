package javascript


// The `devEngines` field in `package.json`.
// See: https://docs.npmjs.com/cli/v10/configuring-npm/package-json#devengines
//
// Experimental.
type DevEngines struct {
	// Supported CPU architectures.
	// Experimental.
	Cpu interface{} `field:"optional" json:"cpu" yaml:"cpu"`
	// Supported C standard libraries.
	// Experimental.
	Libc interface{} `field:"optional" json:"libc" yaml:"libc"`
	// Supported operating systems.
	// Experimental.
	Os interface{} `field:"optional" json:"os" yaml:"os"`
	// Supported package managers.
	// Experimental.
	PackageManager interface{} `field:"optional" json:"packageManager" yaml:"packageManager"`
	// Supported JavaScript runtimes.
	// Experimental.
	Runtime interface{} `field:"optional" json:"runtime" yaml:"runtime"`
}

