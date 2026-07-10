//go:build !no_runtime_type_checking

package vscode

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (v *jsiiProxy_VsCodeRecommendedExtensions) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VsCodeRecommendedExtensions) validateProjectCreationParameters(initProject *projen.InitProject) error {
	if initProject == nil {
		return fmt.Errorf("parameter initProject is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(initProject, func() string { return "parameter initProject" }); err != nil {
		return err
	}

	return nil
}

func validateVsCodeRecommendedExtensions_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateVsCodeRecommendedExtensions_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewVsCodeRecommendedExtensionsParameters(vscode VsCode) error {
	if vscode == nil {
		return fmt.Errorf("parameter vscode is required, but nil was provided")
	}

	return nil
}

