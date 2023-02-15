package workflows


// The permissions available for an access token for a GitHub App.
// Experimental.
type AppPermission string

const (
	// Read-only acccess.
	// Experimental.
	AppPermission_READ AppPermission = "READ"
	// Read-write access.
	// Experimental.
	AppPermission_WRITE AppPermission = "WRITE"
	// Read-write and admin access.
	//
	// Not all permissions support `admin`.
	// Experimental.
	AppPermission_ADMIN AppPermission = "ADMIN"
)

