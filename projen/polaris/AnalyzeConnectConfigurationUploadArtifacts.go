package polaris


// Artifacts to upload following analysis when the analysis location is Connect.
// Experimental.
type AnalyzeConnectConfigurationUploadArtifacts string

const (
	// All.
	// Experimental.
	AnalyzeConnectConfigurationUploadArtifacts_ALL AnalyzeConnectConfigurationUploadArtifacts = "ALL"
	// LogsOnly.
	// Experimental.
	AnalyzeConnectConfigurationUploadArtifacts_LOGS_ONLY AnalyzeConnectConfigurationUploadArtifacts = "LOGS_ONLY"
	// None.
	// Experimental.
	AnalyzeConnectConfigurationUploadArtifacts_NONE AnalyzeConnectConfigurationUploadArtifacts = "NONE"
	// OnFailure.
	// Experimental.
	AnalyzeConnectConfigurationUploadArtifacts_ON_FAILURE AnalyzeConnectConfigurationUploadArtifacts = "ON_FAILURE"
)

