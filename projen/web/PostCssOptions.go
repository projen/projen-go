package web


// Experimental.
type PostCssOptions struct {
	// Default: "postcss.config.json"
	//
	// Experimental.
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Install Tailwind CSS as a PostCSS plugin.
	// Default: true.
	//
	// Experimental.
	Tailwind *bool `field:"optional" json:"tailwind" yaml:"tailwind"`
	// Tailwind CSS options.
	// Experimental.
	TailwindOptions *TailwindConfigOptions `field:"optional" json:"tailwindOptions" yaml:"tailwindOptions"`
}

