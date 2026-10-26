package biomeconfig


// Options that change how the CSS formatter behaves.
// Experimental.
type CssFormatterConfiguration struct {
	// Controls spaces inside CSS parentheses and square brackets when their content fits on one line.
	//
	// When enabled, `rgb(0, 0, 0)` becomes `rgb( 0, 0, 0 )` and `[data-attr]` becomes
	// `[ data-attr ]`. Empty delimiters are unchanged.
	//
	// If unset, inherits the global delimiter spacing setting.
	// Experimental.
	DelimiterSpacing *bool `field:"optional" json:"delimiterSpacing" yaml:"delimiterSpacing"`
	// Enables or disables the formatter for CSS and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The indent style applied to CSS and languages that extend it.
	//
	// If unset, inherits the global
	// indentation style.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The indentation width applied to CSS and languages that extend it.
	//
	// If unset, inherits the
	// global indentation width.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The line ending applied to CSS and languages that extend it.
	//
	// If unset, inherits the global
	// line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// The preferred maximum line width for CSS and languages that extend it.
	//
	// If unset, inherits
	// the global line width.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Selects the preferred quote style for CSS strings.
	//
	// Biome may use the alternate quote when
	// that avoids additional escaping. Quotes in `@charset` rules are currently preserved.
	// Defaults to `double`.
	// Default: double`.
	//
	// Experimental.
	QuoteStyle QuoteStyle `field:"optional" json:"quoteStyle" yaml:"quoteStyle"`
	// Whether to add a trailing newline at the end of the file.
	//
	// If unset, inherits the global
	// trailing newline setting.
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
}

