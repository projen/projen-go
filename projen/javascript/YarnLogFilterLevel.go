package javascript


// https://v3.yarnpkg.com/configuration/yarnrc#logFilters.0.level.
// Experimental.
type YarnLogFilterLevel string

const (
	// Experimental.
	YarnLogFilterLevel_INFO YarnLogFilterLevel = "INFO"
	// Experimental.
	YarnLogFilterLevel_WARNING YarnLogFilterLevel = "WARNING"
	// Experimental.
	YarnLogFilterLevel_ERROR YarnLogFilterLevel = "ERROR"
	// Experimental.
	YarnLogFilterLevel_DISCARD YarnLogFilterLevel = "DISCARD"
)

