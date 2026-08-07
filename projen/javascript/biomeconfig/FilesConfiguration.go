package biomeconfig


// The file handling configuration.
// Experimental.
type FilesConfiguration struct {
	// **Deprecated:** Please use _force-ignore syntax_ in `files.includes` instead: <https://biomejs.dev/reference/configuration/#filesincludes>.
	//
	// Set of file and folder names that should be unconditionally ignored by
	// Biome's scanner.
	// Experimental.
	ExperimentalScannerIgnores *[]*string `field:"optional" json:"experimentalScannerIgnores" yaml:"experimentalScannerIgnores"`
	// Prevents Biome from emitting diagnostics for unrecognized file types.
	// Experimental.
	IgnoreUnknown *bool `field:"optional" json:"ignoreUnknown" yaml:"ignoreUnknown"`
	// A list of glob patterns.
	//
	// Biome handles only files and directories that match these
	// patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// The maximum source file size in bytes.
	//
	// Biome ignores larger files. Defaults to `1 MiB`.
	// Default: 1 MiB`.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
}

