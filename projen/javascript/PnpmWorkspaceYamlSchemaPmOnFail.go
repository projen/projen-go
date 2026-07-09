package javascript


// Overrides the `onFail` behavior of both the `packageManager` field and `devEngines.packageManager` when the running pnpm version does not match the declared one.
// Experimental.
type PnpmWorkspaceYamlSchemaPmOnFail string

const (
	// download.
	// Experimental.
	PnpmWorkspaceYamlSchemaPmOnFail_DOWNLOAD PnpmWorkspaceYamlSchemaPmOnFail = "DOWNLOAD"
	// error.
	// Experimental.
	PnpmWorkspaceYamlSchemaPmOnFail_ERROR PnpmWorkspaceYamlSchemaPmOnFail = "ERROR"
	// warn.
	// Experimental.
	PnpmWorkspaceYamlSchemaPmOnFail_WARN PnpmWorkspaceYamlSchemaPmOnFail = "WARN"
	// ignore.
	// Experimental.
	PnpmWorkspaceYamlSchemaPmOnFail_IGNORE PnpmWorkspaceYamlSchemaPmOnFail = "IGNORE"
)

