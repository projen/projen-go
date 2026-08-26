package sonarqube


// Log level for SonarQube analysis.
// Experimental.
type SonarqubeLogLevel string

const (
	// Standard logging (default).
	// Experimental.
	SonarqubeLogLevel_INFO SonarqubeLogLevel = "INFO"
	// Verbose logging.
	// Experimental.
	SonarqubeLogLevel_DEBUG SonarqubeLogLevel = "DEBUG"
	// Most verbose, includes plugin/library output.
	// Experimental.
	SonarqubeLogLevel_TRACE SonarqubeLogLevel = "TRACE"
)

