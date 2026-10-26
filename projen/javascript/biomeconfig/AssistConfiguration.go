package biomeconfig


// Experimental.
type AssistConfiguration struct {
	// The assist-action configuration.
	// Experimental.
	Actions *Actions `field:"optional" json:"actions" yaml:"actions"`
	// Enables configured assist actions in the CLI and editor integrations.
	//
	// Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of glob patterns selecting files on which assist actions can run.
	//
	// If omitted, all files
	// selected by `files.includes` remain eligible for assist actions. An empty list selects no
	// files. This option can only narrow the files selected by `files.includes`.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

