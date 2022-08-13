package gitlab


// Describes the conditions for when to run the job.
//
// Defaults to 'on_success'.
// See: https://docs.gitlab.com/ee/ci/yaml/#when
//
// Experimental.
type JobWhen string

const (
	// Experimental.
	JobWhen_ALWAYS JobWhen = "ALWAYS"
	// Experimental.
	JobWhen_DELAYED JobWhen = "DELAYED"
	// Experimental.
	JobWhen_MANUAL JobWhen = "MANUAL"
	// Experimental.
	JobWhen_NEVER JobWhen = "NEVER"
	// Experimental.
	JobWhen_ON_FAILURE JobWhen = "ON_FAILURE"
	// Experimental.
	JobWhen_ON_SUCCESS JobWhen = "ON_SUCCESS"
)

