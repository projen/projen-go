package github


// You can use the `allow` option to customize which dependencies are updated.
//
// The allow option supports the following options.
// Experimental.
type DependabotAllow struct {
	// Use to allow updates for dependencies with matching names, optionally using `*` to match zero or more characters.
	//
	// For Java dependencies, the format of the dependency-name attribute is:
	// `groupId:artifactId`, for example: `org.kohsuke:github-api`.
	// Experimental.
	DependencyName *string `field:"required" json:"dependencyName" yaml:"dependencyName"`
}

