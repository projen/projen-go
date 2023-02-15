package gitlab


// Configure when artifacts are uploaded depended on job status.
// See: https://docs.gitlab.com/ee/ci/yaml/#cachewhen
//
// Experimental.
type CacheWhen string

const (
	// Upload artifacts regardless of job status.
	// Experimental.
	CacheWhen_ALWAYS CacheWhen = "ALWAYS"
	// Upload artifacts only when the job fails.
	// Experimental.
	CacheWhen_ON_FAILURE CacheWhen = "ON_FAILURE"
	// Upload artifacts only when the job succeeds (this is the default).
	// Experimental.
	CacheWhen_ON_SUCCESS CacheWhen = "ON_SUCCESS"
)

