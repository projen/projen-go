package javascript

import (
	"github.com/projen/projen-go/projen/javascript/biomeconfig"
)

// Experimental.
type BiomeOptions struct {
	// Full Biome configuration.
	//
	// Note that this configuration dictates the final outcome if value is set.
	//
	// Example:
	//   if linter is disabled on main level, it can be enabled on fullConfiguration.formatter.enabled.
	//
	// Experimental.
	BiomeConfig biomeconfig.IConfiguration `field:"optional" json:"biomeConfig" yaml:"biomeConfig"`
	// Enable code formatter.
	//
	// Replaces mainly Prettier.
	// Default: false.
	//
	// Experimental.
	Formatter *bool `field:"optional" json:"formatter" yaml:"formatter"`
	// Enable linting.
	//
	// Replaces Eslint.
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
	// Enable import sorting/organizing.
	//
	// Replaces mainly Prettier.
	// Default: false.
	//
	// Experimental.
	OrganizeImports *bool `field:"optional" json:"organizeImports" yaml:"organizeImports"`
	// Version of Biome to use.
	// Default: "^1".
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

