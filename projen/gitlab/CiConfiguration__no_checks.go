//go:build no_runtime_type_checking

package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CiConfiguration) validateAddDefaultCachesParameters(caches *[]*Cache) error {
	return nil
}

func (c *jsiiProxy_CiConfiguration) validateAddGlobalVariablesParameters(variables *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CiConfiguration) validateAddIncludesParameters(includes *[]*Include) error {
	return nil
}

func (c *jsiiProxy_CiConfiguration) validateAddJobsParameters(jobs *map[string]*Job) error {
	return nil
}

func (c *jsiiProxy_CiConfiguration) validateAddServicesParameters(services *[]*Service) error {
	return nil
}

func validateCiConfiguration_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCiConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCiConfigurationParameters(project projen.Project, name *string, options *CiConfigurationOptions) error {
	return nil
}

