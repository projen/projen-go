package circleci


// set an inclusive or exclusive filter.
// See: https://circleci.com/docs/2.0/configuration-reference/#filters
//
// Experimental.
type FilterConfig struct {
	// Either a single branch specifier, or a list of branch specifiers.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// Either a single branch specifier, or a list of branch specifiers.
	// Experimental.
	Only *[]*string `field:"optional" json:"only" yaml:"only"`
}

