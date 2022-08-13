package gitlab


// Configure the upload and download behaviour of a cache.
// See: https://docs.gitlab.com/ee/ci/yaml/#cachepolicy
//
// Experimental.
type CachePolicy string

const (
	// Only download the cache when the job starts, but never upload changes when the job finishes.
	// Experimental.
	CachePolicy_PULL CachePolicy = "PULL"
	// Only upload a cache when the job finishes, but never download the cache when the job starts.
	// Experimental.
	CachePolicy_PUSH CachePolicy = "PUSH"
	// The job downloads the cache when the job starts, and uploads changes to the cache when the job ends.
	// Experimental.
	CachePolicy_PULL_PUSH CachePolicy = "PULL_PUSH"
)

