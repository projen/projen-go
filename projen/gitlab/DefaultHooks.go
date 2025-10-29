package gitlab


// Experimental.
type DefaultHooks struct {
	// Specify a list of commands to execute on the runner before cloning the Git repository and any submodules https://docs.gitlab.com/ci/yaml/#hookspre_get_sources_script.
	// Experimental.
	PreGetSourcesScript *[]*string `field:"optional" json:"preGetSourcesScript" yaml:"preGetSourcesScript"`
}

