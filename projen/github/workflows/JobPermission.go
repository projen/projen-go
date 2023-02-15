package workflows


// Access level for workflow permission scopes.
// Experimental.
type JobPermission string

const (
	// Read-only access.
	// Experimental.
	JobPermission_READ JobPermission = "READ"
	// Read-write access.
	// Experimental.
	JobPermission_WRITE JobPermission = "WRITE"
	// No access at all.
	// Experimental.
	JobPermission_NONE JobPermission = "NONE"
)

