//go:build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func validateMavenSample_IsComponentParameters(x interface{}) error {
	return nil
}

func validateMavenSample_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMavenSampleParameters(project projen.Project, options *MavenSampleOptions) error {
	return nil
}

