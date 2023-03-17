package workflows


// Options for push-like events.
// Experimental.
type PushOptions struct {
	// When using the push, pull_request and pull_request_target events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// When using the push, pull_request and pull_request_target events, you can configure a workflow to run when at least one file does not match paths-ignore or at least one modified file matches the configured paths.
	//
	// Path filters are not
	// evaluated for pushes to tags.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// When using the push, pull_request and pull_request_target events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

