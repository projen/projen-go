package github

import (
	"github.com/projen/projen-go/projen"
)

// Options for PullRequestLint.
// Experimental.
type PullRequestLintOptions struct {
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
	// Add a job that combines the results of all checks into a single status check.
	//
	// Use this job as the required status check, so it does not need updating when checks are added or removed.
	// Default: true.
	//
	// Experimental.
	ConsolidatedStatusCheck *bool `field:"optional" json:"consolidatedStatusCheck" yaml:"consolidatedStatusCheck"`
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
	// The pull request event that triggers the workflow.
	// Default: PullRequestLintTrigger.PULL_REQUEST_TARGET
	//
	// Experimental.
	Trigger PullRequestLintTrigger `field:"optional" json:"trigger" yaml:"trigger"`
}

