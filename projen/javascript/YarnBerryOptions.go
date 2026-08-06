package javascript


// Configure Yarn Berry.
// Experimental.
type YarnBerryOptions struct {
	// Packages to deduplicate when the `dedupe` task runs.
	//
	// This will prevent multiple versions of the same package from being
	// installed in the lock file, if a single version could satisfy all requested
	// version ranges. This will prevent version proliferation and reduce the size
	// of the dependency tree.
	//
	// Setting this implies `dedupeDeps: true` on the project, which will run
	// `yarn dedupe` after every mutating install. Set `dedupeDeps: false`
	// explicitly to disable the task.
	//
	// Supports glob patterns, e.g. `@aws-sdk/*`.
	// Default: - all packages are deduplicated.
	//
	// Experimental.
	DedupePackages *[]*string `field:"optional" json:"dedupePackages" yaml:"dedupePackages"`
	// A fully specified version to use for yarn (e.g., x.x.x).
	// Default: - 4.13.0
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The yarnrc configuration.
	// Default: - a blank Yarn RC file.
	//
	// Experimental.
	YarnRcOptions *YarnrcOptions `field:"optional" json:"yarnRcOptions" yaml:"yarnRcOptions"`
	// Should zero-installs be enabled?
	//
	// Learn more at: https://yarnpkg.com/features/caching#zero-installs
	// Default: false.
	//
	// Experimental.
	ZeroInstalls *bool `field:"optional" json:"zeroInstalls" yaml:"zeroInstalls"`
}

