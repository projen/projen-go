package biomeconfig


// Options that change how the JavaScript parser behaves.
// Experimental.
type JsParserConfiguration struct {
	// Enables parsing of Grit metavariables.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	GritMetavariables *bool `field:"optional" json:"gritMetavariables" yaml:"gritMetavariables"`
	// When enabled, files such as `.js`, `.mjs`, and `.cjs` may contain JSX syntax. Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	JsxEverywhere *bool `field:"optional" json:"jsxEverywhere" yaml:"jsxEverywhere"`
	// It enables the experimental and unsafe parsing of parameter decorators.
	//
	// These decorators belong to an old proposal, and they are subject to change.
	// Experimental.
	UnsafeParameterDecoratorsEnabled *bool `field:"optional" json:"unsafeParameterDecoratorsEnabled" yaml:"unsafeParameterDecoratorsEnabled"`
}

