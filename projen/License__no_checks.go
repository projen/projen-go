//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_License) validateSynthesizeContentParameters(_arg IResolver) error {
	return nil
}

func validateLicense_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLicense_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_License) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_License) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewLicenseParameters(project Project, options *LicenseOptions) error {
	return nil
}

