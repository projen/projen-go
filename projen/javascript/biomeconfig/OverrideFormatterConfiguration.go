package biomeconfig


// Experimental.
type OverrideFormatterConfiguration struct {
	// The attribute position style in HTML-like languages.
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Controls the placement of the closing bracket for multiline HTML and JSX opening tags.
	//
	// Biome
	// places the bracket at the end of the last attribute line when enabled and on its own line when
	// disabled. This also affects self-closing HTML elements, but self-closing JSX elements are
	// unaffected.
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Controls spaces inside braces in supported single-line structures.
	//
	// The affected structures
	// vary by language.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Controls spaces immediately inside supported delimiters when their content fits on one line.
	//
	// It doesn't add spaces before opening delimiters or inside empty delimiters.
	//
	// The affected delimiters vary by language. If unset, uses the configured formatter setting.
	// Experimental.
	DelimiterSpacing *bool `field:"optional" json:"delimiterSpacing" yaml:"delimiterSpacing"`
	// Enables or disables the formatter for matched files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Controls whether arrays and objects are formatted on one line or multiple lines.
	//
	// `auto` formats objects on multiple lines if the first property has a newline, and arrays on
	// one line if they fit.
	//
	// `always` formats arrays and objects on multiple lines.
	//
	// `never` formats arrays and objects on one line if they fit.
	//
	// If unset, uses the configured formatter setting.
	//
	// When formatting `package.json`, Biome uses `always` unless configured otherwise.
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// Allows formatting matched files that contain syntax errors.
	// Experimental.
	FormatWithErrors *bool `field:"optional" json:"formatWithErrors" yaml:"formatWithErrors"`
	// The indentation width.
	//
	// Deprecated, use `indentWidth` instead.
	// Experimental.
	IndentSize *float64 `field:"optional" json:"indentSize" yaml:"indentSize"`
	// Uses tabs or spaces for indentation.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// Sets the indentation width.
	//
	// With space indentation, this is the number of spaces emitted per
	// indentation level. With tab indentation, Biome emits one tab per level and uses this value as
	// the tab's display width when calculating line length. Accepted values are `0` through `24`.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// Sets the preferred maximum line width used when deciding where to wrap code.
	//
	// Some content,
	// such as long unbreakable strings, may still exceed this width. Accepted values are `1` through
	// `320`.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Prints trailing commas wherever possible in multiline comma-separated structures.
	//
	// This is a
	// legacy override option. Prefer `javascript.formatter.trailingCommas`.
	// Experimental.
	TrailingCommas JsTrailingCommas `field:"optional" json:"trailingCommas" yaml:"trailingCommas"`
	// Whether to add a trailing newline at the end of matched files.
	//
	// Disabling this option can
	// cause compatibility problems with other tools.
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
}

