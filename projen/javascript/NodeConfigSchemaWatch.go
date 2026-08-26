package javascript


// Experimental.
type NodeConfigSchemaWatch struct {
	// run in watch mode.
	// Experimental.
	Watch *bool `field:"optional" json:"watch" yaml:"watch"`
	// kill signal to send to the process on watch mode restarts(default: SIGTERM).
	// Experimental.
	WatchKillSignal *string `field:"optional" json:"watchKillSignal" yaml:"watchKillSignal"`
	// path to watch.
	// Experimental.
	WatchPath *[]*string `field:"optional" json:"watchPath" yaml:"watchPath"`
	// preserve outputs on watch mode restart.
	// Experimental.
	WatchPreserveOutput *bool `field:"optional" json:"watchPreserveOutput" yaml:"watchPreserveOutput"`
}

