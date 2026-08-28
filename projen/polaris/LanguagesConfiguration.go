package polaris


// Experimental.
type LanguagesConfiguration struct {
	// Specifies the languages for which the source code should be excluded in the capture.
	//
	// This key is mutually exclusive with the "include" key.
	// Experimental.
	Exclude *[]LanguagesConfigurationExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// Specifies the languages for which the source code should be included in the capture.
	//
	// This key is mutually exclusive with the "exclude" key.
	// Experimental.
	Include *[]LanguagesConfigurationInclude `field:"optional" json:"include" yaml:"include"`
}

