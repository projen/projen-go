//go:build no_runtime_type_checking

package polaris

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolarisJavascriptCoverity) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (p *jsiiProxy_PolarisJavascriptCoverity) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validatePolarisJavascriptCoverity_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePolarisJavascriptCoverity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPolarisJavascriptCoverityParameters(project projen.Project, options *PolarisCoverityJavascriptOptions) error {
	return nil
}

