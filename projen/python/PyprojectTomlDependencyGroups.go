package python


// Named groups of dependencies, similar to `requirements.txt` files, which launchers, IDEs, and other tools can find and identify by name. Each item in `[dependency-groups]` is defined as mapping of group name to list of [dependency specifiers](https://packaging.python.org/en/latest/specifications/dependency-specifiers/).
// Experimental.
type PyprojectTomlDependencyGroups struct {
	// Experimental.
	Dev *[]interface{} `field:"optional" json:"dev" yaml:"dev"`
}

