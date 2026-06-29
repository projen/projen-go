package javascript


// Snapshot formatting options.
//
// Mirrors the pretty-format options, with the exceptions of
// `compareKeys` and `plugins`.
// See: https://jestjs.io/docs/configuration#snapshotformat-object
//
// Experimental.
type SnapshotFormatOptions struct {
	// Calls `toJSON` on objects that have such a method.
	// Default: - true.
	//
	// Experimental.
	CallToJSON *bool `field:"optional" json:"callToJSON" yaml:"callToJSON"`
	// Escapes special characters in regular expressions.
	// Default: - false.
	//
	// Experimental.
	EscapeRegex *bool `field:"optional" json:"escapeRegex" yaml:"escapeRegex"`
	// Escapes quotes in strings.
	// Default: - false.
	//
	// Experimental.
	EscapeString *bool `field:"optional" json:"escapeString" yaml:"escapeString"`
	// Highlights syntax with colors in terminal (some plugins).
	// Default: - false.
	//
	// Experimental.
	Highlight *bool `field:"optional" json:"highlight" yaml:"highlight"`
	// Spaces of indentation between levels of nesting.
	// Default: - 2.
	//
	// Experimental.
	Indent *float64 `field:"optional" json:"indent" yaml:"indent"`
	// Maximum number of levels to print.
	// Default: - Infinity.
	//
	// Experimental.
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// Maximum number of elements to print at a given level.
	// Default: - Infinity.
	//
	// Experimental.
	MaxWidth *float64 `field:"optional" json:"maxWidth" yaml:"maxWidth"`
	// Prints objects on a single line when `true`.
	// Default: - false.
	//
	// Experimental.
	Min *bool `field:"optional" json:"min" yaml:"min"`
	// Prints the prototype for basic objects and arrays.
	// Default: - false.
	//
	// Experimental.
	PrintBasicPrototype *bool `field:"optional" json:"printBasicPrototype" yaml:"printBasicPrototype"`
	// Prints the name of functions.
	// Default: - true.
	//
	// Experimental.
	PrintFunctionName *bool `field:"optional" json:"printFunctionName" yaml:"printFunctionName"`
}

