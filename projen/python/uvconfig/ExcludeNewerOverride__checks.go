//go:build !no_runtime_type_checking

package uvconfig

import (
	"fmt"
)

func validateExcludeNewerOverride_FromBooleanParameters(value *bool) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateExcludeNewerOverride_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

