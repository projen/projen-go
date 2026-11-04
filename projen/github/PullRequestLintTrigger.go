package github


// The pull request event that triggers the pull request lint workflow.
// Experimental.
type PullRequestLintTrigger string

const (
	// Run in the context of the base branch, with access to secrets and a read/write token.
	// Experimental.
	PullRequestLintTrigger_PULL_REQUEST_TARGET PullRequestLintTrigger = "PULL_REQUEST_TARGET"
	// Run in the context of the merge commit, with a read-only token for pull requests from forks.
	// Experimental.
	PullRequestLintTrigger_PULL_REQUEST PullRequestLintTrigger = "PULL_REQUEST"
)

