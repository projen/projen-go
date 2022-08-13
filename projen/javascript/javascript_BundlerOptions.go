package javascript


// Options for `Bundler`.
// Experimental.
type BundlerOptions struct {
	// Install the `bundle` command as a pre-compile phase.
	// Experimental.
	AddToPreCompile *bool `field:"optional" json:"addToPreCompile" yaml:"addToPreCompile"`
	// Output directory for all bundles.
	// Experimental.
	AssetsDir *string `field:"optional" json:"assetsDir" yaml:"assetsDir"`
	// The semantic version requirement for `esbuild`.
	// Experimental.
	EsbuildVersion *string `field:"optional" json:"esbuildVersion" yaml:"esbuildVersion"`
}

