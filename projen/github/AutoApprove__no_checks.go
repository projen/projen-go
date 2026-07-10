//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoApprove) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (a *jsiiProxy_AutoApprove) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateAutoApprove_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAutoApprove_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAutoApproveParameters(github GitHub, options *AutoApproveOptions) error {
	return nil
}

