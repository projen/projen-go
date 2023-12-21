package vscode


// Options for a 'VsCodeLaunchPickStringInputEntry' Source: https://code.visualstudio.com/docs/editor/variables-reference#_input-variables.
// Experimental.
type VsCodeLaunchPickStringInputEntry struct {
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Experimental.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
	// Experimental.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

