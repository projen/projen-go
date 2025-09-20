package github


// The semantic versioning levels a group may be limited to.
// Experimental.
type DependabotGroupUpdateType string

const (
	// Include major releases.
	// Experimental.
	DependabotGroupUpdateType_MAJOR DependabotGroupUpdateType = "MAJOR"
	// Include minor releases.
	// Experimental.
	DependabotGroupUpdateType_MINOR DependabotGroupUpdateType = "MINOR"
	// Include patch releases.
	// Experimental.
	DependabotGroupUpdateType_PATCH DependabotGroupUpdateType = "PATCH"
)

