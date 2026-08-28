package polaris


// Specifies how the project should be analyzed.
// Experimental.
type AnalysisConfiguration struct {
	// Specifies the aggressiveness level for the analysis.
	//
	// The aggressiveness level causes the analysis to make more or less aggressive assumptions during the analysis where the higher the aggressiveness level the more defects are reported.
	// Experimental.
	AggressivenessLevel AnalysisConfigurationAggressivenessLevel `field:"optional" json:"aggressivenessLevel" yaml:"aggressivenessLevel"`
	// Enables callgraph metrics output in the intermediate directory.
	// Experimental.
	CallgraphMetrics *bool `field:"optional" json:"callgraphMetrics" yaml:"callgraphMetrics"`
	// Enables analysis of calls to function pointers for defects.
	// Experimental.
	CCppFnptr *bool `field:"optional" json:"cCppFnptr" yaml:"cCppFnptr"`
	// Enables full virtual-call resolution for C++.
	// Experimental.
	CCppVirtual *bool `field:"optional" json:"cCppVirtual" yaml:"cCppVirtual"`
	// If no checker configuration is specified, the CLI will enable a set of checkers based on the files that were captured.
	// Experimental.
	Checkers *CheckerConfiguration `field:"optional" json:"checkers" yaml:"checkers"`
	// If specified, the analysis will scan the code for compliance according to the given coding standard configuration.
	//
	// If this configuration is present, the capture "emit-complementary-info" flag will be set to true.
	// Experimental.
	CodingStandards *CodingStandardConfiguration `field:"optional" json:"codingStandards" yaml:"codingStandards"`
	// Coverity Connect configuration to use when performing analysis in Coverity Connect.
	// Experimental.
	Connect *AnalyzeConnectConfiguration `field:"optional" json:"connect" yaml:"connect"`
	// Enables additional filtering of defects by using an additional false-path pruner.
	//
	// If set to true, the constraint FPP is enabled.
	// Experimental.
	ConstraintFpp *bool `field:"optional" json:"constraintFpp" yaml:"constraintFpp"`
	// Additional arguments to pass to cov-analyze when doing analysis.
	// Experimental.
	CovAnalyzeArgs *[]*string `field:"optional" json:"covAnalyzeArgs" yaml:"covAnalyzeArgs"`
	// Additional arguments to pass to cov-collect-models following analysis when "output-model-file" is specified.
	// Experimental.
	CovCollectModelsArgs *[]*string `field:"optional" json:"covCollectModelsArgs" yaml:"covCollectModelsArgs"`
	// Specifies directives to use for the analysis, including for web application security analysis.
	// Experimental.
	Directives *[]*DirectivesConfiguration `field:"optional" json:"directives" yaml:"directives"`
	// Specifies which files to analyze when the "analyze.mode" setting is "hfi". Analysis will be performed for only these files.
	// Experimental.
	Files *AnalyzeFilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// Specifies analysis worker parallelism.
	// Experimental.
	Jobs *[]*JobsConfiguration `field:"optional" json:"jobs" yaml:"jobs"`
	// Specifies whether the analysis should be done locally, in Coverity Connect, or in Software Risk Manager.
	//
	// The possible values are as follows: connect - Run the analysis in the Coverity Connect job farm; srm - Run the analysis in the Software Risk Manager job farm; local - Run the analysis locally.
	// Experimental.
	Location AnalysisConfigurationLocation `field:"optional" json:"location" yaml:"location"`
	// Analysis mode: "pfi" (perfect fidelity incremental) for complete analysis;
	//
	// or "hfi" (high fidelity incremental) for analysis of only specific files specified by analyze.files settings, omitting any other files which may have been incidentally captured by the build. An "hfi" analysis can be faster but may produce results which are incomplete or inconsistent, due to the lack of context, and should be used only when speed is more important than accuracy.
	// Experimental.
	Mode AnalysisConfigurationMode `field:"optional" json:"mode" yaml:"mode"`
	// File containing function models.
	//
	// This overrides models specified in the default location of "config/user_models.xmldb".
	// Experimental.
	ModelFile *string `field:"optional" json:"modelFile" yaml:"modelFile"`
	// If set to to true, only one TU (translation unit) will be analyzed per source file name.
	//
	// If set to false, all translation units will be analyzed.
	// Experimental.
	OneTuPerPsf *bool `field:"optional" json:"oneTuPerPsf" yaml:"oneTuPerPsf"`
	// Output file to which function models for the project should be written following analysis.
	// Experimental.
	OutputModelFile *string `field:"optional" json:"outputModelFile" yaml:"outputModelFile"`
	// Specifies how parse warnings are handled.
	// Experimental.
	ParseWarnings *ParseWarningsConfiguration `field:"optional" json:"parseWarnings" yaml:"parseWarnings"`
	// Specifies whether to enable the collection of scan transparency data for analysis.
	//
	// This setting must be enabled if the Coverity Connect instance has 'scan.transparency.enabled=true' in its configuration.
	// Experimental.
	ScanTransparency *bool `field:"optional" json:"scanTransparency" yaml:"scanTransparency"`
	// Specifies options for Sigma analysis.
	// Experimental.
	Sigma *SigmaConfiguration `field:"optional" json:"sigma" yaml:"sigma"`
	// This is a map from trust option name to boolean to indicate whether the particular trust property should be trusted or distrusted.
	//
	// The trust option "all" controls whether all trust options should be trusted or distrusted.
	// Experimental.
	Trust interface{} `field:"optional" json:"trust" yaml:"trust"`
}

