//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateAutoQueue_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAutoQueue_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAutoQueueParameters(scope constructs.IConstruct, options *AutoQueueOptions) error {
	return nil
}

