package java


// Options for Maven plugins.
// Experimental.
type PluginOptions struct {
	// Plugin key/value configuration.
	// Experimental.
	Configuration *map[string]interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// You could configure the dependencies for the plugin.
	//
	// Dependencies are in `<groupId>/<artifactId>@<semver>` format.
	// Experimental.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Plugin executions.
	// Experimental.
	Executions *[]*PluginExecution `field:"optional" json:"executions" yaml:"executions"`
}

