package javascript


// https://yarnpkg.com/configuration/yarnrc#logFilters.
// Experimental.
type YarnLogFilter struct {
	// Experimental.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Experimental.
	Level YarnLogFilterLevel `field:"optional" json:"level" yaml:"level"`
	// Experimental.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Experimental.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

