package cdk


// Experimental.
type Catalog struct {
	// Should we announce new versions?
	// Default: true.
	//
	// Experimental.
	Announce *bool `field:"optional" json:"announce" yaml:"announce"`
	// Twitter account to.
	// Experimental.
	Twitter *string `field:"optional" json:"twitter" yaml:"twitter"`
}

