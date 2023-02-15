// CDK for software projects
package projen


// Configure where in the IDE the terminal should be opened.
// Experimental.
type GitpodOpenIn string

const (
	// the bottom panel (default).
	// Experimental.
	GitpodOpenIn_BOTTOM GitpodOpenIn = "BOTTOM"
	// the left panel.
	// Experimental.
	GitpodOpenIn_LEFT GitpodOpenIn = "LEFT"
	// the right panel.
	// Experimental.
	GitpodOpenIn_RIGHT GitpodOpenIn = "RIGHT"
	// the main editor area.
	// Experimental.
	GitpodOpenIn_MAIN GitpodOpenIn = "MAIN"
)

