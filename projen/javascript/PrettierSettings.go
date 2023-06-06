package javascript


// Options to set in Prettier directly or through overrides.
// See: https://prettier.io/docs/en/options.html
//
// Experimental.
type PrettierSettings struct {
	// Include parentheses around a sole arrow function parameter.
	// Experimental.
	ArrowParens ArrowParens `field:"optional" json:"arrowParens" yaml:"arrowParens"`
	// Put > of opening tags on the last line instead of on a new line.
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Print spaces between brackets.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Print (to stderr) where a cursor at the given position would move to after formatting.
	//
	// This option cannot be used with --range-start and --range-end.
	// Experimental.
	CursorOffset *float64 `field:"optional" json:"cursorOffset" yaml:"cursorOffset"`
	// Control how Prettier formats quoted code embedded in the file.
	// Experimental.
	EmbeddedLanguageFormatting EmbeddedLanguageFormatting `field:"optional" json:"embeddedLanguageFormatting" yaml:"embeddedLanguageFormatting"`
	// Which end of line characters to apply.
	// Experimental.
	EndOfLine EndOfLine `field:"optional" json:"endOfLine" yaml:"endOfLine"`
	// Specify the input filepath.
	//
	// This will be used to do parser inference.
	// Experimental.
	Filepath *string `field:"optional" json:"filepath" yaml:"filepath"`
	// How to handle whitespaces in HTML.
	// Experimental.
	HtmlWhitespaceSensitivity HTMLWhitespaceSensitivity `field:"optional" json:"htmlWhitespaceSensitivity" yaml:"htmlWhitespaceSensitivity"`
	// Insert.
	// Experimental.
	InsertPragma *bool `field:"optional" json:"insertPragma" yaml:"insertPragma"`
	// Use single quotes in JSX.
	// Experimental.
	JsxSingleQuote *bool `field:"optional" json:"jsxSingleQuote" yaml:"jsxSingleQuote"`
	// Which parser to use.
	// Experimental.
	Parser *string `field:"optional" json:"parser" yaml:"parser"`
	// Add a plugin.
	//
	// Multiple plugins can be passed as separate `--plugin`s.
	// Experimental.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Custom directory that contains prettier plugins in node_modules subdirectory.
	//
	// Overrides default behavior when plugins are searched relatively to the location of
	// Prettier.
	// Multiple values are accepted.
	// Experimental.
	PluginSearchDirs *[]*string `field:"optional" json:"pluginSearchDirs" yaml:"pluginSearchDirs"`
	// The line length where Prettier will try wrap.
	// Experimental.
	PrintWidth *float64 `field:"optional" json:"printWidth" yaml:"printWidth"`
	// How to wrap prose.
	// Experimental.
	ProseWrap ProseWrap `field:"optional" json:"proseWrap" yaml:"proseWrap"`
	// Change when properties in objects are quoted.
	// Experimental.
	QuoteProps QuoteProps `field:"optional" json:"quoteProps" yaml:"quoteProps"`
	// Format code ending at a given character offset (exclusive).
	//
	// The range will extend forwards to the end of the selected statement.
	// This option cannot be used with --cursor-offset.
	// Experimental.
	RangeEnd *float64 `field:"optional" json:"rangeEnd" yaml:"rangeEnd"`
	// Format code starting at a given character offset.
	//
	// The range will extend backwards to the start of the first line containing the selected
	// statement.
	// This option cannot be used with --cursor-offset.
	// Experimental.
	RangeStart *float64 `field:"optional" json:"rangeStart" yaml:"rangeStart"`
	// Require either '@prettier' or '@format' to be present in the file's first docblock comment in order for it to be formatted.
	// Experimental.
	RequirePragma *bool `field:"optional" json:"requirePragma" yaml:"requirePragma"`
	// Print semicolons.
	// Experimental.
	Semi *bool `field:"optional" json:"semi" yaml:"semi"`
	// Use single quotes instead of double quotes.
	// Experimental.
	SingleQuote *bool `field:"optional" json:"singleQuote" yaml:"singleQuote"`
	// Number of spaces per indentation level.
	// Experimental.
	TabWidth *float64 `field:"optional" json:"tabWidth" yaml:"tabWidth"`
	// Print trailing commas wherever possible when multi-line.
	// Experimental.
	TrailingComma TrailingComma `field:"optional" json:"trailingComma" yaml:"trailingComma"`
	// Indent with tabs instead of spaces.
	// Experimental.
	UseTabs *bool `field:"optional" json:"useTabs" yaml:"useTabs"`
	// Indent script and style tags in Vue files.
	// Experimental.
	VueIndentScriptAndStyle *bool `field:"optional" json:"vueIndentScriptAndStyle" yaml:"vueIndentScriptAndStyle"`
}

