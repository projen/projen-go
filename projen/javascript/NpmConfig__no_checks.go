//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NpmConfig) validateAddConfigParameters(name *string, value *string) error {
	return nil
}

func (n *jsiiProxy_NpmConfig) validateAddRegistryParameters(url *string) error {
	return nil
}

func validateNpmConfig_IsComponentParameters(x interface{}) error {
	return nil
}

func validateNpmConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNpmConfigParameters(project NodeProject, options *NpmConfigOptions) error {
	return nil
}

