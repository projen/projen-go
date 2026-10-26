package biomeconfig


// The file handling configuration.
// Experimental.
type FilesConfiguration struct {
	// **Deprecated:** Please use _force-ignore syntax_ in `files.includes` instead: <https://biomejs.dev/reference/configuration/#filesincludes>.
	//
	// File or directory names that the scanner ignores unconditionally while crawling. Each value
	// matches a complete path component, not a glob. Ignored files are not added to the module graph,
	// so types are not inferred from them.
	// Experimental.
	ExperimentalScannerIgnores *[]*string `field:"optional" json:"experimentalScannerIgnores" yaml:"experimentalScannerIgnores"`
	// Prevents Biome from emitting diagnostics for unrecognized file types.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	IgnoreUnknown *bool `field:"optional" json:"ignoreUnknown" yaml:"ignoreUnknown"`
	// A list of glob patterns selecting which files and directories Biome can process.
	//
	// If omitted,
	// this option does not further restrict the supported files Biome can process. An empty list
	// selects no files. Biome ignores files under `node_modules` even when they match a pattern.
	//
	// The scanner uses these patterns and, when enabled, VCS ignore files while discovering nested
	// configuration and ignore files. When enabled rules use the `project` or `types` domains, it
	// also indexes source files and project dependencies. It may still discover required metadata or
	// index excluded source files needed for analysis. A pattern beginning with `!!` prevents the
	// scanner from indexing matching paths. Patterns are evaluated in order, and a pattern beginning
	// with `!` excludes matches.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// The maximum source-file size in bytes.
	//
	// Biome skips larger files and emits a warning
	// diagnostic. Defaults to `1 MiB`.
	// Default: 1 MiB`.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
}

