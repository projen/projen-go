//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateEdgeLambdaAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateEdgeLambdaAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEdgeLambdaAutoDiscoverParameters(project projen.Project, options *EdgeLambdaAutoDiscoverOptions) error {
	return nil
}

