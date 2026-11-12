package build


// Options for pushing files modified during builds back to pull request branches.
// Experimental.
type MutableBuildOptions struct {
	// Also push changes back to pull requests from forks.
	//
	// The changes are produced by the fork's build and pushed to the fork's
	// branch using the projen credentials. This requires the pull request to
	// allow edits from maintainers and credentials that can push to the fork.
	//
	// The fork author controls the content of these commits, which can change
	// any file. Review them like the rest of the pull request, even though they
	// are authored by the automation user. Their builds also run without the
	// approval GitHub requires for workflow runs from first-time contributors.
	// Default: false.
	//
	// Experimental.
	AllowFromForks *bool `field:"optional" json:"allowFromForks" yaml:"allowFromForks"`
}

