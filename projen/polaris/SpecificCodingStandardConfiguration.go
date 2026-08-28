package polaris


// Experimental.
type SpecificCodingStandardConfiguration struct {
	// This key specifies the coding standard configuration for the given coding standard.
	//
	// The actual type of this key is specific to the particular coding standard. This key is mutually exclusive with the "file" key. A temporary configuration file will be generated containing the in-line configuration and then passed to "cov-analyze" using the "--coding-standard-config <config_file>" option.
	// Experimental.
	Config *ResolvedCodingStandardConfiguration `field:"optional" json:"config" yaml:"config"`
	// This specifies the filename containing the configuration to use for the corresponding coding standard.
	//
	// This key is mutually exclusive with the "config" key.
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// This key specifies the name of a "pre-canned" coding standard configuration to use.
	//
	// The available pre-canned coding standard configurations depend on the coding standard in question. Refer to Coverity's documentation for details on the "pre-canned" configurations.
	// Experimental.
	PreCanned *string `field:"optional" json:"preCanned" yaml:"preCanned"`
}

