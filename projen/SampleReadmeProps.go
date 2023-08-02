package projen


// SampleReadme Properties.
// Experimental.
type SampleReadmeProps struct {
	// The contents.
	// Default: "# replace this".
	//
	// Experimental.
	Contents *string `field:"optional" json:"contents" yaml:"contents"`
	// The name of the README.md file.
	//
	// Example:
	//   "readme.md"
	//
	// Default: "README.md"
	//
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

