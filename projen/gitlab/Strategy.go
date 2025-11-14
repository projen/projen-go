package gitlab


// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend or mirror.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerstrategy
//
// Experimental.
type Strategy string

const (
	// Not recommended, use mirror instead.
	//
	// The trigger job status shows failed, success, or running, depending on the downstream pipeline status.
	// Experimental.
	Strategy_DEPEND Strategy = "DEPEND"
	// Mirrors the status of the downstream pipeline exactly.
	// Experimental.
	Strategy_MIRROR Strategy = "MIRROR"
)

