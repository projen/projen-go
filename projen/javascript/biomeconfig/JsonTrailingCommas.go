package biomeconfig


// Controls trailing commas in multiline JSON arrays and objects.
//
// Use `all` only for JSON variants that allow trailing commas, and ensure `json.parser.allowTrailingCommas` is enabled or automatically detected for those files.
// Experimental.
type JsonTrailingCommas string

const (
	// none.
	// Experimental.
	JsonTrailingCommas_NONE JsonTrailingCommas = "NONE"
	// all.
	// Experimental.
	JsonTrailingCommas_ALL JsonTrailingCommas = "ALL"
)

