package gitlab


// The environment that a job deploys to.
// Experimental.
type Environment struct {
	// The name of the environment, e.g. 'qa', 'staging', 'production'.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies what this job will do.
	//
	// 'start' (default) indicates the job will start the deployment. 'prepare' indicates this will not affect the deployment. 'stop' indicates this will stop the deployment.
	// Experimental.
	Action Action `field:"optional" json:"action" yaml:"action"`
	// The amount of time it should take before Gitlab will automatically stop the environment.
	//
	// Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'.
	// Experimental.
	AutoStopIn *string `field:"optional" json:"autoStopIn" yaml:"autoStopIn"`
	// Explicitly specifies the tier of the deployment environment if non-standard environment name is used.
	// Experimental.
	DeploymentTier DeploymentTier `field:"optional" json:"deploymentTier" yaml:"deploymentTier"`
	// Used to configure the kubernetes deployment for this environment.
	//
	// This is currently not supported for kubernetes clusters that are managed by Gitlab.
	// Experimental.
	Kubernetes *KubernetesConfig `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// The name of a job to execute when the environment is about to be stopped.
	// Experimental.
	OnStop *string `field:"optional" json:"onStop" yaml:"onStop"`
	// When set, this will expose buttons in various places for the current environment in Gitlab, that will take you to the defined URL.
	// Experimental.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

