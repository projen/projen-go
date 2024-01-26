//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Poetry) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (p *jsiiProxy_Poetry) validateAddDevDependencyParameters(spec *string) error {
	return nil
}

func validatePoetry_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePoetry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPoetryParameters(project projen.Project, options *PoetryOptions) error {
	return nil
}

