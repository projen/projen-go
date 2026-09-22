//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodeNativeTestRunner) validateAddReporterParameters(name *string) error {
	return nil
}

func (n *jsiiProxy_NodeNativeTestRunner) validateAddTestMatchParameters(pattern *string) error {
	return nil
}

func (n *jsiiProxy_NodeNativeTestRunner) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (n *jsiiProxy_NodeNativeTestRunner) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (n *jsiiProxy_NodeNativeTestRunner) validateRemoveReporterParameters(name *string) error {
	return nil
}

func (n *jsiiProxy_NodeNativeTestRunner) validateRemoveTestMatchParameters(pattern *string) error {
	return nil
}

func validateNodeNativeTestRunner_IsComponentParameters(x interface{}) error {
	return nil
}

func validateNodeNativeTestRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNodeNativeTestRunner_OfParameters(project projen.Project) error {
	return nil
}

func validateNewNodeNativeTestRunnerParameters(scope constructs.IConstruct, options *NodeNativeTestRunnerOptions) error {
	return nil
}

