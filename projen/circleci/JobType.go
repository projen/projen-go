package circleci


// A job may have a type of approval indicating it must be manually approved before downstream jobs may proceed.
// See: https://circleci.com/docs/2.0/configuration-reference/#type
//
// Experimental.
type JobType string

const (
	// Experimental.
	JobType_APPROVAL JobType = "APPROVAL"
)

