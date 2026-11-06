package javascript


// Additional information about how a peer dependency is used.
// See: https://docs.npmjs.com/cli/v11/configuring-npm/package-json#peerdependenciesmeta
//
// Experimental.
type PeerDependencyMeta struct {
	// Marks the peer dependency as optional.
	//
	// Package managers will not
	// install it or warn if it is missing.
	// Default: false.
	//
	// Experimental.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

