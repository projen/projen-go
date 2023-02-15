package gitlab


// Use this construct to generate a new key when one or two specific files change.
// See: https://docs.gitlab.com/ee/ci/yaml/#cachekeyfiles
//
// Experimental.
type CacheKeyFiles struct {
	// The files that are checked against.
	//
	// If the SHA checksum changes, the cache becomes invalid.
	// Experimental.
	Files *[]*string `field:"required" json:"files" yaml:"files"`
	// Adds a custom prefix to the checksums computed.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

