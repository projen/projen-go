package gitlab


// Describes the conditions for when to pull an image.
// See: https://docs.gitlab.com/ee/ci/yaml/#servicepull_policy
//
// Experimental.
type PullPolicy string

const (
	// Experimental.
	PullPolicy_ALWAYS PullPolicy = "ALWAYS"
	// Experimental.
	PullPolicy_NEVER PullPolicy = "NEVER"
	// Experimental.
	PullPolicy_IF_NOT_PRESENT PullPolicy = "IF_NOT_PRESENT"
)

