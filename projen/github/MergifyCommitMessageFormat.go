package github


// Declarative configuration for Mergify `commit_message_format`.
// See: https://docs.mergify.com/workflow/actions/merge/#customizing-the-commit-message
//
// Experimental.
type MergifyCommitMessageFormat struct {
	// Commit body format.
	//
	// - `inherit`: use the GitHub repository default merge commit body format
	// - `pr-body`: use the pull request body
	// - `empty`: set the commit body to be empty.
	// Experimental.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Commit title format.
	//
	// - `inherit`: use the GitHub repository default merge commit title format
	// - `pr-title`: use the pull request title (with the PR number appended).
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Optional list of trailers to append to the commit message.
	// Experimental.
	Trailers *[]*string `field:"optional" json:"trailers" yaml:"trailers"`
}

