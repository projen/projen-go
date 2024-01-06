package javascript

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `renderWorkflowSetup()`.
// Experimental.
type RenderWorkflowSetupOptions struct {
	// Configure the install step in the workflow setup.
	//
	// Example:
	//   - { env: { NPM_TOKEN: "token" }} for installing from private npm registry.
	//
	// Default: - `{ name: "Install dependencies" }`.
	//
	// Experimental.
	InstallStepConfiguration *workflows.JobStepConfiguration `field:"optional" json:"installStepConfiguration" yaml:"installStepConfiguration"`
	// Should the package lockfile be updated?
	// Default: false.
	//
	// Experimental.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

