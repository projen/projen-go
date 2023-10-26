//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateProjectBuild_IsComponentParameters(x interface{}) error {
	return nil
}

func validateProjectBuild_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewProjectBuildParameters(project Project) error {
	return nil
}

