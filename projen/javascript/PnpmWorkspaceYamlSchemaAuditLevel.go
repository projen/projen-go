package javascript


// Only print advisories with severity greater than or equal to this level.
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

