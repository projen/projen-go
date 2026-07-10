//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdkTasks) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (c *jsiiProxy_CdkTasks) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateCdkTasks_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCdkTasks_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdkTasksParameters(project projen.Project) error {
	return nil
}

