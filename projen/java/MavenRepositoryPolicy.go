package java


// Represents a Maven Repository Policy.
// See: https://maven.apache.org/settings.html#repositories
//
// Experimental.
type MavenRepositoryPolicy struct {
	// Checksum Policy When Maven deploys files to the repository, it also deploys corresponding checksum files.
	// Experimental.
	ChecksumPolicy ChecksumPolicy `field:"optional" json:"checksumPolicy" yaml:"checksumPolicy"`
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Update Policy This element specifies how often updates should attempt to occur.
	//
	// Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote.
	// Default: UpdatePolicy.DAILY
	//
	// Experimental.
	UpdatePolicy UpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
}

