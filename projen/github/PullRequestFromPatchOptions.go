package github


// Experimental.
type PullRequestFromPatchOptions struct {
	// The job credentials used to create the pull request.
	//
	// Provided credentials must have permissions to create a pull request on the repository.
	// Experimental.
	Credentials GithubCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// The git identity used to create the commit.
	// Experimental.
	GitIdentity *GitIdentity `field:"required" json:"gitIdentity" yaml:"gitIdentity"`
	// Information about the patch that is used to create the pull request.
	// Experimental.
	Patch *PullRequestPatchSource `field:"required" json:"patch" yaml:"patch"`
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
	// Assignees to add on the PR.
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// The name of the job displayed on GitHub.
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Labels to apply on the PR.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Branch or tag name.
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Add Signed-off-by line by the committer at the end of the commit log message.
	// Experimental.
	Signoff *bool `field:"optional" json:"signoff" yaml:"signoff"`
}

