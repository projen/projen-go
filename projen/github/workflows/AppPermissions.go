package workflows


// The permissions available to a GitHub App.
//
// Typically a token for a GitHub App has all the available scopes/permissions available to the app
// itself; however, a more limited set of permissions can be specified. When permissions are provided,
// **only** the specified permissions are granted to the token.
// See: https://github.com/actions/create-github-app-token/blob/main/action.yml#L28
//
// Experimental.
type AppPermissions struct {
	// Experimental.
	Actions AppPermission `field:"optional" json:"actions" yaml:"actions"`
	// Experimental.
	Administration AppPermission `field:"optional" json:"administration" yaml:"administration"`
	// Experimental.
	Attestations AppPermission `field:"optional" json:"attestations" yaml:"attestations"`
	// Experimental.
	Checks AppPermission `field:"optional" json:"checks" yaml:"checks"`
	// Experimental.
	Codespaces AppPermission `field:"optional" json:"codespaces" yaml:"codespaces"`
	// Experimental.
	Contents AppPermission `field:"optional" json:"contents" yaml:"contents"`
	// Experimental.
	DependabotSecrets AppPermission `field:"optional" json:"dependabotSecrets" yaml:"dependabotSecrets"`
	// Experimental.
	Deployments AppPermission `field:"optional" json:"deployments" yaml:"deployments"`
	// Experimental.
	EmailAddresses AppPermission `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Experimental.
	Environments AppPermission `field:"optional" json:"environments" yaml:"environments"`
	// Experimental.
	Followers AppPermission `field:"optional" json:"followers" yaml:"followers"`
	// Experimental.
	GitSshKeys AppPermission `field:"optional" json:"gitSshKeys" yaml:"gitSshKeys"`
	// Experimental.
	GpgKeys AppPermission `field:"optional" json:"gpgKeys" yaml:"gpgKeys"`
	// Experimental.
	InteractionLimits AppPermission `field:"optional" json:"interactionLimits" yaml:"interactionLimits"`
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
	OrganizationCopilotSeatManagement AppPermission `field:"optional" json:"organizationCopilotSeatManagement" yaml:"organizationCopilotSeatManagement"`
	// Experimental.
	OrganizationCustomOrgRoles AppPermission `field:"optional" json:"organizationCustomOrgRoles" yaml:"organizationCustomOrgRoles"`
	// Experimental.
	OrganizationCustomProperties AppPermission `field:"optional" json:"organizationCustomProperties" yaml:"organizationCustomProperties"`
	// Experimental.
	OrganizationCustomRoles AppPermission `field:"optional" json:"organizationCustomRoles" yaml:"organizationCustomRoles"`
	// Experimental.
	OrganizationEvents AppPermission `field:"optional" json:"organizationEvents" yaml:"organizationEvents"`
	// Experimental.
	OrganizationHooks AppPermission `field:"optional" json:"organizationHooks" yaml:"organizationHooks"`
	// Experimental.
	OrganizationPackages AppPermission `field:"optional" json:"organizationPackages" yaml:"organizationPackages"`
	// Experimental.
	OrganizationPersonalAccessTokenRequests AppPermission `field:"optional" json:"organizationPersonalAccessTokenRequests" yaml:"organizationPersonalAccessTokenRequests"`
	// Experimental.
	OrganizationPersonalAccessTokens AppPermission `field:"optional" json:"organizationPersonalAccessTokens" yaml:"organizationPersonalAccessTokens"`
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
	Profile AppPermission `field:"optional" json:"profile" yaml:"profile"`
	// Experimental.
	PullRequests AppPermission `field:"optional" json:"pullRequests" yaml:"pullRequests"`
	// Deprecated: removed by GitHub.
	RepositoryAnnouncementBanners AppPermission `field:"optional" json:"repositoryAnnouncementBanners" yaml:"repositoryAnnouncementBanners"`
	// Experimental.
	RepositoryCustomProperties AppPermission `field:"optional" json:"repositoryCustomProperties" yaml:"repositoryCustomProperties"`
	// Experimental.
	RepositoryHooks AppPermission `field:"optional" json:"repositoryHooks" yaml:"repositoryHooks"`
	// Experimental.
	RepositoryProjects AppPermission `field:"optional" json:"repositoryProjects" yaml:"repositoryProjects"`
	// Experimental.
	Secrets AppPermission `field:"optional" json:"secrets" yaml:"secrets"`
	// Experimental.
	SecretScanningAlerts AppPermission `field:"optional" json:"secretScanningAlerts" yaml:"secretScanningAlerts"`
	// Experimental.
	SecurityEvents AppPermission `field:"optional" json:"securityEvents" yaml:"securityEvents"`
	// Experimental.
	SingleFile AppPermission `field:"optional" json:"singleFile" yaml:"singleFile"`
	// Experimental.
	Starring AppPermission `field:"optional" json:"starring" yaml:"starring"`
	// Experimental.
	Statuses AppPermission `field:"optional" json:"statuses" yaml:"statuses"`
	// Experimental.
	TeamDiscussions AppPermission `field:"optional" json:"teamDiscussions" yaml:"teamDiscussions"`
	// Experimental.
	VulnerabilityAlerts AppPermission `field:"optional" json:"vulnerabilityAlerts" yaml:"vulnerabilityAlerts"`
	// Experimental.
	Workflows AppPermission `field:"optional" json:"workflows" yaml:"workflows"`
}

