package javascript


// Determines where the virtual store is located.
//
// When set to project, a separate virtual store is created in each project's node_modules/.pnpm. When set to global, a single store is shared by every project on the machine, with each project's node_modules holding only symlinks into it. Added in pnpm v11.23.0.
// Experimental.
type PnpmWorkspaceYamlSchemaVirtualStoreType string

const (
	// project.
	// Experimental.
	PnpmWorkspaceYamlSchemaVirtualStoreType_PROJECT PnpmWorkspaceYamlSchemaVirtualStoreType = "PROJECT"
	// global.
	// Experimental.
	PnpmWorkspaceYamlSchemaVirtualStoreType_GLOBAL PnpmWorkspaceYamlSchemaVirtualStoreType = "GLOBAL"
)

