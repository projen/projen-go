package polaris


// Specifies that build capture should be used to capture the project and provides the build configuration to use.
//
// If not specified and the project directory contains compiled source files then automatic build capture will be used to capture compiled source files in the project directory.
// Experimental.
type BuildConfiguration struct {
	// The build command will be invoked to use build capture to capture the project.
	//
	// A build command specified on the command-line will override this setting.
	// Experimental.
	BuildCommand *string `field:"required" json:"buildCommand" yaml:"buildCommand"`
	// Specifies whether to enable or disable the automatic invocation of Aspnet_compiler.exe for any ASP.NET 4 and earlier Web applications that are detected in the build. The output of Aspnet_compiler.exe is required by the C# and Visual Basic security checkers.
	// Experimental.
	AspnetCompiler *bool `field:"optional" json:"aspnetCompiler" yaml:"aspnetCompiler"`
	// Specifies whether to enable Bazel capture.
	// Experimental.
	Bazel *bool `field:"optional" json:"bazel" yaml:"bazel"`
	// The clean command will be invoked prior to doing build capture to capture the project.
	// Experimental.
	CleanCommand *string `field:"optional" json:"cleanCommand" yaml:"cleanCommand"`
	// Additional arguments to pass to cov-build when doing build capture.
	// Experimental.
	CovBuildArgs *[]*string `field:"optional" json:"covBuildArgs" yaml:"covBuildArgs"`
	// Specifies whether the build should only record the decompilations of byte code during the build and not attempt to decompile and emit the byte code.
	//
	// During the analysis phase, cov-build will be rerun with --replay-decomp to decompile and emit the byte code.
	// Experimental.
	DeferDecomp *bool `field:"optional" json:"deferDecomp" yaml:"deferDecomp"`
	// Specifies whether to use the instrumentation mode instead of the debugger.
	//
	// For certain builds, this configuration can significantly improve build times. This setting is applicable only on Windows.
	// Experimental.
	Instrument *bool `field:"optional" json:"instrument" yaml:"instrument"`
	// Specifies how to parallelize translation of C and C++ code.
	// Experimental.
	ParallelTranslate *ParallelTranslateConfiguration `field:"optional" json:"parallelTranslate" yaml:"parallelTranslate"`
	// Specifies whether to enable the collection of scan transparency data for build capture.
	//
	// This setting must be enabled if the Coverity Connect instance has 'scan.transparency.enabled=true' in its configuration.
	// Experimental.
	ScanTransparency *bool `field:"optional" json:"scanTransparency" yaml:"scanTransparency"`
}

