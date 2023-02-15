package workflows


// Credentials to use to authenticate to Docker registries.
// Experimental.
type ContainerCredentials struct {
	// The password.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

