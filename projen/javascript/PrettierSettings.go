package javascript


// Options to set in Prettier directly or through overrides.
// See: https://prettier.io/docs/en/options.html
//
// Experimental.
type PrettierSettings struct {
	// Include parentheses around a sole arrow function parameter.
	// Default: ArrowParens.ALWAYS
	//
	// Experimental.
	ArrowParens ArrowParens `field:"optional" json:"arrowParens" yaml:"arrowParens"`
	// Put > of opening tags on the last line instead of on a new line.
	// Default: false.
	//
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Print spaces between brackets.
	// Default: true.
	//
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Print (to stderr) where a cursor at the given position would move to after formatting.
	//
	// This option cannot be used with --range-start and --range-end.
	// Default: -1.
	//
	// Experimental.
	CursorOffset *float64 `field:"optional" json:"cursorOffset" yaml:"cursorOffset"`
	// Control how Prettier formats quoted code embedded in the file.
	// Default: EmbeddedLanguageFormatting.AUTO
	//
	// Experimental.
	EmbeddedLanguageFormatting EmbeddedLanguageFormatting `field:"optional" json:"embeddedLanguageFormatting" yaml:"embeddedLanguageFormatting"`
	// Which end of line characters to apply.
	// Default: EndOfLine.LF
	//
	// Experimental.
	EndOfLine EndOfLine `field:"optional" json:"endOfLine" yaml:"endOfLine"`
	// Specify the input filepath.
	//
	// This will be used to do parser inference.
	// Default: none.
	//
	// Experimental.
	Filepath *string `field:"optional" json:"filepath" yaml:"filepath"`
	// How to handle whitespaces in HTML.
	// Default: HTMLWhitespaceSensitivity.CSS
	//
	// Experimental.
	HtmlWhitespaceSensitivity HTMLWhitespaceSensitivity `field:"optional" json:"htmlWhitespaceSensitivity" yaml:"htmlWhitespaceSensitivity"`
	// Insert.
	// Default: false.
	//
	// Experimental.
	InsertPragma *bool `field:"optional" json:"insertPragma" yaml:"insertPragma"`
	// Use single quotes in JSX.
	// Default: false.
	//
	// Experimental.
	JsxSingleQuote *bool `field:"optional" json:"jsxSingleQuote" yaml:"jsxSingleQuote"`
	// Which parser to use.
	// Default: - Prettier automatically infers the parser from the input file path, so you shouldn’t have to change this setting.
	//
	// Experimental.
	Parser *string `field:"optional" json:"parser" yaml:"parser"`
	// Add a plugin.
	//
	// Multiple plugins can be passed as separate `--plugin`s.
	// Default: [].
	//
	// Experimental.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Custom directory that contains prettier plugins in node_modules subdirectory.
	//
	// Overrides default behavior when plugins are searched relatively to the location of
	// Prettier.
	// Multiple values are accepted.
	// Default: [].
	//
	// Experimental.
	PluginSearchDirs *[]*string `field:"optional" json:"pluginSearchDirs" yaml:"pluginSearchDirs"`
	// The line length where Prettier will try wrap.
	// Default: 80.
	//
	// Experimental.
	PrintWidth *float64 `field:"optional" json:"printWidth" yaml:"printWidth"`
	// How to wrap prose.
	// Default: ProseWrap.PRESERVE
	//
	// Experimental.
	ProseWrap ProseWrap `field:"optional" json:"proseWrap" yaml:"proseWrap"`
	// Change when properties in objects are quoted.
	// Default: QuoteProps.ASNEEDED
	//
	// Experimental.
	QuoteProps QuoteProps `field:"optional" json:"quoteProps" yaml:"quoteProps"`
	// Format code ending at a given character offset (exclusive).
	//
	// The range will extend forwards to the end of the selected statement.
	// This option cannot be used with --cursor-offset.
	// Default: null.
	//
	// Experimental.
	RangeEnd *float64 `field:"optional" json:"rangeEnd" yaml:"rangeEnd"`
	// Format code starting at a given character offset.
	//
	// The range will extend backwards to the start of the first line containing the selected
	// statement.
	// This option cannot be used with --cursor-offset.
	// Default: 0.
	//
	// Experimental.
	RangeStart *float64 `field:"optional" json:"rangeStart" yaml:"rangeStart"`
	// Require either '@prettier' or '@format' to be present in the file's first docblock comment in order for it to be formatted.
	// Default: false.
	//
	// Experimental.
	RequirePragma *bool `field:"optional" json:"requirePragma" yaml:"requirePragma"`
	// Print semicolons.
	// Default: true.
	//
	// Experimental.
	Semi *bool `field:"optional" json:"semi" yaml:"semi"`
	// Use single quotes instead of double quotes.
	// Default: false.
	//
	// Experimental.
	SingleQuote *bool `field:"optional" json:"singleQuote" yaml:"singleQuote"`
	// Number of spaces per indentation level.
	// Default: 2.
	//
	// Experimental.
	TabWidth *float64 `field:"optional" json:"tabWidth" yaml:"tabWidth"`
	// Print trailing commas wherever possible when multi-line.
	// Default: TrailingComma.ES5
	//
	// Experimental.
	TrailingComma TrailingComma `field:"optional" json:"trailingComma" yaml:"trailingComma"`
	// Indent with tabs instead of spaces.
	// Default: false.
	//
	// Experimental.
	UseTabs *bool `field:"optional" json:"useTabs" yaml:"useTabs"`
	// Indent script and style tags in Vue files.
	// Default: false.
	//
	// Experimental.
	VueIndentScriptAndStyle *bool `field:"optional" json:"vueIndentScriptAndStyle" yaml:"vueIndentScriptAndStyle"`
}

