package circleci


// A Workflow is comprised of one or more uniquely named jobs.
//
// Jobs are specified in the jobs map,
// see Sample 2.0 config.yml for two examples of a job map.
// The name of the job is the key in the map, and the value is a map describing the job.
// Each job consists of the job’s name as a key and a map as a value. A name should be case insensitive unique within a current jobs list.
// See: https://circleci.com/docs/2.0/configuration-reference/#job_name
//
// Experimental.
type Job struct {
	// name of dynamic key *.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Experimental.
	Docker *[]*Docker `field:"optional" json:"docker" yaml:"docker"`
	// A map of environment variable names and values.
	// Experimental.
	Environment *map[string]interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Experimental.
	Machine *Machine `field:"optional" json:"machine" yaml:"machine"`
	// Experimental.
	Macos *Macos `field:"optional" json:"macos" yaml:"macos"`
	// Number of parallel instances of this job to run (default: 1).
	// Experimental.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Parameters for making a job explicitly configurable in a workflow.
	// Experimental.
	Parameters *map[string]*PipelineParameter `field:"optional" json:"parameters" yaml:"parameters"`
	// {@link ResourceClass}.
	// Experimental.
	ResourceClass *string `field:"optional" json:"resourceClass" yaml:"resourceClass"`
	// Shell to use for execution command in all steps.
	//
	// Can be overridden by shell in each step.
	// Experimental.
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// no type support here, for syntax {@see https://circleci.com/docs/2.0/configuration-reference/#steps}.
	// Experimental.
	Steps *[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	// In which directory to run the steps.
	//
	// Will be interpreted as an absolute path. Default: `~/project`
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

