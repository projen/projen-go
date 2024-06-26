//go:build !no_runtime_type_checking

package vscode

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (v *jsiiProxy_VsCodeLaunchConfig) validateAddCommandInputParameters(cfg *VsCodeLaunchCommandInputEntry) error {
	if cfg == nil {
		return fmt.Errorf("parameter cfg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cfg, func() string { return "parameter cfg" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VsCodeLaunchConfig) validateAddConfigurationParameters(cfg *VsCodeLaunchConfigurationEntry) error {
	if cfg == nil {
		return fmt.Errorf("parameter cfg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cfg, func() string { return "parameter cfg" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VsCodeLaunchConfig) validateAddPickStringInputParameters(cfg *VsCodeLaunchPickStringInputEntry) error {
	if cfg == nil {
		return fmt.Errorf("parameter cfg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cfg, func() string { return "parameter cfg" }); err != nil {
		return err
	}

	return nil
}

func (v *jsiiProxy_VsCodeLaunchConfig) validateAddPromptStringInputParameters(cfg *VsCodeLaunchPromptStringInputEntry) error {
	if cfg == nil {
		return fmt.Errorf("parameter cfg is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cfg, func() string { return "parameter cfg" }); err != nil {
		return err
	}

	return nil
}

func validateVsCodeLaunchConfig_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateVsCodeLaunchConfig_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewVsCodeLaunchConfigParameters(vscode VsCode) error {
	if vscode == nil {
		return fmt.Errorf("parameter vscode is required, but nil was provided")
	}

	return nil
}

