package biomeconfig


// Configures how Biome integrates with a version control system.
// Experimental.
type VcsConfiguration struct {
	// Selects the version-control client.
	//
	// Currently, only `git` is supported.
	// Experimental.
	ClientKind VcsClientKind `field:"optional" json:"clientKind" yaml:"clientKind"`
	// Sets the base branch used by `--changed` when `--since` is not provided.
	//
	// If neither this option
	// nor `--since` is set, commands using `--changed` fail because Biome cannot determine the
	// comparison base.
	// Experimental.
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Enables or disables version-control integration.
	//
	// Defaults to `false`. Enabling integration
	// requires `vcs.clientKind` to be set.
	// Default: false`. Enabling integration
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Sets the directory where Biome looks for version-control files.
	//
	// A relative value is resolved
	// from the directory containing the current configuration file, while an absolute path is used
	// directly.
	//
	// Defaults to the directory containing `biome.json` or `biome.jsonc`. If no configuration is
	// found, Biome uses the current working directory.
	//
	// If neither directory is available, Biome disables version control integration and emits a
	// diagnostic.
	// Default: the directory containing `biome.json` or `biome.jsonc`. If no configuration is
	//
	// Experimental.
	Root *string `field:"optional" json:"root" yaml:"root"`
	// Controls whether Biome applies patterns from `.gitignore`, Git's local `.git/info/exclude`, and supported `.ignore` files, including nested ignore files. Patterns in the root ignore file are resolved from `vcs.root`. Patterns in nested ignore files are resolved from the directory containing that file. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	UseIgnoreFile *bool `field:"optional" json:"useIgnoreFile" yaml:"useIgnoreFile"`
}

