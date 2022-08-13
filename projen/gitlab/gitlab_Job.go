package gitlab


// Jobs are the most fundamental element of a .gitlab-ci.yml file.
// See: https://docs.gitlab.com/ee/ci/jobs/
//
// Experimental.
type Job struct {
	// Experimental.
	AfterScript *[]*string `field:"optional" json:"afterScript" yaml:"afterScript"`
	// Whether to allow the pipeline to continue running on job failure (Default: false).
	// Experimental.
	AllowFailure interface{} `field:"optional" json:"allowFailure" yaml:"allowFailure"`
	// Experimental.
	Artifacts *Artifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Experimental.
	BeforeScript *[]*string `field:"optional" json:"beforeScript" yaml:"beforeScript"`
	// Experimental.
	Cache *Cache `field:"optional" json:"cache" yaml:"cache"`
	// Must be a regular expression, optionally but recommended to be quoted, and must be surrounded with '/'.
	//
	// Example: '/Code coverage: \d+\.\d+/'
	// Experimental.
	Coverage *string `field:"optional" json:"coverage" yaml:"coverage"`
	// Specify a list of job names from earlier stages from which artifacts should be loaded.
	//
	// By default, all previous artifacts are passed. Use an empty array to skip downloading artifacts.
	// Experimental.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Used to associate environment metadata with a deploy.
	//
	// Environment can have a name and URL attached to it, and will be displayed under /environments under the project.
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Job will run *except* for when these filtering options match.
	// Experimental.
	Except interface{} `field:"optional" json:"except" yaml:"except"`
	// The name of one or more jobs to inherit configuration from.
	// Experimental.
	Extends *[]*string `field:"optional" json:"extends" yaml:"extends"`
	// Experimental.
	Image *Image `field:"optional" json:"image" yaml:"image"`
	// Controls inheritance of globally-defined defaults and variables.
	//
	// Boolean values control inheritance of all default: or variables: keywords. To inherit only a subset of default: or variables: keywords, specify what you wish to inherit. Anything not listed is not inherited.
	// Experimental.
	Inherit *Inherit `field:"optional" json:"inherit" yaml:"inherit"`
	// Experimental.
	Interruptible *bool `field:"optional" json:"interruptible" yaml:"interruptible"`
	// The list of jobs in previous stages whose sole completion is needed to start the current job.
	// Experimental.
	Needs *[]interface{} `field:"optional" json:"needs" yaml:"needs"`
	// Job will run *only* when these filtering options match.
	// Experimental.
	Only interface{} `field:"optional" json:"only" yaml:"only"`
	// Parallel will split up a single job into several, and provide `CI_NODE_INDEX` and `CI_NODE_TOTAL` environment variables for the running jobs.
	// Experimental.
	Parallel interface{} `field:"optional" json:"parallel" yaml:"parallel"`
	// Indicates that the job creates a Release.
	// Experimental.
	Release *Release `field:"optional" json:"release" yaml:"release"`
	// Limit job concurrency.
	//
	// Can be used to ensure that the Runner will not run certain jobs simultaneously.
	// Experimental.
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// Experimental.
	Retry *Retry `field:"optional" json:"retry" yaml:"retry"`
	// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
	// Experimental.
	Rules *[]*IncludeRule `field:"optional" json:"rules" yaml:"rules"`
	// Shell scripts executed by the Runner.
	//
	// The only required property of jobs. Be careful with special characters (e.g. `:`, `{`, `}`, `&`) and use single or double quotes to avoid issues.
	// Experimental.
	Script *[]*string `field:"optional" json:"script" yaml:"script"`
	// CI/CD secrets.
	// Experimental.
	Secrets *map[string]*map[string]*Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Experimental.
	Services *[]*Service `field:"optional" json:"services" yaml:"services"`
	// Define what stage the job will run in.
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Experimental.
	StartIn *string `field:"optional" json:"startIn" yaml:"startIn"`
	// Experimental.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Experimental.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// Trigger allows you to define downstream pipeline trigger.
	//
	// When a job created from trigger definition is started by GitLab, a downstream pipeline gets created. Read more: https://docs.gitlab.com/ee/ci/yaml/README.html#trigger
	// Experimental.
	Trigger interface{} `field:"optional" json:"trigger" yaml:"trigger"`
	// Configurable values that are passed to the Job.
	// Experimental.
	Variables *map[string]interface{} `field:"optional" json:"variables" yaml:"variables"`
	// Describes the conditions for when to run the job.
	//
	// Defaults to 'on_success'.
	// Experimental.
	When JobWhen `field:"optional" json:"when" yaml:"when"`
}

