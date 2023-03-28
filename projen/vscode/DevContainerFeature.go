package vscode


// devcontainer features options.
// See: https://containers.dev/implementors/features/#devcontainer-json-properties
//
// Experimental.
type DevContainerFeature struct {
	// feature name.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// feature version.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

