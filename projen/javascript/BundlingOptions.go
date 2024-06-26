package javascript


// Options for bundling.
// Experimental.
type BundlingOptions struct {
	// You can mark a file or a package as external to exclude it from your build.
	//
	// Instead of being bundled, the import will be preserved (using require for
	// the iife and cjs formats and using import for the esm format) and will be
	// evaluated at run time instead.
	//
	// This has several uses. First of all, it can be used to trim unnecessary
	// code from your bundle for a code path that you know will never be executed.
	// For example, a package may contain code that only runs in node but you will
	// only be using that package in the browser. It can also be used to import
	// code in node at run time from a package that cannot be bundled. For
	// example, the fsevents package contains a native extension, which esbuild
	// doesn't support.
	// Default: [].
	//
	// Experimental.
	Externals *[]*string `field:"optional" json:"externals" yaml:"externals"`
	// Include a source map in the bundle.
	// Default: false.
	//
	// Experimental.
	Sourcemap *bool `field:"optional" json:"sourcemap" yaml:"sourcemap"`
	// In addition to the `bundle:xyz` task, creates `bundle:xyz:watch` task which will invoke the same esbuild command with the `--watch` flag.
	//
	// This can be used
	// to continusouly watch for changes.
	// Default: true.
	//
	// Experimental.
	WatchTask *bool `field:"optional" json:"watchTask" yaml:"watchTask"`
}

