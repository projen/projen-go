package vscode


// Options for a 'VsCodeLaunchPromptStringInputEntry' Source: https://code.visualstudio.com/docs/editor/variables-reference#_input-variables.
// Experimental.
type VsCodeLaunchPromptStringInputEntry struct {
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Experimental.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Experimental.
	Password *bool `field:"optional" json:"password" yaml:"password"`
}

