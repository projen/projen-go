package polaris


// Security directives configuration to use during the analysis.
//
// This key is mutually exclusive with the "file" key and is specified in the case where the user wants to in-line the security directives configuration in the file.
// Experimental.
type DirectivesConfigurationConfig struct {
	// Specify a particular analysis behavior.
	// Experimental.
	Directives *[]interface{} `field:"required" json:"directives" yaml:"directives"`
	// Language or language family to which directives apply.
	// Experimental.
	Language *string `field:"required" json:"language" yaml:"language"`
	// Version of the directives format.
	// Experimental.
	FormatVersion *float64 `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// Must be the string "Coverity analysis configuration".
	// Experimental.
	Type DirectivesConfigurationConfigType `field:"optional" json:"type" yaml:"type"`
}

