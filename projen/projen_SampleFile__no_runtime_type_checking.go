//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateNewSampleFileParameters(project Project, filePath *string, options *SampleFileOptions) error {
	return nil
}

