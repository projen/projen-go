package gitlab


// Cache Definition.
// See: https://docs.gitlab.com/ee/ci/yaml/#cache
//
// Experimental.
type Cache struct {
	// Used the to give each cache a unique identifying key.
	//
	// All jobs that use the same cache key use the same cache.
	// Experimental.
	Key interface{} `field:"optional" json:"key" yaml:"key"`
	// Defines which files or directories to cache.
	// Experimental.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// Defines the upload and download behaviour of the cache.
	// Experimental.
	Policy CachePolicy `field:"optional" json:"policy" yaml:"policy"`
	// If set to true all files that are untracked in your Git repository will be cached.
	// Experimental.
	Untracked *bool `field:"optional" json:"untracked" yaml:"untracked"`
	// Defines when to save the cache, based on the status of the job (Default: Job Success).
	// Experimental.
	When CacheWhen `field:"optional" json:"when" yaml:"when"`
}

