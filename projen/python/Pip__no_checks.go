//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pip) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (p *jsiiProxy_Pip) validateAddDevDependencyParameters(spec *string) error {
	return nil
}

func validatePip_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePip_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPipParameters(project projen.Project, _options *PipOptions) error {
	return nil
}

