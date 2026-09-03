package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaAudit struct {
	// A list of GHSA codes that will be ignored by pnpm audit.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// When `true`, `pnpm audit --fix` removes the `audit.ignore` entries whose GHSA no longer appears in the audit report, so a list of tolerated advisories doesn't accumulate entries for dependencies that are long gone. Added in: v11.25.0 and v12.0.0.
	// Experimental.
	IgnorePrune *bool `field:"optional" json:"ignorePrune" yaml:"ignorePrune"`
	// Only print advisories with severity greater than or equal to this level.
	// Experimental.
	Level PnpmWorkspaceYamlSchemaAuditLevel `field:"optional" json:"level" yaml:"level"`
}

