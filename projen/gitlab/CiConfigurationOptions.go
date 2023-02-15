package gitlab


// Options for `CiConfiguration`.
// Experimental.
type CiConfigurationOptions struct {
	// Default settings for the CI Configuration.
	//
	// Jobs that do not define one or more of the listed keywords use the value defined in the default section.
	// Experimental.
	Default *Default `field:"optional" json:"default" yaml:"default"`
	// An initial set of jobs to add to the configuration.
	// Experimental.
	Jobs *map[string]*Job `field:"optional" json:"jobs" yaml:"jobs"`
	// A special job used to upload static sites to Gitlab pages.
	//
	// Requires a `public/` directory
	// with `artifacts.path` pointing to it.
	// Experimental.
	Pages *Job `field:"optional" json:"pages" yaml:"pages"`
	// Groups jobs into stages.
	//
	// All jobs in one stage must complete before next stage is
	// executed. If no stages are specified. Defaults to ['build', 'test', 'deploy'].
	// Experimental.
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
	// Global variables that are passed to jobs.
	//
	// If the job already has that variable defined, the job-level variable takes precedence.
	// Experimental.
	Variables *map[string]interface{} `field:"optional" json:"variables" yaml:"variables"`
	// Used to control pipeline behavior.
	// Experimental.
	Workflow *Workflow `field:"optional" json:"workflow" yaml:"workflow"`
}

