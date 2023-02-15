package github

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `GithubCredentials.fromApp`.
// Experimental.
type GithubCredentialsAppOptions struct {
	// Experimental.
	AppIdSecret *string `field:"optional" json:"appIdSecret" yaml:"appIdSecret"`
	// The permissions granted to the token.
	// Experimental.
	Permissions *workflows.AppPermissions `field:"optional" json:"permissions" yaml:"permissions"`
	// Experimental.
	PrivateKeySecret *string `field:"optional" json:"privateKeySecret" yaml:"privateKeySecret"`
}

