package github


// Defines a single group for dependency updates.
// Experimental.
type DependabotGroup struct {
	// Define a list of strings (with or without wildcards) that will match package names to form this dependency group.
	// Experimental.
	Patterns *[]*string `field:"required" json:"patterns" yaml:"patterns"`
	// Optionally you can use this to exclude certain dependencies from the group.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
}

