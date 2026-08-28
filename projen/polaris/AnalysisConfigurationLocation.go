package polaris


// Specifies whether the analysis should be done locally, in Coverity Connect, or in Software Risk Manager.
//
// The possible values are as follows: connect - Run the analysis in the Coverity Connect job farm; srm - Run the analysis in the Software Risk Manager job farm; local - Run the analysis locally.
// Experimental.
type AnalysisConfigurationLocation string

const (
	// local.
	// Experimental.
	AnalysisConfigurationLocation_LOCAL AnalysisConfigurationLocation = "LOCAL"
	// connect.
	// Experimental.
	AnalysisConfigurationLocation_CONNECT AnalysisConfigurationLocation = "CONNECT"
	// srm.
	// Experimental.
	AnalysisConfigurationLocation_SRM AnalysisConfigurationLocation = "SRM"
)

