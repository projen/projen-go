//go:build no_runtime_type_checking

package web

// Building without runtime type checking enabled, so all the below just return nil

func validateNewReactComponentParameters(project javascript.NodeProject, options *ReactComponentOptions) error {
	return nil
}

