//go:build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func validateJunit_IsComponentParameters(x interface{}) error {
	return nil
}

func validateJunit_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewJunitParameters(project projen.Project, options *JunitOptions) error {
	return nil
}

