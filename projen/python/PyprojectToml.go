package python


// Experimental.
type PyprojectToml struct {
	// Experimental.
	BuildSystem *BuildSystem `field:"optional" json:"buildSystem" yaml:"buildSystem"`
	// Named groups of dependencies, similar to `requirements.txt` files, which launchers, IDEs, and other tools can find and identify by name. Each item in `[dependency-groups]` is defined as mapping of group name to list of [dependency specifiers](https://packaging.python.org/en/latest/specifications/dependency-specifiers/).
	// Experimental.
	DependencyGroups *PyprojectTomlDependencyGroups `field:"optional" json:"dependencyGroups" yaml:"dependencyGroups"`
	// There are two kinds of metadata: _static_ and _dynamic_.
	//
	// - Static metadata is listed in the `[project]` table directly and cannot be specified or changed by a tool.
	// - Dynamic metadata key names are listed inside the `dynamic` key and represents metadata that a tool will later provide.
	// Experimental.
	Project *PyprojectTomlProject `field:"optional" json:"project" yaml:"project"`
	// Every tool that is used by the project can have users specify configuration data as long as they use a sub-table within `[tool]`.
	//
	// Generally a project can use the subtable `tool.$NAME` if, and only if, they own the entry for `$NAME` in the Cheeseshop/PyPI.
	// Experimental.
	Tool *PyprojectTomlTool `field:"optional" json:"tool" yaml:"tool"`
}

