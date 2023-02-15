package gitlab


// References a local file or an artifact from another job to define the pipeline configuration.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerinclude
//
// Experimental.
type TriggerInclude struct {
	// Relative path to the generated YAML file which is extracted from the artifacts and used as the configuration for triggering the child pipeline.
	// Experimental.
	Artifact *string `field:"optional" json:"artifact" yaml:"artifact"`
	// Relative path from repository root (`/`) to the pipeline configuration YAML file.
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Job name which generates the artifact.
	// Experimental.
	Job *string `field:"optional" json:"job" yaml:"job"`
	// Relative path from local repository root (`/`) to the local YAML file to define the pipeline configuration.
	// Experimental.
	Local *string `field:"optional" json:"local" yaml:"local"`
	// Path to another private project under the same GitLab instance, like `group/project` or `group/sub-group/project`.
	// Experimental.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Branch/Tag/Commit hash for the target project.
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// Name of the template YAML file to use in the pipeline configuration.
	// Experimental.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

