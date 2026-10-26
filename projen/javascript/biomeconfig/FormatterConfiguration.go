package biomeconfig


// Configures formatting for selected files.
// Experimental.
type FormatterConfiguration struct {
	// The attribute position style in HTML-like languages.
	//
	// Defaults to `auto`.
	// Default: auto`.
	//
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Controls the placement of the closing bracket for multiline HTML and JSX opening tags.
	//
	// Biome
	// places the bracket at the end of the last attribute line when enabled and on its own line when
	// disabled. This also affects self-closing HTML elements, but self-closing JSX elements are
	// unaffected. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Controls spaces inside braces in supported single-line structures.
	//
	// The affected structures
	// vary by language. Defaults to `true`.
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
	// Enables or disables the formatter.
	//
	// Defaults to `true`.
	// Default: true`.
	//
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
	// A list of glob patterns selecting files to format.
	//
	// If omitted, all files selected by
	// `files.includes` remain eligible for formatting. An empty list selects no files. This option
	// can only narrow the files selected by `files.includes`.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// Uses tabs or spaces for indentation.
	//
	// Defaults to `tab`.
	// Default: tab`.
	//
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// Sets the indentation width.
	//
	// With space indentation, this is the number of spaces emitted per
	// indentation level. With tab indentation, Biome emits one tab per level and uses this value as
	// the tab's display width when calculating line length. Accepted values are `0` through `24`.
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
	// Sets the preferred maximum line width used when deciding where to wrap code.
	//
	// Some content,
	// such as long unbreakable strings, may still exceed this width. Accepted values are `1` through
	// `320`. Defaults to `80`.
	// Default: 80`.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Whether to add a trailing newline at the end of the file.
	//
	// Defaults to `true`. Disabling
	// this option can cause compatibility problems with other tools.
	// Default: true`. Disabling
	//
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
	// Controls whether Biome loads formatter settings from an `.editorconfig` file. Biome reads `indent_style`, `indent_size`, `end_of_line`, and `insert_final_newline`. Settings in `biome.json` or `biome.jsonc` take precedence. Biome loads only one `.editorconfig`. Nested cascading configurations and the EditorConfig `root` property are not supported. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	UseEditorconfig *bool `field:"optional" json:"useEditorconfig" yaml:"useEditorconfig"`
}

