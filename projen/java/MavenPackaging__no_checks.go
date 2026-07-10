//go:build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MavenPackaging) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (m *jsiiProxy_MavenPackaging) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateMavenPackaging_IsComponentParameters(x interface{}) error {
	return nil
}

func validateMavenPackaging_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMavenPackagingParameters(project projen.Project, pom Pom, options *MavenPackagingOptions) error {
	return nil
}

