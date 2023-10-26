//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateSampleReadme_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSampleReadme_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSampleReadmeParameters(project Project, props *SampleReadmeProps) error {
	return nil
}

