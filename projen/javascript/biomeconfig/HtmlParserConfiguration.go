package biomeconfig


// Options that change how the HTML parser behaves.
// Experimental.
type HtmlParserConfiguration struct {
	// Enables parsing double text expressions such as `{{ expression }}` inside `.html` files.
	// Experimental.
	Interpolation *bool `field:"optional" json:"interpolation" yaml:"interpolation"`
	// Enables parsing Vue syntax (`v-if`, `v-bind`, etc.) in `.html` files. Enabling this option also enables `interpolation` implicitly.
	//
	// Biome will already automatically enable Vue parsing in `.vue` files, so you probably don't need
	// to enable this option. This only affects `.html` files, and does not change how `.vue`, `.svelte`,
	// or `.astro` files are parsed.
	// Experimental.
	Vue *bool `field:"optional" json:"vue" yaml:"vue"`
}

