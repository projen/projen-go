package biomeconfig


// Options that change how the HTML parser behaves.
// Experimental.
type HtmlParserConfiguration struct {
	// Enables double-curly interpolation expressions such as `{{ expression }}` in `.html` files. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	Interpolation *bool `field:"optional" json:"interpolation" yaml:"interpolation"`
	// Enables parsing Vue syntax (`v-if`, `v-bind`, etc.) in `.html` files. Enabling this option also enables `interpolation` implicitly.
	//
	// Biome will already automatically enable Vue parsing in `.vue` files, so you probably don't need
	// to enable this option. This only affects `.html` files, and does not change how `.vue`, `.svelte`,
	// or `.astro` files are parsed. Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	Vue *bool `field:"optional" json:"vue" yaml:"vue"`
}

