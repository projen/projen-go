package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaAudit struct {
	// A list of GHSA codes that will be ignored by pnpm audit.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// Only print advisories with severity greater than or equal to this level.
	// Experimental.
	Level PnpmWorkspaceYamlSchemaAuditLevel `field:"optional" json:"level" yaml:"level"`
}

