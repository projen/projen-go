package polaris


// Experimental.
type FilesConfiguration struct {
	// Specifies whether to enable capture of minified JavaScript files.
	// Experimental.
	EmitMinifiedJs *bool `field:"optional" json:"emitMinifiedJs" yaml:"emitMinifiedJs"`
	// Glob pattern that specifies the set of source files to exclude from capture.
	//
	// Note that any include glob patterns and regular expressions are processed prior to handling exclude glob patterns and regular expressions.
	// Experimental.
	ExcludeGlob *string `field:"optional" json:"excludeGlob" yaml:"excludeGlob"`
	// Regular expression that specifies the set of source files to exclude from capture.
	//
	// Note that any include glob patterns and regular expressions are processed prior to handling exclude glob patterns and regular expressions.
	// Experimental.
	ExcludeRegex *string `field:"optional" json:"excludeRegex" yaml:"excludeRegex"`
	// List of directory basenames to include for capture, which would normally have been excluded.
	//
	// By default, directories named "vendor" or "node_modules" are excluded, as are directories whose names begin with "."
	// Experimental.
	IncludeDirs *[]*string `field:"optional" json:"includeDirs" yaml:"includeDirs"`
	// Glob pattern that specifies the set of source files to capture.
	// Experimental.
	IncludeGlob *string `field:"optional" json:"includeGlob" yaml:"includeGlob"`
	// File containing the paths of source files to capture, one per line.
	//
	// Include and exclude glob patterns and regular expressions are applied to determine which of these files are actually captured.
	// Experimental.
	IncludeListFile *string `field:"optional" json:"includeListFile" yaml:"includeListFile"`
	// Regular expression that specifies the set of source files to capture.
	// Experimental.
	IncludeRegex *string `field:"optional" json:"includeRegex" yaml:"includeRegex"`
	// Specifies the Java version to use when parsing and emitting Java source files with buildless capture.
	// Experimental.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// List of directories to look in for dependencies to use during capture.
	// Experimental.
	LibraryDirs *[]*string `field:"optional" json:"libraryDirs" yaml:"libraryDirs"`
	// List of file dependencies to use during capture.
	// Experimental.
	LibraryFiles *[]*string `field:"optional" json:"libraryFiles" yaml:"libraryFiles"`
	// Specifies information about which web-application archives should be captured.
	//
	// By default all webapp archives are captured.
	// Experimental.
	WebappArchives *[]*WebappArchiveConfiguration `field:"optional" json:"webappArchives" yaml:"webappArchives"`
}

