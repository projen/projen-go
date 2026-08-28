package polaris


// Specifies the aggressiveness level for the analysis.
//
// The aggressiveness level causes the analysis to make more or less aggressive assumptions during the analysis where the higher the aggressiveness level the more defects are reported.
// Experimental.
type AnalysisConfigurationAggressivenessLevel string

const (
	// low.
	// Experimental.
	AnalysisConfigurationAggressivenessLevel_LOW AnalysisConfigurationAggressivenessLevel = "LOW"
	// medium.
	// Experimental.
	AnalysisConfigurationAggressivenessLevel_MEDIUM AnalysisConfigurationAggressivenessLevel = "MEDIUM"
	// high.
	// Experimental.
	AnalysisConfigurationAggressivenessLevel_HIGH AnalysisConfigurationAggressivenessLevel = "HIGH"
)

