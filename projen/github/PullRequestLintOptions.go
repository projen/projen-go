package github


// Options for PullRequestLint.
// Experimental.
type PullRequestLintOptions struct {
	// Require a contributor statement to be included in the PR description.
	//
	// For example confirming that the contribution has been made by the contributor and complies with the project's license.
	//
	// Appends the statement to the end of the Pull Request template.
	// Default: - no contributor statement is required.
	//
	// Experimental.
	ContributorStatement *string `field:"optional" json:"contributorStatement" yaml:"contributorStatement"`
	// Options for requiring a contributor statement on Pull Requests.
	// Default: - none.
	//
	// Experimental.
	ContributorStatementOptions *ContributorStatementOptions `field:"optional" json:"contributorStatementOptions" yaml:"contributorStatementOptions"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Validate that pull request titles follow Conventional Commits.
	// See: https://www.conventionalcommits.org/
	//
	// Default: true.
	//
	// Experimental.
	SemanticTitle *bool `field:"optional" json:"semanticTitle" yaml:"semanticTitle"`
	// Options for validating the conventional commit title linter.
	// Default: - title must start with "feat", "fix", or "chore".
	//
	// Experimental.
	SemanticTitleOptions *SemanticTitleOptions `field:"optional" json:"semanticTitleOptions" yaml:"semanticTitleOptions"`
}

