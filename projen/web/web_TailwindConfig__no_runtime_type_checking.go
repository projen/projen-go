//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package web

// Building without runtime type checking enabled, so all the below just return nil

func validateNewTailwindConfigParameters(project javascript.NodeProject, options *TailwindConfigOptions) error {
	return nil
}

