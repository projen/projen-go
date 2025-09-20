package github


// The type of update a group applies to.
// Experimental.
type DependabotGroupAppliesTo string

const (
	// Apply only to version updates.
	// Experimental.
	DependabotGroupAppliesTo_VERSION_UPDATES DependabotGroupAppliesTo = "VERSION_UPDATES"
	// Apply only to security updates.
	// Experimental.
	DependabotGroupAppliesTo_SECURITY_UPDATES DependabotGroupAppliesTo = "SECURITY_UPDATES"
)

