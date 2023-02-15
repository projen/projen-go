package gitlab


// Used to configure the kubernetes deployment for this environment.
//
// This is currently not
// supported for kubernetes clusters that are managed by Gitlab.
// Experimental.
type KubernetesConfig struct {
	// The kubernetes namespace where this environment should be deployed to.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

