//go:build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JsiiBuild) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_JsiiBuild) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJsiiBuildParameters(options *JsiiBuildOptions) error {
	return nil
}

