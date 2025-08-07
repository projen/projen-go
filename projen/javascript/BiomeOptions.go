package javascript

import (
	"github.com/projen/projen-go/projen/javascript/biomeconfig"
)

// Experimental.
type BiomeOptions struct {
	// Enable code assist with recommended actions.
	// Default: true.
	//
	// Experimental.
	Assist *bool `field:"optional" json:"assist" yaml:"assist"`
	// Full Biome configuration.
	//
	// This configuration dictates the final outcome if value is set.
	// For example, if the linter is disabled at the top-level, it can be enabled with `biomeConfig.linter.enabled`.
	// Experimental.
	BiomeConfig *biomeconfig.BiomeConfiguration `field:"optional" json:"biomeConfig" yaml:"biomeConfig"`
	// Enable code formatter with recommended settings.
	// Default: true.
	//
	// Experimental.
	Formatter *bool `field:"optional" json:"formatter" yaml:"formatter"`
	// Automatically ignore all generated files.
	//
	// This prevents Biome from trying to format or lint files that are marked as generated,
	// which would fail since generated files are typically read-only.
	// Default: true.
	//
	// Experimental.
	IgnoreGeneratedFiles *bool `field:"optional" json:"ignoreGeneratedFiles" yaml:"ignoreGeneratedFiles"`
	// Enable linting with recommended rules.
	// Default: true.
	//
	// Experimental.
	Linter *bool `field:"optional" json:"linter" yaml:"linter"`
	// Should arrays be merged or overwritten when creating Biome configuration.
	//
	// By default arrays are merged and duplicate values are removed.
	// Default: true.
	//
	// Experimental.
	MergeArraysInConfiguration *bool `field:"optional" json:"mergeArraysInConfiguration" yaml:"mergeArraysInConfiguration"`
	// Version of Biome to use.
	// Default: "^2".
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

