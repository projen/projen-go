//go:build no_runtime_type_checking

package web

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReactTypeDef) validateSynthesizeContentParameters(_arg projen.IResolver) error {
	return nil
}

func validateReactTypeDef_IsComponentParameters(x interface{}) error {
	return nil
}

func validateReactTypeDef_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ReactTypeDef) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_ReactTypeDef) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewReactTypeDefParameters(project ReactTypeScriptProject, filePath *string, options *ReactTypeDefOptions) error {
	return nil
}

