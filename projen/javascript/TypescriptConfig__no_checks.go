//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TypescriptConfig) validateAddExcludeParameters(pattern *string) error {
	return nil
}

func (t *jsiiProxy_TypescriptConfig) validateAddIncludeParameters(pattern *string) error {
	return nil
}

func validateNewTypescriptConfigParameters(project projen.Project, options *TypescriptConfigOptions) error {
	return nil
}

