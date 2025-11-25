package uvconfig


// Experimental.
type DependencyGroupSettings struct {
	// Version of python to require when installing this group.
	// Experimental.
	RequiresPython *string `field:"optional" json:"requiresPython" yaml:"requiresPython"`
}

