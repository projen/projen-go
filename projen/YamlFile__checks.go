//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (y *jsiiProxy_YamlFile) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (y *jsiiProxy_YamlFile) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (y *jsiiProxy_YamlFile) validateAddToArrayParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (y *jsiiProxy_YamlFile) validateSynthesizeContentParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func validateYamlFile_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateYamlFile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_YamlFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_YamlFile) validateSetLineWidthParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_YamlFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewYamlFileParameters(scope constructs.IConstruct, filePath *string, options *YamlFileOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

