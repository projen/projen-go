package biomeconfig


// The configuration of the filesystem.
// Experimental.
type FilesConfiguration struct {
	// **Deprecated:** Please use _force-ignore syntax_ in `files.includes` instead: <https://biomejs.dev/reference/configuration/#filesincludes>.
	//
	// Set of file and folder names that should be unconditionally ignored by
	// Biome's scanner.
	// Experimental.
	ExperimentalScannerIgnores *[]*string `field:"optional" json:"experimentalScannerIgnores" yaml:"experimentalScannerIgnores"`
	// Tells Biome to not emit diagnostics when handling files that it doesn't know.
	// Experimental.
	IgnoreUnknown *bool `field:"optional" json:"ignoreUnknown" yaml:"ignoreUnknown"`
	// A list of glob patterns.
	//
	// Biome will handle only those files/folders that will
	// match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// The maximum allowed size for source code files in bytes.
	//
	// Files above
	// this limit will be ignored for performance reasons. Defaults to 1 MiB
	// Default: 1 MiB.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
}

