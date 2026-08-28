package polaris


// Experimental.
type DirectivesConfiguration struct {
	// Security directives configuration to use during the analysis.
	//
	// This key is mutually exclusive with the "file" key and is specified in the case where the user wants to in-line the security directives configuration in the file.
	// Experimental.
	Config *DirectivesConfigurationConfig `field:"optional" json:"config" yaml:"config"`
	// File containing security directives to use during the analysis.
	//
	// This key is mutually exclusive with the "config" key.
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
}

