package gitlab


// Default settings for the CI Configuration.
//
// Jobs that do not define one or more of the listed keywords use the value defined in the default section.
// See: https://docs.gitlab.com/ee/ci/yaml/#default
//
// Experimental.
type Default struct {
	// Experimental.
	AfterScript *[]*string `field:"optional" json:"afterScript" yaml:"afterScript"`
	// Experimental.
	Artifacts *Artifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Experimental.
	BeforeScript *[]*string `field:"optional" json:"beforeScript" yaml:"beforeScript"`
	// Experimental.
	Cache *[]*Cache `field:"optional" json:"cache" yaml:"cache"`
	// Experimental.
	Image *Image `field:"optional" json:"image" yaml:"image"`
	// Experimental.
	Interruptible *bool `field:"optional" json:"interruptible" yaml:"interruptible"`
	// Experimental.
	Retry *Retry `field:"optional" json:"retry" yaml:"retry"`
	// Experimental.
	Services *[]*Service `field:"optional" json:"services" yaml:"services"`
	// Experimental.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Experimental.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

