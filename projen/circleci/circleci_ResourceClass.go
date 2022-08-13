package circleci


// The resource_class feature allows configuring CPU and RAM resources for each job.
//
// Different resource classes are available for different executors, as described in the tables below.
// See: https://circleci.com/docs/2.0/configuration-reference/#resourceclass
//
// Experimental.
type ResourceClass string

const (
	// Experimental.
	ResourceClass_SMALL ResourceClass = "SMALL"
	// Experimental.
	ResourceClass_MEDIUM ResourceClass = "MEDIUM"
	// Experimental.
	ResourceClass_MEDIUM_PLUS ResourceClass = "MEDIUM_PLUS"
	// Experimental.
	ResourceClass_LARGE_X ResourceClass = "LARGE_X"
	// Experimental.
	ResourceClass_LARGE_2X ResourceClass = "LARGE_2X"
	// Experimental.
	ResourceClass_LARGE_2X_PLUS ResourceClass = "LARGE_2X_PLUS"
)

