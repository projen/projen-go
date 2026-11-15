package javascript


// When set to no-downgrade, installation fails if a package's trust level has decreased compared to previous releases.
// Experimental.
type PnpmWorkspaceYamlSchemaBootConfigTrustPolicy string

const (
	// off.
	// Experimental.
	PnpmWorkspaceYamlSchemaBootConfigTrustPolicy_OFF PnpmWorkspaceYamlSchemaBootConfigTrustPolicy = "OFF"
	// no-downgrade.
	// Experimental.
	PnpmWorkspaceYamlSchemaBootConfigTrustPolicy_NO_HYPHEN_DOWNGRADE PnpmWorkspaceYamlSchemaBootConfigTrustPolicy = "NO_HYPHEN_DOWNGRADE"
)

