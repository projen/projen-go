package vscode


// Controls where to launch the debug target Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type Console string

const (
	// Experimental.
	Console_INTERNAL_CONSOLE Console = "INTERNAL_CONSOLE"
	// Experimental.
	Console_INTEGRATED_TERMINAL Console = "INTEGRATED_TERMINAL"
	// Experimental.
	Console_EXTERNAL_TERMINAL Console = "EXTERNAL_TERMINAL"
)

