//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package vscode

import (
	"fmt"
)

func (v *jsiiProxy_VsCodeSettings) validateAddSettingParameters(setting *string, value interface{}) error {
	if setting == nil {
		return fmt.Errorf("parameter setting is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (v *jsiiProxy_VsCodeSettings) validateAddSettingsParameters(settings *map[string]interface{}, languages interface{}) error {
	if settings == nil {
		return fmt.Errorf("parameter settings is required, but nil was provided")
	}

	switch languages.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	default:
		return fmt.Errorf("parameter languages must be one of the allowed types: *string, *[]*string; received %#v (a %T)", languages, languages)
	}

	return nil
}

func validateNewVsCodeSettingsParameters(vscode VsCode) error {
	if vscode == nil {
		return fmt.Errorf("parameter vscode is required, but nil was provided")
	}

	return nil
}

