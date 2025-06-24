package biomeconfig


// Generic options applied to all files.
// Experimental.
type FormatterConfiguration struct {
	// The attribute position style in HTML-ish languages.
	//
	// Defaults to auto.
	// Default: auto.
	//
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
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether to expand arrays and objects on multiple lines.
	//
	// When set to `auto`, object literals are formatted on multiple lines if the first property has a newline, and array literals are formatted on a single line if it fits in the line. When set to `always`, these literals are formatted on multiple lines, regardless of length of the list. When set to `never`, these literals are formatted on a single line if it fits in the line. When formatting `package.json`, Biome will use `always` unless configured otherwise. Defaults to "auto".
	// Default: auto".
	//
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// Stores whether formatting should be allowed to proceed if a given file has syntax errors.
	// Experimental.
	FormatWithErrors *bool `field:"optional" json:"formatWithErrors" yaml:"formatWithErrors"`
	// A list of glob patterns.
	//
	// The formatter will include files/folders that will match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
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
	// Use any `.editorconfig` files to configure the formatter. Configuration in `biome.json` will override `.editorconfig` configuration.
	//
	// Default: `true`.
	// Experimental.
	UseEditorconfig *bool `field:"optional" json:"useEditorconfig" yaml:"useEditorconfig"`
}

