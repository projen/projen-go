package biomeconfig


// A set of options applied to the JavaScript files.
// Experimental.
type JsConfiguration struct {
	// Assist options.
	// Experimental.
	Assist *JsAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Formatting options.
	// Experimental.
	Formatter *JsFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// A list of global bindings that should be ignored by the analyzers.
	//
	// If defined here, they should not emit diagnostics.
	// Experimental.
	Globals *[]*string `field:"optional" json:"globals" yaml:"globals"`
	// Indicates the type of runtime or transformation used for interpreting JSX.
	// Experimental.
	JsxRuntime JsxRuntime `field:"optional" json:"jsxRuntime" yaml:"jsxRuntime"`
	// Linter options.
	// Experimental.
	Linter *JsLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// Parsing options.
	// Experimental.
	Parser *JsParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
}

