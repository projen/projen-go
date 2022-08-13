package gitlab


// Controls inheritance of globally-defined defaults and variables.
//
// Boolean values control
// inheritance of all default: or variables: keywords. To inherit only a subset of default:
// or variables: keywords, specify what you wish to inherit. Anything not listed is not
// inherited.
// Experimental.
type Inherit struct {
	// Whether to inherit all globally-defined defaults or not.
	//
	// Or subset of inherited defaults.
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Whether to inherit all globally-defined variables or not.
	//
	// Or subset of inherited variables.
	// Experimental.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

