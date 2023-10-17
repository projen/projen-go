package gitlab


// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
// See: https://docs.gitlab.com/ee/ci/yaml/includes.html#use-rules-with-include
//
// Experimental.
type IncludeRule struct {
	// Experimental.
	AllowFailure interface{} `field:"optional" json:"allowFailure" yaml:"allowFailure"`
	// Experimental.
	Changes *[]*string `field:"optional" json:"changes" yaml:"changes"`
	// Experimental.
	Exists *[]*string `field:"optional" json:"exists" yaml:"exists"`
	// Experimental.
	If *string `field:"optional" json:"if" yaml:"if"`
	// Experimental.
	StartIn *string `field:"optional" json:"startIn" yaml:"startIn"`
	// Experimental.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
	// Experimental.
	When JobWhen `field:"optional" json:"when" yaml:"when"`
}

