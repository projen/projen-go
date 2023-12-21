package vscode


// Base options for a 'VsCodeLaunchInputEntry' Source: https://code.visualstudio.com/docs/editor/variables-reference#_input-variables.
// Experimental.
type VsCodeLaunchInputEntry struct {
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
}

