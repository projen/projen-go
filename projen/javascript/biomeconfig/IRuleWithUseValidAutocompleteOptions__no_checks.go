//go:build no_runtime_type_checking

package biomeconfig

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IRuleWithUseValidAutocompleteOptions) validateSetLevelParameters(val *string) error {
	return nil
}

