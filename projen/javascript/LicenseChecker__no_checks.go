//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LicenseChecker) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (l *jsiiProxy_LicenseChecker) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateLicenseChecker_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLicenseChecker_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLicenseCheckerParameters(scope constructs.Construct, options *LicenseCheckerOptions) error {
	return nil
}

