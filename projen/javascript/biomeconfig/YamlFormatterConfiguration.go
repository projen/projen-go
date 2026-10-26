package biomeconfig


// Options that change how the Yaml formatter behaves.
// Experimental.
type YamlFormatterConfiguration struct {
	// Whether to insert spaces inside non-empty flow mappings (`{ key: value }`) and sequences (`[ item ]`) that fit on one line.
	//
	// If unset, inherits the
	// global bracket spacing setting. When neither is set, mappings have spaces
	// and sequences do not.
	// Experimental.
	BracketSpacing *bool `field:"optional" json:"bracketSpacing" yaml:"bracketSpacing"`
	// Control the formatter for Yaml (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The size of the indentation applied to Yaml files.
	//
	// Defaults to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending applied to Yaml (and its super languages) files.
	//
	// `auto` uses CRLF on Windows and LF on other platforms.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to Yaml files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// The type of quotes used in YAML code.
	//
	// Defaults to `double`.
	// Default: double`.
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

