// CDK for software projects
package projen


// What to do when a service on a port is detected.
// Experimental.
type GitpodOnOpen string

const (
	// Open a new browser tab.
	// Experimental.
	GitpodOnOpen_OPEN_BROWSER GitpodOnOpen = "OPEN_BROWSER"
	// Open a preview on the right side of the IDE.
	// Experimental.
	GitpodOnOpen_OPEN_PREVIEW GitpodOnOpen = "OPEN_PREVIEW"
	// Show a notification asking the user what to do (default).
	// Experimental.
	GitpodOnOpen_NOTIFY GitpodOnOpen = "NOTIFY"
	// Do nothing.
	// Experimental.
	GitpodOnOpen_IGNORE GitpodOnOpen = "IGNORE"
)

