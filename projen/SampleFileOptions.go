package projen


// Options for the SampleFile object.
// Experimental.
type SampleFileOptions struct {
	// The contents of the file to write.
	// Experimental.
	Contents *string `field:"optional" json:"contents" yaml:"contents"`
	// Absolute path to a file to copy the contents from (does not need to be a text file).
	//
	// If your project is Typescript-based and has configured `testdir` to be a
	// subdirectory of `src`, sample files should outside of the `src` directory,
	// otherwise they may not be copied. For example:
	// ```
	// new SampleFile(this, 'assets/icon.png', { source: path.join(__dirname, '..', 'sample-assets', 'icon.png') });
	// ```.
	// Experimental.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

