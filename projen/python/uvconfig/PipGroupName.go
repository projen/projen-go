package uvconfig


// The pip-compatible variant of a [`GroupName`].
//
// Either <groupname> or <path>:<groupname>.
// If <path> is omitted it defaults to "pyproject.toml".
// Experimental.
type PipGroupName struct {
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

