package projen


// Git configuration options.
// Experimental.
type GitOptions struct {
	// File patterns to mark as stored in Git LFS.
	// Default: - No files stored in LFS.
	//
	// Experimental.
	LfsPatterns *[]*string `field:"optional" json:"lfsPatterns" yaml:"lfsPatterns"`
}

