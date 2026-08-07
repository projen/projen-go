package biomeconfig


// Experimental.
type OverrideFormatterConfiguration struct {
	// The attribute position style.
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Put the `>` of a multi-line HTML or JSX element at the end of the last line instead of being alone on the next line (does not apply to self closing elements).
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Default: true.
	//
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Controls spaces immediately inside supported delimiters when their content fits on one line.
	//
	// It doesn't add spaces before opening delimiters or inside empty delimiters.
	//
	// The affected delimiters vary by language. If unset, uses the configured formatter setting.
	// Experimental.
	DelimiterSpacing *bool `field:"optional" json:"delimiterSpacing" yaml:"delimiterSpacing"`
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
	// Stores whether formatting should be allowed to proceed if a given file has syntax errors.
	// Experimental.
	FormatWithErrors *bool `field:"optional" json:"formatWithErrors" yaml:"formatWithErrors"`
	// The size of the indentation, 2 by default (deprecated, use `indent-width`).
	// Experimental.
	IndentSize *float64 `field:"optional" json:"indentSize" yaml:"indentSize"`
	// The indent style.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation, 2 by default.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Print trailing commas wherever possible in multi-line comma-separated syntactic structures.
	// Experimental.
	TrailingCommas JsTrailingCommas `field:"optional" json:"trailingCommas" yaml:"trailingCommas"`
	// Whether to add a trailing newline at the end of the file.
	//
	// Setting this option to `false` is **highly discouraged** because it could cause many problems with other tools:
	// - https://thoughtbot.com/blog/no-newline-at-end-of-file
	// - https://callmeryan.medium.com/no-newline-at-end-of-file-navigating-gits-warning-for-android-developers-af14e73dd804
	// - https://unix.stackexchange.com/questions/345548/how-to-cat-files-together-adding-missing-newlines-at-end-of-some-files
	//
	// Disable the option at your own risk.
	//
	// Defaults to true.
	// Default: true.
	//
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
}

