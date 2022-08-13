package github


// Options for PullRequestLint.
// Experimental.
type PullRequestLintOptions struct {
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Validate that pull request titles follow Conventional Commits.
	// See: https://www.conventionalcommits.org/
	//
	// Experimental.
	SemanticTitle *bool `field:"optional" json:"semanticTitle" yaml:"semanticTitle"`
	// Options for validating the conventional commit title linter.
	// Experimental.
	SemanticTitleOptions *SemanticTitleOptions `field:"optional" json:"semanticTitleOptions" yaml:"semanticTitleOptions"`
}

