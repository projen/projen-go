//go:build no_runtime_type_checking

package polaris

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolarisJavaCoverity) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (p *jsiiProxy_PolarisJavaCoverity) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validatePolarisJavaCoverity_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePolarisJavaCoverity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPolarisJavaCoverityParameters(project projen.Project, options *PolarisCoverityJavaOptions) error {
	return nil
}

