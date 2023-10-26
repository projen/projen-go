//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UpgradeDependencies) validateAddPostBuildStepsParameters(steps *[]*workflows.JobStep) error {
	return nil
}

func validateUpgradeDependencies_IsComponentParameters(x interface{}) error {
	return nil
}

func validateUpgradeDependencies_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_UpgradeDependencies) validateSetContainerOptionsParameters(val *workflows.ContainerOptions) error {
	return nil
}

func validateNewUpgradeDependenciesParameters(project NodeProject, options *UpgradeDependenciesOptions) error {
	return nil
}

