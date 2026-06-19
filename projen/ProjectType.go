package projen


// Experimental.
type ProjectType struct {
	// Experimental.
	Docsurl *string `field:"required" json:"docsurl" yaml:"docsurl"`
	// Experimental.
	Fqn *string `field:"required" json:"fqn" yaml:"fqn"`
	// Experimental.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// Experimental.
	Options *[]*ProjectOption `field:"required" json:"options" yaml:"options"`
	// Experimental.
	Pjid *string `field:"required" json:"pjid" yaml:"pjid"`
	// Experimental.
	Typename *string `field:"required" json:"typename" yaml:"typename"`
	// Experimental.
	Docs *string `field:"optional" json:"docs" yaml:"docs"`
}

