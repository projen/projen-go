package github


// Experimental.
type CreatePullRequestOptions struct {
	// Description added to the pull request.
	//
	// Providence information are automatically added.
	// Experimental.
	PullRequestDescription *string `field:"required" json:"pullRequestDescription" yaml:"pullRequestDescription"`
	// The full title used to create the pull request.
	//
	// If PR titles are validated in this repo, the title should comply with the respective rules.
	// Experimental.
	PullRequestTitle *string `field:"required" json:"pullRequestTitle" yaml:"pullRequestTitle"`
	// The name of the workflow that will create the PR.
	// Experimental.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// Assignees to add on the PR.
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// Sets the pull request base branch.
	// Experimental.
	BaseBranch *string `field:"optional" json:"baseBranch" yaml:"baseBranch"`
	// The pull request branch name.
	// Experimental.
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// The job credentials used to create the pull request.
	//
	// Provided credentials must have permissions to create a pull request on the repository.
	// Experimental.
	Credentials GithubCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// The git identity used to create the commit.
	// Experimental.
	GitIdentity *GitIdentity `field:"optional" json:"gitIdentity" yaml:"gitIdentity"`
	// Labels to apply on the PR.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Add Signed-off-by line by the committer at the end of the commit log message.
	// Experimental.
	Signoff *bool `field:"optional" json:"signoff" yaml:"signoff"`
	// The step ID which produces the output which indicates if a patch was created.
	// Experimental.
	StepId *string `field:"optional" json:"stepId" yaml:"stepId"`
	// The name of the step displayed on GitHub.
	// Experimental.
	StepName *string `field:"optional" json:"stepName" yaml:"stepName"`
}

