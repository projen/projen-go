package uvconfig


// Cache control configuration for an index.
// Experimental.
type IndexCacheControl struct {
	// Cache control header for Simple API requests.
	// Experimental.
	Api *string `field:"optional" json:"api" yaml:"api"`
	// Cache control header for file downloads.
	// Experimental.
	Files *string `field:"optional" json:"files" yaml:"files"`
}

