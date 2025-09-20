package github


// The type of dependency a group may be limited to.
// Experimental.
type DependabotGroupDependencyType string

const (
	// Include only dependencies in the "Development dependency group".
	// Experimental.
	DependabotGroupDependencyType_DEVELOPMENT DependabotGroupDependencyType = "DEVELOPMENT"
	// Include only dependencies in the "Production dependency group".
	// Experimental.
	DependabotGroupDependencyType_PRODUCTION DependabotGroupDependencyType = "PRODUCTION"
)

