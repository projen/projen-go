//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateDependencyReview_IsComponentParameters(x interface{}) error {
	return nil
}

func validateDependencyReview_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDependencyReviewParameters(github GitHub, options *DependencyReviewOptions) error {
	return nil
}

