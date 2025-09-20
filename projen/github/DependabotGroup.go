package github


// Defines a single group for dependency updates.
// See: https://docs.github.com/en/code-security/dependabot/working-with-dependabot/dependabot-options-reference#groups--
//
// Experimental.
type DependabotGroup struct {
	// Define a list of strings (with or without wildcards) that will match package names to form this dependency group.
	// Experimental.
	Patterns *[]*string `field:"required" json:"patterns" yaml:"patterns"`
	// Specify which type of update the group applies to.
	// Default: - version updates.
	//
	// Experimental.
	AppliesTo DependabotGroupAppliesTo `field:"optional" json:"appliesTo" yaml:"appliesTo"`
	// Limit the group to a type of dependency.
	// See: https://docs.github.com/en/code-security/dependabot/working-with-dependabot/dependabot-options-reference#dependency-type-groups
	//
	// Default: - all types of dependencies.
	//
	// Experimental.
	DependencyType DependabotGroupDependencyType `field:"optional" json:"dependencyType" yaml:"dependencyType"`
	// Optionally you can use this to exclude certain dependencies from the group.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Limit the group to one or more semantic versioning levels.
	//
	// If specified, must contain at least one element and elements must be unique.
	// See: https://docs.github.com/en/code-security/dependabot/working-with-dependabot/dependabot-options-reference#update-types-groups
	//
	// Default: - all semantic versioning levels.
	//
	// Experimental.
	UpdateTypes *[]DependabotGroupUpdateType `field:"optional" json:"updateTypes" yaml:"updateTypes"`
}

