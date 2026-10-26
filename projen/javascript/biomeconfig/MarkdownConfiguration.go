package biomeconfig


// Options applied to Markdown files.
// Experimental.
type MarkdownConfiguration struct {
	// Runs the linter and assist actions on the code inside fenced code blocks, frontmatter, and HTML blocks, with the rules and actions of the code's language.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	AnalyzeEmbeds *bool `field:"optional" json:"analyzeEmbeds" yaml:"analyzeEmbeds"`
	// Formatter options.
	// Experimental.
	Formatter *MarkdownFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Linter options.
	// Experimental.
	Linter *MarkdownLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// Parsing options.
	// Experimental.
	Parser *MarkdownParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
}

