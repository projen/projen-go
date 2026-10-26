package biomeconfig


// Options that change how the Yaml linter behaves.
// Experimental.
type YamlLinterConfiguration struct {
	// Control the linter for Yaml files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

