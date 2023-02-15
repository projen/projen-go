package gitlab


// Describes the conditions for when to run the job.
//
// Defaults to 'on_success'.
// The value can only be 'always' or 'never' when used with workflow.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflowrules
//
// Experimental.
type WorkflowWhen string

const (
	// Experimental.
	WorkflowWhen_ALWAYS WorkflowWhen = "ALWAYS"
	// Experimental.
	WorkflowWhen_NEVER WorkflowWhen = "NEVER"
)

