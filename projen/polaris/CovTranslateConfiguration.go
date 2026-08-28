package polaris


// Command to invoke that will invoke "cov-translate" to capture the project.
// Experimental.
type CovTranslateConfiguration struct {
	// This key specifies a command to invoke that will invoke "cov-translate" in the case where the user is doing a "cov-translate" capture.
	// Experimental.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Additional arguments to pass to cov-build when invoking the provided command.
	// Experimental.
	CovBuildArgs *[]*string `field:"optional" json:"covBuildArgs" yaml:"covBuildArgs"`
	// Specifies whether the build should only record the decompilations of byte code during the build and not attempt to decompile and emit the byte code.
	//
	// During the analysis phase, cov-build will be rerun with --replay-decomp to decompile and emit the byte code.
	// Experimental.
	DeferDecomp *bool `field:"optional" json:"deferDecomp" yaml:"deferDecomp"`
	// Specifies how to parallelize translation of C and C++ code.
	// Experimental.
	ParallelTranslate *ParallelTranslateConfiguration `field:"optional" json:"parallelTranslate" yaml:"parallelTranslate"`
	// Specifies whether to enable the collection of scan transparency data for cov-translate capture.
	//
	// This setting must be enabled if the Coverity Connect instance has 'scan.transparency.enabled=true' in its configuration.
	// Experimental.
	ScanTransparency *bool `field:"optional" json:"scanTransparency" yaml:"scanTransparency"`
}

