package github


// Use to add private registry support for dependabot.
// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#configuration-options-for-private-registries
//
// Experimental.
type DependabotRegistry struct {
	// Registry type e.g. 'npm-registry' or 'docker-registry'.
	// Experimental.
	Type DependabotRegistryType `field:"required" json:"type" yaml:"type"`
	// Url for the registry e.g. 'https://npm.pkg.github.com' or 'registry.hub.docker.com'.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// A reference to a Dependabot secret containing an access key for this registry.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Used with the hex-organization registry type.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#hex-organization
	//
	// Experimental.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A reference to a Dependabot secret containing the password for the specified user.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// For registries with type: python-index, if the boolean value is true, pip esolves dependencies by using the specified URL rather than the base URL of the Python Package Index (by default https://pypi.org/simple).
	// Experimental.
	ReplacesBase *bool `field:"optional" json:"replacesBase" yaml:"replacesBase"`
	// Secret token for dependabot access e.g. '${{ secrets.DEPENDABOT_PACKAGE_TOKEN }}'.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username that Dependabot uses to access the registry.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

