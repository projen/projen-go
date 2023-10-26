//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Dependencies) validateAddDependencyParameters(spec *string, type_ DependencyType) error {
	return nil
}

func (d *jsiiProxy_Dependencies) validateGetDependencyParameters(name *string) error {
	return nil
}

func (d *jsiiProxy_Dependencies) validateRemoveDependencyParameters(name *string) error {
	return nil
}

func (d *jsiiProxy_Dependencies) validateTryGetDependencyParameters(name *string) error {
	return nil
}

func validateDependencies_IsComponentParameters(x interface{}) error {
	return nil
}

func validateDependencies_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDependencies_ParseDependencyParameters(spec *string) error {
	return nil
}

func validateNewDependenciesParameters(project Project) error {
	return nil
}

