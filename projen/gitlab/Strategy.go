package gitlab


// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerstrategy
//
// Experimental.
type Strategy string

const (
	// Experimental.
	Strategy_DEPEND Strategy = "DEPEND"
)

