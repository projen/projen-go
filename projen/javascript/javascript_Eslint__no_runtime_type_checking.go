//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Eslint) validateAddIgnorePatternParameters(pattern *string) error {
	return nil
}

func (e *jsiiProxy_Eslint) validateAddOverrideParameters(override *EslintOverride) error {
	return nil
}

func (e *jsiiProxy_Eslint) validateAddRulesParameters(rules *map[string]interface{}) error {
	return nil
}

func (e *jsiiProxy_Eslint) validateAllowDevDepsParameters(pattern *string) error {
	return nil
}

func validateEslint_OfParameters(project projen.Project) error {
	return nil
}

func validateNewEslintParameters(project NodeProject, options *EslintOptions) error {
	return nil
}

