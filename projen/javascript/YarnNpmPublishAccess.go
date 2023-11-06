package javascript


// https://yarnpkg.com/configuration/yarnrc#npmPublishAccess.
// Experimental.
type YarnNpmPublishAccess string

const (
	// Experimental.
	YarnNpmPublishAccess_PUBLIC YarnNpmPublishAccess = "PUBLIC"
	// Experimental.
	YarnNpmPublishAccess_RESTRICTED YarnNpmPublishAccess = "RESTRICTED"
)

