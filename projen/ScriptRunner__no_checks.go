//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScriptRunner) validateAttachParameters(scope constructs.IConstruct) error {
	return nil
}

func (s *jsiiProxy_ScriptRunner) validateConfigForParameters(entrypoint *string) error {
	return nil
}

func (s *jsiiProxy_ScriptRunner) validateTryAttachParameters(scope constructs.IConstruct) error {
	return nil
}

func validateScriptRunner_IsComponentParameters(x interface{}) error {
	return nil
}

func validateScriptRunner_IsConstructParameters(x interface{}) error {
	return nil
}

