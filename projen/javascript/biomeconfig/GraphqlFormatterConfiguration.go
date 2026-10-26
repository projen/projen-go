package biomeconfig


// Options that change how the GraphQL formatter behaves.
// Experimental.
type GraphqlFormatterConfiguration struct {
	// Controls spaces inside the braces of single-line GraphQL object values.
	//
	// Biome formats
	// `{name: "Biome"}` when disabled and `{ name: "Biome" }` when enabled. If unset, inherits the
	// global bracket spacing setting.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Enables or disables the formatter for GraphQL.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The indent style applied to GraphQL files.
	//
	// If unset, inherits the global indentation style.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The indentation width applied to GraphQL files.
	//
	// If unset, inherits the global indentation
	// width.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The line ending applied to GraphQL files.
	//
	// If unset, inherits the global line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// The preferred maximum line width for GraphQL files.
	//
	// If unset, inherits the global line
	// width.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Reserved for future formatter support.
	//
	// The current formatter preserves standard
	// double-quoted GraphQL strings and does not use this setting.
	// Experimental.
	QuoteStyle QuoteStyle `field:"optional" json:"quoteStyle" yaml:"quoteStyle"`
	// Whether to add a trailing newline at the end of the file.
	//
	// If unset, inherits the global
	// trailing newline setting.
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
}

