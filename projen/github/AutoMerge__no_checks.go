//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoMerge) validateAddConditionsLaterParameters(later IAddConditionsLater) error {
	return nil
}

func validateAutoMerge_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAutoMerge_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAutoMergeParameters(github GitHub, options *AutoMergeOptions) error {
	return nil
}

