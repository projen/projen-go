package biomeconfig


// Options that changes how the GraphQL formatter behaves.
// Experimental.
type GraphqlFormatterConfiguration struct {
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Default: true.
	//
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Control the formatter for GraphQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The indent style applied to GraphQL files.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation applied to GraphQL files.
	//
	// Default to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending applied to GraphQL files.
	//
	// `auto` uses CRLF on Windows and LF on other platforms.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to GraphQL files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// The type of quotes used in GraphQL code.
	//
	// Defaults to double.
	// Default: double.
	//
	// Experimental.
	QuoteStyle QuoteStyle `field:"optional" json:"quoteStyle" yaml:"quoteStyle"`
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

