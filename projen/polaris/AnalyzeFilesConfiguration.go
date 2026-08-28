package polaris


// Experimental.
type AnalyzeFilesConfiguration struct {
	// Glob pattern that specifies the set of source files to exclude from analysis.
	//
	// Note that any include glob patterns and regular expressions are processed prior to handling exclude glob patterns and regular expressions.
	// Experimental.
	ExcludeGlob *string `field:"optional" json:"excludeGlob" yaml:"excludeGlob"`
	// Regular expression that specifies the set of source files to exclude from analysis.
	//
	// Note that any include glob patterns and regular expressions are processed prior to handling exclude glob patterns and regular expressions.
	// Experimental.
	ExcludeRegex *string `field:"optional" json:"excludeRegex" yaml:"excludeRegex"`
	// Paths of source files to analyze.
	//
	// Include and exclude glob patterns and regular expressions are applied to determine which of these files are actually analyzed.
	// Experimental.
	IncludeFiles *string `field:"optional" json:"includeFiles" yaml:"includeFiles"`
	// Glob pattern that specifies the set of source files to analyze.
	// Experimental.
	IncludeGlob *string `field:"optional" json:"includeGlob" yaml:"includeGlob"`
	// File containing the paths of source files to analyze, one per line.
	//
	// Include and exclude glob patterns and regular expressions are applied to determine which of these files are actually analyzed.
	// Experimental.
	IncludeListFile *string `field:"optional" json:"includeListFile" yaml:"includeListFile"`
	// Regular expression that specifies the set of source files to analyze.
	// Experimental.
	IncludeRegex *string `field:"optional" json:"includeRegex" yaml:"includeRegex"`
}

