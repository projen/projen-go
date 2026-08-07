package biomeconfig


// Formatting options specific to the JavaScript files.
// Experimental.
type JsFormatterConfiguration struct {
	// Whether to add parentheses around arrow function parameters.
	//
	// Defaults to `always`.
	// Default: always`.
	//
	// Experimental.
	ArrowParentheses ArrowParentheses `field:"optional" json:"arrowParentheses" yaml:"arrowParentheses"`
	// The attribute position style in JSX elements.
	//
	// If unset, inherits the global attribute
	// position setting.
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Whether to hug the closing bracket of multiline HTML/JSX tags to the end of the last line, rather than being alone on the following line.
	//
	// If unset, inherits the global bracket
	// placement setting.
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Whether to insert spaces inside braces in object literals.
	//
	// If unset, inherits the global
	// bracket spacing setting.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Controls spaces immediately inside supported JavaScript and TypeScript delimiters when their content fits on one line.
	//
	// It doesn't add spaces before opening delimiters or inside empty
	// delimiters.
	//
	// It affects parentheses, square brackets, template interpolations, TypeScript angle brackets,
	// JSX expression braces, and logical NOT. In operator chains, only the final operator receives
	// a following space.
	//
	// If unset, inherits the global delimiter spacing setting.
	// Experimental.
	DelimiterSpacing *bool `field:"optional" json:"delimiterSpacing" yaml:"delimiterSpacing"`
	// Controls the formatter for JavaScript and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Uses the same `auto`, `always`, and `never` behavior as the global expansion setting.
	//
	// If unset, inherits the global expansion setting.
	// Experimental.
	Expand Expand `field:"optional" json:"expand" yaml:"expand"`
	// The indent style applied to JavaScript and languages that extend it.
	//
	// If unset, inherits the
	// global indentation style.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The indentation width applied to JavaScript and languages that extend it.
	//
	// If unset,
	// inherits the global indentation width.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of quotes used in JSX.
	//
	// Defaults to `double`.
	// Default: double`.
	//
	// Experimental.
	JsxQuoteStyle QuoteStyle `field:"optional" json:"jsxQuoteStyle" yaml:"jsxQuoteStyle"`
	// The line ending applied to JavaScript and languages that extend it.
	//
	// If unset, inherits the
	// global line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// The maximum line width applied to JavaScript and languages that extend it.
	//
	// If unset,
	// inherits the global line width.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// When breaking binary expressions into multiple lines, whether to break them before or after the binary operator.
	//
	// Defaults to `after`.
	// Default: after`.
	//
	// Experimental.
	OperatorLinebreak OperatorLinebreak `field:"optional" json:"operatorLinebreak" yaml:"operatorLinebreak"`
	// Controls when object properties are quoted.
	//
	// Defaults to `asNeeded` in configuration
	// (`as-needed` on the CLI).
	// Default: asNeeded` in configuration.
	//
	// Experimental.
	QuoteProperties QuoteProperties `field:"optional" json:"quoteProperties" yaml:"quoteProperties"`
	// The type of quotes used in JavaScript code.
	//
	// Defaults to `double`.
	// Default: double`.
	//
	// Experimental.
	QuoteStyle QuoteStyle `field:"optional" json:"quoteStyle" yaml:"quoteStyle"`
	// Prints semicolons after every statement or only where needed to avoid automatic semicolon insertion hazards.
	//
	// Defaults to `always`.
	// Default: always`.
	//
	// Experimental.
	Semicolons Semicolons `field:"optional" json:"semicolons" yaml:"semicolons"`
	// Prints trailing commas wherever possible in multiline comma-separated structures.
	//
	// Defaults
	// to `all`.
	// Default: all`.
	//
	// Experimental.
	TrailingCommas JsTrailingCommas `field:"optional" json:"trailingCommas" yaml:"trailingCommas"`
	// Whether to add a trailing newline at the end of the file.
	//
	// If unset, inherits the global
	// trailing newline setting.
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
}

