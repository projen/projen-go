package biomeconfig


// Options applied to HTML and languages that extend it.
//
// Full HTML support and the HTML formatter are experimental. Biome aims to minimize breaking
// changes, but bug fixes and new features may change formatting output or diagnostics.
//
// Language-specific settings take precedence over corresponding global settings. Global settings
// apply when their language-specific counterparts are omitted, unless stated otherwise.
// Experimental.
type HtmlConfiguration struct {
	// HTML assist options.
	// Experimental.
	Assist *HtmlAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Enables Biome's experimental full support for `.html`, `.vue`, `.svelte`, and `.astro` files. In this mode, Biome parses the complete document and can analyze or format its markup and supported embedded languages.
	//
	// When disabled, `.vue`, `.svelte`, and `.astro` files use legacy handling, which extracts their
	// JavaScript or TypeScript portions and leaves the rest unchanged. This option selects how
	// these files are processed. It does not enable the HTML formatter. Set
	// `html.formatter.enabled` to `true` to format complete HTML, Vue, Svelte, and Astro files.
	// Defaults to `false`.
	// Default: false`.
	//
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

