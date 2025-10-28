package biomeconfig


// Experimental.
type JsonFormatterConfiguration struct {
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Default: true.
	//
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Control the formatter for JSON (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether to expand arrays and objects on multiple lines.
	//
	// When set to `auto`, object literals are formatted on multiple lines if the first property has a newline, and array literals are formatted on a single line if it fits in the line. When set to `always`, these literals are formatted on multiple lines, regardless of length of the list. When set to `never`, these literals are formatted on a single line if it fits in the line. When formatting `package.json`, Biome will use `always` unless configured otherwise. Defaults to "auto".
	// Default: auto".
	//
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// The indent style applied to JSON (and its super languages) files.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation applied to JSON (and its super languages) files.
	//
	// Default to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending applied to JSON (and its super languages) files.
	//
	// `auto` uses CRLF on Windows and LF on other platforms.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to JSON (and its super languages) files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Print trailing commas wherever possible in multi-line comma-separated syntactic structures.
	//
	// Defaults to "none".
	// Default: none".
	//
	// Experimental.
	TrailingCommas TrailingCommas2 `field:"optional" json:"trailingCommas" yaml:"trailingCommas"`
}

