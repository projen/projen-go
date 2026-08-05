package javascript


// Changelog storage mode.
//
// `registry` composes changelog entries at publish time, while `repository` commits CHANGELOG.md files in packages.
// Experimental.
type PnpmWorkspaceYamlSchemaVersioningChangelogStorage string

const (
	// registry.
	// Experimental.
	PnpmWorkspaceYamlSchemaVersioningChangelogStorage_REGISTRY PnpmWorkspaceYamlSchemaVersioningChangelogStorage = "REGISTRY"
	// repository.
	// Experimental.
	PnpmWorkspaceYamlSchemaVersioningChangelogStorage_REPOSITORY PnpmWorkspaceYamlSchemaVersioningChangelogStorage = "REPOSITORY"
)

