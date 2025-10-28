package biomeconfig


// Options that changes how the CSS formatter behaves.
// Experimental.
type CssFormatterConfiguration struct {
	// Control the formatter for CSS (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The indent style applied to CSS (and its super languages) files.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation applied to CSS (and its super languages) files.
	//
	// Default to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending applied to CSS (and its super languages) files.
	//
	// `auto` uses CRLF on Windows and LF on other platforms.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to CSS (and its super languages) files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// The type of quotes used in CSS code.
	//
	// Defaults to double.
	// Default: double.
	//
	// Experimental.
	QuoteStyle QuoteStyle `field:"optional" json:"quoteStyle" yaml:"quoteStyle"`
}

