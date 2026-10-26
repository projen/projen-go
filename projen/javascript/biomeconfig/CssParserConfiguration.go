package biomeconfig


// Options that change how the CSS parser behaves.
// Experimental.
type CssParserConfiguration struct {
	// Enables `//` line comments in plain CSS.
	//
	// Standard CSS treats `//` as delimiter characters
	// rather than as a comment, and SCSS accepts `//` comments independently of this option.
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	AllowWrongLineComments *bool `field:"optional" json:"allowWrongLineComments" yaml:"allowWrongLineComments"`
	// Enables CSS Modules-specific syntax such as `:local`, `:global`, `composes`, and `@value`.
	//
	// When unset, Biome enables this syntax automatically for files whose names end in
	// `.module.css` or `.module.scss; otherwise, it defaults to `false`. Enable it explicitly when CSS Module files
	// use another naming convention.
	// Experimental.
	CssModules *bool `field:"optional" json:"cssModules" yaml:"cssModules"`
	// Enables parsing Tailwind CSS 4.0 directives and functions, including `@theme`, `@utility`, `@variant`, `@source`, and `@apply`. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	TailwindDirectives *bool `field:"optional" json:"tailwindDirectives" yaml:"tailwindDirectives"`
}

