package web


// Experimental.
type PostCssOptions struct {
	// Experimental.
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Install Tailwind CSS as a PostCSS plugin.
	// Experimental.
	Tailwind *bool `field:"optional" json:"tailwind" yaml:"tailwind"`
	// Tailwind CSS options.
	// Experimental.
	TailwindOptions *TailwindConfigOptions `field:"optional" json:"tailwindOptions" yaml:"tailwindOptions"`
}

