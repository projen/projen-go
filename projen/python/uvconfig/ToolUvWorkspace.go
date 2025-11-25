package uvconfig


// Experimental.
type ToolUvWorkspace struct {
	// Packages to exclude as workspace members. If a package matches both `members` and `exclude`, it will be excluded.
	//
	// Supports both globs and explicit paths.
	//
	// For more information on the glob syntax, refer to the [`glob` documentation](https://docs.rs/glob/latest/glob/struct.Pattern.html).
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Packages to include as workspace members.
	//
	// Supports both globs and explicit paths.
	//
	// For more information on the glob syntax, refer to the [`glob` documentation](https://docs.rs/glob/latest/glob/struct.Pattern.html).
	// Experimental.
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
}

