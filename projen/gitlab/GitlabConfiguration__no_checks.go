//go:build no_runtime_type_checking

package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitlabConfiguration) validateAddDefaultCachesParameters(caches *[]*Cache) error {
	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddGlobalVariablesParameters(variables *map[string]interface{}) error {
	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddIncludesParameters(includes *[]*Include) error {
	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddJobsParameters(jobs *map[string]*Job) error {
	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateAddServicesParameters(services *[]*Service) error {
	return nil
}

func (g *jsiiProxy_GitlabConfiguration) validateCreateNestedTemplatesParameters(config *map[string]*CiConfigurationOptions) error {
	return nil
}

func validateGitlabConfiguration_IsComponentParameters(x interface{}) error {
	return nil
}

func validateGitlabConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGitlabConfigurationParameters(project projen.Project, options *CiConfigurationOptions) error {
	return nil
}

