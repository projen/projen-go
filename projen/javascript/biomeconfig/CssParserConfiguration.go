package biomeconfig


// Options that changes how the CSS parser behaves.
// Experimental.
type CssParserConfiguration struct {
	// Allow comments to appear on incorrect lines in `.css` files.
	// Experimental.
	AllowWrongLineComments *bool `field:"optional" json:"allowWrongLineComments" yaml:"allowWrongLineComments"`
	// Enables parsing of CSS Modules specific features.
	// Experimental.
	CssModules *bool `field:"optional" json:"cssModules" yaml:"cssModules"`
}

