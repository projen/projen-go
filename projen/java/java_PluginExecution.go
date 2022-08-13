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
}

