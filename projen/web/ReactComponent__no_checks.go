//go:build no_runtime_type_checking

package web

// Building without runtime type checking enabled, so all the below just return nil

func validateReactComponent_IsComponentParameters(x interface{}) error {
	return nil
}

func validateReactComponent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewReactComponentParameters(project javascript.NodeProject, options *ReactComponentOptions) error {
	return nil
}

