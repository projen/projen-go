package biomeconfig


// Experimental.
type AssistConfiguration struct {
	// Whether Biome should fail in CLI if the assist were not applied to the code.
	// Experimental.
	Actions *Actions `field:"optional" json:"actions" yaml:"actions"`
	// Whether Biome should enable assist via LSP and CLI.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of glob patterns.
	//
	// Biome will include files/folders that will
	// match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

