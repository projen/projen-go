//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (l *jsiiProxy_Logger) validateLogParameters(level LogLevel) error {
	if level == "" {
		return fmt.Errorf("parameter level is required, but nil was provided")
	}

	return nil
}

func validateNewLoggerParameters(project Project, options *LoggerOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

