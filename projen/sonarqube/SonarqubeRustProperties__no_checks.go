//go:build no_runtime_type_checking

package sonarqube

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SonarqubeRustProperties) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (s *jsiiProxy_SonarqubeRustProperties) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateSonarqubeRustProperties_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSonarqubeRustProperties_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSonarqubeRustPropertiesParameters(scope constructs.IConstruct, options *SonarqubeRustPropertiesOptions) error {
	return nil
}

