package cdktn

import (
	"github.com/projen/projen-go/projen"
)

// Experimental.
type CdktnDepsOptions struct {
	// Minimum version of the CDKTN to depend on.
	// Default: "^0.24.0"
	//
	// Experimental.
	CdktnVersion *string `field:"required" json:"cdktnVersion" yaml:"cdktnVersion"`
	// Version range of the CDKTN CLI to depend on.
	//
	// Can be either a specific version, or an NPM version range.
	//
	// By default, the latest version will be installed; you can use this
	// option to restrict it to a specific version or version range.
	// Default: - no specific version.
	//
	// Experimental.
	CdktnCliVersion *string `field:"optional" json:"cdktnCliVersion" yaml:"cdktnCliVersion"`
	// Use pinned version instead of caret version for CDKTN.
	//
	// You can use this to prevent mixed versions for your CDKTN dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdktnVersionPinning *bool `field:"optional" json:"cdktnVersionPinning" yaml:"cdktnVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Default: "^10.7.2"
	//
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
	// The type of dependency to use for the runtime `cdktn` module.
	//
	// For libraries, use peer dependencies and for apps use runtime dependencies.
	// Experimental.
	DependencyType projen.DependencyType `field:"required" json:"dependencyType" yaml:"dependencyType"`
}

