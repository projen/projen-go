//go:build no_runtime_type_checking

package typescript

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TypeScriptRunner) validateAttachParameters(scope constructs.IConstruct) error {
	return nil
}

func (t *jsiiProxy_TypeScriptRunner) validateConfigForParameters(entrypoint *string) error {
	return nil
}

func (t *jsiiProxy_TypeScriptRunner) validateCopyParameters(overrides *TypeScriptRunnerOptions) error {
	return nil
}

func (t *jsiiProxy_TypeScriptRunner) validateTryAttachParameters(scope constructs.IConstruct) error {
	return nil
}

func validateTypeScriptRunner_IsComponentParameters(x interface{}) error {
	return nil
}

func validateTypeScriptRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTypeScriptRunner_NodejsParameters(options *NodeRunnerOptions) error {
	return nil
}

func validateTypeScriptRunner_TsNodeParameters(options *TsNodeRunnerOptions) error {
	return nil
}

func validateTypeScriptRunner_TsxParameters(options *TsxRunnerOptions) error {
	return nil
}

