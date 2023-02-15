package javascript


// Npm package access level.
// Experimental.
type NpmAccess string

const (
	// Package is public.
	// Experimental.
	NpmAccess_PUBLIC NpmAccess = "PUBLIC"
	// Package can only be accessed with credentials.
	// Experimental.
	NpmAccess_RESTRICTED NpmAccess = "RESTRICTED"
)

