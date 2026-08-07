package biomeconfig


// Options that change how the JSON parser behaves.
// Experimental.
type JsonParserConfiguration struct {
	// Allows parsing comments in `.json` files.
	// Experimental.
	AllowComments *bool `field:"optional" json:"allowComments" yaml:"allowComments"`
	// Allows parsing trailing commas in `.json` files.
	// Experimental.
	AllowTrailingCommas *bool `field:"optional" json:"allowTrailingCommas" yaml:"allowTrailingCommas"`
}

