//go:build no_runtime_type_checking

package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NestedConfiguration) validateAddDefaultCachesParameters(caches *[]*Cache) error {
	return nil
}

func (n *jsiiProxy_NestedConfiguration) validateAddGlobalVariablesParameters(variables *map[string]interface{}) error {
	return nil
}

func (n *jsiiProxy_NestedConfiguration) validateAddIncludesParameters(includes *[]*Include) error {
	return nil
}

func (n *jsiiProxy_NestedConfiguration) validateAddJobsParameters(jobs *map[string]*Job) error {
	return nil
}

func (n *jsiiProxy_NestedConfiguration) validateAddServicesParameters(services *[]*Service) error {
	return nil
}

func validateNestedConfiguration_IsComponentParameters(x interface{}) error {
	return nil
}

func validateNestedConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNestedConfigurationParameters(project projen.Project, parent GitlabConfiguration, name *string, options *CiConfigurationOptions) error {
	return nil
}

