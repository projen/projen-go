package web


// Experimental.
type NextComponentOptions struct {
	// Setup Tailwind as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Default: true.
	//
	// Experimental.
	Tailwind *bool `field:"optional" json:"tailwind" yaml:"tailwind"`
	// Whether to apply options specific for TypeScript Next.js projects.
	// Default: false.
	//
	// Experimental.
	Typescript *bool `field:"optional" json:"typescript" yaml:"typescript"`
}

