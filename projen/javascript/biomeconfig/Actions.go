package biomeconfig


// Experimental.
type Actions struct {
	// The actions preset to use.
	// Experimental.
	Preset PresetConfig `field:"optional" json:"preset" yaml:"preset"`
	// It enables the assist actions recommended by Biome.
	//
	// `true` by default.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Experimental.
	Source *Source `field:"optional" json:"source" yaml:"source"`
}

