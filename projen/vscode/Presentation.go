package vscode


// VSCode launch configuration Presentation interface "using the order, group, and hidden attributes in the presentation object you can sort, group, and hide configurations and compounds in the Debug configuration dropdown and in the Debug quick pick." Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type Presentation struct {
	// Experimental.
	Group *string `field:"required" json:"group" yaml:"group"`
	// Experimental.
	Hidden *bool `field:"required" json:"hidden" yaml:"hidden"`
	// Experimental.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

