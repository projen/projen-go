//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_JsonFile) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JsonFile) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JsonFile) validateAddToArrayParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JsonFile) validateSynthesizeContentParameters(resolver IResolver) error {
	if resolver == nil {
		return fmt.Errorf("parameter resolver is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JsonFile) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JsonFile) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewJsonFileParameters(project Project, filePath *string, options *JsonFileOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
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

