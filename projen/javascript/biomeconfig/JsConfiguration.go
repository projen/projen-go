package biomeconfig


// Options applied to JavaScript, TypeScript, JSX, TSX, and supported languages that embed JavaScript.
//
// Language-specific settings take precedence over corresponding global settings. Global settings
// apply when their language-specific counterparts are omitted, unless stated otherwise.
// Experimental.
type JsConfiguration struct {
	// JavaScript assist options.
	// Experimental.
	Assist *JsAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Enables experimental parsing, formatting, linting, diagnostics, and fixes for CSS and GraphQL snippets embedded in JavaScript and TypeScript template literals.
	//
	// Biome recognizes CSS in `css` and `styled` templates and GraphQL in `gql` and `graphql`
	// templates or calls. Templates containing interpolations are not currently supported.
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	ExperimentalEmbeddedSnippetsEnabled *bool `field:"optional" json:"experimentalEmbeddedSnippetsEnabled" yaml:"experimentalEmbeddedSnippetsEnabled"`
	// JavaScript formatter options.
	// Experimental.
	Formatter *JsFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// A list of additional names that Biome's JavaScript linter treats as predefined global bindings.
	//
	// Use this for globals supplied by a runtime, framework, or external script that are not
	// declared in the source file.
	// Experimental.
	Globals *[]*string `field:"optional" json:"globals" yaml:"globals"`
	// Configures how the analyzer accounts for imports used by JSX.
	//
	// This option does not transform
	// JSX or select a runtime for a build tool. Defaults to `transparent`.
	// Default: transparent`.
	//
	// Experimental.
	JsxRuntime JsxRuntime `field:"optional" json:"jsxRuntime" yaml:"jsxRuntime"`
	// JavaScript linter options.
	// Experimental.
	Linter *JsLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// JavaScript parser options.
	// Experimental.
	Parser *JsParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
	// JavaScript module and dependency resolver options.
	// Experimental.
	Resolver *JsResolverConfiguration `field:"optional" json:"resolver" yaml:"resolver"`
}

