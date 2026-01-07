package uvconfig


// When to use authentication.
// Experimental.
type AuthPolicy string

const (
	// Authenticate when necessary.
	//
	// If credentials are provided, they will be used. Otherwise, an unauthenticated request will
	// be attempted first. If the request fails, uv will search for credentials. If credentials are
	// found, an authenticated request will be attempted. (auto)
	// Experimental.
	AuthPolicy_AUTO AuthPolicy = "AUTO"
	// Always authenticate.
	//
	// If credentials are not provided, uv will eagerly search for credentials. If credentials
	// cannot be found, uv will error instead of attempting an unauthenticated request. (always)
	// Experimental.
	AuthPolicy_ALWAYS AuthPolicy = "ALWAYS"
	// Never authenticate.
	//
	// If credentials are provided, uv will error. uv will not search for credentials. (never)
	// Experimental.
	AuthPolicy_NEVER AuthPolicy = "NEVER"
)

