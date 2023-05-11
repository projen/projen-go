package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Project-level logging utilities.
// Experimental.
type Logger interface {
	Component
	// Experimental.
	Project() Project
	// Log a message to stderr with DEBUG severity.
	// Experimental.
	Debug(text ...interface{})
	// Log a message to stderr with ERROR severity.
	// Experimental.
	Error(text ...interface{})
	// Log a message to stderr with INFO severity.
	// Experimental.
	Info(text ...interface{})
	// Log a message to stderr with a given logging level.
	//
	// The message will be
	// printed as long as `logger.level` is set to the message's severity or higher.
	// Experimental.
	Log(level LogLevel, text ...interface{})
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Log a message to stderr with VERBOSE severity.
	// Experimental.
	Verbose(text ...interface{})
	// Log a message to stderr with WARN severity.
	// Experimental.
	Warn(text ...interface{})
}

// The jsii proxy struct for Logger
type jsiiProxy_Logger struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Logger) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogger(project Project, options *LoggerOptions) Logger {
	_init_.Initialize()

	if err := validateNewLoggerParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Logger{}

	_jsii_.Create(
		"projen.Logger",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLogger_Override(l Logger, project Project, options *LoggerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Logger",
		[]interface{}{project, options},
		l,
	)
}

func (l *jsiiProxy_Logger) Debug(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"debug",
		args,
	)
}

func (l *jsiiProxy_Logger) Error(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"error",
		args,
	)
}

func (l *jsiiProxy_Logger) Info(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"info",
		args,
	)
}

func (l *jsiiProxy_Logger) Log(level LogLevel, text ...interface{}) {
	if err := l.validateLogParameters(level); err != nil {
		panic(err)
	}
	args := []interface{}{level}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"log",
		args,
	)
}

func (l *jsiiProxy_Logger) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Logger) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Logger) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Logger) Verbose(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"verbose",
		args,
	)
}

func (l *jsiiProxy_Logger) Warn(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"warn",
		args,
	)
}

