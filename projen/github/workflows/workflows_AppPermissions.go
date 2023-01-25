package workflows


// The permissions available to a GitHub App.
//
// Typically a token for a GitHub App has all the available scopes/permissions available to the app
// itself; however, a more limited set of permissions can be specified. When permissions are provided,
// **only** the specified permissions are granted to the token.
// See: https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28#create-an-installation-access-token-for-an-app
//
// Experimental.
type AppPermissions struct {
	// Experimental.
	Actions AppPermission `field:"optional" json:"actions" yaml:"actions"`
	// Experimental.
	Administration AppPermission `field:"optional" json:"administration" yaml:"administration"`
	// Experimental.
	Checks AppPermission `field:"optional" json:"checks" yaml:"checks"`
	// Experimental.
	Contents AppPermission `field:"optional" json:"contents" yaml:"contents"`
	// Experimental.
	Deployments AppPermission `field:"optional" json:"deployments" yaml:"deployments"`
	// Experimental.
	Environments AppPermission `field:"optional" json:"environments" yaml:"environments"`
	// Experimental.
	Issues AppPermission `field:"optional" json:"issues" yaml:"issues"`
	// Experimental.
	Members AppPermission `field:"optional" json:"members" yaml:"members"`
	// Experimental.
	Metadata AppPermission `field:"optional" json:"metadata" yaml:"metadata"`
	// Experimental.
	OrganizationAdministration AppPermission `field:"optional" json:"organizationAdministration" yaml:"organizationAdministration"`
	// Experimental.
	OrganizationAnnouncementBanners AppPermission `field:"optional" json:"organizationAnnouncementBanners" yaml:"organizationAnnouncementBanners"`
	// Experimental.
	OrganizationCustomRoles AppPermission `field:"optional" json:"organizationCustomRoles" yaml:"organizationCustomRoles"`
	// Experimental.
	OrganizationHooks AppPermission `field:"optional" json:"organizationHooks" yaml:"organizationHooks"`
	// Experimental.
	OrganizationPackages AppPermission `field:"optional" json:"organizationPackages" yaml:"organizationPackages"`
	// Experimental.
	OrganizationPlan AppPermission `field:"optional" json:"organizationPlan" yaml:"organizationPlan"`
	// Experimental.
	OrganizationProjects AppPermission `field:"optional" json:"organizationProjects" yaml:"organizationProjects"`
	// Experimental.
	OrganizationSecrets AppPermission `field:"optional" json:"organizationSecrets" yaml:"organizationSecrets"`
	// Experimental.
	OrganizationSelfHostedRunners AppPermission `field:"optional" json:"organizationSelfHostedRunners" yaml:"organizationSelfHostedRunners"`
	// Experimental.
	OrgnaizationUserBlocking AppPermission `field:"optional" json:"orgnaizationUserBlocking" yaml:"orgnaizationUserBlocking"`
	// Experimental.
	Packages AppPermission `field:"optional" json:"packages" yaml:"packages"`
	// Experimental.
	Pages AppPermission `field:"optional" json:"pages" yaml:"pages"`
	// Experimental.
	PullRequests AppPermission `field:"optional" json:"pullRequests" yaml:"pullRequests"`
	// Experimental.
	RepositoryAnnouncementBanners AppPermission `field:"optional" json:"repositoryAnnouncementBanners" yaml:"repositoryAnnouncementBanners"`
	// Experimental.
	RepositoryHooks AppPermission `field:"optional" json:"repositoryHooks" yaml:"repositoryHooks"`
	// Experimental.
	RepositoryProject AppPermission `field:"optional" json:"repositoryProject" yaml:"repositoryProject"`
	// Experimental.
	Secrets AppPermission `field:"optional" json:"secrets" yaml:"secrets"`
	// Experimental.
	SecretScanningAlerts AppPermission `field:"optional" json:"secretScanningAlerts" yaml:"secretScanningAlerts"`
	// Experimental.
	SecurityEvents AppPermission `field:"optional" json:"securityEvents" yaml:"securityEvents"`
	// Experimental.
	SingleFile AppPermission `field:"optional" json:"singleFile" yaml:"singleFile"`
	// Experimental.
	Statuses AppPermission `field:"optional" json:"statuses" yaml:"statuses"`
	// Experimental.
	TeamDiscussions AppPermission `field:"optional" json:"teamDiscussions" yaml:"teamDiscussions"`
	// Experimental.
	VulnerabilityAlerts AppPermission `field:"optional" json:"vulnerabilityAlerts" yaml:"vulnerabilityAlerts"`
	// Experimental.
	Workflows AppPermission `field:"optional" json:"workflows" yaml:"workflows"`
}

