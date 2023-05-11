package projen


// Which type of project this is.
// Deprecated: no longer supported at the base project level.
type ProjectType string

const (
	// This module may be a either a library or an app.
	// Deprecated: no longer supported at the base project level.
	ProjectType_UNKNOWN ProjectType = "UNKNOWN"
	// This is a library, intended to be published to a package manager and consumed by other projects.
	// Deprecated: no longer supported at the base project level.
	ProjectType_LIB ProjectType = "LIB"
	// This is an app (service, tool, website, etc).
	//
	// Its artifacts are intended to
	// be deployed or published for end-user consumption.
	// Deprecated: no longer supported at the base project level.
	ProjectType_APP ProjectType = "APP"
)

