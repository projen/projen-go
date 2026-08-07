package biomeconfig


// Generic options applied to all files.
// Experimental.
type FormatterConfiguration struct {
	// The attribute position style in HTML-like languages.
	//
	// Defaults to `auto`.
	// Default: auto`.
	//
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Places the `>` of a multiline HTML or JSX element at the end of the last line instead of on the next line.
	//
	// Self-closing elements are unaffected. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Whether to insert spaces inside braces in object literals.
	//
	// Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Controls spaces immediately inside supported delimiters when their content fits on one line.
	//
	// It doesn't add spaces before opening delimiters or inside empty delimiters.
	//
	// The affected delimiters vary by language. Defaults to `false`.
	// Default: false`.
	//
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
	// Defaults to `auto`.
	//
	// When formatting `package.json`, Biome uses `always` unless configured otherwise.
	// Default: auto`.
	//
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// Allows formatting files that contain syntax errors when set to `true`.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	FormatWithErrors *bool `field:"optional" json:"formatWithErrors" yaml:"formatWithErrors"`
	// A list of glob patterns.
	//
	// The formatter will include files/folders that will
	// match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// Uses tabs or spaces for indentation.
	//
	// Defaults to `tab`.
	// Default: tab`.
	//
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The indentation width.
	//
	// Defaults to `2`.
	// Default: 2`.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// Selects the line ending.
	//
	// `auto` uses the platform convention. Defaults to `lf`.
	// Default: lf`.
	//
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// The maximum line width.
	//
	// Defaults to `80`.
	// Default: 80`.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Whether to add a trailing newline at the end of the file.
	//
	// Defaults to `true`; disabling
	// this option can cause compatibility problems with other tools.
	// Default: true`; disabling.
	//
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
	// Uses `.editorconfig` files to configure the formatter. Settings in `biome.json` or `biome.jsonc` override `.editorconfig` settings. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	UseEditorconfig *bool `field:"optional" json:"useEditorconfig" yaml:"useEditorconfig"`
}

