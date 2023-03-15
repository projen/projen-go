package javascript

import (
	"github.com/projen/projen-go/projen"
)

// Options for Prettier.
// Experimental.
type PrettierOptions struct {
	// Defines an .prettierIgnore file.
	// Experimental.
	IgnoreFile *bool `field:"optional" json:"ignoreFile" yaml:"ignoreFile"`
	// Configuration options for .prettierignore file.
	// Experimental.
	IgnoreFileOptions *projen.IgnoreFileOptions `field:"optional" json:"ignoreFileOptions" yaml:"ignoreFileOptions"`
	// Provide a list of patterns to override prettier configuration.
	// See: https://prettier.io/docs/en/configuration.html#configuration-overrides
	//
	// Experimental.
	Overrides *[]*PrettierOverride `field:"optional" json:"overrides" yaml:"overrides"`
	// Prettier settings.
	// Experimental.
	Settings *PrettierSettings `field:"optional" json:"settings" yaml:"settings"`
	// Write prettier configuration as YAML instead of JSON.
	// Experimental.
	Yaml *bool `field:"optional" json:"yaml" yaml:"yaml"`
}

