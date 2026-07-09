package javascript


// Controls the way packages are imported from the store (if you want to disable symlinks inside node_modules, then you need to change the nodeLinker setting, not this one).
// Experimental.
type PnpmWorkspaceYamlSchemaPackageImportMethod string

const (
	// auto.
	// Experimental.
	PnpmWorkspaceYamlSchemaPackageImportMethod_AUTO PnpmWorkspaceYamlSchemaPackageImportMethod = "AUTO"
	// hardlink.
	// Experimental.
	PnpmWorkspaceYamlSchemaPackageImportMethod_HARDLINK PnpmWorkspaceYamlSchemaPackageImportMethod = "HARDLINK"
	// copy.
	// Experimental.
	PnpmWorkspaceYamlSchemaPackageImportMethod_COPY PnpmWorkspaceYamlSchemaPackageImportMethod = "COPY"
	// clone.
	// Experimental.
	PnpmWorkspaceYamlSchemaPackageImportMethod_CLONE PnpmWorkspaceYamlSchemaPackageImportMethod = "CLONE"
	// clone-or-copy.
	// Experimental.
	PnpmWorkspaceYamlSchemaPackageImportMethod_CLONE_HYPHEN_OR_HYPHEN_COPY PnpmWorkspaceYamlSchemaPackageImportMethod = "CLONE_HYPHEN_OR_HYPHEN_COPY"
)

