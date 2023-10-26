//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateCdkConfig_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCdkConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdkConfigParameters(project projen.Project, options *CdkConfigOptions) error {
	return nil
}

