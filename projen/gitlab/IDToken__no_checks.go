//go:build no_runtime_type_checking

package gitlab

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IDToken) validateSetAudParameters(val interface{}) error {
	return nil
}

