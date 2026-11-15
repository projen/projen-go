//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func validatePnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromBooleanParameters(value *bool) error {
	return nil
}

func validatePnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromStringParameters(value *string) error {
	return nil
}

