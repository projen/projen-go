//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AiInstructionsFile) validateSynthesizeContentParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func validateAiInstructionsFile_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateAiInstructionsFile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AiInstructionsFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AiInstructionsFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAiInstructionsFileParameters(scope constructs.IConstruct, filePath *string, options *FileBaseOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

