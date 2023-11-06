package javascript


// Configure Yarn Berry.
// Experimental.
type YarnBerryOptions struct {
	// A fully specified version to use for yarn (e.g., x.x.x).
	// Default: - 4.0.1
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The yarnrc configuration.
	// Default: - a blank Yarn RC file.
	//
	// Experimental.
	YarnRcOptions *YarnrcOptions `field:"optional" json:"yarnRcOptions" yaml:"yarnRcOptions"`
	// Should zero-installs be enabled?
	//
	// Learn more at: https://yarnpkg.com/features/caching#zero-installs
	// Default: false.
	//
	// Experimental.
	ZeroInstalls *bool `field:"optional" json:"zeroInstalls" yaml:"zeroInstalls"`
}

