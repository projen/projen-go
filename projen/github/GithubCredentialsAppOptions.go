package github

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `GithubCredentials.fromApp`.
// Experimental.
type GithubCredentialsAppOptions struct {
	// The secret containing the GitHub App ID.
	// Default: "PROJEN_APP_ID".
	//
	// Experimental.
	AppIdSecret *string `field:"optional" json:"appIdSecret" yaml:"appIdSecret"`
	// The GitHub Actions environment the secrets are added to.
	//
	// This can be used to add explicit approval steps to access the secrets.
	// Default: - no environment used.
	//
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The owner of the GitHub App installation.
	// Default: - if empty, defaults to the current repository owner.
	//
	// Experimental.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The permissions granted to the token.
	// Default: - all permissions granted to the app.
	//
	// Experimental.
	Permissions *workflows.AppPermissions `field:"optional" json:"permissions" yaml:"permissions"`
	// The secret containing the GitHub App private key.
	//
	// Escaped newlines (\\n) will be automatically replaced with actual newlines.
	// Default: "PROJEN_APP_PRIVATE_KEY".
	//
	// Experimental.
	PrivateKeySecret *string `field:"optional" json:"privateKeySecret" yaml:"privateKeySecret"`
	// List of repositories to grant access to.
	// Default: - if owner is set and repositories is empty, access will be scoped to all repositories in the provided repository owner's installation.
	// If owner and repositories are empty, access will be scoped to only the current repository.
	//
	// Experimental.
	Repositories *[]*string `field:"optional" json:"repositories" yaml:"repositories"`
}

