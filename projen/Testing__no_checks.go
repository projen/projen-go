//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateTesting_SynthParameters(project Project, options *SnapshotOptions) error {
	return nil
}

