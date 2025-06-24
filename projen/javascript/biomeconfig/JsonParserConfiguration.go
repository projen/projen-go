package biomeconfig


// Options that changes how the JSON parser behaves.
// Experimental.
type JsonParserConfiguration struct {
	// Allow parsing comments in `.json` files.
	// Experimental.
	AllowComments *bool `field:"optional" json:"allowComments" yaml:"allowComments"`
	// Allow parsing trailing commas in `.json` files.
	// Experimental.
	AllowTrailingCommas *bool `field:"optional" json:"allowTrailingCommas" yaml:"allowTrailingCommas"`
}

