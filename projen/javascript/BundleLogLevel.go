package javascript


// Log levels for esbuild and package managers' install commands.
// Experimental.
type BundleLogLevel string

const (
	// Show everything.
	// Experimental.
	BundleLogLevel_VERBOSE BundleLogLevel = "VERBOSE"
	// Show everything from info and some additional messages for debugging.
	// Experimental.
	BundleLogLevel_DEBUG BundleLogLevel = "DEBUG"
	// Show warnings, errors, and an output file summary.
	// Experimental.
	BundleLogLevel_INFO BundleLogLevel = "INFO"
	// Show warnings and errors.
	// Experimental.
	BundleLogLevel_WARNING BundleLogLevel = "WARNING"
	// Show errors only.
	// Experimental.
	BundleLogLevel_ERROR BundleLogLevel = "ERROR"
	// Show nothing.
	// Experimental.
	BundleLogLevel_SILENT BundleLogLevel = "SILENT"
)

