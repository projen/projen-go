package github


// Options for `GithubCredentials.fromApp`.
// Experimental.
type GithubCredentialsAppOptions struct {
	// Experimental.
	AppIdSecret *string `field:"optional" json:"appIdSecret" yaml:"appIdSecret"`
	// Experimental.
	PrivateKeySecret *string `field:"optional" json:"privateKeySecret" yaml:"privateKeySecret"`
}

