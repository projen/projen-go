package biomeconfig


// Options that change how the HTML formatter behaves.
// Experimental.
type HtmlFormatterConfiguration struct {
	// The attribute position style in HTML elements.
	//
	// If unset, inherits the global attribute
	// position setting.
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Controls the placement of the closing bracket for multiline HTML opening tags.
	//
	// Biome places
	// the bracket at the end of the last attribute line when enabled and on its own line after the
	// last attribute when disabled. This option also affects self-closing HTML elements. If unset,
	// inherits the global `bracketSameLine` setting.
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Enables or disables the formatter for HTML and languages that extend it.
	//
	// The formatter is
	// experimental and disabled by default. Formatting complete HTML, Vue, Svelte, and Astro files
	// requires `html.experimentalFullSupportEnabled` to be `true`.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Controls whether the content of `<script>` and `<style>` tags is indented by one level in HTML, Vue, Svelte, and Astro files.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	IndentScriptAndStyle *bool `field:"optional" json:"indentScriptAndStyle" yaml:"indentScriptAndStyle"`
	// The indent style applied to HTML and languages that extend it.
	//
	// If unset, inherits the global
	// indentation style.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The indentation width applied to HTML and languages that extend it.
	//
	// If unset, inherits the
	// global indentation width.
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The line ending applied to HTML and languages that extend it.
	//
	// If unset, inherits the global
	// line ending.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// The preferred maximum line width for HTML and languages that extend it.
	//
	// If unset, inherits
	// the global line width.
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Controls whether void elements are self-closed.
	//
	// Defaults to `never`.
	// Default: never`.
	//
	// Experimental.
	SelfCloseVoidElements SelfCloseVoidElements `field:"optional" json:"selfCloseVoidElements" yaml:"selfCloseVoidElements"`
	// Whether to add a trailing newline at the end of the file.
	//
	// Unlike other language-specific
	// trailing newline settings, this option defaults to `true` instead of inheriting the global
	// setting.
	// Experimental.
	TrailingNewline *bool `field:"optional" json:"trailingNewline" yaml:"trailingNewline"`
	// Controls how the formatter treats whitespace around text and child elements in HTML, Vue, Svelte, and Astro markup.
	//
	// The `ignore` setting should be used only when whitespace cannot
	// affect rendered output. Defaults to `css`.
	// Default: css`.
	//
	// Experimental.
	WhitespaceSensitivity WhitespaceSensitivity `field:"optional" json:"whitespaceSensitivity" yaml:"whitespaceSensitivity"`
}

