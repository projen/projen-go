//go:build no_runtime_type_checking

package sonarqube

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SonarqubeTypescriptProperties) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (s *jsiiProxy_SonarqubeTypescriptProperties) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateSonarqubeTypescriptProperties_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSonarqubeTypescriptProperties_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSonarqubeTypescriptPropertiesParameters(scope constructs.IConstruct, options *SonarqubeTypescriptPropertiesOptions) error {
	return nil
}

