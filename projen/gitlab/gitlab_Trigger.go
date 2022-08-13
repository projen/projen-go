package gitlab


// Trigger a multi-project or a child pipeline.
//
// Read more:.
// See: https://docs.gitlab.com/ee/ci/yaml/README.html#trigger-syntax-for-child-pipeline
//
// Experimental.
type Trigger struct {
	// The branch name that a downstream pipeline will use.
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// A list of local files or artifacts from other jobs to define the pipeline.
	// Experimental.
	Include *[]*TriggerInclude `field:"optional" json:"include" yaml:"include"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project`.
	// Experimental.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend.
	// Experimental.
	Strategy Strategy `field:"optional" json:"strategy" yaml:"strategy"`
}

