package polaris


// Format in which to save defects.
//
// Either "html" or "json".
// Experimental.
type CommitConfigurationLocalFormat string

const (
	// html.
	// Experimental.
	CommitConfigurationLocalFormat_HTML CommitConfigurationLocalFormat = "HTML"
	// json.
	// Experimental.
	CommitConfigurationLocalFormat_JSON CommitConfigurationLocalFormat = "JSON"
)

