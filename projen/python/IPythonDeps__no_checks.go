//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPythonDeps) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (i *jsiiProxy_IPythonDeps) validateAddDevDependencyParameters(spec *string) error {
	return nil
}

