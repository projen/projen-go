package vscode


// Options for a 'VsCodeLaunchCommandInputEntry' Source: https://code.visualstudio.com/docs/editor/variables-reference#_input-variables.
// Experimental.
type VsCodeLaunchCommandInputEntry struct {
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Experimental.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Experimental.
	Args interface{} `field:"optional" json:"args" yaml:"args"`
}

