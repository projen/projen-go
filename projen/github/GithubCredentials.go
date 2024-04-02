package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen/github/workflows"
)

// Represents a method of providing GitHub API access for projen workflows.
// Experimental.
type GithubCredentials interface {
	// Setup steps to obtain GitHub credentials.
	// Experimental.
	SetupSteps() *[]*workflows.JobStep
	// The value to use in a workflow when a GitHub token is expected.
	//
	// This
	// typically looks like "${{ some.path.to.a.value }}".
	// Experimental.
	TokenRef() *string
}

// The jsii proxy struct for GithubCredentials
type jsiiProxy_GithubCredentials struct {
	_ byte // padding
}

func (j *jsiiProxy_GithubCredentials) SetupSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"setupSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubCredentials) TokenRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenRef",
		&returns,
	)
	return returns
}


// Provide API access through a GitHub App.
//
// The GitHub App must be installed on the GitHub repo, its App ID and a
// private key must be added as secrets to the repo. The name of the secrets
// can be specified here.
// See: https://projen.io/docs/integrations/github/#github-app
//
// Default: - app id stored in "PROJEN_APP_ID" and private key stored in "PROJEN_APP_PRIVATE_KEY" with all permissions attached to the app.
//
// Experimental.
func GithubCredentials_FromApp(options *GithubCredentialsAppOptions) GithubCredentials {
	_init_.Initialize()

	if err := validateGithubCredentials_FromAppParameters(options); err != nil {
		panic(err)
	}
	var returns GithubCredentials

	_jsii_.StaticInvoke(
		"projen.github.GithubCredentials",
		"fromApp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Provide API access through a GitHub personal access token.
//
// The token must be added as a secret to the GitHub repo, and the name of the
// secret can be specified here.
// See: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
//
// Default: - a secret named "PROJEN_GITHUB_TOKEN".
//
// Experimental.
func GithubCredentials_FromPersonalAccessToken(options *GithubCredentialsPersonalAccessTokenOptions) GithubCredentials {
	_init_.Initialize()

	if err := validateGithubCredentials_FromPersonalAccessTokenParameters(options); err != nil {
		panic(err)
	}
	var returns GithubCredentials

	_jsii_.StaticInvoke(
		"projen.github.GithubCredentials",
		"fromPersonalAccessToken",
		[]interface{}{options},
		&returns,
	)

	return returns
}

