package biomeconfig


// Options that changes how the HTML parser behaves.
// Experimental.
type HtmlParserConfiguration struct {
	// Enables the parsing of double text expressions such as `{{ expression }}` inside `.html` files.
	// Experimental.
	Interpolation *bool `field:"optional" json:"interpolation" yaml:"interpolation"`
}

