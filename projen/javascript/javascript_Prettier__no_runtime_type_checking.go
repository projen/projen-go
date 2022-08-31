//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Prettier) validateAddIgnorePatternParameters(pattern *string) error {
	return nil
}

func (p *jsiiProxy_Prettier) validateAddOverrideParameters(override *PrettierOverride) error {
	return nil
}

func validatePrettier_OfParameters(project projen.Project) error {
	return nil
}

func validateNewPrettierParameters(project NodeProject, options *PrettierOptions) error {
	return nil
}

