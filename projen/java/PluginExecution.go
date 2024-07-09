package java


// Plugin execution definition.
// Experimental.
type PluginExecution struct {
	// Which Maven goals this plugin should be associated with.
	// Experimental.
	Goals *[]*string `field:"required" json:"goals" yaml:"goals"`
	// The ID.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Execution key/value configuration.
	// Default: {}.
	//
	// Experimental.
	Configuration *map[string]interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The phase in which the plugin should execute.
	// Experimental.
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
}

