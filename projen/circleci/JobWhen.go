package circleci


// Specify when to enable or disable the step.
// See: https://circleci.com/docs/2.0/configuration-reference/#steps
//
// Experimental.
type JobWhen string

const (
	// Experimental.
	JobWhen_ALWAYS JobWhen = "ALWAYS"
	// Experimental.
	JobWhen_ON_SUCCESS JobWhen = "ON_SUCCESS"
	// Experimental.
	JobWhen_ON_FAIL JobWhen = "ON_FAIL"
)

