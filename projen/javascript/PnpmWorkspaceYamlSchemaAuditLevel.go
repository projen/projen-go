package javascript


// Controls the level of issues reported by `pnpm audit`.
//
// When set to 'low', all vulnerabilities are reported. When set to 'moderate', 'high', or 'critical', only vulnerabilities with that severity or higher are reported.
// Experimental.
type PnpmWorkspaceYamlSchemaAuditLevel string

const (
	// low.
	// Experimental.
	PnpmWorkspaceYamlSchemaAuditLevel_LOW PnpmWorkspaceYamlSchemaAuditLevel = "LOW"
	// moderate.
	// Experimental.
	PnpmWorkspaceYamlSchemaAuditLevel_MODERATE PnpmWorkspaceYamlSchemaAuditLevel = "MODERATE"
	// high.
	// Experimental.
	PnpmWorkspaceYamlSchemaAuditLevel_HIGH PnpmWorkspaceYamlSchemaAuditLevel = "HIGH"
	// critical.
	// Experimental.
	PnpmWorkspaceYamlSchemaAuditLevel_CRITICAL PnpmWorkspaceYamlSchemaAuditLevel = "CRITICAL"
)

