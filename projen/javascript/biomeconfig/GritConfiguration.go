package biomeconfig


// Options applied to GritQL files.
// Experimental.
type GritConfiguration struct {
	// Assist options.
	// Experimental.
	Assist *GritAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Formatting options.
	// Experimental.
	Formatter *GritFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Formatting options.
	// Experimental.
	Linter *GritLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
}

