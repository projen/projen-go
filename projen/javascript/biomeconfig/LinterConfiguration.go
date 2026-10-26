package biomeconfig


// Experimental.
type LinterConfiguration struct {
	// Configures rules associated with a framework, library, project-wide analysis, or type analysis.
	//
	// Keys are domain names and values are `all`, `recommended`, or `none`.
	// Enabling rules from the `project` or `types` domains activates project and dependency
	// scanning. The `types` domain also enables type-aware resolution.
	// Experimental.
	Domains *RuleDomains `field:"optional" json:"domains" yaml:"domains"`
	// Enables or disables the linter.
	//
	// Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of glob patterns selecting files to lint.
	//
	// If omitted, all files selected by
	// `files.includes` remain eligible for linting. An empty list selects no files. This option can
	// only narrow the files selected by `files.includes`.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// The lint-rule configuration.
	// Experimental.
	Rules *Rules `field:"optional" json:"rules" yaml:"rules"`
}

