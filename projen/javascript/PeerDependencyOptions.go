package javascript


// Experimental.
type PeerDependencyOptions struct {
	// Automatically add a pinned dev dependency.
	// Default: true.
	//
	// Experimental.
	PinnedDevDependency *bool `field:"optional" json:"pinnedDevDependency" yaml:"pinnedDevDependency"`
}

