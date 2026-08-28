package polaris


// Experimental.
type SigmaConfiguration struct {
	// List of check sets to enable.
	// Experimental.
	EnableCheckSet *[]SigmaConfigurationEnableCheckSet `field:"optional" json:"enableCheckSet" yaml:"enableCheckSet"`
	// List of files containing malicious URL patterns.
	// Experimental.
	MaliciousUrlPatternsFile *[]*string `field:"optional" json:"maliciousUrlPatternsFile" yaml:"maliciousUrlPatternsFile"`
}

