package biomeconfig


// Formatting options specific to the JavaScript files.
// Experimental.
type JsFormatterConfiguration struct {
	// Whether to add non-necessary parentheses to arrow functions.
	//
	// Defaults to "always".
	// Default: always".
	//
	// Experimental.
	ArrowParentheses ArrowParentheses `field:"optional" json:"arrowParentheses" yaml:"arrowParentheses"`
	// The attribute position style in JSX elements.
	//
	// Defaults to auto.
	// Default: auto.
	//
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Whether to hug the closing bracket of multiline HTML/JSX tags to the end of the last line, rather than being alone on the following line.
	//
	// Defaults to false.
	// Default: false.
	//
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Default: true.
	//
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Control the formatter for JavaScript (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether to expand arrays and objects on multiple lines.
	//
	// When set to `auto`, object literals are formatted on multiple lines if the first property has a newline, and array literals are formatted on a single line if it fits in the line. When set to `always`, these literals are formatted on multiple lines, regardless of length of the list. When set to `never`, these literals are formatted on a single line if it fits in the line. When formatting `package.json`, Biome will use `always` unless configured otherwise. Defaults to "auto".
	// Default: auto".
	//
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// The indent style applied to JavaScript (and its super languages) files.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation applied to JavaScript (and its super languages) files.
	//
	// Default to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of quotes used in JSX.
	//
	// Defaults to double.
	// Default: double.
	//
	// Experimental.
	JsxQuoteStyle QuoteStyle `field:"optional" json:"jsxQuoteStyle" yaml:"jsxQuoteStyle"`
	// The type of line ending applied to JavaScript (and its super languages) files.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to JavaScript (and its super languages) files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// When breaking binary expressions into multiple lines, whether to break them before or after the binary operator.
	//
	// Defaults to "after".
	// Default: after".
	//
	// Experimental.
	OperatorLinebreak OperatorLinebreak `field:"optional" json:"operatorLinebreak" yaml:"operatorLinebreak"`
	// When properties in objects are quoted.
	//
	// Defaults to asNeeded.
	// Default: asNeeded.
	//
	// Experimental.
	QuoteProperties QuoteProperties `field:"optional" json:"quoteProperties" yaml:"quoteProperties"`
	// The type of quotes used in JavaScript code.
	//
	// Defaults to double.
	// Default: double.
	//
	// Experimental.
	QuoteStyle QuoteStyle `field:"optional" json:"quoteStyle" yaml:"quoteStyle"`
	// Whether the formatter prints semicolons for all statements or only in for statements where it is necessary because of ASI.
	// Experimental.
	Semicolons Semicolons `field:"optional" json:"semicolons" yaml:"semicolons"`
	// Print trailing commas wherever possible in multi-line comma-separated syntactic structures.
	//
	// Defaults to "all".
	// Default: all".
	//
	// Experimental.
	TrailingCommas TrailingCommas `field:"optional" json:"trailingCommas" yaml:"trailingCommas"`
}

