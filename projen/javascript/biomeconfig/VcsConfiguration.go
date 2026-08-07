package biomeconfig


// Settings for integrating Biome with version control.
// Experimental.
type VcsConfiguration struct {
	// The version control client.
	// Experimental.
	ClientKind VcsClientKind `field:"optional" json:"clientKind" yaml:"clientKind"`
	// The project's default branch.
	// Experimental.
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Whether Biome should integrate with the version control client.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Sets the directory where Biome checks for version control files.
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
	// When `true`, Biome ignores files listed in `.gitignore`, `.ignore`, and Git's local exclude file.
	// Experimental.
	UseIgnoreFile *bool `field:"optional" json:"useIgnoreFile" yaml:"useIgnoreFile"`
}

