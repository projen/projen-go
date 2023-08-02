package web


// Experimental.
type NextJsCommonProjectOptions struct {
	// Assets directory.
	// Default: "public".
	//
	// Experimental.
	Assetsdir *string `field:"optional" json:"assetsdir" yaml:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Default: true.
	//
	// Experimental.
	Tailwind *bool `field:"optional" json:"tailwind" yaml:"tailwind"`
}

