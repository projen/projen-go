package gitlab


// Specify what to forward to the downstream pipeline.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerforward
//
// Experimental.
type TriggerForward struct {
	// Forward pipeline variables, such as manual, scheduled or API variables, to the downstream pipeline.
	// Default: false.
	//
	// Experimental.
	PipelineVariables *bool `field:"optional" json:"pipelineVariables" yaml:"pipelineVariables"`
	// Forward the trigger job's variables defined in the YAML to the downstream pipeline.
	// Default: true.
	//
	// Experimental.
	YamlVariables *bool `field:"optional" json:"yamlVariables" yaml:"yamlVariables"`
}

