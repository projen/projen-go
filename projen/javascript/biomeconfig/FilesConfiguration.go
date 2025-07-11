package biomeconfig


// The configuration of the filesystem.
// Experimental.
type FilesConfiguration struct {
	// Set of file and folder names that should be unconditionally ignored by Biome's scanner.
	//
	// Biome maintains an internal list of default ignore entries, which is based on user feedback and which may change in any release. This setting allows overriding this internal list completely.
	//
	// This is considered an advanced feature that users _should_ not need to tweak themselves, but they can as a last resort. This setting can only be configured in root configurations, and is ignored in nested configs.
	//
	// Entries must be file or folder *names*. Specific paths and globs are not supported.
	//
	// Examples where this may be useful:
	//
	// ```jsonc { "files": { "experimentalScannerIgnores": [ // You almost certainly don't want to scan your `.git` // folder, which is why it's already ignored by default: ".git",
	//
	// // But the scanner does scan `node_modules` by default. If // you *really* don't want this, you can ignore it like // this: "node_modules",
	//
	// // But it's probably better to ignore a specific dependency. // For instance, one that happens to be particularly slow to // scan: "RedisCommander.d.ts", ], } } ```
	//
	// Please be aware that rules relying on the module graph or type inference information may be negatively affected if dependencies of your project aren't (fully) scanned.
	// Experimental.
	ExperimentalScannerIgnores *[]*string `field:"optional" json:"experimentalScannerIgnores" yaml:"experimentalScannerIgnores"`
	// Tells Biome to not emit diagnostics when handling files that doesn't know.
	// Experimental.
	IgnoreUnknown *bool `field:"optional" json:"ignoreUnknown" yaml:"ignoreUnknown"`
	// A list of glob patterns.
	//
	// Biome will handle only those files/folders that will match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// The maximum allowed size for source code files in bytes.
	//
	// Files above this limit will be ignored for performance reasons. Defaults to 1 MiB
	// Default: 1 MiB.
	//
	// Experimental.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
}

