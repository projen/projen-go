package polaris


// Analysis mode: "pfi" (perfect fidelity incremental) for complete analysis;
//
// or "hfi" (high fidelity incremental) for analysis of only specific files specified by analyze.files settings, omitting any other files which may have been incidentally captured by the build. An "hfi" analysis can be faster but may produce results which are incomplete or inconsistent, due to the lack of context, and should be used only when speed is more important than accuracy.
// Experimental.
type AnalysisConfigurationMode string

const (
	// hfi.
	// Experimental.
	AnalysisConfigurationMode_HFI AnalysisConfigurationMode = "HFI"
	// pfi.
	// Experimental.
	AnalysisConfigurationMode_PFI AnalysisConfigurationMode = "PFI"
)

