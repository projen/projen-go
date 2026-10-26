package biomeconfig


// Formatting options specific to JavaScript and languages that extend it.
// Experimental.
type JsFormatterConfiguration struct {
	// Controls parentheses around arrow-function parameters.
	//
	// `always` always prints parentheses,
	// while `asNeeded` omits them when an arrow function has one parameter and the syntax permits
	// it. Defaults to `always`.
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
	// Controls the placement of the closing bracket for multiline JSX opening tags.
	//
	// Biome places
	// the bracket at the end of the last attribute line when enabled and on its own line after the
	// last attribute when disabled. Self-closing JSX elements are unaffected. If unset, inherits
	// the global bracket placement setting.
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Controls spaces inside the braces of single-line object literals.
	//
	// Biome formats `{value: 1}`
	// when disabled and `{ value: 1 }` when enabled. If unset, inherits the global bracket spacing
	// setting.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Controls spaces immediately inside supported JavaScript and TypeScript delimiters when their content fits on one line.
	//
	// This affects parentheses, square brackets, template
	// interpolations, TypeScript angle brackets, and JSX expression braces. Empty delimiters are
	// unchanged.
	//
	// It also controls spacing after logical NOT operators. In a chain of logical NOT operators,
	// only the final operator receives a following space. Biome doesn't add spaces before opening
	// delimiters.
	//
	// If unset, inherits the global delimiter spacing setting.
	// Experimental.
	DelimiterSpacing *bool `field:"optional" json:"delimiterSpacing" yaml:"delimiterSpacing"`
	// Enables or disables the formatter for JavaScript and languages that extend it.
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
	// Selects the preferred quote style for JSX attribute values.
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
	// The preferred maximum line width applied to JavaScript and languages that extend it.
	//
	// If
	// unset, inherits the global line width.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// When breaking binary expressions into multiple lines, whether to break them before or after the binary operator.
	//
	// Defaults to `after`.
	// Default: after`.
	//
	// Experimental.
	OperatorLinebreak OperatorLinebreak `field:"optional" json:"operatorLinebreak" yaml:"operatorLinebreak"`
	// Controls whether quotes around object property names are preserved.
	//
	// `asNeeded` removes
	// quotes when the property name is valid without them, while `preserve` keeps quotes around
	// property names that were quoted in the input. Defaults to `asNeeded` in configuration
	// (`as-needed` on the CLI).
	// Default: asNeeded` in configuration.
	//
	// Experimental.
	QuoteProperties QuoteProperties `field:"optional" json:"quoteProperties" yaml:"quoteProperties"`
	// Selects the preferred quote style for JavaScript and TypeScript string literals.
	//
	// Biome may
	// use the alternate quote when that avoids additional escaping. Defaults to `double`.
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
	// Controls trailing commas in multiline comma-separated structures.
	//
	// `all` adds trailing commas
	// wherever syntax permits, including function parameters and calls. `es5` adds them only in
	// constructs supported by ES5, such as array and object literals, and excludes function
	// parameters, function calls, and TypeScript type parameters. `none` removes trailing commas.
	// Defaults to `all`.
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

