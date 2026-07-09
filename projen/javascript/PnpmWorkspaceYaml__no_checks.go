//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func validatePnpmWorkspaceYaml_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePnpmWorkspaceYaml_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePnpmWorkspaceYaml_OfParameters(project projen.Project) error {
	return nil
}

func validateNewPnpmWorkspaceYamlParameters(project projen.Project, options *PnpmWorkspaceYamlOptions) error {
	return nil
}

