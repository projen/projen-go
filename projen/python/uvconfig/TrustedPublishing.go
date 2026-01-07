package uvconfig


// Experimental.
type TrustedPublishing string

const (
	// always.
	// Experimental.
	TrustedPublishing_ALWAYS TrustedPublishing = "ALWAYS"
	// never.
	// Experimental.
	TrustedPublishing_NEVER TrustedPublishing = "NEVER"
	// Attempt trusted publishing when we're in a supported environment, continue if that fails.
	//
	// Supported environments include GitHub Actions and GitLab CI/CD. (automatic)
	// Experimental.
	TrustedPublishing_AUTOMATIC TrustedPublishing = "AUTOMATIC"
)

