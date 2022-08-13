// CDK for software projects
package projen


// SampleDir options.
// Experimental.
type SampleDirOptions struct {
	// The files to render into the directory.
	//
	// These files get added after
	// any files from `source` if that option is specified (replacing if names
	// overlap).
	// Experimental.
	Files *map[string]*string `field:"optional" json:"files" yaml:"files"`
	// Absolute path to a directory to copy files from (does not need to be text files).
	//
	// If your project is typescript-based and has configured `testdir` to be a
	// subdirectory of `src`, sample files should outside of the `src` directory
	// otherwise they may not be copied. For example:
	// ```
	// new SampleDir(this, 'public', { source: path.join(__dirname, '..', 'sample-assets') });
	// ```.
	// Experimental.
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
}

