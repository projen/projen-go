//go:build no_runtime_type_checking

package vscode

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VsCodeLaunchConfig) validateAddConfigurationParameters(cfg *VsCodeLaunchConfigurationEntry) error {
	return nil
}

func validateVsCodeLaunchConfig_IsComponentParameters(x interface{}) error {
	return nil
}

func validateVsCodeLaunchConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewVsCodeLaunchConfigParameters(vscode VsCode) error {
	return nil
}

