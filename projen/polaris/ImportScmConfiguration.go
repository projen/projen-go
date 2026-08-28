package polaris


// Experimental.
type ImportScmConfiguration struct {
	// Additional arguments to pass to cov-import-scm following capture.
	// Experimental.
	CovImportScmArgs *[]*string `field:"optional" json:"covImportScmArgs" yaml:"covImportScmArgs"`
	// Regular expression that specifies the set of files for which to import change information.
	// Experimental.
	FilenameRegex *string `field:"optional" json:"filenameRegex" yaml:"filenameRegex"`
	// Delay in milliseconds between calls to the underlying SCM.
	// Experimental.
	MsDelay *float64 `field:"optional" json:"msDelay" yaml:"msDelay"`
	// The name of the source control management system.
	// Experimental.
	Scm *string `field:"optional" json:"scm" yaml:"scm"`
}

