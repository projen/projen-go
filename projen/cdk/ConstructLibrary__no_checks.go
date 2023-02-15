//go:build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConstructLibrary) validateAddBinsParameters(bins *map[string]*string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateAddFieldsParameters(fields *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateAddGitIgnoreParameters(pattern *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateAddPackageIgnoreParameters(pattern *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateAddTipParameters(message *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateAnnotateGeneratedParameters(glob *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateHasScriptParameters(name *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateRemoveScriptParameters(name *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateRemoveTaskParameters(name *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateRenderWorkflowSetupParameters(options *javascript.RenderWorkflowSetupOptions) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateRunTaskCommandParameters(task projen.Task) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateSetScriptParameters(name *string, command *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateTryFindFileParameters(filePath *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateTryFindJsonFileParameters(filePath *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateTryFindObjectFileParameters(filePath *string) error {
	return nil
}

func (c *jsiiProxy_ConstructLibrary) validateTryRemoveFileParameters(filePath *string) error {
	return nil
}

func validateNewConstructLibraryParameters(options *ConstructLibraryOptions) error {
	return nil
}

