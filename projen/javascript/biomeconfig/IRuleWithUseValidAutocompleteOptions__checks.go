//go:build !no_runtime_type_checking

package biomeconfig

import (
	"fmt"
)

func (j *jsiiProxy_IRuleWithUseValidAutocompleteOptions) validateSetLevelParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

