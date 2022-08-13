package javascript


// Options for Prettier.
// Experimental.
type PrettierOptions struct {
	// Defines an .prettierIgnore file.
	// Experimental.
	IgnoreFile *bool `field:"optional" json:"ignoreFile" yaml:"ignoreFile"`
	// Provide a list of patterns to override prettier configuration.
	// See: https://prettier.io/docs/en/configuration.html#configuration-overrides
	//
	// Experimental.
	Overrides *[]*PrettierOverride `field:"optional" json:"overrides" yaml:"overrides"`
	// Prettier settings.
	// Experimental.
	Settings *PrettierSettings `field:"optional" json:"settings" yaml:"settings"`
}

