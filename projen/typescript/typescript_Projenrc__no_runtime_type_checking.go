//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package typescript

// Building without runtime type checking enabled, so all the below just return nil

func validateNewProjenrcParameters(project TypeScriptProject, options *ProjenrcOptions) error {
	return nil
}

