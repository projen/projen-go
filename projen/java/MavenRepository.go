package java


// Represents a Maven repository.
// See: https://maven.apache.org/guides/introduction/introduction-to-repositories.html
//
// Experimental.
type MavenRepository struct {
	// The identifier for the repository.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The url of the repository.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The layout of the repository.
	// Experimental.
	Layout *string `field:"optional" json:"layout" yaml:"layout"`
	// The name of the repository.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Repository Policy for Releases.
	// Experimental.
	Releases *MavenRepositoryPolicy `field:"optional" json:"releases" yaml:"releases"`
	// Repository Policy for Snapshots.
	// Experimental.
	Snapshots *MavenRepositoryPolicy `field:"optional" json:"snapshots" yaml:"snapshots"`
}

