package uvconfig


// Experimental.
type PythonDownloads string

const (
	// Automatically download managed Python installations when needed.
	//
	// (automatic).
	// Experimental.
	PythonDownloads_AUTOMATIC PythonDownloads = "AUTOMATIC"
	// Do not automatically download managed Python installations;
	//
	// require explicit installation. (manual)
	// Experimental.
	PythonDownloads_MANUAL PythonDownloads = "MANUAL"
	// Do not ever allow Python downloads.
	//
	// (never).
	// Experimental.
	PythonDownloads_NEVER PythonDownloads = "NEVER"
)

