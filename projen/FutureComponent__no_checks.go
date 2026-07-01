//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FutureComponent) validateAttachParameters(scope constructs.IConstruct) error {
	return nil
}

func (f *jsiiProxy_FutureComponent) validateTryAttachParameters(scope constructs.IConstruct) error {
	return nil
}

func validateFutureComponent_IsComponentParameters(x interface{}) error {
	return nil
}

func validateFutureComponent_IsConstructParameters(x interface{}) error {
	return nil
}

