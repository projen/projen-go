package github


// Each configuration type requires you to provide particular settings.
//
// Some types allow more than one way to connect.
// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#configuration-options-for-private-registries
//
// Experimental.
type DependabotRegistryType string

const (
	// The composer-repository type supports username and password.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#composer-repository
	//
	// Experimental.
	DependabotRegistryType_COMPOSER_REGISTRY DependabotRegistryType = "COMPOSER_REGISTRY"
	// The docker-registry type supports username and password.
	//
	// The docker-registry type can also be used to pull from Amazon ECR using static AWS credentials.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#docker-registry
	//
	// Experimental.
	DependabotRegistryType_DOCKER_REGISTRY DependabotRegistryType = "DOCKER_REGISTRY"
	// The git type supports username and password.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#git
	//
	// Experimental.
	DependabotRegistryType_GIT DependabotRegistryType = "GIT"
	// The hex-organization type supports organization and key.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#hex-organization
	//
	// Experimental.
	DependabotRegistryType_HEX_ORGANIZATION DependabotRegistryType = "HEX_ORGANIZATION"
	// The maven-repository type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#maven-repository
	//
	// Experimental.
	DependabotRegistryType_MAVEN_REPOSITORY DependabotRegistryType = "MAVEN_REPOSITORY"
	// The npm-registry type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#npm-registry
	//
	// Experimental.
	DependabotRegistryType_NPM_REGISTRY DependabotRegistryType = "NPM_REGISTRY"
	// The nuget-feed type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#nuget-feed
	//
	// Experimental.
	DependabotRegistryType_NUGET_FEED DependabotRegistryType = "NUGET_FEED"
	// The python-index type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#python-index
	//
	// Experimental.
	DependabotRegistryType_PYTHON_INDEX DependabotRegistryType = "PYTHON_INDEX"
	// The rubygems-server type supports username and password, or token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#rubygems-server
	//
	// Experimental.
	DependabotRegistryType_RUBYGEMS_SERVER DependabotRegistryType = "RUBYGEMS_SERVER"
	// The terraform-registry type supports a token.
	// See: https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically/configuration-options-for-dependency-updates#terraform-registry
	//
	// Experimental.
	DependabotRegistryType_TERRAFORM_REGISTRY DependabotRegistryType = "TERRAFORM_REGISTRY"
)

