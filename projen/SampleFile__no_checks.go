//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateSampleFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSampleFile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSampleFileParameters(project Project, filePath *string, options *SampleFileOptions) error {
	return nil
}

