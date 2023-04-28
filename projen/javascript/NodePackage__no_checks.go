//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodePackage) validateAddBinParameters(bins *map[string]*string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateAddEngineParameters(engine *string, version *string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateAddFieldParameters(name *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateAddVersionParameters(version *string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateHasScriptParameters(name *string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateRemoveScriptParameters(name *string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateRenderUpgradePackagesCommandParameters(exclude *[]*string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateSetScriptParameters(name *string, command *string) error {
	return nil
}

func (n *jsiiProxy_NodePackage) validateTryResolveDependencyVersionParameters(dependencyName *string) error {
	return nil
}

func validateNewNodePackageParameters(project projen.Project, options *NodePackageOptions) error {
	return nil
}

