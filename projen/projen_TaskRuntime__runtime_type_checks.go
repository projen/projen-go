//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func (t *jsiiProxy_TaskRuntime) validateRunTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskRuntime) validateTryFindTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewTaskRuntimeParameters(workdir *string) error {
	if workdir == nil {
		return fmt.Errorf("parameter workdir is required, but nil was provided")
	}

	return nil
}

