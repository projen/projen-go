//go:build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func validateJsiiDocgen_IsComponentParameters(x interface{}) error {
	return nil
}

func validateJsiiDocgen_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewJsiiDocgenParameters(scope constructs.IConstruct, options *JsiiDocgenOptions) error {
	return nil
}

