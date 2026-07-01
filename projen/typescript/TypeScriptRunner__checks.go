//go:build !no_runtime_type_checking

package typescript

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TypeScriptRunner) validateAttachParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TypeScriptRunner) validateConfigForParameters(entrypoint *string) error {
	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TypeScriptRunner) validateCopyParameters(overrides *TypeScriptRunnerOptions) error {
	if err := _jsii_.ValidateStruct(overrides, func() string { return "parameter overrides" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TypeScriptRunner) validateTryAttachParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateTypeScriptRunner_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTypeScriptRunner_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTypeScriptRunner_NodejsParameters(options *NodeRunnerOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateTypeScriptRunner_TsNodeParameters(options *TsNodeRunnerOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateTypeScriptRunner_TsxParameters(options *TsxRunnerOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

