//go:build no_runtime_type_checking

package circleci

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Circleci) validateAddOrbParameters(name *string, orb *string) error {
	return nil
}

func (c *jsiiProxy_Circleci) validateAddWorkflowParameters(workflow *Workflow) error {
	return nil
}

func validateNewCircleciParameters(project projen.Project, options *CircleCiProps) error {
	return nil
}

