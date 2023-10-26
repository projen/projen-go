//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateProjectTree_IsComponentParameters(x interface{}) error {
	return nil
}

func validateProjectTree_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectTree) validateSetFileParameters(val JsonFile) error {
	return nil
}

func validateNewProjectTreeParameters(project Project) error {
	return nil
}

