//go:build no_runtime_type_checking

package biomeconfig

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ISecurity) validateSetNoDangerouslySetInnerHtmlParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ISecurity) validateSetNoDangerouslySetInnerHtmlWithChildrenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ISecurity) validateSetNoGlobalEvalParameters(val interface{}) error {
	return nil
}

