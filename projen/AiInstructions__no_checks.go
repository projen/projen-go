//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AiInstructions) validateAddAgentSpecificInstructionsParameters(agent AiAgent) error {
	return nil
}

func validateAiInstructions_BestPracticesParameters(project Project) error {
	return nil
}

func validateAiInstructions_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAiInstructions_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAiInstructions_ProjenParameters(project Project) error {
	return nil
}

func validateNewAiInstructionsParameters(project Project, options *AiInstructionsOptions) error {
	return nil
}

