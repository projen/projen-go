package projen


// Constructor options for the Gitpod component.
//
// By default, Gitpod uses the 'gitpod/workspace-full' docker image.
// See: https://github.com/gitpod-io/workspace-images/blob/master/full/Dockerfile
//
// By default, all tasks will be run in parallel. To run the tasks in sequence,
// create a new task and specify the other tasks as subtasks.
//
// Experimental.
type GitpodOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage DevEnvironmentDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks *[]Task `field:"optional" json:"tasks" yaml:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions *[]*string `field:"optional" json:"vscodeExtensions" yaml:"vscodeExtensions"`
	// Optional Gitpod's Github App integration for prebuilds If this is not set and Gitpod's Github App is installed, then Gitpod will apply these defaults: https://www.gitpod.io/docs/prebuilds/#configure-the-github-app.
	// Default: undefined.
	//
	// Experimental.
	Prebuilds *GitpodPrebuilds `field:"optional" json:"prebuilds" yaml:"prebuilds"`
}

