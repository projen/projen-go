//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package web

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (r *jsiiProxy_ReactTypeDef) validateSynthesizeContentParameters(_arg projen.IResolver) error {
	if _arg == nil {
		return fmt.Errorf("parameter _arg is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ReactTypeDef) validateSetExecutableParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ReactTypeDef) validateSetReadonlyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewReactTypeDefParameters(project ReactTypeScriptProject, filePath *string, options *ReactTypeDefOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

