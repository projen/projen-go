//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cdk8sDepsPy) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (c *jsiiProxy_Cdk8sDepsPy) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateCdk8sDepsPy_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCdk8sDepsPy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdk8sDepsPyParameters(project projen.Project, options *Cdk8sDepsOptions) error {
	return nil
}

