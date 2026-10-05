package github

import (
	"github.com/projen/projen-go/projen"
)

// Options for 'AutoApprove'.
// Experimental.
type AutoApproveOptions struct {
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
	// Only pull requests authored by these Github usernames will be auto-approved.
	//
	// An empty list approves pull requests from any user.
	// Default: ['github-actions[bot]'].
	//
	// Deprecated: Use `sources` with `PullRequestSource.fromUsers({ logins })` instead.
	AllowedUsernames *[]*string `field:"optional" json:"allowedUsernames" yaml:"allowedUsernames"`
	// The credentials used to approve pull requests.
	//
	// Github forbids an identity to approve its own pull request.
	// These credentials must belong to a different identity than the one creating the pull requests.
	// Use `environment` on the credentials to restrict which branches can access them.
	// Default: - the workflow's `GITHUB_TOKEN`.
	//
	// Experimental.
	Credentials GithubCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Only pull requests with this label will be auto-approved.
	//
	// This is required in addition to matching one of the `sources`.
	// Default: 'auto-approve'.
	//
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// A GitHub secret name which contains a GitHub Access Token with write permissions for the `pull_request` scope.
	//
	// This token is used to approve pull requests.
	//
	// Github forbids an identity to approve its own pull request.
	// If your project produces automated pull requests using the Github default token -
	// {@link https://docs.github.com/en/actions/reference/authentication-in-a-workflow `GITHUB_TOKEN` }
	// - that you would like auto approved, such as when using the `depsUpgrade` property in
	// `NodeProjectOptions`, then you must use a different token here.
	// Default: "GITHUB_TOKEN".
	//
	// Deprecated: Use `credentials` with `GithubCredentials.fromPersonalAccessToken({ secret })` instead.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Only pull requests from one of these sources will be auto-approved.
	//
	// Components that create pull requests can add more sources with `AutoApprove.addSource()`.
	// Providing sources, either here or with `addSource()`, replaces the default source.
	// Default: [PullRequestSource.fromUsers({ logins: ["github-actions[bot]"] })]
	//
	// Experimental.
	Sources *[]PullRequestSource `field:"optional" json:"sources" yaml:"sources"`
}

