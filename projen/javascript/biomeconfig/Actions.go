package biomeconfig


// Experimental.
type Actions struct {
	// It enables the assist actions recommended by Biome.
	//
	// `true` by default.
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Experimental.
	Source *Source `field:"optional" json:"source" yaml:"source"`
}

