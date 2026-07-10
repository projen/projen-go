//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (p *jsiiProxy_ProjectTree) validatePostProjectCreationParameters(initProject *InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_ProjectTree) validateProjectCreationParameters(initProject *InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func validateProjectTree_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateProjectTree_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ProjectTree) validateSetFileParameters(val JsonFile) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewProjectTreeParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

