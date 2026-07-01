package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A script runner that can produce the configuration to execute a file of a particular type.
// Experimental.
type IScriptRunner interface {
	// Produce the configuration to run the given entrypoint.
	// Experimental.
	ConfigFor(entrypoint *string) *RunScriptConfig
}

// The jsii proxy for IScriptRunner
type jsiiProxy_IScriptRunner struct {
	_ byte // padding
}

func (i *jsiiProxy_IScriptRunner) ConfigFor(entrypoint *string) *RunScriptConfig {
	if err := i.validateConfigForParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns *RunScriptConfig

	_jsii_.Invoke(
		i,
		"configFor",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

