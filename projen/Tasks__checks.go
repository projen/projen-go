//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_Tasks) validateAddEnvironmentParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateAddTaskParameters(name *string, options *TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateRemoveTaskParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Tasks) validateTryFindParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewTasksParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

