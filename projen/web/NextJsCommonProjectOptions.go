package web


// Experimental.
type NextJsCommonProjectOptions struct {
	// Assets directory.
	// Experimental.
	Assetsdir *string `field:"optional" json:"assetsdir" yaml:"assetsdir"`
	// Setup Tailwind CSS as a PostCSS plugin.
	// See: https://tailwindcss.com/docs/installation
	//
	// Experimental.
	Tailwind *bool `field:"optional" json:"tailwind" yaml:"tailwind"`
}

