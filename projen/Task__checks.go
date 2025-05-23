//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_Task) validateBuiltinParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Task) validateEnvParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Task) validateExecParameters(command *string, options *TaskStepOptions) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validateInsertStepParameters(index *float64, steps *[]*TaskStep) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	for idx_b7595e, v := range *steps {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter steps[%#v]", idx_b7595e) }); err != nil {
			return err
		}
	}

	return nil
}

func (t *jsiiProxy_Task) validatePrependParameters(shell *string, options *TaskStepOptions) error {
	if shell == nil {
		return fmt.Errorf("parameter shell is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validatePrependExecParameters(shell *string, options *TaskStepOptions) error {
	if shell == nil {
		return fmt.Errorf("parameter shell is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validatePrependSayParameters(message *string, options *TaskStepOptions) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validatePrependSpawnParameters(subtask Task, options *TaskStepOptions) error {
	if subtask == nil {
		return fmt.Errorf("parameter subtask is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validateRemoveStepParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_Task) validateResetParameters(options *TaskStepOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validateSayParameters(message *string, options *TaskStepOptions) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validateSpawnParameters(subtask Task, options *TaskStepOptions) error {
	if subtask == nil {
		return fmt.Errorf("parameter subtask is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_Task) validateUpdateStepParameters(index *float64, step *TaskStep) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	if step == nil {
		return fmt.Errorf("parameter step is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(step, func() string { return "parameter step" }); err != nil {
		return err
	}

	return nil
}

func validateNewTaskParameters(name *string, props *TaskOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

