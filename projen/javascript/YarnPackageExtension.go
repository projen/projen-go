package javascript


// https://yarnpkg.com/configuration/yarnrc#packageExtensions.
// Experimental.
type YarnPackageExtension struct {
	// Experimental.
	Dependencies *map[string]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Experimental.
	PeerDependencies *map[string]*string `field:"optional" json:"peerDependencies" yaml:"peerDependencies"`
	// Experimental.
	PeerDependenciesMeta *map[string]*map[string]*YarnPeerDependencyMeta `field:"optional" json:"peerDependenciesMeta" yaml:"peerDependenciesMeta"`
}

