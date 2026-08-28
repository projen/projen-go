package polaris


// Artifacts to upload following analysis when the analysis location is Connect.
// Experimental.
type CommitConfigurationConnectUploadArtifacts string

const (
	// All.
	// Experimental.
	CommitConfigurationConnectUploadArtifacts_ALL CommitConfigurationConnectUploadArtifacts = "ALL"
	// LogsOnly.
	// Experimental.
	CommitConfigurationConnectUploadArtifacts_LOGS_ONLY CommitConfigurationConnectUploadArtifacts = "LOGS_ONLY"
	// None.
	// Experimental.
	CommitConfigurationConnectUploadArtifacts_NONE CommitConfigurationConnectUploadArtifacts = "NONE"
	// OnFailure.
	// Experimental.
	CommitConfigurationConnectUploadArtifacts_ON_FAILURE CommitConfigurationConnectUploadArtifacts = "ON_FAILURE"
)

