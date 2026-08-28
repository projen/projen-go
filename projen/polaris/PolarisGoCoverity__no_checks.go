//go:build no_runtime_type_checking

package polaris

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolarisGoCoverity) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (p *jsiiProxy_PolarisGoCoverity) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validatePolarisGoCoverity_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePolarisGoCoverity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPolarisGoCoverityParameters(project projen.Project, options *PolarisCoverityGoOptions) error {
	return nil
}

