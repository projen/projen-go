package biomeconfig


// Options that change how the Markdown formatter behaves.
// Experimental.
type MarkdownFormatterConfiguration struct {
	// Control the formatter for Markdown (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Formats the code inside fenced code blocks, frontmatter, and HTML blocks with the formatter of the code's language.
	//
	// Defaults to `true`.
	//
	// A block with a syntax error stays as written, and Biome still formats the rest of the file.
	//
	// When this option is `false`, Biome keeps every block as written.
	// Default: true`.
	//
	// Experimental.
	FormatEmbeds *bool `field:"optional" json:"formatEmbeds" yaml:"formatEmbeds"`
	// The indent style applied to Markdown files.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation applied to Markdown files.
	//
	// Defaults to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending applied to Markdown (and its super languages) files.
	//
	// `auto` uses CRLF on Windows and LF on other platforms.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to Markdown files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Controls whether Biome keeps, adds, or removes line breaks in Markdown paragraphs.
	//
	// Manual line breaks are always kept. In Markdown, a manual line break is created by ending a
	// line with two spaces or a backslash.
	// Experimental.
	ProseWrap ProseWrap `field:"optional" json:"proseWrap" yaml:"proseWrap"`
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

