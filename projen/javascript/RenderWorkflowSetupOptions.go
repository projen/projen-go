package javascript


// Options for `renderInstallSteps()`.
// Experimental.
type RenderWorkflowSetupOptions struct {
	// Should the pacakge lockfile be updated?
	// Experimental.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

