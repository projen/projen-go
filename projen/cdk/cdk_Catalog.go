package cdk


// Experimental.
type Catalog struct {
	// Should we announce new versions?
	// Experimental.
	Announce *bool `field:"optional" json:"announce" yaml:"announce"`
	// Twitter account to @mention in announcement tweet.
	// Experimental.
	Twitter *string `field:"optional" json:"twitter" yaml:"twitter"`
}

