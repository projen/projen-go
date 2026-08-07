package biomeconfig


// Options applied to GritQL files.
// Experimental.
type GritConfiguration struct {
	// GritQL assist options.
	// Experimental.
	Assist *GritAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// GritQL formatter options.
	// Experimental.
	Formatter *GritFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// GritQL linter options.
	// Experimental.
	Linter *GritLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
}

