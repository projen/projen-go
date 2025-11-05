//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitAttributesFile) validateAddAttributesParameters(glob *string) error {
	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateAddLfsPatternParameters(glob *string) error {
	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateRemoveAttributesParameters(glob *string) error {
	return nil
}

func (g *jsiiProxy_GitAttributesFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validateGitAttributesFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateGitAttributesFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GitAttributesFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_GitAttributesFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewGitAttributesFileParameters(scope constructs.IConstruct, options *GitAttributesFileOptions) error {
	return nil
}

