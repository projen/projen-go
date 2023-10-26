//go:build !no_runtime_type_checking

package vscode

import (
	"fmt"
)

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

