package javascript


// Options for `Bundler`.
// Experimental.
type BundlerOptions struct {
	// Install the `bundle` command as a pre-compile phase.
	// Default: true.
	//
	// Deprecated: Use `runBundleTask` instead.
	AddToPreCompile *bool `field:"optional" json:"addToPreCompile" yaml:"addToPreCompile"`
	// Output directory for all bundles.
	// Default: "assets".
	//
	// Experimental.
	AssetsDir *string `field:"optional" json:"assetsDir" yaml:"assetsDir"`
	// The semantic version requirement for `esbuild`.
	// Default: - no specific version (implies latest).
	//
	// Experimental.
	EsbuildVersion *string `field:"optional" json:"esbuildVersion" yaml:"esbuildVersion"`
	// Map of file extensions (without dot) and loaders to use for this file type.
	//
	// Loaders are appended to the esbuild command by `--loader:.extension=loader`
	// Experimental.
	Loaders *map[string]*string `field:"optional" json:"loaders" yaml:"loaders"`
	// Choose which phase (if any) to add the `bundle` command to.
	//
	// Note: If using `addBundle()` with the `bundleCompiledResults`, this option
	// must be set to `RunBundleTask.POST_COMPILE` or `RunBundleTask.MANUAL`.
	// See: AddBundleOptions.bundleCompiledResults *
	//
	// Default: RunBundleTask.PRE_COMPILE
	//
	// Experimental.
	RunBundleTask RunBundleTask `field:"optional" json:"runBundleTask" yaml:"runBundleTask"`
}

