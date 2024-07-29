//go:build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pom) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (p *jsiiProxy_Pom) validateAddPluginParameters(spec *string, options *PluginOptions) error {
	return nil
}

func (p *jsiiProxy_Pom) validateAddPluginRepositoryParameters(repository *MavenRepository) error {
	return nil
}

func (p *jsiiProxy_Pom) validateAddPropertyParameters(key *string, value *string) error {
	return nil
}

func (p *jsiiProxy_Pom) validateAddRepositoryParameters(repository *MavenRepository) error {
	return nil
}

func (p *jsiiProxy_Pom) validateAddTestDependencyParameters(spec *string) error {
	return nil
}

func validatePom_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePom_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPomParameters(project projen.Project, options *PomOptions) error {
	return nil
}

