//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Logger) validateLogParameters(level LogLevel) error {
	return nil
}

func validateNewLoggerParameters(project Project, options *LoggerOptions) error {
	return nil
}

