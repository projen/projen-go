package biomeconfig


// Experimental.
type JsonFormatterConfiguration struct {
	// Whether to insert spaces inside braces in object literals.
	//
	// If unset, inherits the global
	// bracket spacing setting.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Controls spaces inside JSON square brackets when their content fits on one line.
	//
	// When
	// enabled, `[1, 2, 3]` becomes `[ 1, 2, 3 ]`. Empty brackets are unchanged.
	//
	// If unset, inherits the global delimiter spacing setting.
	// Experimental.
	DelimiterSpacing *bool `field:"optional" json:"delimiterSpacing" yaml:"delimiterSpacing"`
	// Controls the formatter for JSON and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Uses the same `auto`, `always`, and `never` behavior as the global expansion setting.
	//
	// If unset, inherits the global expansion setting.
	//
	// When formatting `package.json`, Biome uses `always` unless configured otherwise.
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// The indent style applied to JSON and languages that extend it.
	//
	// If unset, inherits the global
	// indentation style.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The indentation width applied to JSON and languages that extend it.
	//
	// If unset, inherits the
	// global indentation width.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The line ending applied to JSON and languages that extend it.
	//
	// If unset, inherits the global
	// line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// The maximum line width applied to JSON and languages that extend it.
	//
	// If unset, inherits the
	// global line width.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Prints trailing commas wherever possible in multiline comma-separated structures.
	//
	// Defaults
	// to `none`.
	// Default: none`.
	//
	// Experimental.
	TrailingCommas JsonTrailingCommas `field:"optional" json:"trailingCommas" yaml:"trailingCommas"`
	// Whether to add a trailing newline at the end of the file.
	//
	// If unset, inherits the global
	// trailing newline setting.
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
}

