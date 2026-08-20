//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdktnTasks) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (c *jsiiProxy_CdktnTasks) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateCdktnTasks_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCdktnTasks_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdktnTasksParameters(project projen.Project) error {
	return nil
}

