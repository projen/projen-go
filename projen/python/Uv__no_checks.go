//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_Uv) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (u *jsiiProxy_Uv) validateAddDevDependencyParameters(spec *string) error {
	return nil
}

func validateUv_IsComponentParameters(x interface{}) error {
	return nil
}

func validateUv_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewUvParameters(scope constructs.IConstruct, options *UvOptions) error {
	return nil
}

