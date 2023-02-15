package workflows


// The set of available triggers for GitHub Workflows.
// See: https://docs.github.com/en/actions/reference/events-that-trigger-workflows
//
// Experimental.
type Triggers struct {
	// Runs your workflow anytime the branch_protection_rule event occurs.
	// Experimental.
	BranchProtectionRule *BranchProtectionRuleOptions `field:"optional" json:"branchProtectionRule" yaml:"branchProtectionRule"`
	// Runs your workflow anytime the check_run event occurs.
	// Experimental.
	CheckRun *CheckRunOptions `field:"optional" json:"checkRun" yaml:"checkRun"`
	// Runs your workflow anytime the check_suite event occurs.
	// Experimental.
	CheckSuite *CheckSuiteOptions `field:"optional" json:"checkSuite" yaml:"checkSuite"`
	// Runs your workflow anytime someone creates a branch or tag, which triggers the create event.
	// Experimental.
	Create *CreateOptions `field:"optional" json:"create" yaml:"create"`
	// Runs your workflow anytime someone deletes a branch or tag, which triggers the delete event.
	// Experimental.
	Delete *DeleteOptions `field:"optional" json:"delete" yaml:"delete"`
	// Runs your workflow anytime someone creates a deployment, which triggers the deployment event.
	//
	// Deployments created with a commit SHA may not have
	// a Git ref.
	// Experimental.
	Deployment *DeploymentOptions `field:"optional" json:"deployment" yaml:"deployment"`
	// Runs your workflow anytime a third party provides a deployment status, which triggers the deployment_status event.
	//
	// Deployments created with a
	// commit SHA may not have a Git ref.
	// Experimental.
	DeploymentStatus *DeploymentStatusOptions `field:"optional" json:"deploymentStatus" yaml:"deploymentStatus"`
	// Runs your workflow anytime the discussion event occurs.
	//
	// More than one activity type triggers this event.
	// See: https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	//
	// Experimental.
	Discussion *DiscussionOptions `field:"optional" json:"discussion" yaml:"discussion"`
	// Runs your workflow anytime the discussion_comment event occurs.
	//
	// More than one activity type triggers this event.
	// See: https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	//
	// Experimental.
	DiscussionComment *DiscussionCommentOptions `field:"optional" json:"discussionComment" yaml:"discussionComment"`
	// Runs your workflow anytime when someone forks a repository, which triggers the fork event.
	// Experimental.
	Fork *ForkOptions `field:"optional" json:"fork" yaml:"fork"`
	// Runs your workflow when someone creates or updates a Wiki page, which triggers the gollum event.
	// Experimental.
	Gollum *GollumOptions `field:"optional" json:"gollum" yaml:"gollum"`
	// Runs your workflow anytime the issue_comment event occurs.
	// Experimental.
	IssueComment *IssueCommentOptions `field:"optional" json:"issueComment" yaml:"issueComment"`
	// Runs your workflow anytime the issues event occurs.
	// Experimental.
	Issues *IssuesOptions `field:"optional" json:"issues" yaml:"issues"`
	// Runs your workflow anytime the label event occurs.
	// Experimental.
	Label *LabelOptions `field:"optional" json:"label" yaml:"label"`
	// Runs your workflow when a pull request is added to a merge queue, which adds the pull request to a merge group.
	// Experimental.
	MergeGroup *MergeGroupOptions `field:"optional" json:"mergeGroup" yaml:"mergeGroup"`
	// Runs your workflow anytime the milestone event occurs.
	// Experimental.
	Milestone *MilestoneOptions `field:"optional" json:"milestone" yaml:"milestone"`
	// Runs your workflow anytime someone pushes to a GitHub Pages-enabled branch, which triggers the page_build event.
	// Experimental.
	PageBuild *PageBuildOptions `field:"optional" json:"pageBuild" yaml:"pageBuild"`
	// Runs your workflow anytime the project event occurs.
	// Experimental.
	Project *ProjectOptions `field:"optional" json:"project" yaml:"project"`
	// Runs your workflow anytime the project_card event occurs.
	// Experimental.
	ProjectCard *ProjectCardOptions `field:"optional" json:"projectCard" yaml:"projectCard"`
	// Runs your workflow anytime the project_column event occurs.
	// Experimental.
	ProjectColumn *ProjectColumnOptions `field:"optional" json:"projectColumn" yaml:"projectColumn"`
	// Runs your workflow anytime someone makes a private repository public, which triggers the public event.
	// Experimental.
	Public *PublicOptions `field:"optional" json:"public" yaml:"public"`
	// Runs your workflow anytime the pull_request event occurs.
	// Experimental.
	PullRequest *PullRequestOptions `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// Runs your workflow anytime the pull_request_review event occurs.
	// Experimental.
	PullRequestReview *PullRequestReviewOptions `field:"optional" json:"pullRequestReview" yaml:"pullRequestReview"`
	// Runs your workflow anytime a comment on a pull request's unified diff is modified, which triggers the pull_request_review_comment event.
	// Experimental.
	PullRequestReviewComment *PullRequestReviewCommentOptions `field:"optional" json:"pullRequestReviewComment" yaml:"pullRequestReviewComment"`
	// This event runs in the context of the base of the pull request, rather than in the merge commit as the pull_request event does.
	//
	// This prevents
	// executing unsafe workflow code from the head of the pull request that
	// could alter your repository or steal any secrets you use in your workflow.
	// This event allows you to do things like create workflows that label and
	// comment on pull requests based on the contents of the event payload.
	//
	// WARNING: The `pull_request_target` event is granted read/write repository
	// token and can access secrets, even when it is triggered from a fork.
	// Although the workflow runs in the context of the base of the pull request,
	// you should make sure that you do not check out, build, or run untrusted
	// code from the pull request with this event. Additionally, any caches
	// share the same scope as the base branch, and to help prevent cache
	// poisoning, you should not save the cache if there is a possibility that
	// the cache contents were altered.
	// See: https://securitylab.github.com/research/github-actions-preventing-pwn-requests
	//
	// Experimental.
	PullRequestTarget *PullRequestTargetOptions `field:"optional" json:"pullRequestTarget" yaml:"pullRequestTarget"`
	// Runs your workflow when someone pushes to a repository branch, which triggers the push event.
	// Experimental.
	Push *PushOptions `field:"optional" json:"push" yaml:"push"`
	// Runs your workflow anytime a package is published or updated.
	// Experimental.
	RegistryPackage *RegistryPackageOptions `field:"optional" json:"registryPackage" yaml:"registryPackage"`
	// Runs your workflow anytime the release event occurs.
	// Experimental.
	Release *ReleaseOptions `field:"optional" json:"release" yaml:"release"`
	// You can use the GitHub API to trigger a webhook event called repository_dispatch when you want to trigger a workflow for activity that happens outside of GitHub.
	// Experimental.
	RepositoryDispatch *RepositoryDispatchOptions `field:"optional" json:"repositoryDispatch" yaml:"repositoryDispatch"`
	// You can schedule a workflow to run at specific UTC times using POSIX cron syntax.
	//
	// Scheduled workflows run on the latest commit on the default or
	// base branch. The shortest interval you can run scheduled workflows is
	// once every 5 minutes.
	// See: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07
	//
	// Experimental.
	Schedule *[]*CronScheduleOptions `field:"optional" json:"schedule" yaml:"schedule"`
	// Runs your workflow anytime the status of a Git commit changes, which triggers the status event.
	// Experimental.
	Status *StatusOptions `field:"optional" json:"status" yaml:"status"`
	// Runs your workflow anytime the watch event occurs.
	// Experimental.
	Watch *WatchOptions `field:"optional" json:"watch" yaml:"watch"`
	// Can be called from another workflow.
	// See: https://docs.github.com/en/actions/learn-github-actions/reusing-workflows
	//
	// Experimental.
	WorkflowCall *WorkflowCallOptions `field:"optional" json:"workflowCall" yaml:"workflowCall"`
	// You can configure custom-defined input properties, default input values, and required inputs for the event directly in your workflow.
	//
	// When the
	// workflow runs, you can access the input values in the github.event.inputs
	// context.
	// Experimental.
	WorkflowDispatch *WorkflowDispatchOptions `field:"optional" json:"workflowDispatch" yaml:"workflowDispatch"`
	// This event occurs when a workflow run is requested or completed, and allows you to execute a workflow based on the finished result of another workflow.
	//
	// A workflow run is triggered regardless of the result of the
	// previous workflow.
	// Experimental.
	WorkflowRun *WorkflowRunOptions `field:"optional" json:"workflowRun" yaml:"workflowRun"`
}

