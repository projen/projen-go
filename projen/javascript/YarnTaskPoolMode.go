package javascript


// https://yarnpkg.com/configuration/yarnrc#taskPoolMode.
// Experimental.
type YarnTaskPoolMode string

const (
	// Experimental.
	YarnTaskPoolMode_ASYNC YarnTaskPoolMode = "ASYNC"
	// Experimental.
	YarnTaskPoolMode_WORKERS YarnTaskPoolMode = "WORKERS"
)

