//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (r *jsiiProxy_Renovatebot) validatePostProjectCreationParameters(initProject *InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_Renovatebot) validateProjectCreationParameters(initProject *InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func validateRenovatebot_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateRenovatebot_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewRenovatebotParameters(project Project, options *RenovatebotOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

