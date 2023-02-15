package gitlab


// How many times a job is retried if it fails.
//
// If not defined, defaults to 0 and jobs do not retry.
// See: https://docs.gitlab.com/ee/ci/yaml/#retry
//
// Experimental.
type Retry struct {
	// 0 (default), 1, or 2.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Either a single or array of error types to trigger job retry.
	// Experimental.
	When interface{} `field:"optional" json:"when" yaml:"when"`
}

