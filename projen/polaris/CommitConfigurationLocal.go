package polaris


// Local configuration to use when saving defects to the local file system.
// Experimental.
type CommitConfigurationLocal struct {
	// Directory (for "html" format) or file (for "json" format) in which to save defects.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Format in which to save defects.
	//
	// Either "html" or "json".
	// Experimental.
	Format CommitConfigurationLocalFormat `field:"optional" json:"format" yaml:"format"`
}

