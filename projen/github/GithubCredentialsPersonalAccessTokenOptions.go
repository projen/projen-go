package github


// Options for `GithubCredentials.fromPersonalAccessToken`.
// Experimental.
type GithubCredentialsPersonalAccessTokenOptions struct {
	// The GitHub Actions environment the secrets is added to.
	//
	// This can be used to add explicit approval steps to access the secret.
	// Default: - no environment used.
	//
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The name of the secret that holds the GitHub personal access token.
	// Default: "PROJEN_GITHUB_TOKEN".
	//
	// Experimental.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

