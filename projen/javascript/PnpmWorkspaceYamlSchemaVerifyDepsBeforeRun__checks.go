//go:build !no_runtime_type_checking

package javascript

import (
	"fmt"
)

func validatePnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromBooleanParameters(value *bool) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

