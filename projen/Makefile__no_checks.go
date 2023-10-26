//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Makefile) validateAddAllParameters(target *string) error {
	return nil
}

func (m *jsiiProxy_Makefile) validateAddRuleParameters(rule *Rule) error {
	return nil
}

func (m *jsiiProxy_Makefile) validateAddRulesParameters(rules *[]*Rule) error {
	return nil
}

func (m *jsiiProxy_Makefile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validateMakefile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateMakefile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Makefile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_Makefile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewMakefileParameters(project Project, filePath *string, options *MakefileOptions) error {
	return nil
}

