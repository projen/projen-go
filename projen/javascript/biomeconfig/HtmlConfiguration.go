package biomeconfig


// Options applied to HTML files.
// Experimental.
type HtmlConfiguration struct {
	// HTML formatter options.
	// Experimental.
	Formatter *HtmlFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// HTML parsing options.
	// Experimental.
	Parser interface{} `field:"optional" json:"parser" yaml:"parser"`
}

