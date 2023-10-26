//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_Mergify) validateAddQueueParameters(queue *MergifyQueue) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(queue, func() string { return "parameter queue" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_Mergify) validateAddRuleParameters(rule *MergifyRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func validateMergify_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateMergify_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewMergifyParameters(github GitHub, options *MergifyOptions) error {
	if github == nil {
		return fmt.Errorf("parameter github is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

