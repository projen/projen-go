package javascript


// Overrides the `onFail` field of `devEngines.runtime` (and `engines.runtime`) in the root project's `package.json`. This is useful when you want a different local behavior than what is written in the manifest — for instance, forcing pnpm to download the declared runtime even when the manifest sets `onFail: "warn"`.
// Experimental.
type PnpmWorkspaceYamlSchemaRuntimeOnFail string

const (
	// download.
	// Experimental.
	PnpmWorkspaceYamlSchemaRuntimeOnFail_DOWNLOAD PnpmWorkspaceYamlSchemaRuntimeOnFail = "DOWNLOAD"
	// error.
	// Experimental.
	PnpmWorkspaceYamlSchemaRuntimeOnFail_ERROR PnpmWorkspaceYamlSchemaRuntimeOnFail = "ERROR"
	// warn.
	// Experimental.
	PnpmWorkspaceYamlSchemaRuntimeOnFail_WARN PnpmWorkspaceYamlSchemaRuntimeOnFail = "WARN"
	// ignore.
	// Experimental.
	PnpmWorkspaceYamlSchemaRuntimeOnFail_IGNORE PnpmWorkspaceYamlSchemaRuntimeOnFail = "IGNORE"
)

