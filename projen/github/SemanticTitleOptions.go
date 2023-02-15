package github


// Options for linting that PR titles follow Conventional Commits.
// See: https://www.conventionalcommits.org/
//
// Experimental.
type SemanticTitleOptions struct {
	// Configure that a scope must always be provided.
	//
	// e.g. feat(ui), fix(core)
	// Experimental.
	RequireScope *bool `field:"optional" json:"requireScope" yaml:"requireScope"`
	// Configure a list of commit types that are allowed.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

