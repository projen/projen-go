//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateMergeQueue_IsComponentParameters(x interface{}) error {
	return nil
}

func validateMergeQueue_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMergeQueueParameters(scope constructs.IConstruct, options *MergeQueueOptions) error {
	return nil
}

