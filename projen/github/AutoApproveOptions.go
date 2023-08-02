package github


// Options for 'AutoApprove'.
// Experimental.
type AutoApproveOptions struct {
	// Only pull requests authored by these Github usernames will be auto-approved.
	// Default: ['github-bot'].
	//
	// Experimental.
	AllowedUsernames *[]*string `field:"optional" json:"allowedUsernames" yaml:"allowedUsernames"`
	// Only pull requests with this label will be auto-approved.
	// Default: 'auto-approve'.
	//
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
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
	// Experimental.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

