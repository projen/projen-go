package biomeconfig


// Experimental.
type Rules struct {
	// Experimental.
	A11Y interface{} `field:"optional" json:"a11Y" yaml:"a11Y"`
	// Experimental.
	Complexity interface{} `field:"optional" json:"complexity" yaml:"complexity"`
	// Experimental.
	Correctness interface{} `field:"optional" json:"correctness" yaml:"correctness"`
	// Experimental.
	Nursery interface{} `field:"optional" json:"nursery" yaml:"nursery"`
	// Experimental.
	Performance interface{} `field:"optional" json:"performance" yaml:"performance"`
	// Selects the baseline set of lint rules.
	//
	// `recommended` enables Biome's recommended
	// non-nursery rules, `all` enables all non-nursery rules, and `none` starts with no
	// rules enabled. Group-level settings and explicit rule settings override this
	// preset. Defaults to `recommended`.
	// Default: recommended`.
	//
	// Experimental.
	Preset PresetConfig `field:"optional" json:"preset" yaml:"preset"`
	// Enables or disables Biome's recommended non-nursery rules. Defaults to `true`.
	//
	// **Deprecated:** This option will be removed in the next major version. Use
	// `linter.rules.preset` instead, or run `biome migrate` to update the configuration.
	// Default: true`.
	//
	// Experimental.
	Recommended *bool `field:"optional" json:"recommended" yaml:"recommended"`
	// Experimental.
	Security interface{} `field:"optional" json:"security" yaml:"security"`
	// Experimental.
	Style interface{} `field:"optional" json:"style" yaml:"style"`
	// Experimental.
	Suspicious interface{} `field:"optional" json:"suspicious" yaml:"suspicious"`
}

