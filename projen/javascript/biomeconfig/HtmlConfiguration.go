package biomeconfig


// Options applied to HTML files.
// Experimental.
type HtmlConfiguration struct {
	// Experimental.
	Assist *HtmlAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Enables full support for HTML, Vue, Svelte and Astro files.
	// Experimental.
	ExperimentalFullSupportEnabled *bool `field:"optional" json:"experimentalFullSupportEnabled" yaml:"experimentalFullSupportEnabled"`
	// HTML formatter options.
	// Experimental.
	Formatter *HtmlFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// HTML linter options.
	// Experimental.
	Linter *HtmlLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// HTML parsing options.
	// Experimental.
	Parser *HtmlParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
}

