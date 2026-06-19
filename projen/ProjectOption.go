package projen


// Experimental.
type ProjectOption struct {
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	Type interface{} `field:"required" json:"type" yaml:"type"`
	// Experimental.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Experimental.
	Deprecated *bool `field:"optional" json:"deprecated" yaml:"deprecated"`
	// Experimental.
	Docs *string `field:"optional" json:"docs" yaml:"docs"`
	// Experimental.
	Featured *bool `field:"optional" json:"featured" yaml:"featured"`
	// Experimental.
	Fqn *string `field:"optional" json:"fqn" yaml:"fqn"`
	// Experimental.
	InitialValue *string `field:"optional" json:"initialValue" yaml:"initialValue"`
	// Experimental.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

