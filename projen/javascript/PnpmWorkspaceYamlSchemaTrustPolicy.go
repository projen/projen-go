package javascript


// When set to no-downgrade, pnpm will fail if a package's trust level has decreased compared to previous releases.
//
// For example, if a package was previously published by a trusted publisher but now only has provenance or no trust evidence, installation will fail. This helps prevent installing potentially compromised versions.
// Experimental.
type PnpmWorkspaceYamlSchemaTrustPolicy string

const (
	// off.
	// Experimental.
	PnpmWorkspaceYamlSchemaTrustPolicy_OFF PnpmWorkspaceYamlSchemaTrustPolicy = "OFF"
	// no-downgrade.
	// Experimental.
	PnpmWorkspaceYamlSchemaTrustPolicy_NO_HYPHEN_DOWNGRADE PnpmWorkspaceYamlSchemaTrustPolicy = "NO_HYPHEN_DOWNGRADE"
)

