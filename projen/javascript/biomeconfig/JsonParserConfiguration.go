package biomeconfig


// Options that change how the JSON parser behaves.
// Experimental.
type JsonParserConfiguration struct {
	// Controls whether comments are allowed in files parsed as JSON.
	//
	// When unset, Biome follows the
	// file type's default behavior. An explicit value overrides that behavior, except that
	// `biome.jsonc` always allows comments.
	// Experimental.
	AllowComments *bool `field:"optional" json:"allowComments" yaml:"allowComments"`
	// Controls whether trailing commas are allowed in files parsed as JSON.
	//
	// When unset, Biome
	// follows the file type's default behavior. An explicit value overrides that behavior, except
	// that `biome.jsonc` always allows trailing commas.
	// Experimental.
	AllowTrailingCommas *bool `field:"optional" json:"allowTrailingCommas" yaml:"allowTrailingCommas"`
}

