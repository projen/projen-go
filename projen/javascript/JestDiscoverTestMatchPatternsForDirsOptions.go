package javascript


// Options for discoverTestMatchPatternsForDirs.
// Experimental.
type JestDiscoverTestMatchPatternsForDirsOptions struct {
	// The file extension pattern to use.
	//
	// Defaults to "[jt]s?(x)".
	// Experimental.
	FileExtensionPattern *string `field:"optional" json:"fileExtensionPattern" yaml:"fileExtensionPattern"`
}

