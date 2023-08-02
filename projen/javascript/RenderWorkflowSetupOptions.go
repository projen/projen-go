package javascript


// Options for `renderInstallSteps()`.
// Experimental.
type RenderWorkflowSetupOptions struct {
	// Should the pacakge lockfile be updated?
	// Default: false.
	//
	// Experimental.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

