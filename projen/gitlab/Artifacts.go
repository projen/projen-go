package gitlab


// Used to specify a list of files and directories that should be attached to the job if it succeeds.
//
// Artifacts are sent to Gitlab where they can be downloaded.
// See: https://docs.gitlab.com/ee/ci/yaml/#artifacts
//
// Experimental.
type Artifacts struct {
	// A list of paths to files/folders that should be excluded in the artifact.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// How long artifacts should be kept.
	//
	// They are saved 30 days by default. Artifacts that have expired are removed periodically via cron job. Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'.
	// Experimental.
	ExpireIn *string `field:"optional" json:"expireIn" yaml:"expireIn"`
	// Can be used to expose job artifacts in the merge request UI.
	//
	// GitLab will add a link <expose_as> to the relevant merge request that points to the artifact.
	// Experimental.
	ExposeAs *string `field:"optional" json:"exposeAs" yaml:"exposeAs"`
	// Name for the archive created on job success.
	//
	// Can use variables in the name, e.g. '$CI_JOB_NAME'
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of paths to files/folders that should be included in the artifact.
	// Experimental.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests.
	// Experimental.
	Reports *Reports `field:"optional" json:"reports" yaml:"reports"`
	// Whether to add all untracked files (along with 'artifacts.paths') to the artifact.
	// Experimental.
	Untracked *bool `field:"optional" json:"untracked" yaml:"untracked"`
	// Configure when artifacts are uploaded depended on job status.
	// Experimental.
	When CacheWhen `field:"optional" json:"when" yaml:"when"`
}

