package javascript


// Sets the release-age and trust policy pnpm applies when installing config dependencies and the pnpm version pinned by the packageManager field, before the regular configuration is loaded.
//
// Added in pnpm v12.15.0.
// Experimental.
type PnpmWorkspaceYamlSchemaBootConfig struct {
	// The minimum number of minutes that must pass after a version is published before pnpm will install it.
	// Experimental.
	MinimumReleaseAge *float64 `field:"optional" json:"minimumReleaseAge" yaml:"minimumReleaseAge"`
	// Packages that are exempt from the minimumReleaseAge check.
	// Experimental.
	MinimumReleaseAgeExclude *[]*string `field:"optional" json:"minimumReleaseAgeExclude" yaml:"minimumReleaseAgeExclude"`
	// When set to no-downgrade, installation fails if a package's trust level has decreased compared to previous releases.
	// Experimental.
	TrustPolicy PnpmWorkspaceYamlSchemaBootConfigTrustPolicy `field:"optional" json:"trustPolicy" yaml:"trustPolicy"`
	// Packages or versions that are allowed to install even if they don't satisfy the trust policy requirement.
	// Experimental.
	TrustPolicyExclude *[]*string `field:"optional" json:"trustPolicyExclude" yaml:"trustPolicyExclude"`
}

