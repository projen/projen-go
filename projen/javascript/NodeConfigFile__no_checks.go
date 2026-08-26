//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodeConfigFile) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (n *jsiiProxy_NodeConfigFile) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateNodeConfigFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateNodeConfigFile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNodeConfigFile_OfParameters(project projen.Project) error {
	return nil
}

func validateNewNodeConfigFileParameters(project projen.Project, options *NodeConfigFileOptions) error {
	return nil
}

