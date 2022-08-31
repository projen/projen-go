//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package web

// Building without runtime type checking enabled, so all the below just return nil

func validateNewNextComponentParameters(project javascript.NodeProject, options *NextComponentOptions) error {
	return nil
}

