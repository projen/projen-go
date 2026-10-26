package biomeconfig


// Options that change how the JavaScript parser behaves.
// Experimental.
type JsParserConfiguration struct {
	// Controls whether `.js`, `.mjs`, and `.cjs` files may contain JSX syntax. Disabling this option causes JSX in those files to raise a diagnostic. Defaults to `true`.
	// Default: true`.
	//
	// Experimental.
	JsxEverywhere *bool `field:"optional" json:"jsxEverywhere" yaml:"jsxEverywhere"`
	// Enables parsing decorators on class parameters.
	//
	// This syntax belongs to an old experimental
	// proposal and may change. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	UnsafeParameterDecoratorsEnabled *bool `field:"optional" json:"unsafeParameterDecoratorsEnabled" yaml:"unsafeParameterDecoratorsEnabled"`
}

