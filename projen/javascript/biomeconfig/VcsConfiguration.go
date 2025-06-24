package biomeconfig


// Set of properties to integrate Biome with a VCS software.
// Experimental.
type VcsConfiguration struct {
	// The kind of client.
	// Experimental.
	ClientKind VcsClientKind `field:"optional" json:"clientKind" yaml:"clientKind"`
	// The main branch of the project.
	// Experimental.
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Whether Biome should integrate itself with the VCS client.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The folder where Biome should check for VCS files.
	//
	// By default, Biome will use the same folder where `biome.json` was found.
	//
	// If Biome can't find the configuration, it will attempt to use the current working directory. If no current working directory can't be found, Biome won't use the VCS integration, and a diagnostic will be emitted
	// Experimental.
	Root *string `field:"optional" json:"root" yaml:"root"`
	// Whether Biome should use the VCS ignore file.
	//
	// When [true], Biome will ignore the files specified in the ignore file.
	// Experimental.
	UseIgnoreFile *bool `field:"optional" json:"useIgnoreFile" yaml:"useIgnoreFile"`
}

