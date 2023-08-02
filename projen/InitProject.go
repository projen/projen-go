package projen


// Information passed from `projen new` to the project object when the project is first created.
//
// It is used to generate projenrc files in various languages.
// Experimental.
type InitProject struct {
	// Initial arguments passed to `projen new`.
	// Experimental.
	Args *map[string]interface{} `field:"required" json:"args" yaml:"args"`
	// Include commented out options.
	//
	// Does not apply to projenrc.json files.
	// Default: InitProjectOptionHints.FEATURED
	//
	// Experimental.
	Comments InitProjectOptionHints `field:"required" json:"comments" yaml:"comments"`
	// The JSII FQN of the project type.
	// Experimental.
	Fqn *string `field:"required" json:"fqn" yaml:"fqn"`
	// Project metadata.
	// Experimental.
	Type ProjectType `field:"required" json:"type" yaml:"type"`
}

