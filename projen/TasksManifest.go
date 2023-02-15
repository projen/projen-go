// CDK for software projects
package projen


// Schema for `tasks.json`.
// Experimental.
type TasksManifest struct {
	// Environment for all tasks.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// All tasks available for this project.
	// Experimental.
	Tasks *map[string]*TaskSpec `field:"optional" json:"tasks" yaml:"tasks"`
}

