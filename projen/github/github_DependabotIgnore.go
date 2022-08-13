package github


// You can use the `ignore` option to customize which dependencies are updated.
//
// The ignore option supports the following options.
// Experimental.
type DependabotIgnore struct {
	// Use to ignore updates for dependencies with matching names, optionally using `*` to match zero or more characters.
	//
	// For Java dependencies, the format of the dependency-name attribute is:
	// `groupId:artifactId`, for example: `org.kohsuke:github-api`.
	// Experimental.
	DependencyName *string `field:"required" json:"dependencyName" yaml:"dependencyName"`
	// Use to ignore specific versions or ranges of versions.
	//
	// If you want to
	// define a range, use the standard pattern for the package manager (for
	// example: `^1.0.0` for npm, or `~> 2.0` for Bundler).
	// Experimental.
	Versions *[]*string `field:"optional" json:"versions" yaml:"versions"`
}

