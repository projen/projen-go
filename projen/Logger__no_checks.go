//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Logger) validateLogParameters(level LogLevel) error {
	return nil
}

func validateLogger_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLogger_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLoggerParameters(scope constructs.IConstruct, options *LoggerOptions) error {
	return nil
}

