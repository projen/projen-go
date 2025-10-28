package biomeconfig


// Options that changes how the HTML formatter behaves.
// Experimental.
type HtmlFormatterConfiguration struct {
	// The attribute position style in HTML elements.
	//
	// Defaults to auto.
	// Default: auto.
	//
	// Experimental.
	AttributePosition AttributePosition `field:"optional" json:"attributePosition" yaml:"attributePosition"`
	// Whether to hug the closing bracket of multiline HTML tags to the end of the last line, rather than being alone on the following line.
	//
	// Defaults to false.
	// Default: false.
	//
	// Experimental.
	BracketSameLine *bool `field:"optional" json:"bracketSameLine" yaml:"bracketSameLine"`
	// Control the formatter for HTML (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether to indent the `<script>` and `<style>` tags for HTML (and its super languages).
	//
	// Defaults to false.
	// Default: false.
	//
	// Experimental.
	IndentScriptAndStyle *bool `field:"optional" json:"indentScriptAndStyle" yaml:"indentScriptAndStyle"`
	// The indent style applied to HTML (and its super languages) files.
	// Experimental.
	IndentStyle IndentStyle `field:"optional" json:"indentStyle" yaml:"indentStyle"`
	// The size of the indentation applied to HTML (and its super languages) files.
	//
	// Default to 2.
	// Default: 2.
	//
	// Experimental.
	IndentWidth *float64 `field:"optional" json:"indentWidth" yaml:"indentWidth"`
	// The type of line ending applied to HTML (and its super languages) files.
	//
	// `auto` uses CRLF on Windows and LF on other platforms.
	// Experimental.
	LineEnding LineEnding `field:"optional" json:"lineEnding" yaml:"lineEnding"`
	// What's the max width of a line applied to HTML (and its super languages) files.
	//
	// Defaults to 80.
	// Default: 80.
	//
	// Experimental.
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// Whether void elements should be self-closed.
	//
	// Defaults to never.
	// Default: never.
	//
	// Experimental.
	SelfCloseVoidElements SelfCloseVoidElements `field:"optional" json:"selfCloseVoidElements" yaml:"selfCloseVoidElements"`
	// Whether to account for whitespace sensitivity when formatting HTML (and its super languages).
	//
	// Defaults to "css".
	// Default: css".
	//
	// Experimental.
	WhitespaceSensitivity WhitespaceSensitivity `field:"optional" json:"whitespaceSensitivity" yaml:"whitespaceSensitivity"`
}

