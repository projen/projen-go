package javascript


// How the store records which files each package has.
//
// sqlite - one index.db database; SQLite's locking only works between processes on one machine, so this format is for a store on a local disk. files - one file per package in the store's index directory, for a store on a network filesystem that several machines use at once, such as a store shared by CI jobs; each entry is written to a temporary file and renamed into place, so concurrent jobs need no locks. The format belongs to the store: once the index directory exists, every pnpm process uses it whatever its own setting says, and the packages recorded in an existing index.db are not carried over when switching. pnpm v11 reads only index.db. Added in pnpm v12.13.0.
// Experimental.
type PnpmWorkspaceYamlSchemaStoreIndexFormat string

const (
	// sqlite.
	// Experimental.
	PnpmWorkspaceYamlSchemaStoreIndexFormat_SQLITE PnpmWorkspaceYamlSchemaStoreIndexFormat = "SQLITE"
	// files.
	// Experimental.
	PnpmWorkspaceYamlSchemaStoreIndexFormat_FILES PnpmWorkspaceYamlSchemaStoreIndexFormat = "FILES"
)

