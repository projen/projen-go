//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskRuntime) validateRunTaskParameters(name *string, args *[]interface{}) error {
	return nil
}

func (t *jsiiProxy_TaskRuntime) validateTryFindTaskParameters(name *string) error {
	return nil
}

func validateNewTaskRuntimeParameters(workdir *string) error {
	return nil
}

