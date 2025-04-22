package workflows


// The available scopes and access values for workflow permissions.
//
// If you
// specify the access for any of these scopes, all those that are not
// specified are set to `JobPermission.NONE`, instead of the default behavior
// when none is specified.
// See: https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/controlling-permissions-for-github_token#defining-access-for-the-github_token-permissions
//
// Experimental.
type JobPermissions struct {
	// Experimental.
	Actions JobPermission `field:"optional" json:"actions" yaml:"actions"`
	// Experimental.
	Attestations JobPermission `field:"optional" json:"attestations" yaml:"attestations"`
	// Experimental.
	Checks JobPermission `field:"optional" json:"checks" yaml:"checks"`
	// Experimental.
	Contents JobPermission `field:"optional" json:"contents" yaml:"contents"`
	// Experimental.
	Deployments JobPermission `field:"optional" json:"deployments" yaml:"deployments"`
	// Experimental.
	Discussions JobPermission `field:"optional" json:"discussions" yaml:"discussions"`
	// Experimental.
	IdToken JobPermission `field:"optional" json:"idToken" yaml:"idToken"`
	// Experimental.
	Issues JobPermission `field:"optional" json:"issues" yaml:"issues"`
	// Experimental.
	Models JobPermission `field:"optional" json:"models" yaml:"models"`
	// Experimental.
	Packages JobPermission `field:"optional" json:"packages" yaml:"packages"`
	// Experimental.
	Pages JobPermission `field:"optional" json:"pages" yaml:"pages"`
	// Experimental.
	PullRequests JobPermission `field:"optional" json:"pullRequests" yaml:"pullRequests"`
	// Deprecated: removed by GitHub.
	RepositoryProjects JobPermission `field:"optional" json:"repositoryProjects" yaml:"repositoryProjects"`
	// Experimental.
	SecurityEvents JobPermission `field:"optional" json:"securityEvents" yaml:"securityEvents"`
	// Experimental.
	Statuses JobPermission `field:"optional" json:"statuses" yaml:"statuses"`
}

