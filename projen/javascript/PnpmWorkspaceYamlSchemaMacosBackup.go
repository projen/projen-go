package javascript


// On macOS, exclude newly created pnpm directories from Time Machine backups by marking them with the com.apple.metadata:com_apple_backup_excludeItem extended attribute. Read from the global configuration file or the PNPM_CONFIG_MACOS_BACKUP_* environment variables; ignored in project pnpm-workspace.yaml files. Added in pnpm v12.6.0.
// Experimental.
type PnpmWorkspaceYamlSchemaMacosBackup struct {
	// When true on macOS, newly created node_modules, virtual-store, and similar directories are marked so Time Machine skips them.
	// Experimental.
	ExcludeModulesDir *bool `field:"optional" json:"excludeModulesDir" yaml:"excludeModulesDir"`
	// When true on macOS, newly created package-store directories are marked so Time Machine skips them.
	// Experimental.
	ExcludeStoreDir *bool `field:"optional" json:"excludeStoreDir" yaml:"excludeStoreDir"`
}

