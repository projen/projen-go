package projen


// Options for logging utilities.
// Experimental.
type LoggerOptions struct {
	// The logging verbosity.
	//
	// The levels available (in increasing verbosity) are
	// OFF, ERROR, WARN, INFO, DEBUG, and VERBOSE.
	// Default: LogLevel.INFO
	//
	// Experimental.
	Level LogLevel `field:"optional" json:"level" yaml:"level"`
	// Include a prefix for all logging messages with the project name.
	// Default: false.
	//
	// Experimental.
	UsePrefix *bool `field:"optional" json:"usePrefix" yaml:"usePrefix"`
}

