package javascript


// Automatic bump modes.
// Experimental.
type AutoRelease string

const (
	// Automatically bump & release a new version for every commit to "main".
	// Experimental.
	AutoRelease_EVERY_COMMIT AutoRelease = "EVERY_COMMIT"
	// Automatically bump & release a new version on a daily basis.
	// Experimental.
	AutoRelease_DAILY AutoRelease = "DAILY"
)

