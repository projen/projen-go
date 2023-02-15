package vscode


// VSCode launch configuration ServerReadyAction interface "if you want to open a URL in a web browser whenever the program under debugging outputs a specific message to the debug console or integrated terminal." Source: https://code.visualstudio.com/docs/editor/debugging#_launchjson-attributes.
// Experimental.
type ServerReadyAction struct {
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Experimental.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Experimental.
	UriFormat *string `field:"optional" json:"uriFormat" yaml:"uriFormat"`
}

