package github


// Represents the git identity.
// Experimental.
type GitIdentity struct {
	// The email address of the git user.
	// Experimental.
	Email *string `field:"required" json:"email" yaml:"email"`
	// The name of the user.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

