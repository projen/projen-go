package java


// Experimental.
type ParentPom struct {
	// Parent Pom Artifact ID.
	// Experimental.
	ArtifactId *string `field:"optional" json:"artifactId" yaml:"artifactId"`
	// Parent Pom Group ID.
	// Experimental.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Parent Pom Relative path from the current pom.
	// Experimental.
	RelativePath *string `field:"optional" json:"relativePath" yaml:"relativePath"`
	// Parent Pom Version.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

