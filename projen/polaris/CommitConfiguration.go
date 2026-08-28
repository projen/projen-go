package polaris


// Specifies where the analysis results should be sent.
// Experimental.
type CommitConfiguration struct {
	// Coverity Connect configuration to use when committing defects to Coverity Connect.
	// Experimental.
	Connect *CommitConfigurationConnect `field:"optional" json:"connect" yaml:"connect"`
	// Local configuration to use when saving defects to the local file system.
	// Experimental.
	Local *CommitConfigurationLocal `field:"optional" json:"local" yaml:"local"`
	// Software Risk Manager configuration to use when storing defects in Software Risk Manager.
	// Experimental.
	Srm *CommitConfigurationSrm `field:"optional" json:"srm" yaml:"srm"`
}

