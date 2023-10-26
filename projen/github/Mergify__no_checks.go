//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Mergify) validateAddQueueParameters(queue *MergifyQueue) error {
	return nil
}

func (m *jsiiProxy_Mergify) validateAddRuleParameters(rule *MergifyRule) error {
	return nil
}

func validateMergify_IsComponentParameters(x interface{}) error {
	return nil
}

func validateMergify_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMergifyParameters(github GitHub, options *MergifyOptions) error {
	return nil
}

