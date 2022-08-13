package vscode


// Controls the visibility of the VSCode Debug Console panel during a debugging session Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type InternalConsoleOptions string

const (
	// Experimental.
	InternalConsoleOptions_NEVER_OPEN InternalConsoleOptions = "NEVER_OPEN"
	// Experimental.
	InternalConsoleOptions_OPEN_ON_FIRST_SESSION_START InternalConsoleOptions = "OPEN_ON_FIRST_SESSION_START"
	// Experimental.
	InternalConsoleOptions_OPEN_ON_SESSION_START InternalConsoleOptions = "OPEN_ON_SESSION_START"
)

