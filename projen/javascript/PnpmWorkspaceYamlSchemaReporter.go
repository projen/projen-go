package javascript


// Allows you to customize the output style of the logs.
//
// https://pnpm.io/cli/install#--reportername
// Experimental.
type PnpmWorkspaceYamlSchemaReporter string

const (
	// silent.
	// Experimental.
	PnpmWorkspaceYamlSchemaReporter_SILENT PnpmWorkspaceYamlSchemaReporter = "SILENT"
	// default.
	// Experimental.
	PnpmWorkspaceYamlSchemaReporter_DEFAULT PnpmWorkspaceYamlSchemaReporter = "DEFAULT"
	// append-only.
	// Experimental.
	PnpmWorkspaceYamlSchemaReporter_APPEND_HYPHEN_ONLY PnpmWorkspaceYamlSchemaReporter = "APPEND_HYPHEN_ONLY"
	// ndjson.
	// Experimental.
	PnpmWorkspaceYamlSchemaReporter_NDJSON PnpmWorkspaceYamlSchemaReporter = "NDJSON"
)

