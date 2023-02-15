//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func validateNewPythonSampleParameters(project projen.Project, options *PythonSampleOptions) error {
	return nil
}

