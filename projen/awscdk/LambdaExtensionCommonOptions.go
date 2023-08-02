package awscdk

import (
	"github.com/projen/projen-go/projen/javascript"
)

// Common options for creating lambda extensions.
// Experimental.
type LambdaExtensionCommonOptions struct {
	// Bundling options for this AWS Lambda extension.
	//
	// If not specified the default bundling options specified for the project
	// `Bundler` instance will be used.
	// Default: - defaults.
	//
	// Experimental.
	BundlingOptions *javascript.BundlingOptions `field:"optional" json:"bundlingOptions" yaml:"bundlingOptions"`
	// The extension's compatible runtimes.
	// Experimental.
	CompatibleRuntimes *[]LambdaRuntime `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
}

