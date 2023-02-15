//go:build no_runtime_type_checking

package vscode

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VsCodeSettings) validateAddSettingParameters(setting *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_VsCodeSettings) validateAddSettingsParameters(settings *map[string]interface{}, languages interface{}) error {
	return nil
}

func validateNewVsCodeSettingsParameters(vscode VsCode) error {
	return nil
}

