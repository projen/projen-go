package polaris


// Specifies how the project should be captured.
// Experimental.
type CaptureConfiguration struct {
	// Experimental.
	BuildCapture *BuildConfiguration `field:"optional" json:"buildCapture" yaml:"buildCapture"`
	// Specifies whether to enable or disable build command inference.
	//
	// If build command inference is disabled and no build command is provided then no attempt at build capture will be made.
	// Experimental.
	BuildCommandInference *bool `field:"optional" json:"buildCommandInference" yaml:"buildCommandInference"`
	// Specifies which compilers to configure.
	//
	// By default, template compilers are configured.
	// Experimental.
	CompilerConfiguration *CompilerConfiguration `field:"optional" json:"compilerConfiguration" yaml:"compilerConfiguration"`
	// Experimental.
	CovTranslate *CovTranslateConfiguration `field:"optional" json:"covTranslate" yaml:"covTranslate"`
	// Records additional information during the emit process needed for the compliance checkers.
	//
	// If a "coding-standards" configuration is present then this flag will automatically be set to true.
	// Experimental.
	EmitComplementaryInfo *bool `field:"optional" json:"emitComplementaryInfo" yaml:"emitComplementaryInfo"`
	// Specifies the encoding to use when parsing and emitting the source files.
	// Experimental.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Specifies the minimum percentage of files that must be captured in order to proceed with the analysis.
	// Experimental.
	FailureThresholdPercent *float64 `field:"optional" json:"failureThresholdPercent" yaml:"failureThresholdPercent"`
	// Specifies which non-compiled files to capture.
	//
	// By default, all files are captured.
	// Experimental.
	Files *FilesConfiguration `field:"optional" json:"files" yaml:"files"`
	// Force resolution of Maven, Gradle and MSBuild dependencies even if this is not needed based on the detected source languages in the project.
	// Experimental.
	ForceDependencyResolution *bool `field:"optional" json:"forceDependencyResolution" yaml:"forceDependencyResolution"`
	// Specifies how to import data about source file changes from the source control management system.
	// Experimental.
	ImportScm *ImportScmConfiguration `field:"optional" json:"importScm" yaml:"importScm"`
	// Specifies which languages to include or exclude for capture.
	//
	// By default, all languages are captured.
	// Experimental.
	Languages *LanguagesConfiguration `field:"optional" json:"languages" yaml:"languages"`
	// Specifies whether to limit the group of emitted JAR files to those needed for compilation of the Java files.
	//
	// The default behavior without this option is to emit all the JAR files in the classpath regardless of whether they are referenced by a Java file in the compilation.
	// Experimental.
	MinimalClasspathEmit *bool `field:"optional" json:"minimalClasspathEmit" yaml:"minimalClasspathEmit"`
	// Specifies whether to do a complete capture or a record with source capture.
	// Experimental.
	RecordWithSource *bool `field:"optional" json:"recordWithSource" yaml:"recordWithSource"`
	// Enables or disables security dynamic analysis.
	//
	// If set to true (the default), security dynamic analysis is run as part of the capture step. If set to false, security dynamic analysis is not run.
	// Experimental.
	SecurityDa *bool `field:"optional" json:"securityDa" yaml:"securityDa"`
}

