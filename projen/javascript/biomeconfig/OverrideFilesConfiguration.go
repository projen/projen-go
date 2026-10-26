package biomeconfig


// Experimental.
type OverrideFilesConfiguration struct {
	// The maximum source-file size in bytes for matched files.
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
}

