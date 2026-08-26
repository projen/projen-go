//go:build no_runtime_type_checking

package sonarqube

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SonarqubeJavascriptProperties) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (s *jsiiProxy_SonarqubeJavascriptProperties) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateSonarqubeJavascriptProperties_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSonarqubeJavascriptProperties_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSonarqubeJavascriptPropertiesParameters(scope constructs.IConstruct, options *SonarqubeJavascriptPropertiesOptions) error {
	return nil
}

