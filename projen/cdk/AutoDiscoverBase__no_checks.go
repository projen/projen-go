//go:build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func validateAutoDiscoverBase_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAutoDiscoverBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAutoDiscoverBaseParameters(project projen.Project, options *AutoDiscoverBaseOptions) error {
	return nil
}

