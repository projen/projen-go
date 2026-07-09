package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaAuditConfig struct {
	// A list of CVE IDs that will be ignored by "pnpm audit".
	// Experimental.
	IgnoreCves *[]*string `field:"optional" json:"ignoreCves" yaml:"ignoreCves"`
	// A list of GHSA Codes that will be ignored by "pnpm audit".
	// Experimental.
	IgnoreGhsas *[]*string `field:"optional" json:"ignoreGhsas" yaml:"ignoreGhsas"`
}

